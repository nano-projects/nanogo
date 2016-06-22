package io

import (
	"gopkg.in/yaml.v2"
	"github.com/nano-projects/nanogo/models"
	"log"
	"fmt"
	"encoding/xml"
	"os"
	"strings"
	"bufio"
)

func GeneralDefaultWebapp(_arg *models.Argument) {
	arg := *_arg
	schema := models.Schema{}
	if err := yaml.Unmarshal([]byte(GeneralWebapp(_arg)), &schema); err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	for module, project := range schema.Projects {
		project.Xmlns = XMLNS
		project.XmlnsXsi = XMLNS_XSI
		project.XsiSchemaLocation = XSI_SCHEMA_LOCATION

		if value, err := xml.MarshalIndent(project, "", "    "); err != nil {
			log.Fatalf("error: %v", err)
			return
		} else {
			if module == *arg.ArtifactId {
				mkdirBase(*arg.Path + *arg.ArtifactId)
				writeFile(*arg.Path + *arg.ArtifactId + "/pom.xml", xml.Header + XML_LICENSE + string(value))
			} else {
				var moduleType string
				if project.Packaging == "war" || project.Packaging == "ear" || project.ModuleType == "web" {
					moduleType = "web"
				}

				mkdir(*arg.Path + *arg.ArtifactId + "/" + module, *arg.GroupId, *arg.ArtifactId, moduleType)
				writeFile(*arg.Path + *arg.ArtifactId + "/" + module + "/pom.xml", xml.Header + XML_LICENSE + string(value))
			}
		}
	}
}

func mkdirBase(absolutePath string) {
	os.MkdirAll(absolutePath, 0755)

	writeFile(absolutePath + "/.gitignore", GeneralGitIgnore())
	writeFile(absolutePath + "/findbugs-rules.xml", GeneralFindBugs())
	writeFile(absolutePath + "/checkstyle-rules.xml", GeneralCheckstyleRules())
	writeFile(absolutePath + "/checkstyle-suppressions.xml", GeneralCheckstyleSuppressions())
}

func mkdir(absolutePath, groupId, artifactId, moduleType string) {
	pack := "/" + strings.Replace(groupId, ".", "/", -1) + "/" + artifactId
	os.MkdirAll(absolutePath, 0755)
	os.MkdirAll(absolutePath + "/src/main/java" + pack, 0755)
	os.MkdirAll(absolutePath + "/src/main/resources", 0755)
	os.MkdirAll(absolutePath + "/src/test/java" + pack, 0755)
	os.MkdirAll(absolutePath + "/src/test/resources", 0755)

	writeFile(absolutePath + "/src/main/java" + pack + "/.gitkeep", "")
	writeFile(absolutePath + "/src/main/resources/.gitkeep", "")
	writeFile(absolutePath + "/src/test/java" + pack + "/.gitkeep", "")
	writeFile(absolutePath + "/src/test/resources/.gitkeep", "")

	if moduleType == "web" {
		os.MkdirAll(absolutePath + "/src/main/webapp/WEB-INF", 0755)
		os.MkdirAll(absolutePath + "/bin", 0755)
		os.MkdirAll(absolutePath + "/configure/public", 0755)
		os.MkdirAll(absolutePath + "/configure/sit", 0755)
		os.MkdirAll(absolutePath + "/configure/uat", 0755)
		os.MkdirAll(absolutePath + "/configure/release", 0755)

		writeFile(absolutePath + "/bin/bootstrap.sh", GeneralBootstrapShell(groupId, artifactId))
		writeFile(absolutePath + "/src/main/java" + pack + "/Bootstrap.java", GeneralBootstrap(groupId, artifactId))
		writeFile(absolutePath + "/src/main/resources/assembly.xml", GeneralAssembly())
		writeFile(absolutePath + "/src/main/resources/context.properties", GeneralContext(groupId, artifactId))
		writeFile(absolutePath + "/src/main/webapp/WEB-INF/jetty.xml", GeneralJettyXml("8080"))
		writeFile(absolutePath + "/src/main/webapp/WEB-INF/web.xml", GeneralWebXml(artifactId))
		writeFile(absolutePath + "/src/main/webapp/WEB-INF/webdefault.xml", GeneralWebDefaultXml())
		writeFile(absolutePath + "/src/main/webapp/index.jsp", GeneralIndexJsp())
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
