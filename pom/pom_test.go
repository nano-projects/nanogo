package pom

import (
	"encoding/xml"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/test"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMakeProject(t *testing.T) {
	test.DebugMode(t)

	if err := io.WriteFile("pom.xml", pom()); err != nil {
		t.Error(err)
		return
	} else {
		t.Log("Initial pom.xml")
	}

	defer func() {
		if err := os.Remove("pom.xml"); err != nil {
			t.Error(err)
			return
		} else {
			t.Log("Removed pom.xml")
		}
	}()

	data, err := ioutil.ReadFile("pom.xml")
	if err != nil {
		t.Error(err)
		return
	}

	project := (&Schema{}).MakeProject()
	if err := xml.Unmarshal(data, project); err != nil {
		t.Error(err)
		return
	}

	t.Logf("Project: %v", project)
	t.Logf("deps: %v.%v", project.GroupId, strings.Replace(project.ArtifactId, "-", ".", -1))
}

func pom() string {
	return `
<?xml version="1.0" encoding="UTF-8"?>
<!--
    Copyright © 2015-2016 the original author or authors.

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
-->
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>org.nanoframework</groupId>
        <artifactId>super</artifactId>
        <version>0.0.11-SNAPSHOT</version>
    </parent>
    <groupId>org.nanoframework.nanogo</groupId>
    <artifactId>testgo-ap</artifactId>
    <version>0.0.1-SNAPSHOT</version>
</project>
	`
}
