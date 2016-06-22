package main

import (
	"gopkg.in/yaml.v2"
	"github.com/nano-projects/nanogo/models"
	"github.com/nano-projects/nanogo/formats"
	"log"
	"fmt"
	"encoding/xml"
	"os"
	"strings"
	"bufio"
)

func main() {
	path := "/Users/yanghe/Works/____Go_Project____/____Workspaces____"
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	groupId := "cn.net.yto.test"
	artifactId := "nanogo"
	version := "0.0.1"

	general(path, groupId, artifactId, version)

}

func general(path, groupId, artifactId, version string) {
	schema := models.Schema{}
	if err := yaml.Unmarshal([]byte(formats.GeneralWebapp(groupId, artifactId, version)), &schema); err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	for module, project := range schema.Projects {
		project.Xmlns = formats.XMLNS
		project.XmlnsXsi = formats.XMLNS_XSI
		project.XsiSchemaLocation = formats.XSI_SCHEMA_LOCATION

		if value, err := xml.MarshalIndent(project, "", "    "); err != nil {
			log.Fatalf("error: %v", err)
			return
		} else {
			if module == artifactId {
				mkdirBase(path + artifactId)
				writeFile(path + artifactId + "/pom.xml", xml.Header + formats.XML_LICENSE + string(value))
			} else {
				mkdir(path + artifactId + "/" + module, groupId, artifactId, project.ModuleType)
				writeFile(path + artifactId + "/" + module + "/pom.xml", xml.Header + formats.XML_LICENSE + string(value))
			}
		}
	}
}

func mkdirBase(absolutePath string) {
	os.MkdirAll(absolutePath, 0755)

	writeFile(absolutePath + "/.gitignore", formats.GeneralGitIgnore())
	writeFile(absolutePath + "/findbugs-rules.xml", formats.GeneralFindBugs())
	writeFile(absolutePath + "/checkstyle-rules.xml", formats.GeneralCheckstyleRules())
	writeFile(absolutePath + "/checkstyle-suppressions.xml", formats.GeneralCheckstyleSuppressions())
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

		writeFile(absolutePath + "/bin/bootstrap.sh", formats.GeneralBootstrapShell(groupId, artifactId))
		writeFile(absolutePath + "/src/main/java" + pack + "/Bootstrap.java", formats.GeneralBootstrap(groupId, artifactId))
		writeFile(absolutePath + "/src/main/resources/assembly.xml", formats.GeneralAssembly())
		writeFile(absolutePath + "/src/main/resources/context.properties", formats.GeneralContext(groupId, artifactId))
		writeFile(absolutePath + "/src/main/webapp/WEB-INF/jetty.xml", formats.GeneralJettyXml("8080"))
		writeFile(absolutePath + "/src/main/webapp/WEB-INF/web.xml", formats.GeneralWebXml(artifactId))
		writeFile(absolutePath + "/src/main/webapp/WEB-INF/webdefault.xml", formats.GeneralWebDefaultXml())
		writeFile(absolutePath + "/src/main/webapp/index.jsp", formats.GeneralIndexJsp())
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
