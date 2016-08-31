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
	"github.com/nano-projects/nanogo/initial/conf"
	"github.com/nano-projects/nanogo/log"
	"github.com/nano-projects/nanogo/pom"
	"github.com/nano-projects/nanogo/template"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	textTemplate "text/template"
)

type ExecutorYml struct {
	*ExecutorWebapp
	n *New
}

func (e *ExecutorYml) Exec() error {
	log.Logger.Debugf("Executor yaml make mode")
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

func (e *ExecutorYml) loadYml() ([]byte, error) {
	ymlFilePath := filepath.Join(e.n.Conf.Path, e.n.Conf.Name.ArtifactId, "src/yml", "nanogo.yml")
	data, err := e.n.Conf.LoadYaml()
	if err != nil {
		return nil, err
	}

	tmp, err := textTemplate.New("DefinitionPom").Parse(string(data))
	if err != nil {
		return nil, err
	}

	if err := conf.WriteTemplate(ymlFilePath, tmp, e.n.Tmp); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(ymlFilePath)
}

func (e *ExecutorYml) makeModuleContext(schema *pom.Schema) error {
	for module, project := range (*schema).Projects {
		projectPath := filepath.Join(e.n.Conf.Path, e.n.Conf.Name.ArtifactId)
		if project.ArtifactId != e.n.Conf.Name.ArtifactId {
			if project.Packaging == "war" || project.Packaging == "ear" || project.ModuleType == "web" {
				modulePath := filepath.Join(projectPath, module)
				if tmp, err := template.Context(); err != nil {
					return err
				} else {
					if err := conf.WriteTemplate(filepath.Join(modulePath, "src/main/resources/context.properties"), tmp, e.n.Tmp); err != nil {
						return err
					}
				}

				return nil
			}
		}
	}

	return nil
}
