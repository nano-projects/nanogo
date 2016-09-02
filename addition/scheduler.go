package addition

import (
	"github.com/nano-projects/nanogo/addition/conf"
	"github.com/nano-projects/nanogo/addition/template"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/log"
	"os"
	"path/filepath"
)

type SchedulerExecutor struct {
	*WebappExecutor
}

func (e *SchedulerExecutor) Exec() error {
	e.makeTemplateConfig()
	if err := e.makeService(); err != nil {
		return err
	}

	if err := e.makeScheduler(); err != nil {
		return err
	}

	return nil
}

func (e *SchedulerExecutor) makeScheduler() error {
	schedulerSupportPath := filepath.Join(e.addConf.ProjectPath, e.addConf.ProjectName+conf.SchedulerSupport, "src/main/java", e.addConf.BaseDirectory)
	cmpPath := filepath.Join(schedulerSupportPath, "scheduler")
	if !io.IsDirExists(cmpPath) {
		if err := os.MkdirAll(cmpPath, io.FILE_MODE); err != nil {
			return err
		}
	}

	cmpItfPath := filepath.Join(schedulerSupportPath, "scheduler", e.tmpConf.SourceName+"Scheduler.java")
	if !io.IsFileExists(cmpItfPath) {
		tmp, err := template.Scheduler()
		if err != nil {
			return err
		}

		if err := io.WriteTemplate(cmpItfPath, tmp, e.tmpConf); err != nil {
			return err
		}
	} else {
		log.Logger.Warnf("Scheduler file already exists")
	}

	return nil
}
