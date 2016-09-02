// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conf

import (
	"encoding/xml"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/log"
	"github.com/nano-projects/nanogo/pom"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	Webapp           = "-webapp"
	WebappSupport    = "-webapp-support"
	Scheduler        = "-scheduler"
	SchedulerSupport = "-scheduler-support"
	Core             = "-core"
	Mapper           = "-mapper"
	Common           = "-common"
)

type AdditionConfig struct {
	Name          string
	ProjectPath   string
	ProjectName   string
	BasePackage   string
	BaseDirectory string
	IsWebapp      bool
	IsScheduler   bool
	Author        string
	Version       string
}

func Make(path, author, name string) (*AdditionConfig, error) {
	idx := strings.LastIndex(path, "/")
	projectName := path[idx+1:]

	isWebapp := false
	if io.IsDirExists(filepath.Join(path, projectName+Webapp)) {
		contextPath := filepath.Join(path, projectName+Webapp, "src/main/resources/context.properties")
		if io.IsFileExists(contextPath) {
			isWebapp = true
		} else {
			return nil, errors.New("Not found context properties file")
		}

		if !io.IsDirExists(filepath.Join(path, projectName+WebappSupport)) {
			return nil, errors.New("Not found webapp support module")
		}
	}

	isScheduler := false
	if io.IsDirExists(filepath.Join(path, projectName+Scheduler)) {
		contextPath := filepath.Join(path, projectName+Scheduler, "src/main/resources/context.properties")
		if io.IsFileExists(contextPath) {
			isScheduler = true
		} else {
			return nil, errors.New("Not found context properties file")
		}

		if !io.IsDirExists(filepath.Join(path, projectName+SchedulerSupport)) {
			return nil, errors.New("Not found scheduler support module")
		}
	}

	if !io.IsDirExists(filepath.Join(path, projectName+Core)) {
		return nil, errors.New("Not found core module")
	}

	if !io.IsDirExists(filepath.Join(path, projectName+Mapper)) {
		return nil, errors.New("Not found mapper module")
	}

	if !io.IsDirExists(filepath.Join(path, projectName+Common)) {
		return nil, errors.New("Not found common module")
	}

	if (isWebapp && isScheduler) || (!isWebapp && !isScheduler) {
		return nil, errors.New("There are and can only be make in a way")
	}

	data, err := ioutil.ReadFile(filepath.Join(path, "src/yml/nanogo.yml"))
	if err != nil {
		return nil, err
	}

	schema := &pom.Schema{}
	if err := yaml.Unmarshal(data, schema); err != nil {
		return nil, err
	}

	var suffix string
	if isWebapp {
		suffix = Webapp
	} else if isScheduler {
		suffix = Scheduler
	}

	project := schema.Project(projectName + suffix)
	basePackage := strings.Replace(project.Parent.GroupId, "-", ".", -1) + "." + strings.Replace(project.Parent.ArtifactId, "-", ".", -1)
	log.Logger.Debugf("Package: %v", basePackage)

	baseDirectory := filepath.Join(strings.Replace(strings.Replace(project.Parent.GroupId, "-", "/", -1), ".", "/", -1),
		strings.Replace(project.Parent.ArtifactId, "-", "/", -1))
	log.Logger.Debugf("Package Directory: %v", baseDirectory)

	xmlData, err := ioutil.ReadFile(filepath.Join(path, "pom.xml"))
	if err != nil {
		return nil, err
	}

	parentPom := &pom.Project{}
	if err := xml.Unmarshal(xmlData, parentPom); err != nil {
		return nil, err
	}

	return &AdditionConfig{
		Name:          name,
		ProjectPath:   path,
		ProjectName:   projectName,
		BasePackage:   basePackage,
		BaseDirectory: baseDirectory,
		IsWebapp:      isWebapp,
		IsScheduler:   isScheduler,
		Author:        author,
		Version:       strings.Replace(parentPom.Version, pom.Snapshot, "", -1),
	}, nil
}
