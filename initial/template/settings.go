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

package template

import (
	"github.com/nano-projects/nanogo/initial/template/license"
	"text/template"
)

func Settings() (*template.Template, error) {
	settings := `<?xml version="1.0" encoding="UTF-8"?>
` + license.Xml() + `
<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0 http://maven.apache.org/xsd/settings-1.0.0.xsd">

    <localRepository>${user.home}/.m2/repository</localRepository>
    <pluginGroups></pluginGroups>
    <proxies></proxies>
    <servers></servers>

    <mirrors>
        <mirror>
            <id>oss-snapshots</id>
            <mirrorOf>oss-snapshots</mirrorOf>
            <name>oss-snapshots</name>
            <url>https://oss.sonatype.org/content/repositories/snapshots</url>
        </mirror>
        <mirror>
            <id>oss-releases</id>
            <mirrorOf>oss-releases</mirrorOf>
            <name>oss-releases</name>
            <url>https://oss.sonatype.org/content/repositories/releases</url>
        </mirror>
    </mirrors>

    <profiles>
        <profile>
            <id>jdk-1.8</id>
            <activation>
                <activeByDefault>true</activeByDefault>
                <jdk>1.8</jdk>
            </activation>
            <properties>
                <maven.compiler.source>1.8</maven.compiler.source>
                <maven.compiler.target>1.8</maven.compiler.target>
                <maven.compiler.compilerVersion>3.1</maven.compiler.compilerVersion>
            </properties>
        </profile>
    </profiles>
</settings>
`

	return template.New("Settings").Parse(settings)
}
