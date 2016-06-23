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
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

func GeneralDefaultWebapp(_arg *models.Argument) {
	schema := models.Schema{}
	yml := GeneralWebapp(_arg)
	if err := yaml.Unmarshal([]byte(yml), &schema); err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	generalDefault(_arg, &schema, &yml)
}

func GeneralDefaultScheduler(_arg *models.Argument) {
	schema := models.Schema{}
	yml := GeneralScheduler(_arg)
	if err := yaml.Unmarshal([]byte(yml), &schema); err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	generalDefault(_arg, &schema, &yml)
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
			if project.ArtifactId == *(*arg).ArtifactId {
				mkdirBase(arg, &project.ArtifactId)
				absolutePath := *(*arg).Path + project.ArtifactId
				writeFile(absolutePath+"/src/yml/nanogo.yml", PROPERTIES_LICENSE+*yml)
				writeFile(absolutePath+"/pom.xml", xml.Header+XML_LICENSE+string(value))
			} else {
				var moduleType string
				if project.Packaging == "war" || project.Packaging == "ear" || project.ModuleType == "web" {
					moduleType = "web"
				}

				absolutePath := *(*arg).Path + *(*arg).ArtifactId + "/" + module
				mkdir(&absolutePath, arg, (*arg).GroupId, (*arg).ArtifactId, &moduleType)
				writeFile(*(*arg).Path+*(*arg).ArtifactId+"/"+module+"/pom.xml", xml.Header+XML_LICENSE+string(value))
			}
		}
	}
}

func mkdirBase(arg *models.Argument, artifactId *string) {
	absolutePath := *(*arg).Path + *artifactId

	os.MkdirAll(absolutePath, 0755)
	os.MkdirAll(absolutePath+"/src/eclipse", 0755)
	os.MkdirAll(absolutePath+"/src/yml", 0755)

	writeFile(absolutePath+"/src/eclipse/eclipse-code-template.xml", GeneralCodeTemplates())
	writeFile(absolutePath+"/src/eclipse/eclipse-formatter.xml", GeneralEclipseCheckstyle())

	writeFile(absolutePath+"/.gitignore", GeneralGitIgnore())
	if !*(*arg).Findbugs {
		writeFile(absolutePath+"/findbugs-rules.xml", GeneralFindBugs())
	}

	if !*(*arg).Checkstyle {
		writeFile(absolutePath+"/checkstyle-rules.xml", GeneralCheckstyleRules())
		writeFile(absolutePath+"/checkstyle-suppressions.xml", GeneralCheckstyleSuppressions())
	}

	if !*(*arg).License {
		os.MkdirAll(absolutePath+"/src/licensing", 0755)
		writeFile(absolutePath+"/src/licensing/header-definitions.xml", GeneralLicenseHeaderDefinitions())
		writeFile(absolutePath+"/src/licensing/header.txt", GeneralLicenseHeader())
	}
}

func mkdir(absolutePath *string, arg *models.Argument, groupId *string, artifactId *string, moduleType *string) {
	pack := "/" + strings.Replace(strings.Replace(*groupId, ".", "/", -1), "-", "/", -1) + "/" + strings.Replace(*artifactId, "-", "/", -1)
	os.MkdirAll(*absolutePath, 0755)
	os.MkdirAll(*absolutePath+"/src/main/java"+pack, 0755)
	os.MkdirAll(*absolutePath+"/src/main/resources", 0755)
	os.MkdirAll(*absolutePath+"/src/test/java"+pack, 0755)
	os.MkdirAll(*absolutePath+"/src/test/resources", 0755)

	writeFile(*absolutePath+"/src/main/java"+pack+"/.gitkeep", "")
	writeFile(*absolutePath+"/src/main/resources/.gitkeep", "")
	writeFile(*absolutePath+"/src/test/java"+pack+"/.gitkeep", "")
	writeFile(*absolutePath+"/src/test/resources/.gitkeep", "")

	if *moduleType == "web" {
		os.MkdirAll(*absolutePath+"/src/main/webapp/WEB-INF", 0755)
		os.MkdirAll(*absolutePath+"/bin", 0755)
		os.MkdirAll(*absolutePath+"/configure/public", 0755)
		os.MkdirAll(*absolutePath+"/configure/sit", 0755)
		os.MkdirAll(*absolutePath+"/configure/uat", 0755)
		os.MkdirAll(*absolutePath+"/configure/release", 0755)

		writeFile(*absolutePath+"/configure/public/.gitkeep", "")
		writeFile(*absolutePath+"/configure/sit/.gitkeep", "")
		writeFile(*absolutePath+"/configure/uat/.gitkeep", "")
		writeFile(*absolutePath+"/configure/release/.gitkeep", "")

		newGroupId := strings.Replace(*groupId, "-", ".", -1)
		newArtifactId := strings.Replace(*artifactId, "-", ".", -1)

		writeFile(*absolutePath+"/bin/bootstrap.sh", GeneralBootstrapShell(newGroupId, newArtifactId))
		writeFile(*absolutePath+"/src/main/java"+pack+"/Bootstrap.java", GeneralBootstrap(newGroupId, newArtifactId))
		writeFile(*absolutePath+"/src/main/resources/assembly.xml", GeneralAssembly())

		if *(*arg).NewWebapp {
			writeFile(*absolutePath+"/src/main/resources/context.properties", GeneralWebappContext(newGroupId, newArtifactId))
		}

		if *(*arg).NewScheduler {
			writeFile(*absolutePath+"/src/main/resources/context.properties", GeneralSchedulerContext(newGroupId, newArtifactId))
		}

		writeFile(*absolutePath+"/src/main/webapp/WEB-INF/jetty.xml", GeneralJettyXml(*(*arg).Port))
		writeFile(*absolutePath+"/src/main/webapp/WEB-INF/web.xml", GeneralWebXml(*artifactId))
		writeFile(*absolutePath+"/src/main/webapp/WEB-INF/webdefault.xml", GeneralWebDefaultXml())
		writeFile(*absolutePath+"/src/main/webapp/index.jsp", GeneralIndexJsp())
	}
}

func writeFile(fileName, data string) {
	fmt.Println("create file: ", fileName)
	outputFile, outputError := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(data)
	outputWriter.Flush()
}
