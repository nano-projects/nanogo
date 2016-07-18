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
	"github.com/nano-projects/nanogo/models"
	"github.com/nano-projects/nanogo/resources"
	"github.com/nano-projects/nanogo/resources/license"
)

const (
	XMLNS               = "http://maven.apache.org/POM/4.0.0"
	XMLNS_XSI           = "http://www.w3.org/2001/XMLSchema-instance"
	XSI_SCHEMA_LOCATION = "http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd"

)

func GeneralWebapp(_arg *models.Argument) (pom string) {
	arg := *_arg
	pom = `
    projects:
      ` + *arg.ArtifactId + `:
        modelVersion: 4.0.0
        parent:
          groupId: ` + (*arg.Parent).GroupId + `
          artifactId: ` + (*arg.Parent).ArtifactId + `
          version: ` + (*arg.Parent).Version + `
        groupId: ` + *arg.GroupId + `
        artifactId: ` + *arg.ArtifactId + `
        version: ` + *arg.Version + `
        packaging: pom
        name: ` + *arg.ArtifactId + `
        url: http://maven.apache.org
        properties:
          project.build.sourceEncoding: UTF-8
        modules:
          module:
            - ` + *arg.ArtifactId + `-common
            - ` + *arg.ArtifactId + `-mapper
            - ` + *arg.ArtifactId + `-core
            - ` + *arg.ArtifactId + `-webapp-support
            - ` + *arg.ArtifactId + `-webapp
        dependencyManagement:
          dependencies:
            dependency:
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-common
                version: ${project.version}
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-mapper
                version: ${project.version}
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-core
                version: ${project.version}
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-webapp-support
                version: ${project.version}
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-webapp
                version: ${project.version}

        build:
          plugins:
            plugin:
              - groupId: org.apache.maven.plugins
                artifactId: maven-compiler-plugin
              ` + appendSourcePlugin(arg.Source) + `
              ` + appendJavadocPlugin(arg.Javadoc) + `
              ` + appendFindBugsPlugin(arg.Findbugs) + `
              ` + appendCheckstylePlugin(arg.Checkstyle) + `
              ` + appendLicensePlugin(arg.License) + `
              ` + appendPmdPlugin(arg.Pmd) + `

      ` + *arg.ArtifactId + `-common:
        modelVersion: 4.0.0
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-common
        name: ` + *arg.ArtifactId + ` Common
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

      ` + *arg.ArtifactId + `-mapper:
        modelVersion: 4.0.0
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-mapper
        name: ` + *arg.ArtifactId + ` Mapper
        url: http://maven.apache.org
        properties:
          cs.dir: ${project.parent.basedir}
        dependencies:
          dependency:
            - groupId: junit
              artifactId: junit
              scope: test
            - groupId: ` + *arg.GroupId + `
              artifactId: ` + *arg.ArtifactId + `-common
            - groupId: org.nanoframework
              artifactId: nano-orm-mybatis
            - groupId: mysql
              artifactId: mysql-connector-java
            - groupId: com.alibaba
              artifactId: druid

      ` + *arg.ArtifactId + `-core:
        modelVersion: 4.0.0
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-core
        name: ` + *arg.ArtifactId + ` Core
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
            - groupId: ` + *arg.GroupId + `
              artifactId: ` + *arg.ArtifactId + `-mapper

      ` + *arg.ArtifactId + `-webapp-support:
        modelVersion: 4.0.0
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-webapp-support
        name: ` + *arg.ArtifactId + ` Webapp Support
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
            - groupId: ` + *arg.GroupId + `
              artifactId: ` + *arg.ArtifactId + `-core

      ` + *arg.ArtifactId + `-webapp:
        modelVersion: 4.0.0
        moduleType: web
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-webapp
        packaging: jar
        name: ` + *arg.ArtifactId + ` Webapp
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
            - groupId: ` + *arg.GroupId + `
              artifactId: ` + *arg.ArtifactId + `-webapp-support
        build:
          finalName: ` + *arg.ArtifactId + `-webapp
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

	return pom
}

func GeneralScheduler(_arg *models.Argument) (pom string) {
	arg := *_arg
	pom = `
    projects:
      ` + *arg.ArtifactId + `:
        modelVersion: 4.0.0
        parent:
          groupId: ` + (*arg.Parent).GroupId + `
          artifactId: ` + (*arg.Parent).ArtifactId + `
          version: ` + (*arg.Parent).Version + `
        groupId: ` + *arg.GroupId + `
        artifactId: ` + *arg.ArtifactId + `
        version: ` + *arg.Version + `
        packaging: pom
        name: ` + *arg.ArtifactId + `
        url: http://maven.apache.org
        properties:
          project.build.sourceEncoding: UTF-8
        modules:
          module:
            - ` + *arg.ArtifactId + `-common
            - ` + *arg.ArtifactId + `-mapper
            - ` + *arg.ArtifactId + `-core
            - ` + *arg.ArtifactId + `-scheduler-support
            - ` + *arg.ArtifactId + `-scheduler
        dependencyManagement:
          dependencies:
            dependency:
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-common
                version: ${project.version}
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-mapper
                version: ${project.version}
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-core
                version: ${project.version}
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-scheduler-support
                version: ${project.version}
              - groupId: ` + *arg.GroupId + `
                artifactId: ` + *arg.ArtifactId + `-scheduler
                version: ${project.version}

        build:
          plugins:
            plugin:
              - groupId: org.apache.maven.plugins
                artifactId: maven-compiler-plugin
              ` + appendSourcePlugin(arg.Source) + `
              ` + appendJavadocPlugin(arg.Javadoc) + `
              ` + appendFindBugsPlugin(arg.Findbugs) + `
              ` + appendCheckstylePlugin(arg.Checkstyle) + `
              ` + appendLicensePlugin(arg.License) + `
              ` + appendPmdPlugin(arg.Pmd) + `

      ` + *arg.ArtifactId + `-common:
        modelVersion: 4.0.0
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-common
        name: ` + *arg.ArtifactId + ` Common
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

      ` + *arg.ArtifactId + `-mapper:
        modelVersion: 4.0.0
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-mapper
        name: ` + *arg.ArtifactId + ` Mapper
        url: http://maven.apache.org
        properties:
          cs.dir: ${project.parent.basedir}
        dependencies:
          dependency:
            - groupId: junit
              artifactId: junit
              scope: test
            - groupId: ` + *arg.GroupId + `
              artifactId: ` + *arg.ArtifactId + `-common
            - groupId: org.nanoframework
              artifactId: nano-orm-mybatis
            - groupId: mysql
              artifactId: mysql-connector-java
            - groupId: com.alibaba
              artifactId: druid

      ` + *arg.ArtifactId + `-core:
        modelVersion: 4.0.0
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-core
        name: ` + *arg.ArtifactId + ` Core
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
            - groupId: ` + *arg.GroupId + `
              artifactId: ` + *arg.ArtifactId + `-mapper

      ` + *arg.ArtifactId + `-scheduler-support:
        modelVersion: 4.0.0
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-scheduler-support
        name: ` + *arg.ArtifactId + ` Scheduler Support
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
              artifactId: nano-ext-concurrent
            - groupId: ` + *arg.GroupId + `
              artifactId: ` + *arg.ArtifactId + `-core

      ` + *arg.ArtifactId + `-scheduler:
        modelVersion: 4.0.0
        moduleType: web
        parent:
          groupId: ` + *arg.GroupId + `
          artifactId: ` + *arg.ArtifactId + `
          version: ` + *arg.Version + `
        artifactId: ` + *arg.ArtifactId + `-scheduler
        packaging: jar
        name: ` + *arg.ArtifactId + ` Scheduler
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
            - groupId: ` + *arg.GroupId + `
              artifactId: ` + *arg.ArtifactId + `-scheduler-support
        build:
          finalName: ` + *arg.ArtifactId + `-scheduler
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

	return pom
}

func appendSourcePlugin(remove *bool) (source string) {
	if !*remove {
		source = `- groupId: org.apache.maven.plugins
                artifactId: maven-source-plugin`
	}

	return
}

func appendJavadocPlugin(remove *bool) (javadoc string) {
	if !*remove {
		javadoc = `- groupId: org.apache.maven.plugins
                artifactId: maven-javadoc-plugin`
	}

	return
}

func appendFindBugsPlugin(remove *bool) (findbugs string) {
	if !*remove {
		findbugs = `- groupId: org.codehaus.mojo
                artifactId: findbugs-maven-plugin`
	}

	return
}

func appendCheckstylePlugin(remove *bool) (checkstyle string) {
	if !*remove {
		checkstyle = `- groupId: org.apache.maven.plugins
                artifactId: maven-checkstyle-plugin`
	}

	return
}

func appendLicensePlugin(remove *bool) (license string) {
	if !*remove {
		license = `- groupId: com.mycila
                artifactId: license-maven-plugin`
	}

	return
}

func appendPmdPlugin(remove *bool) (pmd string) {
	if !*remove {
		pmd = `- groupId: org.apache.maven.plugins
                artifactId: maven-pmd-plugin`
	}

	return
}

func GeneralJettyXml(port string) (jetty string) {
	jetty = resources.JettyXml(port)
	return
}

func GeneralWebXml(displayName string) (web string) {
	web = resources.WebXml(displayName)
	return
}

func GeneralWebDefaultXml() (webdefault string) {
	webdefault = resources.WebDefaultXml()
	return
}

func GeneralWebappContext(groupId, artifactId string) (context string) {
	pack := groupId + "." + artifactId + ".component"
	context = license.Properties() + `
# 组件服务上下文属性文件列表
context=

# 服务根
context.root=/` + artifactId + `

context.component-scan.base-package=` + pack + `
`
	return
}

func GeneralSchedulerContext(groupId, artifactId string) (context string) {
	pack := groupId + "." + artifactId + ".scheduler"
	context = license.Properties() + `
# 组件服务上下文属性文件列表
context=

# 服务根
context.root=/` + artifactId + `

context.scheduler-scan.base-package=` + pack + `
`
	return
}

func GeneralIndexJsp() (index string) {
	index = resources.IndexJsp()
	return
}

func GeneralAssembly() (assembly string) {
	assembly = resources.Assembly()
	return
}

func GeneralBootstrap(groupId, artifactId string) (bootstrap string) {
	pack := groupId + "." + artifactId
	bootstrap = resources.BootstrapClass(pack)
	return
}

func GeneralBootstrapShell(groupId, artifactId string) (shell string) {
	bootstrap := groupId + "." + artifactId + ".Bootstrap"
	shell = resources.Bootstrap(bootstrap)
	return
}

func GeneralFindBugs() (findbugs string) {
	findbugs = resources.Findbugs()
	return
}

func GeneralCheckstyleRules() (rules string) {
	rules = resources.CheckstyleRules()
	return
}

func GeneralCheckstyleSuppressions() (suppressions string) {
	suppressions = resources.CheckstyleSuppressions()
	return
}

func GeneralGitIgnore() (ignore string) {
	ignore = resources.IGNORE
	return
}

func GeneralCodeTemplates() (template string) {
	template = resources.CodeTemplate()
	return
}

func GeneralEclipseCheckstyle() (checkstyle string) {
	checkstyle = resources.CodeStyle()
	return
}

func GeneralLicenseHeader() (header string) {
	header = resources.Header()
	return
}

func GeneralLicenseHeaderDefinitions() (definitions string) {
	definitions = resources.Definitions()
	return
}

func GeneralMavenSettings() (settings string) {
	settings = resources.Settings()
	return
}