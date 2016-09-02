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

package initial

import (
	"github.com/nano-projects/nanogo/initial/template"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/log"
	"github.com/nano-projects/nanogo/pom"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type ExecutorScheduler struct {
	*ExecutorWebapp
}

func (e *ExecutorScheduler) Exec() error {
	log.Logger.Debugf("Executor webapp make mode")
	if err := e.makeBaseDirectory(); err != nil {
		return err
	}

	data, err := e.loadYml()
	if err != nil {
		return err
	}

	schema := &pom.Schema{}
	if err := yaml.Unmarshal(data, schema); err != nil {
		return err
	}

	if err := e.makeBaseSource(); err != nil {
		return err
	}

	if err := e.makeModule(schema); err != nil {
		return err
	}

	if err := e.makeModuleContext(schema); err != nil {
		return err
	}

	return nil
}

func (e *ExecutorScheduler) loadYml() ([]byte, error) {
	ymlFilePath := filepath.Join(e.n.Conf.Path, e.n.Conf.Name.ArtifactId, "src/yml", "nanogo.yml")
	tmp, err := template.SchedulerPom()
	if err != nil {
		return nil, err
	}

	if err := io.WriteTemplate(ymlFilePath, tmp, e.n.Tmp); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(ymlFilePath)
}

func (e *ExecutorScheduler) makeModuleContext(schema *pom.Schema) error {
	for module, project := range (*schema).Projects {
		projectPath := filepath.Join(e.n.Conf.Path, e.n.Conf.Name.ArtifactId)
		if project.ArtifactId != e.n.Conf.Name.ArtifactId {
			if project.Packaging == "war" || project.Packaging == "ear" || project.ModuleType == "web" {
				modulePath := filepath.Join(projectPath, module)
				if tmp, err := template.SchedulerContext(); err != nil {
					return err
				} else {
					if err := io.WriteTemplate(filepath.Join(modulePath, "src/main/resources/context.properties"), tmp, e.n.Tmp); err != nil {
						return err
					}
				}

				return nil
			}
		}
	}

	return nil
}
