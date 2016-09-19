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

import "text/template"

const (
	XMLNS               = "http://maven.apache.org/POM/4.0.0"
	XMLNS_XSI           = "http://www.w3.org/2001/XMLSchema-instance"
	XSI_SCHEMA_LOCATION = "http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd"
)

func WebappPom() (*template.Template, error) {
	pom := `
projects:
  {{.ArtifactId}}:
    modelVersion: 4.0.0
    parent:
      groupId: {{.ParentGroupId}}
      artifactId: {{.ParentArtifactId}}
      version: {{.ParentVersion}}
    groupId: {{.GroupId}}
    artifactId: {{.ArtifactId}}
    version: {{.Version}}
    packaging: pom
    name: {{.ArtifactId}}
    url: http://maven.apache.org
    properties:
      project.build.sourceEncoding: UTF-8
    modules:
      module:
        - {{.ArtifactId}}-common
        - {{.ArtifactId}}-mapper
        - {{.ArtifactId}}-core
        - {{.ArtifactId}}-webapp-support
        - {{.ArtifactId}}-webapp
    dependencyManagement:
      dependencies:
        dependency:
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-common
            version: ${project.version}
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-mapper
            version: ${project.version}
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-core
            version: ${project.version}
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-webapp-support
            version: ${project.version}
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-webapp
            version: ${project.version}

    build:
      plugins:
        plugin:
          - groupId: org.apache.maven.plugins
            artifactId: maven-compiler-plugin
          - groupId: org.apache.maven.plugins
            artifactId: maven-source-plugin
          - groupId: org.apache.maven.plugins
            artifactId: maven-javadoc-plugin
          - groupId: org.codehaus.mojo
            artifactId: findbugs-maven-plugin
          - groupId: org.apache.maven.plugins
            artifactId: maven-checkstyle-plugin
          - groupId: com.mycila
            artifactId: license-maven-plugin
          - groupId: org.apache.maven.plugins
            artifactId: maven-pmd-plugin

  {{.ArtifactId}}-common:
    modelVersion: 4.0.0
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-common
    name: {{.ArtifactId}} Common
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: org.slf4j
          artifactId: slf4j-api
        - groupId: org.slf4j
          artifactId: jcl-over-slf4j
        - groupId: org.apache.logging.log4j
          artifactId: log4j-api
          scope: runtime
        - groupId: org.apache.logging.log4j
          artifactId: log4j-core
          scope: runtime
        - groupId: org.apache.logging.log4j
          artifactId: log4j-slf4j-impl
          scope: runtime
        - groupId: org.nanoframework
          artifactId: nano-commons

  {{.ArtifactId}}-mapper:
    modelVersion: 4.0.0
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-mapper
    name: {{.ArtifactId}} Mapper
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: {{.GroupId}}
          artifactId: {{.ArtifactId}}-common
        - groupId: org.nanoframework
          artifactId: nano-orm-mybatis
        - groupId: mysql
          artifactId: mysql-connector-java
        - groupId: com.alibaba
          artifactId: druid

  {{.ArtifactId}}-core:
    modelVersion: 4.0.0
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-core
    name: {{.ArtifactId}} Core
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: javax.servlet
          artifactId: javax.servlet-api
          scope: provided
        - groupId: org.nanoframework
          artifactId: nano-core
        - groupId: {{.GroupId}}
          artifactId: {{.ArtifactId}}-mapper

  {{.ArtifactId}}-webapp-support:
    modelVersion: 4.0.0
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-webapp-support
    name: {{.ArtifactId}} Webapp Support
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: javax.servlet
          artifactId: javax.servlet-api
          scope: provided
        - groupId: org.nanoframework
          artifactId: nano-webmvc
        - groupId: {{.GroupId}}
          artifactId: {{.ArtifactId}}-core

  {{.ArtifactId}}-webapp:
    modelVersion: 4.0.0
    moduleType: web
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-webapp
    packaging: jar
    name: {{.ArtifactId}} Webapp
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: org.eclipse.jetty.orbit
          artifactId: javax.servlet.jsp
        - groupId: javax.servlet
          artifactId: jstl
        - groupId: org.nanoframework
          artifactId: nano-server
          exclusions:
            exclusion:
              - groupId: javax.servlet.jsp
                artifactId: javax.servlet.jsp-api
        - groupId: {{.GroupId}}
          artifactId: {{.ArtifactId}}-webapp-support
    build:
      finalName: {{.ArtifactId}}-webapp
      resources:
        resource:
          - directory: src/main/webapp
            targetPath: ${project.basedir}/webRoot/
      testResources:
        testResource:
          - directory: src/test/resources
      plugins:
        plugin:
          - groupId: org.apache.maven.plugins
            artifactId: maven-assembly-plugin
            configuration:
              finalName: ${project.artifactId}-${project.version}
              descriptors:
                descriptor:
                  - src/main/resources/assembly.xml
            executions:
              execution:
                - phase: package
                  goals:
                    goal:
                      - single
    profiles:
      profile:
        - id: dev
          activation:
            activeByDefault: true
          build:
            resources:
              resource:
                - directory: src/main/resources
                  excludes:
                    exclude:
                      - assembly.xml
        - id: sit
          build:
            resources:
              resource:
                - directory: configure/public
                  targetPath: ${project.basedir}/target/conf/
                - directory: configure/sit
                  targetPath: ${project.basedir}/target/conf/
                  excludes:
                    exclude:
                      - assembly.xml
        - id: uat
          build:
            resources:
              resource:
                - directory: configure/public
                  targetPath: ${project.basedir}/target/conf/
                - directory: configure/uat
                  targetPath: ${project.basedir}/target/conf/
                  excludes:
                    exclude:
                      - assembly.xml
        - id: release
          build:
            resources:
              resource:
                - directory: configure/public
                  targetPath: ${project.basedir}/target/conf/
                - directory: configure/release
                  targetPath: ${project.basedir}/target/conf/
                  excludes:
                    exclude:
                      - assembly.xml

`

	return template.New("WebappPom").Parse(pom)
}

func SchedulerPom() (*template.Template, error) {
	pom := `
projects:
  {{.ArtifactId}}:
    modelVersion: 4.0.0
    parent:
      groupId: {{.ParentGroupId}}
      artifactId: {{.ParentArtifactId}}
      version: {{.ParentVersion}}
    groupId: {{.GroupId}}
    artifactId: {{.ArtifactId}}
    version: {{.Version}}
    packaging: pom
    name: {{.ArtifactId}}
    url: http://maven.apache.org
    properties:
      project.build.sourceEncoding: UTF-8
    modules:
      module:
        - {{.ArtifactId}}-common
        - {{.ArtifactId}}-mapper
        - {{.ArtifactId}}-core
        - {{.ArtifactId}}-scheduler-support
        - {{.ArtifactId}}-scheduler
    dependencyManagement:
      dependencies:
        dependency:
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-common
            version: ${project.version}
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-mapper
            version: ${project.version}
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-core
            version: ${project.version}
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-scheduler-support
            version: ${project.version}
          - groupId: {{.GroupId}}
            artifactId: {{.ArtifactId}}-scheduler
            version: ${project.version}

    build:
      plugins:
        plugin:
          - groupId: org.apache.maven.plugins
            artifactId: maven-compiler-plugin
          - groupId: org.apache.maven.plugins
            artifactId: maven-source-plugin
          - groupId: org.apache.maven.plugins
            artifactId: maven-javadoc-plugin
          - groupId: org.codehaus.mojo
            artifactId: findbugs-maven-plugin
          - groupId: org.apache.maven.plugins
            artifactId: maven-checkstyle-plugin
          - groupId: com.mycila
            artifactId: license-maven-plugin
          - groupId: org.apache.maven.plugins
            artifactId: maven-pmd-plugin

  {{.ArtifactId}}-common:
    modelVersion: 4.0.0
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-common
    name: {{.ArtifactId}} Common
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: org.slf4j
          artifactId: slf4j-api
        - groupId: org.slf4j
          artifactId: jcl-over-slf4j
        - groupId: org.apache.logging.log4j
          artifactId: log4j-api
          scope: runtime
        - groupId: org.apache.logging.log4j
          artifactId: log4j-core
          scope: runtime
        - groupId: org.apache.logging.log4j
          artifactId: log4j-slf4j-impl
          scope: runtime
        - groupId: org.nanoframework
          artifactId: nano-commons

  {{.ArtifactId}}-mapper:
    modelVersion: 4.0.0
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-mapper
    name: {{.ArtifactId}} Mapper
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: {{.GroupId}}
          artifactId: {{.ArtifactId}}-common
        - groupId: org.nanoframework
          artifactId: nano-orm-mybatis
        - groupId: mysql
          artifactId: mysql-connector-java
        - groupId: com.alibaba
          artifactId: druid

  {{.ArtifactId}}-core:
    modelVersion: 4.0.0
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-core
    name: {{.ArtifactId}} Core
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: javax.servlet
          artifactId: javax.servlet-api
          scope: provided
        - groupId: org.nanoframework
          artifactId: nano-core
        - groupId: {{.GroupId}}
          artifactId: {{.ArtifactId}}-mapper

  {{.ArtifactId}}-scheduler-support:
    modelVersion: 4.0.0
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-scheduler-support
    name: {{.ArtifactId}} Scheduler Support
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: javax.servlet
          artifactId: javax.servlet-api
          scope: provided
        - groupId: org.nanoframework
          artifactId: nano-webmvc
        - groupId: org.nanoframework
          artifactId: nano-concurrent
        - groupId: {{.GroupId}}
          artifactId: {{.ArtifactId}}-core

  {{.ArtifactId}}-scheduler:
    modelVersion: 4.0.0
    moduleType: web
    parent:
      groupId: {{.GroupId}}
      artifactId: {{.ArtifactId}}
      version: {{.Version}}
    artifactId: {{.ArtifactId}}-scheduler
    packaging: jar
    name: {{.ArtifactId}} Scheduler
    url: http://maven.apache.org
    properties:
      cs.dir: ${project.parent.basedir}
    dependencies:
      dependency:
        - groupId: junit
          artifactId: junit
          scope: test
        - groupId: org.eclipse.jetty.orbit
          artifactId: javax.servlet.jsp
        - groupId: javax.servlet
          artifactId: jstl
        - groupId: org.nanoframework
          artifactId: nano-server
          exclusions:
            exclusion:
              - groupId: javax.servlet.jsp
                artifactId: javax.servlet.jsp-api
        - groupId: {{.GroupId}}
          artifactId: {{.ArtifactId}}-scheduler-support
    build:
      finalName: {{.ArtifactId}}-scheduler
      resources:
        resource:
          - directory: src/main/webapp
            targetPath: ${project.basedir}/webRoot/
      testResources:
        testResource:
          - directory: src/test/resources
      plugins:
        plugin:
          - groupId: org.apache.maven.plugins
            artifactId: maven-assembly-plugin
            configuration:
              finalName: ${project.artifactId}-${project.version}
              descriptors:
                descriptor:
                  - src/main/resources/assembly.xml
            executions:
              execution:
                - phase: package
                  goals:
                    goal:
                      - single
    profiles:
      profile:
        - id: dev
          activation:
            activeByDefault: true
          build:
            resources:
              resource:
                - directory: src/main/resources
                  excludes:
                    exclude:
                      - assembly.xml
        - id: sit
          build:
            resources:
              resource:
                - directory: configure/public
                  targetPath: ${project.basedir}/target/conf/
                - directory: configure/sit
                  targetPath: ${project.basedir}/target/conf/
                  excludes:
                    exclude:
                      - assembly.xml
        - id: uat
          build:
            resources:
              resource:
                - directory: configure/public
                  targetPath: ${project.basedir}/target/conf/
                - directory: configure/uat
                  targetPath: ${project.basedir}/target/conf/
                  excludes:
                    exclude:
                      - assembly.xml
        - id: release
          build:
            resources:
              resource:
                - directory: configure/public
                  targetPath: ${project.basedir}/target/conf/
                - directory: configure/release
                  targetPath: ${project.basedir}/target/conf/
                  excludes:
                    exclude:
                      - assembly.xml

`

	return template.New("SchedulerPom").Parse(pom)
}
