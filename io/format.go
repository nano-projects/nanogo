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
	"strconv"
	"time"
)

var (
	XMLNS               = "http://maven.apache.org/POM/4.0.0"
	XMLNS_XSI           = "http://www.w3.org/2001/XMLSchema-instance"
	XSI_SCHEMA_LOCATION = "http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd"

	YEAR = strconv.Itoa(time.Now().Year())

	JSP_LICENSE = `<%--
    Copyright 2015-` + YEAR + ` the original author or authors.

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
--%>
`

	XML_LICENSE = `<!--
    Copyright 2015-` + YEAR + ` the original author or authors.

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
`

	PROPERTIES_LICENSE = `#
# Copyright 2015-` + YEAR + ` the original author or authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
`

	CLASS_LICENSE = `/*
 * Copyright 2015-` + YEAR + ` the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
 `
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
        distributionManagement:
          repository:
            id: releases
            name: Releases
            url: http://10.1.195.225:8081/nexus/content/repositories/releases
          snapshotRepository:
            id: snapshots
            name: Snapshots
            url: http://10.1.195.225:8081/nexus/content/repositories/snapshots
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
        packaging: war
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
        distributionManagement:
          repository:
            id: releases
            name: Releases
            url: http://10.1.195.225:8081/nexus/content/repositories/releases
          snapshotRepository:
            id: snapshots
            name: Snapshots
            url: http://10.1.195.225:8081/nexus/content/repositories/snapshots
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
	jetty = `<?xml version="1.0" encoding="UTF-8"?>
` + XML_LICENSE + `
<!DOCTYPE Configure PUBLIC "-//Jetty//Configure//EN" "http://www.eclipse.org/jetty/configure.dtd">

<Configure id="Server" class="org.eclipse.jetty.server.Server">
    <Set name="ThreadPool">
        <!-- Default queued blocking threadpool -->
        <New class="org.eclipse.jetty.util.thread.QueuedThreadPool">
            <Set name="minThreads">5</Set>
            <Set name="maxThreads">200</Set>
            <Set name="detailedDump">false</Set>
        </New>
    </Set>

    <Call name="addConnector">
        <Arg>
            <New class="org.eclipse.jetty.server.nio.SelectChannelConnector">
                <Set name="host">
                    <Property name="jetty.host" />
                </Set>
                <Set name="port">
                    <Property name="jetty.port" default="` + port + `" />
                </Set>
                <Set name="maxIdleTime">300000</Set>
                <Set name="Acceptors">2</Set>
                <Set name="statsOn">false</Set>
                <Set name="confidentialPort">8443</Set>
                <Set name="lowResourcesConnections">20000</Set>
                <Set name="lowResourcesMaxIdleTime">5000</Set>
            </New>
        </Arg>
    </Call>
</Configure>
`

	return
}

func GeneralWebXml(displayName string) (web string) {
	web = XML_LICENSE + `
<web-app xmlns="http://java.sun.com/xml/ns/javaee" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://java.sun.com/xml/ns/javaee  http://java.sun.com/xml/ns/javaee/web-app_3_0.xsd"
    version="3.0">
    <display-name>` + displayName + `</display-name>

    <filter>
        <filter-name>httpRequestFilter</filter-name>
        <filter-class>org.nanoframework.web.server.filter.HttpRequestFilter</filter-class>
    </filter>

    <filter-mapping>
        <filter-name>httpRequestFilter</filter-name>
        <url-pattern>/*</url-pattern>
    </filter-mapping>

    <servlet>
        <servlet-name>Dispatcher-Servlet</servlet-name>
        <servlet-class>org.nanoframework.web.server.servlet.DispatcherServlet</servlet-class>
        <init-param>
            <param-name>context</param-name>
            <param-value>/context.properties</param-value>
        </init-param>
        <load-on-startup>1</load-on-startup>
    </servlet>

    <servlet-mapping>
        <servlet-name>Dispatcher-Servlet</servlet-name>
        <url-pattern>/dispatcher/*</url-pattern>
    </servlet-mapping>

    <welcome-file-list>
        <welcome-file>index.jsp</welcome-file>
    </welcome-file-list>
</web-app>
`
	return
}

func GeneralWebDefaultXml() (webdefault string) {
	webdefault = `<?xml version="1.0" encoding="ISO-8859-1"?>
` + XML_LICENSE + `

<!-- ===================================================================== -->
<!-- This file contains the default descriptor for web applications.       -->
<!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -->
<!-- The intent of this descriptor is to include jetty specific or common  -->
<!-- configuration for all webapps.   If a context has a webdefault.xml    -->
<!-- descriptor, it is applied before the contexts own web.xml file        -->
<!--                                                                       -->
<!-- A context may be assigned a default descriptor by:                    -->
<!--  + Calling WebApplicationContext.setDefaultsDescriptor                -->
<!--  + Passed an arg to addWebApplications                                -->
<!--                                                                       -->
<!-- This file is used both as the resource within the jetty.jar (which is -->
<!-- used as the default if no explicit defaults descriptor is set) and it -->
<!-- is copied to the etc directory of the Jetty distro and explicitly     -->
<!-- by the jetty.xml file.                                                -->
<!--                                                                       -->
<!-- ===================================================================== -->
<web-app
   xmlns="http://java.sun.com/xml/ns/javaee"
   xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
   xsi:schemaLocation="http://java.sun.com/xml/ns/javaee http://java.sun.com/xml/ns/javaee/web-app_2_5.xsd"
   metadata-complete="true"
   version="2.5">

  <description>
    Default web.xml file.
    This file is applied to a Web application before it's own WEB_INF/web.xml file
  </description>


  <!-- ==================================================================== -->
  <!-- Context params to control Session Cookies                            -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <!-- UNCOMMENT TO ACTIVATE
  <context-param>
    <param-name>org.mortbay.jetty.servlet.SessionDomain</param-name>
    <param-value>127.0.0.1</param-value>
  </context-param>

  <context-param>
    <param-name>org.mortbay.jetty.servlet.SessionPath</param-name>
    <param-value>/</param-value>
  </context-param>

  <context-param>
    <param-name>org.mortbay.jetty.servlet.MaxAge</param-name>
    <param-value>-1</param-value>
  </context-param>
  -->

  <context-param>
    <param-name>org.mortbay.jetty.webapp.NoTLDJarPattern</param-name>
    <param-value>start.jar|ant-.*\.jar|dojo-.*\.jar|jetty-.*\.jar|jsp-api-.*\.jar|junit-.*\.jar|servlet-api-.*\.jar|dnsns\.jar|rt\.jar|jsse\.jar|tools\.jar|sunpkcs11\.jar|sunjce_provider\.jar|xerces.*\.jar</param-value>
  </context-param>



  <!-- ==================================================================== -->
  <!-- The default servlet.                                                 -->
  <!-- This servlet, normally mapped to /, provides the handling for static -->
  <!-- content, OPTIONS and TRACE methods for the context.                  -->
  <!-- The following initParameters are supported:                          -->
  <!--                                                                      -->
  <!--   acceptRanges     If true, range requests and responses are         -->
  <!--                    supported                                         -->
  <!--                                                                      -->
  <!--   dirAllowed       If true, directory listings are returned if no    -->
  <!--                    welcome file is found. Else 403 Forbidden.        -->
  <!--                                                                      -->
  <!--   welcomeServlets  If true, attempt to dispatch to welcome files     -->
  <!--                    that are servlets, if no matching static          -->
  <!--                    resources can be found.                           -->
  <!--                                                                      -->
  <!--   redirectWelcome  If true, redirect welcome file requests           -->
  <!--                    else use request dispatcher forwards              -->
  <!--                                                                      -->
  <!--   gzip             If set to true, then static content will be served-->
  <!--                    as gzip content encoded if a matching resource is -->
  <!--                    found ending with ".gz"                           -->
  <!--                                                                      -->
  <!--   resoureBase      Can be set to replace the context resource base   -->
  <!--                                                                      -->
  <!--   relativeResourceBase                                               -->
  <!--                    Set with a pathname relative to the base of the   -->
  <!--                    servlet context root. Useful for only serving     -->
  <!--                    static content from only specific subdirectories. -->
  <!--                                                                      -->
  <!--   useFileMappedBuffer                                                -->
  <!--                    If set to true (the default), a  memory mapped    -->
  <!--                    file buffer will be used to serve static content  -->
  <!--                    when using an NIO connector. Setting this value   -->
  <!--                    to false means that a direct buffer will be used  -->
  <!--                    instead. If you are having trouble with Windows   -->
  <!--                    file locking, set this to false.                  -->
  <!--                                                                      -->
  <!--  cacheControl      If set, all static content will have this value   -->
  <!--                    set as the cache-control header.                  -->
  <!--                                                                      -->
  <!--  maxCacheSize      Maximum size of the static resource cache         -->
  <!--                                                                      -->
  <!--  maxCachedFileSize Maximum size of any single file in the cache      -->
  <!--                                                                      -->
  <!--  maxCachedFiles    Maximum number of files in the cache              -->
  <!--                                                                      -->
  <!--  cacheType         "nio", "bio" or "both" to determine the type(s)   -->
  <!--                    of resource cache. A bio cached buffer may be used-->
  <!--                    by nio but is not as efficient as a nio buffer.   -->
  <!--                    An nio cached buffer may not be used by bio.      -->
  <!--                                                                      -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <servlet>
    <servlet-name>default</servlet-name>
    <servlet-class>org.eclipse.jetty.servlet.DefaultServlet</servlet-class>
    <init-param>
      <param-name>acceptRanges</param-name>
      <param-value>true</param-value>
    </init-param>
    <init-param>
      <param-name>dirAllowed</param-name>
      <param-value>true</param-value>
    </init-param>
    <init-param>
      <param-name>welcomeServlets</param-name>
      <param-value>false</param-value>
    </init-param>
    <init-param>
      <param-name>redirectWelcome</param-name>
      <param-value>false</param-value>
    </init-param>
    <init-param>
      <param-name>maxCacheSize</param-name>
      <param-value>256000000</param-value>
    </init-param>
    <init-param>
      <param-name>maxCachedFileSize</param-name>
      <param-value>10000000</param-value>
    </init-param>
    <init-param>
      <param-name>maxCachedFiles</param-name>
      <param-value>1000</param-value>
    </init-param>
    <init-param>
      <param-name>cacheType</param-name>
      <param-value>both</param-value>
    </init-param>
    <init-param>
      <param-name>gzip</param-name>
      <param-value>true</param-value>
    </init-param>
    <init-param>
      <param-name>useFileMappedBuffer</param-name>
      <param-value>false</param-value>
    </init-param>
    <!--
    <init-param>
      <param-name>cacheControl</param-name>
      <param-value>max-age=3600,public</param-value>
    </init-param>
    -->
    <load-on-startup>0</load-on-startup>
  </servlet>

  <servlet-mapping> <servlet-name>default</servlet-name> <url-pattern>/</url-pattern> </servlet-mapping>


  <!-- ==================================================================== -->
  <!-- JSP Servlet                                                          -->
  <!-- This is the jasper JSP servlet from the jakarta project              -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <!-- The JSP page compiler and execution servlet, which is the mechanism  -->
  <!-- used by Glassfish to support JSP pages.  Traditionally, this servlet -->
  <!-- is mapped to URL patterh "*.jsp".  This servlet supports the         -->
  <!-- following initialization parameters (default values are in square    -->
  <!-- brackets):                                                           -->
  <!--                                                                      -->
  <!--   checkInterval       If development is false and reloading is true, -->
  <!--                       background compiles are enabled. checkInterval -->
  <!--                       is the time in seconds between checks to see   -->
  <!--                       if a JSP page needs to be recompiled. [300]    -->
  <!--                                                                      -->
  <!--   compiler            Which compiler Ant should use to compile JSP   -->
  <!--                       pages.  See the Ant documenation for more      -->
  <!--                       information. [javac]                           -->
  <!--                                                                      -->
  <!--   classdebuginfo      Should the class file be compiled with         -->
  <!--                       debugging information?  [true]                 -->
  <!--                                                                      -->
  <!--   classpath           What class path should I use while compiling   -->
  <!--                       generated servlets?  [Created dynamically      -->
  <!--                       based on the current web application]          -->
  <!--                       Set to ? to make the container explicitly set  -->
  <!--                       this parameter.                                -->
  <!--                                                                      -->
  <!--   development         Is Jasper used in development mode (will check -->
  <!--                       for JSP modification on every access)?  [true] -->
  <!--                                                                      -->
  <!--   enablePooling       Determines whether tag handler pooling is      -->
  <!--                       enabled  [true]                                -->
  <!--                                                                      -->
  <!--   fork                Tell Ant to fork compiles of JSP pages so that -->
  <!--                       a separate JVM is used for JSP page compiles   -->
  <!--                       from the one Tomcat is running in. [true]      -->
  <!--                                                                      -->
  <!--   ieClassId           The class-id value to be sent to Internet      -->
  <!--                       Explorer when using <jsp:plugin> tags.         -->
  <!--                       [clsid:8AD9C840-044E-11D1-B3E9-00805F499D93]   -->
  <!--                                                                      -->
  <!--   javaEncoding        Java file encoding to use for generating java  -->
  <!--                       source files. [UTF-8]                          -->
  <!--                                                                      -->
  <!--   keepgenerated       Should we keep the generated Java source code  -->
  <!--                       for each page instead of deleting it? [true]   -->
  <!--                                                                      -->
  <!--   logVerbosityLevel   The level of detailed messages to be produced  -->
  <!--                       by this servlet.  Increasing levels cause the  -->
  <!--                       generation of more messages.  Valid values are -->
  <!--                       FATAL, ERROR, WARNING, INFORMATION, and DEBUG. -->
  <!--                       [WARNING]                                      -->
  <!--                                                                      -->
  <!--   mappedfile          Should we generate static content with one     -->
  <!--                       print statement per input line, to ease        -->
  <!--                       debugging?  [false]                            -->
  <!--                                                                      -->
  <!--                                                                      -->
  <!--   reloading           Should Jasper check for modified JSPs?  [true] -->
  <!--                                                                      -->
  <!--   suppressSmap        Should the generation of SMAP info for JSR45   -->
  <!--                       debugging be suppressed?  [false]              -->
  <!--                                                                      -->
  <!--   dumpSmap            Should the SMAP info for JSR45 debugging be    -->
  <!--                       dumped to a file? [false]                      -->
  <!--                       False if suppressSmap is true                  -->
  <!--                                                                      -->
  <!--   scratchdir          What scratch directory should we use when      -->
  <!--                       compiling JSP pages?  [default work directory  -->
  <!--                       for the current web application]               -->
  <!--                                                                      -->
  <!--   tagpoolMaxSize      The maximum tag handler pool size  [5]         -->
  <!--                                                                      -->
  <!--   xpoweredBy          Determines whether X-Powered-By response       -->
  <!--                       header is added by generated servlet  [false]  -->
  <!--                                                                      -->
  <!-- If you wish to use Jikes to compile JSP pages:                       -->
  <!--   Set the init parameter "compiler" to "jikes".  Define              -->
  <!--   the property "-Dbuild.compiler.emacs=true" when starting Jetty     -->
  <!--   to cause Jikes to emit error messages in a format compatible with  -->
  <!--   Jasper.                                                            -->
  <!--   If you get an error reporting that jikes can't use UTF-8 encoding, -->
  <!--   try setting the init parameter "javaEncoding" to "ISO-8859-1".     -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <servlet id="jsp">
    <servlet-name>jsp</servlet-name>
    <servlet-class>org.apache.jasper.servlet.JspServlet</servlet-class>
    <init-param>
        <param-name>logVerbosityLevel</param-name>
        <param-value>DEBUG</param-value>
    </init-param>
    <init-param>
        <param-name>fork</param-name>
        <param-value>false</param-value>
    </init-param>
    <init-param>
        <param-name>xpoweredBy</param-name>
        <param-value>false</param-value>
    </init-param>
    <!--
    <init-param>
        <param-name>classpath</param-name>
        <param-value>?</param-value>
    </init-param>
    -->
    <load-on-startup>0</load-on-startup>
  </servlet>

  <servlet-mapping>
    <servlet-name>jsp</servlet-name>
    <url-pattern>*.jsp</url-pattern>
    <url-pattern>*.jspf</url-pattern>
    <url-pattern>*.jspx</url-pattern>
    <url-pattern>*.xsp</url-pattern>
    <url-pattern>*.JSP</url-pattern>
    <url-pattern>*.JSPF</url-pattern>
    <url-pattern>*.JSPX</url-pattern>
    <url-pattern>*.XSP</url-pattern>
  </servlet-mapping>

  <!-- ==================================================================== -->
  <!-- Dynamic Servlet Invoker.                                             -->
  <!-- This servlet invokes anonymous servlets that have not been defined   -->
  <!-- in the web.xml or by other means. The first element of the pathInfo  -->
  <!-- of a request passed to the envoker is treated as a servlet name for  -->
  <!-- an existing servlet, or as a class name of a new servlet.            -->
  <!-- This servlet is normally mapped to /servlet/*                        -->
  <!-- This servlet support the following initParams:                       -->
  <!--                                                                      -->
  <!--  nonContextServlets       If false, the invoker can only load        -->
  <!--                           servlets from the contexts classloader.    -->
  <!--                           This is false by default and setting this  -->
  <!--                           to true may have security implications.    -->
  <!--                                                                      -->
  <!--  verbose                  If true, log dynamic loads                 -->
  <!--                                                                      -->
  <!--  *                        All other parameters are copied to the     -->
  <!--                           each dynamic servlet as init parameters    -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <!-- Uncomment for dynamic invocation
  <servlet>
    <servlet-name>invoker</servlet-name>
    <servlet-class>org.mortbay.jetty.servlet.Invoker</servlet-class>
    <init-param>
      <param-name>verbose</param-name>
      <param-value>false</param-value>
    </init-param>
    <init-param>
      <param-name>nonContextServlets</param-name>
      <param-value>false</param-value>
    </init-param>
    <init-param>
      <param-name>dynamicParam</param-name>
      <param-value>anyValue</param-value>
    </init-param>
    <load-on-startup>0</load-on-startup>
  </servlet>

  <servlet-mapping> <servlet-name>invoker</servlet-name> <url-pattern>/servlet/*</url-pattern> </servlet-mapping>
  -->



  <!-- ==================================================================== -->
  <session-config>
    <session-timeout>30</session-timeout>
  </session-config>

  <!-- ==================================================================== -->
  <!-- Default MIME mappings                                                -->
  <!-- The default MIME mappings are provided by the mime.properties        -->
  <!-- resource in the org.mortbay.jetty.jar file.  Additional or modified  -->
  <!-- mappings may be specified here                                       -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <!-- UNCOMMENT TO ACTIVATE
  <mime-mapping>
    <extension>mysuffix</extension>
    <mime-type>mymime/type</mime-type>
  </mime-mapping>
  -->

  <!-- ==================================================================== -->
  <welcome-file-list>
    <welcome-file>index.html</welcome-file>
    <welcome-file>index.htm</welcome-file>
    <welcome-file>index.jsp</welcome-file>
  </welcome-file-list>

  <!-- ==================================================================== -->
  <locale-encoding-mapping-list>
    <locale-encoding-mapping><locale>ar</locale><encoding>ISO-8859-6</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>be</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>bg</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ca</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>cs</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>da</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>de</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>el</locale><encoding>ISO-8859-7</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>en</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>es</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>et</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>fi</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>fr</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>hr</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>hu</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>is</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>it</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>iw</locale><encoding>ISO-8859-8</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ja</locale><encoding>Shift_JIS</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ko</locale><encoding>EUC-KR</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>lt</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>lv</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>mk</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>nl</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>no</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>pl</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>pt</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ro</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ru</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sh</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sk</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sl</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sq</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sr</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sv</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>tr</locale><encoding>ISO-8859-9</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>uk</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>zh</locale><encoding>GB2312</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>zh_TW</locale><encoding>Big5</encoding></locale-encoding-mapping>
  </locale-encoding-mapping-list>

  <security-constraint>
    <web-resource-collection>
      <web-resource-name>Disable TRACE</web-resource-name>
      <url-pattern>/</url-pattern>
      <http-method>TRACE</http-method>
    </web-resource-collection>
    <auth-constraint/>
  </security-constraint>

</web-app>

`
	return
}

func GeneralWebappContext(groupId, artifactId string) (context string) {
	pack := groupId + "." + artifactId + ".component"
	context = PROPERTIES_LICENSE + `
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
	context = PROPERTIES_LICENSE + `
# 组件服务上下文属性文件列表
context=

# 服务根
context.root=/` + artifactId + `

context.scheduler-scan.base-package=` + pack + `
`
	return
}

func GeneralIndexJsp() (index string) {
	index = JSP_LICENSE + `
<html>
<body>
<h2>Hello World!</h2>
</body>
</html>
`
	return
}

func GeneralAssembly() (assembly string) {
	assembly = XML_LICENSE + `
<assembly xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSch ema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/assembly-1.0.0.xsd">

  <formats>
    <format>tar.gz</format>
  </formats>
  <includeBaseDirectory>true</includeBaseDirectory>
  <files>
  </files>
  <fileSets>
    <fileSet>
      <directory>bin</directory>
      <outputDirectory>bin</outputDirectory>
      <fileMode>0755</fileMode>
      <lineEnding>keep</lineEnding>
    </fileSet>
    <fileSet>
      <directory>webRoot</directory>
      <outputDirectory>bin/webRoot</outputDirectory>
    </fileSet>
    <fileSet>
      <directory>lib</directory>
      <outputDirectory>lib</outputDirectory>
    </fileSet>
    <fileSet>
        <directory>target/conf</directory>
        <outputDirectory>conf</outputDirectory>
    </fileSet>
  </fileSets>
  <dependencySets>
    <dependencySet>
      <outputDirectory>lib</outputDirectory>
      <scope>runtime</scope>
    </dependencySet>
  </dependencySets>

</assembly>
`
	return
}

func GeneralBootstrap(groupId, artifactId string) (bootstrap string) {
	pack := groupId + "." + artifactId

	bootstrap = CLASS_LICENSE + `
package ` + pack + `;

import org.nanoframework.server.JettyCustomServer;

/**
 *
 * @author yanghe
 * @since 0.0.1
 */
public final class Bootstrap {
    private Bootstrap() { }

    /**
     *
     * @param args bootstrap parameters.
     */
    public static void main(final String[] args) {
        JettyCustomServer.DEFAULT.bootstrap(args);
    }
}
`
	return
}

func GeneralBootstrapShell(groupId, artifactId string) (shell string) {
	bootstrap := groupId + "." + artifactId + ".Bootstrap"

	shell = `#!/bin/sh

#check JAVA_HOME & java
#noJavaHome=false
#if [ -z "$JAVA_HOME" ] ; then
#    noJavaHome=true
#fi
#if [ ! -e "$JAVA_HOME/bin/java" ] ; then
#    noJavaHome=true
#fi
#if $noJavaHome ; then
#    echo
#    echo "Error: JAVA_HOME environment variable is not set."
#    echo
#    exit 1
#fi
#==============================================================================
CURR_DIR=` + "`pwd`" + `

# set JAVA_HOME
#cd ` + "`dirname \"$0\"`/.." + `
# cd ..
# JAVA_HOME=` + "`pwd`" + `/jdk
# cd ./bin

#set JAVA_OPTS
JAVA_OPTS="-server -Xms512m -Xmx512m -Xmn64m -Xss256k"
#performance Options
JAVA_OPTS="$JAVA_OPTS -XX:+AggressiveOpts"
JAVA_OPTS="$JAVA_OPTS -XX:+UseBiasedLocking"
JAVA_OPTS="$JAVA_OPTS -XX:+UseFastAccessorMethods"
JAVA_OPTS="$JAVA_OPTS -XX:+DisableExplicitGC"
JAVA_OPTS="$JAVA_OPTS -XX:+UseParNewGC"
JAVA_OPTS="$JAVA_OPTS -XX:+UseConcMarkSweepGC"
JAVA_OPTS="$JAVA_OPTS -XX:+CMSParallelRemarkEnabled"
JAVA_OPTS="$JAVA_OPTS -XX:+UseCMSCompactAtFullCollection"
JAVA_OPTS="$JAVA_OPTS -XX:+UseCMSInitiatingOccupancyOnly"
JAVA_OPTS="$JAVA_OPTS -XX:CMSInitiatingOccupancyFraction=75"
#==============================================================================

#set HOME
# CURR_DIR=` + "`pwd`" + `
cd ` + "`dirname \"$0\"`/.." + `
APP_HOME=` + "`pwd`" + `
cd $CURR_DIR
if [ -z "$APP_HOME" ] ; then
echo
echo "Error: APP_HOME environment variable is not defined correctly."
echo
exit 1
fi
#==============================================================================

#set CLASSPATH
#APP_CLASSPATH="$APP_HOME/app:$APP_HOME/app/lib"

for i in "$APP_HOME"/lib/*.jar
do
APP_CLASSPATH="$APP_CLASSPATH:$i"
done

APP_CLASSPATH="$APP_CLASSPATH:$APP_HOME/conf"

LOGS_HOME="$APP_HOME/bin/logs"
echo $LOGS_HOME
if [ ! -d "$LOGS_HOME" ]; then
mkdir "$LOGS_HOME"
fi

#==============================================================================

#startup Server
RUN_CMD="\"$JAVA_HOME/bin/java\""
RUN_CMD="$RUN_CMD -Dlogic.home=\"$APP_HOME\""
RUN_CMD="$RUN_CMD -classpath \"$APP_CLASSPATH\""
RUN_CMD="$RUN_CMD $JAVA_OPTS"
# replace Bootstrap class name
RUN_CMD="$RUN_CMD ` + bootstrap + ` $@"
# if not run to docker
# RUN_CMD="$RUN_CMD > /dev/null 2>&1 &"
echo $RUN_CMD
eval $RUN_CMD
#==============================================================================

`

	return
}

func GeneralFindBugs() (findbugs string) {
	findbugs = `<?xml version="1.0" encoding="UTF-8"?>
` + XML_LICENSE + `
<!-- See http://findbugs.sourceforge.net/manual/filter.html -->
<FindBugsFilter>
    <Match>
        <Confidence value="2" />
        <Rank value="15" />
        <Bug category="SECURITY,PERFORMANCE,MALICIOUS_CODE" />
    </Match>
</FindBugsFilter>
`

	return
}

func GeneralCheckstyleRules() (rules string) {
	rules = `<?xml version="1.0" encoding="UTF-8"?>
` + XML_LICENSE + `
<!DOCTYPE module PUBLIC "-//Puppy Crawl//DTD Check Configuration 1.3//EN" "http://www.puppycrawl.com/dtds/configuration_1_3.dtd">
<module name="Checker">
    <property name="severity" value="warning"/>
    <property name="charset" value="UTF-8"/>

    <!-- TreeWalker module checks -->
    <module name="TreeWalker">
        <property name="tabWidth" value="4"/>
        <module name="JavadocMethod">
            <property name="scope" value="protected"/>
            <property name="severity" value="error"/>
            <property name="allowUndeclaredRTE" value="true"/>
            <property name="allowMissingPropertyJavadoc" value="true"/>
            <property name="allowThrowsTagsForSubclasses" value="true"/>
        </module>
        <module name="JavadocType">
            <property name="severity" value="error"/>
            <property name="scope" value="public"/>
            <property name="authorFormat" value="\w+"/>
            <property name="allowMissingParamTags" value="true"/>
        </module>
        <module name="JavadocVariable">
            <property name="severity" value="error"/>
            <property name="scope" value="protected"/>
            <property name="ignoreNamePattern" value="log|logger|LOG|LOGGER"/>
        </module>
        <module name="JavadocStyle">
            <property name="severity" value="error"/>
        </module>
        <module name="ClassFanOutComplexity">
            <property name="max" value="50"/>
            <property name="severity" value="error"/>
        </module>
        <!-- <module name="CommentsIndentation">
            <property name="severity" value="error"/>
        </module> -->
        <!-- <module name="Indentation">
            <property name="severity" value="error"/>
        </module> -->
        <module name="CyclomaticComplexity">
            <property name="max" value="15"/>
            <property name="severity" value="error"/>
        </module>
        <module name="DefaultComesLast">
            <property name="severity" value="error"/>
        </module>

        <module name="AnnotationLocation">
            <property name="severity" value="error"/>
            <property name="allowSamelineSingleParameterlessAnnotation" value="false"/>
        </module>
        <module name="ConstantName">
            <property name="severity" value="error"/>
        </module>
        <module name="GenericWhitespaceCheck">
            <property name="severity" value="error"/>
        </module>
        <module name="ModifiedControlVariable">
            <property name="severity" value="error"/>
        </module>
        <module name="MagicNumber">
            <property name="ignoreAnnotation" value="true"/>
            <property name="ignoreHashCodeMethod" value="true"/>
            <property name="ignoreFieldDeclaration" value="true"/>
            <property name="severity" value="error"/>
        </module>
        <module name="LocalFinalVariableName">
            <property name="severity" value="error"/>
        </module>
        <module name="LocalVariableName">
            <property name="severity" value="error"/>
        </module>
        <module name="AbstractClassName">
            <property name="ignoreModifier" value="true"/>
            <property name="severity" value="error"/>
            <property name="format" value="^Abstract.*$|^.*Factory$|^Base.*$|^Root.*$"/>
        </module>
        <module name="MemberName">
            <property name="severity" value="error"/>
        </module>
        <module name="MethodName">
            <property name="severity" value="error"/>
        </module>
        <module name="GenericWhitespace">
            <property name="severity" value="error"/>
        </module>
        <module name="PackageName">
            <property name="severity" value="error"/>
        </module>
        <module name="ParameterName">
            <property name="severity" value="error"/>
            <!-- <property name="ignoreOverridden" value="true" /> -->
        </module>
        <module name="StaticVariableName">
            <property name="severity" value="error"/>
            <property name="format" value="(^[A-Z][A-Z0-9]*(_[A-Z0-9]+)*$)"/>
        </module>
        <module name="AbbreviationAsWordInName"/>
        <module name="TypeName">
            <property name="severity" value="error"/>
        </module>
        <module name="AvoidStarImport">
            <property name="severity" value="error"/>
            <property name="excludes"
                      value="java.io,java.net,java.lang.Math,org.junit.Assert,org.mockito.Mockito,
                      org.mockito.Matchers,org.springframework.test.web.servlet.request.MockMvcRequestBuilders,
                      org.springframework.test.web.servlet.result.MockMvcResultMatchers,
                      java.nio.file.StandardWatchEventKinds"/>
            <property name="allowStaticMemberImports" value="false"/>
        </module>
        <module name="SingleLineJavadoc">
            <property name="severity" value="warning"/>
        </module>
        <module name="IllegalImport">
            <property name="severity" value="error"/>
        </module>
        <module name="RedundantImport">
            <property name="severity" value="error"/>
        </module>
        <module name="UnusedImports">
            <property name="severity" value="error"/>
        </module>
        <module name="SuperClone">
            <property name="severity" value="info"/>
        </module>
        <module name="SuperFinalize">
            <property name="severity" value="error"/>
        </module>
        <module name="MethodLength"/>
        <module name="ParameterNumber">
            <property name="max" value="10"/>
            <property name="severity" value="error"/>
            <property name="tokens" value="METHOD_DEF"/>
        </module>
        <module name="LineLength">
            <property name="max" value="200"/>
            <property name="tabWidth" value="4"/>
            <property name="severity" value="error"/>
        </module>
        <module name="EmptyForIteratorPad"/>
        <module name="MethodParamPad">
            <property name="severity" value="error"/>
        </module>
        <module name="NoWhitespaceAfter">
            <property name="severity" value="error"/>
        </module>
        <module name="NoWhitespaceBefore">
            <property name="severity" value="error"/>
        </module>
        <module name="OperatorWrap">
            <property name="severity" value="error"/>
        </module>
        <module name="ParenPad"/>
        <module name="TypecastParenPad"/>
        <module name="WhitespaceAfter">
            <property name="severity" value="error"/>
        </module>
        <module name="ModifierOrder">
            <property name="severity" value="error"/>
        </module>
        <module name="RedundantModifier">
            <property name="severity" value="error"/>
        </module>
        <module name="AvoidNestedBlocks"/>
        <module name="EmptyBlock">
            <property name="severity" value="error"/>
        </module>
        <module name="FallThrough">
            <property name="severity" value="error"/>
        </module>
        <module name="DeclarationOrder">
            <property name="severity" value="error"/>
        </module>
        <module name="CovariantEquals">
            <property name="severity" value="error"/>
        </module>
        <module name="ExplicitInitialization">
            <property name="severity" value="error"/>
        </module>
        <module name="LeftCurly">
            <property name="severity" value="error"/>
        </module>
        <module name="NeedBraces">
            <property name="severity" value="error"/>
        </module>
        <module name="EqualsAvoidNull">
            <property name="severity" value="error"/>
        </module>
        <module name="RightCurly">
            <property name="severity" value="error"/>
        </module>
        <module name="NoFinalizer">
            <property name="severity" value="error"/>
        </module>
        <module name="EmptyStatement"/>
        <module name="EqualsHashCode"/>
        <module name="IllegalInstantiation"/>
        <module name="InnerAssignment"/>
        <module name="MissingSwitchDefault">
            <property name="severity" value="info"/>
        </module>
        <module name="SimplifyBooleanExpression">
            <property name="severity" value="error"/>
        </module>
        <module name="SimplifyBooleanReturn">
            <property name="severity" value="error"/>
        </module>
        <module name="FinalClass">
            <property name="severity" value="error"/>
        </module>
        <module name="HideUtilityClassConstructor">
            <property name="severity" value="error"/>
        </module>
        <module name="InterfaceIsType">
            <property name="severity" value="info"/>
        </module>
        <module name="VisibilityModifier">
            <property name="severity" value="error"/>
            <property name="protectedAllowed" value="true"/>
        </module>
        <module name="AtclauseOrder">
            <property name="severity" value="error"/>
        </module>
        <module name="BooleanExpressionComplexity">
            <property name="severity" value="error"/>
            <property name="max" value="4"/>
        </module>
        <module name="ArrayTypeStyle">
            <property name="severity" value="error"/>
        </module>
        <module name="FinalParameters">
            <property name="severity" value="error"/>
            <property name="tokens" value="METHOD_DEF, CTOR_DEF, LITERAL_CATCH"/>
        </module>
        <module name="FinalLocalVariable">
            <property name="severity" value="ignore"/>
            <property name="tokens" value="PARAMETER_DEF, VARIABLE_DEF"/>
            <property name="validateEnhancedForLoopVariable" value="true"/>
        </module>
        <module name="TodoComment">
            <property name="severity" value="error"/>
        </module>
        <module name="MutableException">
            <property name="severity" value="error"/>
        </module>
        <module name="UpperEll"/>
        <module name="MissingOverride">
            <property name="severity" value="error"/>
        </module>
        <module name="MissingDeprecated">
            <property name="severity" value="error"/>
        </module>

        <!-- Custom checks based on regular expressions -->
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Parameteric log messages"/>
            <property name="severity" value="error"/>
            <property name="format" value="log\.\w+\(((\&quot;.+\&quot;\s*\+)|(.*\s*\+\s*\&quot;))"/>
            <property name="message"
                      value="Avoid string concatenation for constructing log messages. Use parametric messages instead"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="JavaDoc @version tag"/>
            <property name="severity" value="error"/>
            <property name="format" value="@version\s+(.+)*(\$Revision|\$Date)"/>
            <property name="message" value="Invalid JavaDoc @version tag."/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Logging framework selection"/>
            <property name="severity" value="error"/>
            <property name="format" value="LogFactory\.getLog"/>
            <property name="message" value="CAS uses the slf4j logging framework."/>
        </module>
        <module name="RegexpSinglelineJava">
            <property name="id" value="sysOutConsoleLogs"/>
            <metadata name="net.sf.eclipsecs.core.comment" value="Console output messages"/>
            <property name="severity" value="error"/>
            <property name="format" value="System\.(out|err)"/>
            <property name="message"
                      value="Avoid sending messages to the console directly. Use a logger object instead"/>
        </module>
        <module name="RegexpSinglelineJava">
            <property name="id" value="stackTraceConsoleLogs"/>
            <metadata name="net.sf.eclipsecs.core.comment" value="Printing stack traces to the console"/>
            <property name="severity" value="error"/>
            <property name="format" value="\.printStackTrace\(\)"/>
            <property name="message"
                      value="Avoid sending stack traces to the console directly. Use a logger object instead"/>
        </module>
        <module name="RegexpSinglelineJava">
            <property name="id" value="junitTestMethodName"/>
            <metadata name="net.sf.eclipsecs.core.comment" value="Using 'test' prefix for JUnit Tests"/>
            <property name="severity" value="error"/>
            <property name="format" value="(public|protected)\s+void\s+test\w+\(.+\{$"/>
            <property name="message"
                      value="JUnit test methods should not begin with the 'test' prefix. Use annotations instead and/or rename the method"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="The logger object must be named &quot;logger&quot;"/>
            <property name="severity" value="error"/>
            <property name="format" value="\s+(static\s)*(final\s)*(static\s)*Logger\s+(log|LOG)\b"/>
            <property name="message"
                      value="The Logger object must only be called &quot;logger&quot; or &quot;LOGGER&quot;"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Qualifying logger with the &quot;this&quot; keyword"/>
            <property name="severity" value="error"/>
            <property name="format" value="((this\.logger)|(super\.logger))\.\w+\("/>
            <property name="message" value="The Logger object need not be qualified with the &quot;this&quot; keyword"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Using @Test instead of extending TestCase"/>
            <property name="severity" value="error"/>
            <property name="format" value="class\s+\w+\s+extends\s+(junit\.framework\.)*TestCase"/>
            <property name="message"
                      value="All testcase must use annotations (@Test) instead of extending junit.framework.TestCase"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Using the junit.framework package"/>
            <property name="severity" value="error"/>
            <property name="format" value="junit.framework"/>
            <property name="message" value="The package junit.framework belongs to JUnit v3. Use org.junit instead."/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Checking for logging level"/>
            <property name="severity" value="warning"/>
            <property name="format" value="log\.is\w+Enabled\("/>
            <property name="message"
                      value="If the construction of the log message is cheap, consider not checking for logging levels."/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Non-static inner class"/>
            <property name="severity" value="error"/>
            <property name="format" value="\s+(private|public|protected)*\s+(abstract\s)*class\s+\w+"/>
            <property name="message"
                      value="Non-static nested classes are a security compromise. Consider using a static class instead"/>
        </module>
    </module>

    <!-- Checker module checks -->
    <module name="UniqueProperties">
        <property name="severity" value="error"/>
    </module>
    <module name="NewlineAtEndOfFile">
        <property name="fileExtensions" value="java, xml, properties, txt"/>
        <property name="lineSeparator" value="lf"/>
        <property name="severity" value="error"/>
    </module>
    <module name="Translation">
        <property name="severity" value="ignore"/>
    </module>
    <module name="FileLength"/>
    <module name="FileTabCharacter">
        <property name="severity" value="error"/>
    </module>

    <!-- Custom checks based on regular expressions -->
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Trailing spaces"/>
        <property name="severity" value="error"/>
        <property name="format" value="\w+\s+$"/>
        <property name="message" value="Line has trailing spaces."/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Space instead of tabs"/>
        <property name="severity" value="error"/>
        <property name="format" value="^\t+"/>
        <property name="message" value="Tabs should never be used for indentation. Use spaces instead"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Space after cast"/>
        <property name="severity" value="error"/>
        <property name="format" value="\(\w+\)\w+"/>
        <property name="message" value="There are no spaces after cast."/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Usage of java.util.Random"/>
        <property name="severity" value="error"/>
        <property name="format" value="(java.util.Random)|(new Random\()"/>
        <property name="message" value="For security purposes, use 'java.security.SecureRandom' instead"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="final clone() method"/>
        <property name="severity" value="error"/>
        <property name="format" value="public\s+\w+\s+clone\(\)"/>
        <property name="message"
                  value="Consider marking the clone() method as final to reduce chances of data corruption"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="JUnit Assert static import"/>
        <property name="severity" value="error"/>
        <property name="format" value="import\s+static\s+org\.junit\.Assert\.\w+"/>
        <property name="message" value="JUnit Assert methods MUST be imported statically with a *"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Mockito static import"/>
        <property name="severity" value="error"/>
        <property name="format" value="import\s+static\s+org\.mockito\.Mockito\.\w+"/>
        <property name="message" value="Mockito methods MUST be imported statically with a *"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Missing @since tag"/>
        <property name="severity" value="error"/>
        <property name="format" value="\s+\*+\s+@since\s+\d+\.\d+.*"/>
        <property name="maximum" value="200"/>
        <property name="minimum" value="1"/>
        <property name="fileExtensions" value="java,groovy"/>
        <property name="message" value="There are no @since tags defined for this component's Javadocs."/>
    </module>
    <module name="SuppressionFilter">
        <property name="file" value="${checkstyle.suppressions.file}"/>
    </module>
</module>
`

	return
}

func GeneralCheckstyleSuppressions() (suppressions string) {
	suppressions = `<?xml version="1.0" encoding="UTF-8"?>
` + XML_LICENSE + `
<!DOCTYPE suppressions PUBLIC "-//Puppy Crawl//DTD Suppressions 1.1//EN" "http://www.puppycrawl.com/dtds/suppressions_1_1.dtd">

<suppressions>
    <suppress checks="JavadocStyleCheck" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="JavadocMethod" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="JavadocType" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="JavadocVariable" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="MagicNumber" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="DesignForExtension" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress id="stackTraceConsoleLogs" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress id="sysOutConsoleLogs" files="\.xml" />
    <suppress checks="\w+" files="(\.(crt|crl|class|keystore))|rebel.xml" />
</suppressions>
`

	return
}

func GeneralGitIgnore() (ignore string) {
	ignore = `.DS_Store
.project
.classpath
.settings/
target/
tags
logs/
assembly/
conf/
*.log
webRoot/
*.zip
*.tar.gz
jetty.pid
.checkstyle
.fbIncludeFilterFile
`

	return
}

func GeneralCodeTemplates() (template string) {
	template = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
` + XML_LICENSE + `
<templates>
<template autoinsert="true" context="filecomment_context" deleted="false" description="Comment for created Java files" enabled="true" id="org.eclipse.jdt.ui.text.codetemplates.filecomment" name="filecomment">/*
 * Copyright 2015-${year} the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */</template>
 <template autoinsert="true" context="typecomment_context" deleted="false" description="Comment for create Java type files" enabled="true" id="org.eclipse.jdt.ui.text.codetemplates.typecomment" name="typecomment">/**
 *
 * @author ${user}
 * @since
 * ${tags}
 */</template>
 </templates>
 `

	return
}

func GeneralEclipseCheckstyle() (checkstyle string) {
	checkstyle = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
` + XML_LICENSE + `
<profiles version="12">
<profile kind="CodeFormatterProfile" name="Eclipse-checkstyle" version="12">
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_ellipsis" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_enum_declarations" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_in_empty_annotation_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_allocation_expression" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_at_in_annotation_type_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.new_lines_at_block_boundaries" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_constructor_declaration_parameters" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.insert_new_line_for_parameter" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_after_annotation_on_package" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_between_empty_parens_in_enum_constant" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_after_imports" value="1"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_while" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.insert_new_line_before_root_tags" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_between_empty_parens_in_annotation_type_member_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_method_declaration_throws" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.format_javadoc_comments" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.indentation.size" value="4"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_postfix_operator" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_for_increments" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_type_arguments" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_for_inits" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_in_empty_anonymous_type_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_semicolon_in_for" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.disabling_tag" value="@formatter:off"/>
<setting id="org.eclipse.jdt.core.formatter.continuation_indentation" value="2"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_enum_constants" value="0"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_before_imports" value="1"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_after_package" value="1"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_binary_operator" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_multiple_local_declarations" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_arguments_in_enum_constant" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_angle_bracket_in_parameterized_type_reference" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.indent_root_tags" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.wrap_before_or_operator_multicatch" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.enabling_tag" value="@formatter:on"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_closing_brace_in_block" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_parenthesized_expression_in_return" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_throws_clause_in_method_declaration" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_after_annotation_on_parameter" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.keep_then_statement_on_same_line" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_after_annotation_on_field" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_explicitconstructorcall_arguments" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_in_empty_block" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_prefix_operator" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_between_type_declarations" value="1"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_brace_in_array_initializer" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_for" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_catch" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_angle_bracket_in_type_arguments" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_after_annotation_on_method" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_switch" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_anonymous_type_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_parenthesized_expression" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.never_indent_line_comments_on_first_column" value="false"/>
<setting id="org.eclipse.jdt.core.compiler.problem.enumIdentifier" value="error"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_and_in_type_parameter" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_for_inits" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.indent_statements_compare_to_block" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_anonymous_type_declaration" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_question_in_wildcard" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_annotation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_method_invocation_arguments" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_switch" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.line_length" value="80"/>
<setting id="org.eclipse.jdt.core.formatter.use_on_off_tags" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_between_empty_brackets_in_array_allocation_expression" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_enum_constant" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_between_empty_parens_in_method_invocation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_assignment_operator" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_type_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_for" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.preserve_white_space_between_code_and_line_comments" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_after_annotation_on_local_variable" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_method_declaration" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_method_invocation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_union_type_in_multicatch" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_colon_in_for" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.number_of_blank_lines_at_beginning_of_method_body" value="0"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_closing_angle_bracket_in_type_arguments" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.keep_else_statement_on_same_line" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_binary_expression" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_parameterized_type_reference" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_array_initializer" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_multiple_field_declarations" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_annotation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_arguments_in_explicit_constructor_call" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.indent_body_declarations_compare_to_annotation_declaration_header" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_superinterfaces" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_colon_in_default" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_question_in_conditional" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_block" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_constructor_declaration" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_lambda_body" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.compact_else_if" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_type_parameters" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_catch" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_method_invocation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.put_empty_statement_on_new_line" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_parameters_in_constructor_declaration" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_method_invocation_arguments" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_arguments_in_method_invocation" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_throws_clause_in_constructor_declaration" value="16"/>
<setting id="org.eclipse.jdt.core.compiler.problem.assertIdentifier" value="error"/>
<setting id="org.eclipse.jdt.core.formatter.comment.clear_blank_lines_in_block_comment" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_before_catch_in_try_statement" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_try" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_at_end_of_file_if_missing" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.clear_blank_lines_in_javadoc_comment" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_array_initializer" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_binary_operator" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_unary_operator" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_expressions_in_array_initializer" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.format_line_comment_starting_on_first_column" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.number_of_empty_lines_to_preserve" value="1"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_colon_in_case" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_ellipsis" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_semicolon_in_try_resources" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_colon_in_assert" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_if" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_type_arguments" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_and_in_type_parameter" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_in_empty_type_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_parenthesized_expression" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.format_line_comments" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_colon_in_labeled_statement" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.align_type_members_on_columns" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_assignment" value="0"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_in_empty_method_body" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.indent_body_declarations_compare_to_type_header" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_between_empty_parens_in_method_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_enum_constant" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_superinterfaces_in_type_declaration" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_before_first_class_body_declaration" value="0"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_conditional_expression" value="80"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_before_closing_brace_in_array_initializer" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_constructor_declaration_parameters" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.format_guardian_clause_on_one_line" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_if" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_after_annotation_on_type" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_block" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_enum_declaration" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_block_in_case" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_constructor_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.format_header" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_arguments_in_allocation_expression" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_method_invocation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_while" value="insert"/>
<setting id="org.eclipse.jdt.core.compiler.codegen.inlineJsrBytecode" value="enabled"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_switch" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_method_declaration" value="0"/>
<setting id="org.eclipse.jdt.core.formatter.join_wrapped_lines" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_between_empty_parens_in_constructor_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.indent_switchstatements_compare_to_cases" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_bracket_in_array_allocation_expression" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_synchronized" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.new_lines_at_javadoc_boundaries" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_annotation_type_declaration" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_colon_in_for" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_resources_in_try" value="80"/>
<setting id="org.eclipse.jdt.core.formatter.use_tabs_only_for_leading_indentations" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_selector_in_method_invocation" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.never_indent_block_comments_on_first_column" value="false"/>
<setting id="org.eclipse.jdt.core.compiler.source" value="1.8"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_synchronized" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_constructor_declaration_throws" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.tabulation.size" value="4"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_in_empty_enum_constant" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_allocation_expression" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_bracket_in_array_reference" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_colon_in_conditional" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.format_source_code" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_array_initializer" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_try" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_semicolon_in_try_resources" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_before_field" value="0"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_at_in_annotation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.continuation_indentation_for_array_initializer" value="2"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_question_in_wildcard" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_before_method" value="1"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_superclass_in_type_declaration" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_superinterfaces_in_enum_declaration" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_parenthesized_expression_in_throw" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_colon_in_labeled_statement" value="do not insert"/>
<setting id="org.eclipse.jdt.core.compiler.codegen.targetPlatform" value="1.8"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_switch" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_superinterfaces" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_method_declaration_parameters" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_after_type_annotation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_brace_in_array_initializer" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_parenthesized_expression" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.format_html" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_at_in_annotation_type_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_closing_angle_bracket_in_type_parameters" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_compact_if" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.indent_empty_lines" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_parameterized_type_reference" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_unary_operator" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_enum_constant" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_arguments_in_annotation" value="0"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_enum_declarations" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.keep_empty_array_initializer_on_one_line" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.indent_switchstatements_compare_to_switch" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_before_else_in_if_statement" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_assignment_operator" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_constructor_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_before_new_chunk" value="1"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_after_label" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.indent_body_declarations_compare_to_enum_declaration_header" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_bracket_in_array_allocation_expression" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_constructor_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_colon_in_conditional" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_angle_bracket_in_parameterized_type_reference" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_method_declaration_parameters" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_angle_bracket_in_type_arguments" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_cast" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_colon_in_assert" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_before_member_type" value="1"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_before_while_in_do_statement" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_bracket_in_array_type_reference" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_angle_bracket_in_parameterized_type_reference" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_arguments_in_qualified_allocation_expression" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_after_opening_brace_in_array_initializer" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_in_empty_enum_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.indent_breaks_compare_to_cases" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_method_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_if" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_semicolon" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_postfix_operator" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_try" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_angle_bracket_in_type_arguments" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_cast" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.format_block_comments" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_lambda_arrow" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_method_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.keep_imple_if_on_one_line" value="false"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_enum_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_parameters_in_method_declaration" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_between_brackets_in_array_type_reference" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_angle_bracket_in_type_parameters" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_semicolon_in_for" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_method_declaration_throws" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_bracket_in_array_allocation_expression" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.indent_statements_compare_to_body" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.alignment_for_multiple_fields" value="16"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_enum_constant_arguments" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_prefix_operator" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_array_initializer" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.wrap_before_binary_operator" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_method_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_type_parameters" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_catch" value="do not insert"/>
<setting id="org.eclipse.jdt.core.compiler.compliance" value="1.8"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_bracket_in_array_reference" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_comma_in_annotation" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_enum_constant_arguments" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_between_empty_braces_in_array_initializer" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_colon_in_case" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_multiple_local_declarations" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_annotation_type_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_bracket_in_array_reference" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_method_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.wrap_outer_expressions_when_nested" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_closing_paren_in_cast" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_enum_constant" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.brace_position_for_type_declaration" value="end_of_line"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_before_package" value="0"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_for" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_synchronized" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_for_increments" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_annotation_type_member_declaration" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_while" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_enum_constant" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_explicitconstructorcall_arguments" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_paren_in_annotation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_angle_bracket_in_type_parameters" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.indent_body_declarations_compare_to_enum_constant_header" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_lambda_arrow" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_brace_in_constructor_declaration" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_constructor_declaration_throws" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.join_lines_in_comments" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_closing_angle_bracket_in_type_parameters" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_question_in_conditional" value="insert"/>
<setting id="org.eclipse.jdt.core.formatter.comment.indent_parameter_description" value="true"/>
<setting id="org.eclipse.jdt.core.formatter.insert_new_line_before_finally_in_try_statement" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.tabulation.char" value="space"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_comma_in_multiple_field_declarations" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.blank_lines_between_import_groups" value="1"/>
<setting id="org.eclipse.jdt.core.formatter.lineSplit" value="120"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_after_opening_paren_in_annotation" value="do not insert"/>
<setting id="org.eclipse.jdt.core.formatter.insert_space_before_opening_paren_in_switch" value="insert"/>
</profile>
</profiles>
`

	return
}

func GeneralLicenseHeader() (header string) {
	header = `Copyright 2015-2016 the original author or authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.`

	return
}

func GeneralLicenseHeaderDefinitions() (definitions string) {
	definitions = `<?xml version="1.0" encoding="UTF-8"?>
<additionalHeaders>
  <JAVADOC_STYLE>
    <firstLine>/*</firstLine>
    <beforeEachLine> * </beforeEachLine>
    <endLine> */</endLine>
    <firstLineDetectionPattern>( |\t)*/\*( |\t)*$</firstLineDetectionPattern>
    <lastLineDetectionPattern>( |\t)*\*/( |\t)*$</lastLineDetectionPattern>
    <allowBlankLines>true</allowBlankLines>
    <isMultiline>true</isMultiline>
  </JAVADOC_STYLE>
</additionalHeaders>
`
	return
}

func GeneralMavenSettings() (settings string) {
	settings = `<?xml version="1.0" encoding="UTF-8"?>
<!--
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
 -->

<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0 http://maven.apache.org/xsd/settings-1.0.0.xsd">

    <localRepository>${user.home}/.m2/repository</localRepository>
    <pluginGroups></pluginGroups>
    <proxies></proxies>
    <servers></servers>

    <mirrors>
        <mirror>
            <id>local-nexus</id>
            <mirrorOf>local-nexus</mirrorOf>
            <name>local-nexus</name>
            <url>http://10.1.195.225:8081/nexus/content/groups/public</url>
        </mirror>
    </mirrors>

    <profiles>
        <profile>
            <id>nexus</id>
            <repositories>
                <repository>
                    <id>local-nexus</id>
                    <url>http://local-nexus</url>
                    <releases>
                        <enabled>true</enabled>
                    </releases>
                    <snapshots>
                        <enabled>true</enabled>
                    </snapshots>
                </repository>
            </repositories>
        </profile>
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

    <activeProfiles>
        <activeProfile>nexus</activeProfile>
    </activeProfiles>
</settings>
`

	return
}