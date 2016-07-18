package resources

import "github.com/nano-projects/nanogo/resources/license"

func Settings() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
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
            <id>local-nexus</id>
            <mirrorOf>local-nexus</mirrorOf>
            <name>local-nexus</name>
            <url>http://10.1.195.225:8081/nexus/content/groups/public</url>
        </mirror>
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
            <id>nexus</id>
            <activation>
                <activeByDefault>true</activeByDefault>
            </activation>
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
            <id>oss-snapshots</id>
            <activation>
                <activeByDefault>true</activeByDefault>
            </activation>
            <repositories>
                <repository>
                    <id>oss-snapshots</id>
                    <url>http://oss-snapshots</url>
                    <releases>
                        <enabled>false</enabled>
                    </releases>
                    <snapshots>
                        <enabled>true</enabled>
                    </snapshots>
                </repository>
            </repositories>
        </profile>
        <profile>
            <id>oss-releases</id>
            <activation>
                <activeByDefault>true</activeByDefault>
            </activation>
            <repositories>
                <repository>
                    <id>oss-releases</id>
                    <url>http://oss-releases</url>
                    <releases>
                        <enabled>true</enabled>
                    </releases>
                    <snapshots>
                        <enabled>false</enabled>
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
</settings>
`

}