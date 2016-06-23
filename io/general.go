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
	yml := GeneralWebapp(_arg)
	if err := yaml.Unmarshal([]byte(yml), &schema); err != nil {
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
				mkdirBase(_arg)
				path := *arg.Path + *arg.ArtifactId
				writeFile(path + "/nanogo.yml", yml)
				writeFile(path + "/pom.xml", xml.Header + XML_LICENSE + string(value))
				writeFile(path + "/eclipse-code-template.xml", GeneralCodeTemplates())
				writeFile(path + "/eclipse-formatter.xml", GeneralEclipseCheckstyle())
			} else {
				var moduleType string
				if project.Packaging == "war" || project.Packaging == "ear" || project.ModuleType == "web" {
					moduleType = "web"
				}

				mkdir(*arg.Path + *arg.ArtifactId + "/" + module, _arg, moduleType)
				writeFile(*arg.Path + *arg.ArtifactId + "/" + module + "/pom.xml", xml.Header + XML_LICENSE + string(value))
			}
		}
	}
}

func mkdirBase(_arg *models.Argument) {
	arg := *_arg
	absolutePath := *arg.Path + *arg.ArtifactId

	os.MkdirAll(absolutePath, 0755)

	writeFile(absolutePath + "/.gitignore", GeneralGitIgnore())
	if !*arg.Findbugs {
		writeFile(absolutePath + "/findbugs-rules.xml", GeneralFindBugs())
	}

	if !*arg.Checkstyle {
		writeFile(absolutePath + "/checkstyle-rules.xml", GeneralCheckstyleRules())
		writeFile(absolutePath + "/checkstyle-suppressions.xml", GeneralCheckstyleSuppressions())
	}
}

func mkdir(absolutePath string, arg *models.Argument, moduleType string) {
	pack := "/" + strings.Replace(strings.Replace(*arg.GroupId, ".", "/", -1), "-", "/", -1) + "/" + strings.Replace(*arg.ArtifactId, "-", "/", -1)
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

		newGroupId := strings.Replace(*arg.GroupId, "-", ".", -1)
		newArtifactId := strings.Replace(*arg.ArtifactId, "-", ".", -1)

		writeFile(absolutePath + "/bin/bootstrap.sh", GeneralBootstrapShell(newGroupId, newArtifactId))
		writeFile(absolutePath + "/src/main/java" + pack + "/Bootstrap.java", GeneralBootstrap(newGroupId, newArtifactId))
		writeFile(absolutePath + "/src/main/resources/assembly.xml", GeneralAssembly())
		writeFile(absolutePath + "/src/main/resources/context.properties", GeneralContext(newGroupId, newArtifactId))
		writeFile(absolutePath + "/src/main/webapp/WEB-INF/jetty.xml", GeneralJettyXml(*arg.Port))
		writeFile(absolutePath + "/src/main/webapp/WEB-INF/web.xml", GeneralWebXml(*arg.ArtifactId))
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
