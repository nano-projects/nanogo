/*
	Copyright 2015-2016 the original author or authors.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package io

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"github.com/nano-projects/nanogo/models"
	"github.com/nano-projects/nanogo/resources"
	"github.com/nano-projects/nanogo/resources/license"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

const (
	FILE_MODE = 0755
)

func GeneralDefaultWebapp(arg *models.Argument) {
	schema := models.Schema{}
	yml := GeneralWebapp(arg)
	if err := yaml.Unmarshal([]byte(yml), &schema); err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	generalDefault(arg, &schema, &yml)
}

func GeneralDefaultScheduler(arg *models.Argument) {
	schema := models.Schema{}
	yml := GeneralScheduler(arg)
	if err := yaml.Unmarshal([]byte(yml), &schema); err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	generalDefault(arg, &schema, &yml)
}

func generalDefault(arg *models.Argument, schema *models.Schema, yml *string) {
	for module, project := range (*schema).Projects {
		project.Xmlns = XMLNS
		project.XmlnsXsi = XMLNS_XSI
		project.XsiSchemaLocation = XSI_SCHEMA_LOCATION

		if value, err := xml.MarshalIndent(project, "", "    "); err != nil {
			log.Fatalf("error: %v", err)
			return
		} else {
			groupId := *(*arg).GroupId
			artifact := *(*arg).ArtifactId
			path := *(*arg).Path
			if project.ArtifactId == artifact {
				mkdirBase(arg, artifact)
				absolutePath := path + artifact
				writeFile(absolutePath+"/src/yml/nanogo.yml", license.Properties()+*yml)
				writeFile(absolutePath+"/pom.xml", xml.Header+license.Xml()+string(value))
			} else {
				var moduleType string
				if project.Packaging == "war" || project.Packaging == "ear" || project.ModuleType == "web" {
					moduleType = "web"
				}

				absolutePath := path + artifact + "/" + module
				mkdir(absolutePath, arg, groupId, artifact, moduleType)
				writeFile(path+artifact+"/"+module+"/pom.xml", xml.Header+license.Xml()+string(value))
			}
		}
	}
}

func mkdirBase(arg *models.Argument, artifactId string) {
	absolutePath := *(*arg).Path + artifactId

	os.MkdirAll(absolutePath, FILE_MODE)
	os.MkdirAll(absolutePath+"/src/eclipse", FILE_MODE)
	os.MkdirAll(absolutePath+"/src/mvn", FILE_MODE)
	os.MkdirAll(absolutePath+"/src/yml", FILE_MODE)

	writeFile(absolutePath+"/src/eclipse/eclipse-code-template.xml", resources.CodeTemplate())
	writeFile(absolutePath+"/src/eclipse/eclipse-formatter.xml", resources.CodeStyle())
	writeFile(absolutePath+"/src/mvn/settings.xml", resources.Settings())

	writeFile(absolutePath+"/.gitignore", resources.IGNORE)
	if !*(*arg).Findbugs {
		writeFile(absolutePath+"/findbugs-rules.xml", resources.Findbugs())
	}

	if !*(*arg).Checkstyle {
		writeFile(absolutePath+"/checkstyle-rules.xml", resources.CheckstyleRules())
		writeFile(absolutePath+"/checkstyle-suppressions.xml", resources.CheckstyleSuppressions())
	}

	if !*(*arg).License {
		os.MkdirAll(absolutePath+"/src/licensing", FILE_MODE)
		writeFile(absolutePath+"/src/licensing/header-definitions.xml", resources.Definitions())
		writeFile(absolutePath+"/src/licensing/header.txt", resources.Header())
	}
}

func mkdir(absolutePath string, arg *models.Argument, groupId string, artifactId string, moduleType string) {
	pack := "/" + strings.Replace(strings.Replace(groupId, ".", "/", -1), "-", "/", -1) + "/" + strings.Replace(artifactId, "-", "/", -1)
	os.MkdirAll(absolutePath, FILE_MODE)
	os.MkdirAll(absolutePath+"/src/main/java"+pack, FILE_MODE)
	os.MkdirAll(absolutePath+"/src/main/resources", FILE_MODE)
	os.MkdirAll(absolutePath+"/src/test/java"+pack, FILE_MODE)
	os.MkdirAll(absolutePath+"/src/test/resources", FILE_MODE)

	writeFile(absolutePath+"/src/main/java"+pack+"/.gitkeep", "")
	writeFile(absolutePath+"/src/main/resources/.gitkeep", "")
	writeFile(absolutePath+"/src/test/java"+pack+"/.gitkeep", "")
	writeFile(absolutePath+"/src/test/resources/.gitkeep", "")

	if moduleType == "web" {
		os.MkdirAll(absolutePath+"/src/main/webapp/WEB-INF", FILE_MODE)
		os.MkdirAll(absolutePath+"/bin", FILE_MODE)
		os.MkdirAll(absolutePath+"/configure/public", FILE_MODE)
		os.MkdirAll(absolutePath+"/configure/sit", FILE_MODE)
		os.MkdirAll(absolutePath+"/configure/uat", FILE_MODE)
		os.MkdirAll(absolutePath+"/configure/release", FILE_MODE)

		writeFile(absolutePath+"/configure/public/.gitkeep", "")
		writeFile(absolutePath+"/configure/sit/.gitkeep", "")
		writeFile(absolutePath+"/configure/uat/.gitkeep", "")
		writeFile(absolutePath+"/configure/release/.gitkeep", "")

		newGroupId := strings.Replace(groupId, "-", ".", -1)
		newArtifactId := strings.Replace(artifactId, "-", ".", -1)
		writeFile(absolutePath+"/bin/bootstrap.sh", resources.Bootstrap(newGroupId + "." + newArtifactId + ".Bootstrap"))
		writeFile(absolutePath+"/src/main/java"+pack+"/Bootstrap.java", resources.BootstrapClass(newGroupId + "." + newArtifactId))
		writeFile(absolutePath+"/src/main/resources/assembly.xml", resources.Assembly())

		if *(*arg).NewWebapp {
			writeFile(absolutePath+"/src/main/resources/context.properties", GeneralWebappContext(newGroupId, newArtifactId))
		}

		if *(*arg).NewScheduler {
			writeFile(absolutePath+"/src/main/resources/context.properties", GeneralSchedulerContext(newGroupId, newArtifactId))
		}

		writeFile(absolutePath+"/src/main/webapp/WEB-INF/jetty.xml", resources.JettyXml(*(*arg).Port))
		writeFile(absolutePath+"/src/main/webapp/WEB-INF/web.xml", resources.WebXml(artifactId))
		writeFile(absolutePath+"/src/main/webapp/WEB-INF/webdefault.xml", resources.WebDefaultXml())
		writeFile(absolutePath+"/src/main/webapp/index.jsp", resources.IndexJsp())
	}
}

func writeFile(fileName, data string) {
	fmt.Println("create file: ", fileName)
	outputFile, outputError := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, FILE_MODE)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(data)
	outputWriter.Flush()
}
