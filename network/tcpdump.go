package network

import (
	"fmt"
	"os"
	"os/exec"
)

func StartTcpdump(pcap_filepath string) *os.Process {
	fmt.Printf("%s\n", pcap_filepath)
	ls, err := exec.Command("ls").Output()
	fmt.Printf("ls:\n%s :Error:\n%v\n", ls, err)
	tcpdump := exec.Command("sh", "-c", fmt.Sprintf("su & tcpdump -w %s", pcap_filepath)) // TODO: command injection handling
	tcpdump.Start()
	fmt.Printf("tcpdump shell PID:\n%d", tcpdump.Process.Pid)
	return tcpdump.Process
}
