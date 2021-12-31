package proc

import (
	"fmt"
	"os"
)

func KillProcess(pid int) {
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = process.Kill()
	if err != nil {
		fmt.Println(err)
		return
	}
}
