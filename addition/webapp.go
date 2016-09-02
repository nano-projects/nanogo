package addition

import (
	"github.com/nano-projects/nanogo/addition/conf"
	"github.com/nano-projects/nanogo/addition/template"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type WebappExecutor struct {
	addConf conf.AdditionConfig
	tmpConf conf.TmpConfig
}

func (e *WebappExecutor) Exec() error {
	e.makeTemplateConfig()
	if err := e.makeService(); err != nil {
		return err
	}

	if err := e.makeComponent(); err != nil {
		return err
	}

	return nil
}

func (e *WebappExecutor) makeTemplateConfig() {
	sourceName := strings.ToUpper(e.addConf.Name[:1]) + e.addConf.Name[1:]
	aliasSourceName := strings.ToLower(e.addConf.Name[:1]) + e.addConf.Name[1:]
	e.tmpConf = conf.TmpConfig{
		Package:         e.addConf.BasePackage,
		SourceName:      sourceName,
		AliasSourceName: aliasSourceName,
		Author:          e.addConf.Author,
		Version:         e.addConf.Version,
		Year:            strconv.Itoa(time.Now().Year()),
	}
}

func (e *WebappExecutor) makeComponent() error {
	webappSupportPath := filepath.Join(e.addConf.ProjectPath, e.addConf.ProjectName+conf.WebappSupport, "src/main/java", e.addConf.BaseDirectory)
	cmpPath := filepath.Join(webappSupportPath, "component/impl")
	if !io.IsDirExists(cmpPath) {
		if err := os.MkdirAll(cmpPath, io.FILE_MODE); err != nil {
			return err
		}
	}

	cmpItfPath := filepath.Join(webappSupportPath, "component", e.tmpConf.SourceName+"Component.java")
	if !io.IsFileExists(cmpItfPath) {
		tmp, err := template.Component()
		if err != nil {
			return err
		}

		if err := io.WriteTemplate(cmpItfPath, tmp, e.tmpConf); err != nil {
			return err
		}
	} else {
		log.Logger.Warnf("Component file already exists")
	}

	cmpImplPath := filepath.Join(webappSupportPath, "component/impl", e.tmpConf.SourceName+"ComponentImpl.java")
	if !io.IsFileExists(cmpImplPath) {
		tmp, err := template.ComponentImpl()
		if err != nil {
			return err
		}

		if err := io.WriteTemplate(cmpImplPath, tmp, e.tmpConf); err != nil {
			return err
		}
	} else {
		log.Logger.Warnf("Component Implement file already exists")
	}

	return nil
}

func (e *WebappExecutor) makeService() error {
	corePath := filepath.Join(e.addConf.ProjectPath, e.addConf.ProjectName+conf.Core, "src/main/java", e.addConf.BaseDirectory)
	servPath := filepath.Join(corePath, "service/impl")
	if !io.IsDirExists(servPath) {
		if err := os.MkdirAll(servPath, io.FILE_MODE); err != nil {
			return err
		}
	}

	servItfPath := filepath.Join(corePath, "service", e.tmpConf.SourceName+"Service.java")
	if !io.IsFileExists(servItfPath) {
		tmp, err := template.Service()
		if err != nil {
			return err
		}

		if err := io.WriteTemplate(servItfPath, tmp, e.tmpConf); err != nil {
			return err
		}
	} else {
		log.Logger.Warnf("Service file already exists")
	}

	servImplPath := filepath.Join(corePath, "service/impl", e.tmpConf.SourceName+"ServiceImpl.java")
	if !io.IsFileExists(servImplPath) {
		tmp, err := template.ServiceImpl()
		if err != nil {
			return err
		}

		if err := io.WriteTemplate(servImplPath, tmp, e.tmpConf); err != nil {
			return err
		}
	} else {
		log.Logger.Warnf("Service Implement file already exists")
	}

	return nil
}
