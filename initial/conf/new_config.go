// Copyright © 2015-2016 River Yang <comicme_yanghe@nanoframework.org>
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
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/log"
	"github.com/nano-projects/nanogo/pom"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

const (
	DefaultYamlFileName = "nanogo.yml"
)

type NewConfig struct {
	Web bool

	Scheduler bool

	Path string

	Template string

	Parent pom.Dependency

	Name pom.Dependency

	Publish uint
}

func (conf *NewConfig) existsYaml() bool {
	var path string
	if conf.Template != "" {
		path = conf.Template
	} else {
		path = filepath.Join(conf.Path, DefaultYamlFileName)
	}

	return io.IsFileExists(path)
}

func (conf *NewConfig) LoadYaml() ([]byte, error) {
	var path string
	if conf.Template != "" {
		path = conf.Template
	} else {
		path = filepath.Join(conf.Path, DefaultYamlFileName)
	}

	return ioutil.ReadFile(path)
}

func (conf *NewConfig) Valid() error {
	if conf.Web && conf.Scheduler {
		return errors.New(`Cannot specify both "web" and "scheduler"`)
	}

	if !conf.Web && !conf.Scheduler {
		if !conf.existsYaml() {
			return errors.New("Not found nanogo.yml file")
		}
	}

	if conf.Publish < 1024 {
		return errors.New("Publish cannot be less than 1024")
	}

	return nil
}

func WriteTemplate(path string, tmp *template.Template, conf TmpConfig) error {
	log.Logger.Infof("create file: %v", path)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, io.FILE_MODE)
	if err != nil {
		return err
	}

	defer file.Close()
	tmp.Execute(file, conf)
	return nil
}
