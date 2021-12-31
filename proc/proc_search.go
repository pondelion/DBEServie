package proc

import (
	"github.com/mitchellh/go-ps"
	"github.com/pondelion/dbeservice/model"
)

func SearchTcpdumpProcess() model.Proc {
	processes, err := ps.Processes()

	if err != nil {
		panic(err)
	}

	var procs []model.Proc
	for _, p := range processes {
		procs = append(procs, model.Proc{
			PID:          p.Pid(),
			PPID:         p.PPid(),
			PROCESS_NAME: p.Executable(),
		})
	}

	var tcpdumpProc model.Proc
	for _, p := range procs {
		if p.PROCESS_NAME == "tcpdump" {
			tcpdumpProc = p
		}
	}

	return tcpdumpProc
}
