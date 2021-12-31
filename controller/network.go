package controller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pondelion/dbeservice/network"
	"github.com/pondelion/dbeservice/proc"
)

var tcpdumpShellProcess *os.Process
var pcap_filepath string

func StartTcpdump(c *gin.Context) {
	if tcpdumpShellProcess != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "tcpdump is already started",
		})
	}
	save_dir := c.Param("save_dir")
	log.Printf("save_dir : %s", save_dir)
	filename := strconv.FormatInt(time.Now().Unix(), 10) + ".pcap"
	pcap_filepath = filepath.Join(save_dir, filename)
	// util.WriteProcMemInt64(pid, addr, value)
	tcpdumpShellProcess = network.StartTcpdump(pcap_filepath)

	c.JSON(http.StatusOK, gin.H{
		"message":       "successfully started tcpdump",
		"pcap_filepath": pcap_filepath,
	})
}

func StopTcpdump(c *gin.Context) {
	if tcpdumpShellProcess != nil {
		// Kill tcpdump process
		tcpdumpProc := proc.SearchTcpdumpProcess()
		proc.KillProcess(tcpdumpProc.PID)
		// Kill tcpdump's parent shell process
		pid := tcpdumpShellProcess.Pid
		proc.KillProcess(pid)
		tcpdumpShellProcess = nil
		c.JSON(http.StatusOK, gin.H{
			"message":       "killed tcpdump process : " + strconv.Itoa(pid),
			"pcap_filepath": pcap_filepath,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "tcpdump is not started",
		})
	}
}
