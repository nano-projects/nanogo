package addition

import (
	"github.com/nano-projects/nanogo/addition/conf"
	"github.com/nano-projects/nanogo/exec"
	"github.com/pkg/errors"
)

type Addition struct {
	AddConf conf.AdditionConfig
}

func (add *Addition) Run() error {
	exec, err := add.withExecutor()
	if err != nil {
		return err
	}

	return exec.Exec()
}

func (add *Addition) withExecutor() (exec.Executor, error) {
	if add.AddConf.IsWebapp {
		return &WebappExecutor{addConf: add.AddConf}, nil
	}

	if add.AddConf.IsScheduler {
		return &SchedulerExecutor{&WebappExecutor{addConf: add.AddConf}}, nil
	}

	return nil, errors.New("There are and can only be make in a way")
}
