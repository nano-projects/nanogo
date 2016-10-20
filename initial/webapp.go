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
	"encoding/xml"
	"github.com/nano-projects/nanogo/initial/template"
	"github.com/nano-projects/nanogo/initial/template/license"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/log"
	"github.com/nano-projects/nanogo/pom"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	textTemplate "text/template"
)

type ExecutorWebapp struct {
	n *Initial
}

func (e *ExecutorWebapp) Exec() error {
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

func (e *ExecutorWebapp) loadYml() ([]byte, error) {
	ymlFilePath := filepath.Join(e.n.Conf.Path, e.n.Conf.Name.ArtifactId, "src/yml", "nanogo.yml")
	tmp, err := template.WebappPom()
	if err != nil {
		return nil, err
	}

	if err := io.WriteTemplate(ymlFilePath, tmp, e.n.Tmp); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(ymlFilePath)
}

func (e *ExecutorWebapp) makeBaseDirectory() error {
	projectPath := filepath.Join(e.n.Conf.Path, e.n.Conf.Name.ArtifactId)
	if io.IsDirExists(projectPath) {
		return errors.New("Project already exists. Please delete it and try again.")
	}

	if err := os.MkdirAll(projectPath, io.FILE_MODE); err != nil {
		return err
	}

	eclipsePath := filepath.Join(projectPath, "src/eclipse")
	if err := os.MkdirAll(eclipsePath, io.FILE_MODE); err != nil {
		return err
	}

	mvnPath := filepath.Join(projectPath, "src/mvn")
	if err := os.MkdirAll(mvnPath, io.FILE_MODE); err != nil {
		return err
	}

	ymlPath := filepath.Join(projectPath, "src/yml")
	if err := os.MkdirAll(ymlPath, io.FILE_MODE); err != nil {
		return err
	}

	licensePath := filepath.Join(projectPath, "src/licensing")
	if err := os.MkdirAll(licensePath, io.FILE_MODE); err != nil {
		return err
	}

	return nil
}

func (e *ExecutorWebapp) makeBaseSource() error {
	projectPath := filepath.Join(e.n.Conf.Path, e.n.Conf.Name.ArtifactId)
	if tmp, err := template.CodeTemplate(); err != nil {
		return err
	} else {
		if err := io.WriteTemplate(filepath.Join(projectPath, "src/eclipse/eclipse-code-template.xml"), tmp, e.n.Tmp); err != nil {
			return err
		}
	}

	if tmp, err := template.CodeStyle(); err != nil {
		return err
	} else {
		if err := io.WriteTemplate(filepath.Join(projectPath, "src/eclipse/eclipse-code-style.xml"), tmp, e.n.Tmp); err != nil {
			return err
		}
	}

	if tmp, err := template.Settings(); err != nil {
		return err
	} else {
		if err := io.WriteTemplate(filepath.Join(projectPath, "src/mvn/settings.xml"), tmp, e.n.Tmp); err != nil {
			return err
		}
	}

	if err := io.WriteFile(filepath.Join(projectPath, ".gitignore"), template.IGNORE); err != nil {
		return err
	}

	if tmp, err := template.Findbugs(); err != nil {
		return err
	} else {
		if err := io.WriteTemplate(filepath.Join(projectPath, "findbugs-rules.xml"), tmp, e.n.Tmp); err != nil {
			return err
		}
	}

	if tmp, err := template.CheckstyleRules(); err != nil {
		return err
	} else {
		if err := io.WriteTemplate(filepath.Join(projectPath, "checkstyle-rules.xml"), tmp, e.n.Tmp); err != nil {
			return err
		}
	}

	if tmp, err := template.CheckstyleSuppressions(); err != nil {
		return err
	} else {
		if err := io.WriteTemplate(filepath.Join(projectPath, "checkstyle-suppressions.xml"), tmp, e.n.Tmp); err != nil {
			return err
		}
	}

	if tmp, err := template.Header(); err != nil {
		return err
	} else {
		if err := io.WriteTemplate(filepath.Join(projectPath, "src/licensing/header.txt"), tmp, e.n.Tmp); err != nil {
			return err
		}
	}

	if tmp, err := template.Definitions(); err != nil {
		return err
	} else {
		if err := io.WriteTemplate(filepath.Join(projectPath, "src/licensing/header-definitions.xml"), tmp, e.n.Tmp); err != nil {
			return err
		}
	}

	return nil
}

func (e *ExecutorWebapp) makeModule(schema *pom.Schema) error {
	for module, project := range (*schema).Projects {
		project.Xmlns = template.XMLNS
		project.XmlnsXsi = template.XMLNS_XSI
		project.XsiSchemaLocation = template.XSI_SCHEMA_LOCATION

		projectPath := filepath.Join(e.n.Conf.Path, e.n.Conf.Name.ArtifactId)
		if data, err := xml.MarshalIndent(project, "", "    "); err != nil {
			return err
		} else {
			if project.ArtifactId == e.n.Conf.Name.ArtifactId {
				if err := e.makePom(data, projectPath); err != nil {
					return err
				}
			} else {
				var moduleType string
				if project.Packaging == "war" || project.Packaging == "ear" || project.ModuleType == "web" {
					moduleType = "web"
				}

				modulePath := filepath.Join(projectPath, module)
				if err := e.makeModuleDirectory(modulePath, moduleType); err != nil {
					return err
				}

				if err := e.makePom(data, modulePath); err != nil {
					return err
				}

				if err := e.makeModuleSource(modulePath, moduleType); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (e *ExecutorWebapp) makePom(data []byte, projectPath string) error {
	tmp, err := textTemplate.New("Pom").Parse(xml.Header + license.Xml() + string(data))
	if err != nil {
		return err
	}

	return io.WriteTemplate(filepath.Join(projectPath, "pom.xml"), tmp, e.n.Tmp)
}

func (e *ExecutorWebapp) makeModuleDirectory(modulePath string, moduleType string) error {
	if err := os.MkdirAll(modulePath, io.FILE_MODE); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(modulePath, "src/main/java", e.n.Tmp.Package), io.FILE_MODE); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(modulePath, "src/main/resources"), io.FILE_MODE); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(modulePath, "src/test/java", e.n.Tmp.Package), io.FILE_MODE); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(modulePath, "src/test/resources"), io.FILE_MODE); err != nil {
		return err
	}

	if moduleType == "web" {
		if err := os.MkdirAll(filepath.Join(modulePath, "src/main/webapp/WEB-INF"), io.FILE_MODE); err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Join(modulePath, "bin"), io.FILE_MODE); err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Join(modulePath, "configure/public"), io.FILE_MODE); err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Join(modulePath, "configure/sit"), io.FILE_MODE); err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Join(modulePath, "configure/uat"), io.FILE_MODE); err != nil {
			return err
		}

		if err := os.MkdirAll(filepath.Join(modulePath, "configure/release"), io.FILE_MODE); err != nil {
			return err
		}
	}

	return nil
}

func (e *ExecutorWebapp) makeModuleSource(modulePath string, moduleType string) error {
	if err := io.WriteFile(filepath.Join(modulePath, "src/main/java", e.n.Tmp.Package, ".gitkeep"), ""); err != nil {
		return err
	}

	if err := io.WriteFile(filepath.Join(modulePath, "src/main/resources/.gitkeep"), ""); err != nil {
		return err
	}

	if err := io.WriteFile(filepath.Join(modulePath, "src/test/java", e.n.Tmp.Package, ".gitkeep"), ""); err != nil {
		return err
	}

	if err := io.WriteFile(filepath.Join(modulePath, "src/test/resources/.gitkeep"), ""); err != nil {
		return err
	}

	if moduleType == "web" {
		if err := io.WriteFile(filepath.Join(modulePath, "configure/public/.gitkeep"), ""); err != nil {
			return err
		}

		if err := io.WriteFile(filepath.Join(modulePath, "configure/sit/.gitkeep"), ""); err != nil {
			return err
		}

		if err := io.WriteFile(filepath.Join(modulePath, "configure/uat/.gitkeep"), ""); err != nil {
			return err
		}

		if err := io.WriteFile(filepath.Join(modulePath, "configure/release/.gitkeep"), ""); err != nil {
			return err
		}

		if tmp, err := template.Bootstrap(); err != nil {
			return err
		} else {
			if err := io.WriteTemplate(filepath.Join(modulePath, "bin/bootstrap.sh"), tmp, e.n.Tmp); err != nil {
				return err
			}
		}

		if tmp, err := template.BootstrapClass(); err != nil {
			return err
		} else {
			if err := io.WriteTemplate(filepath.Join(modulePath, "src/main/java", e.n.Tmp.Package, e.n.Tmp.BootstrapClassName+".java"), tmp, e.n.Tmp); err != nil {
				return err
			}
		}

		if tmp, err := template.Assembly(); err != nil {
			return err
		} else {
			if err := io.WriteTemplate(filepath.Join(modulePath, "src/main/resources/assembly.xml"), tmp, e.n.Tmp); err != nil {
				return err
			}
		}

		if e.n.Tmp.Server == "Jetty" {
			if tmp, err := template.JettyXml(); err != nil {
				return err
			} else {
				if err := io.WriteTemplate(filepath.Join(modulePath, "src/main/webapp/WEB-INF/jetty.xml"), tmp, e.n.Tmp); err != nil {
					return err
				}
			}
		}

		if tmp, err := template.WebXml(); err != nil {
			return err
		} else {
			if err := io.WriteTemplate(filepath.Join(modulePath, "src/main/webapp/WEB-INF/web.xml"), tmp, e.n.Tmp); err != nil {
				return err
			}
		}

		if e.n.Tmp.Server == "Jetty" {
			if tmp, err := template.WebDefaultXml(); err != nil {
				return err
			} else {
				if err := io.WriteTemplate(filepath.Join(modulePath, "src/main/webapp/WEB-INF/webdefault.xml"), tmp, e.n.Tmp); err != nil {
					return err
				}
			}
		}

		if tmp, err := template.IndexJsp(); err != nil {
			return err
		} else {
			if err := io.WriteTemplate(filepath.Join(modulePath, "src/main/webapp/index.jsp"), tmp, e.n.Tmp); err != nil {
				return err
			}
		}

	}

	return nil
}

func (e *ExecutorWebapp) makeModuleContext(schema *pom.Schema) error {
	for module, project := range (*schema).Projects {
		projectPath := filepath.Join(e.n.Conf.Path, e.n.Conf.Name.ArtifactId)
		if project.ArtifactId != e.n.Conf.Name.ArtifactId {
			if project.Packaging == "war" || project.Packaging == "ear" || project.ModuleType == "web" {
				modulePath := filepath.Join(projectPath, module)
				if tmp, err := template.WebappContext(); err != nil {
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
