[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](https://github.com/nano-projects/nanogo/releases)

NanoGo
======

    NanoGo是基于GO语言开发的一款小工具, 主要为NanoFramework框架项目提供快速的项目构建功能.
    用户可以通过NanoGo快速构建基于NanoFramework的Web项目和Scheduler项目.
    
    NanoGo还提供了自定义项目构建的方法, 通过编写yaml文件对项目进行定义与编排.

下载安装
----

    前往 https://github.com/nano-projects/nanogo/releases 下载二进制包
    根据自身操作系统选择对应的版本进行下载, 提供Mac, Linux, Windows等主流操作系统的32位和64位环境.
    
配置环境变量
------

### Mac / Linux
##### 方式一

    直接复制 nanogo 至 /usr/local/bin 目录下
    
##### 方式二

    配置环境变量:
    ~$ vim ~/.bash_profile
        export NANOGO_HOME=<Your NanoGo DIR Path>
        export PATH = $NANOGO_HOME:$PATH
        
    ~$ source ~/.bash_profile
        
### Windows

    进入 控制面板 -> 系统 -> 高级系统设置
    点击 环境变量 按钮
    新建 系统变量
        NANOGO_HOME: <Your NanoGo DIR Path>
        Path: 在最尾端追加 ;%NANOGO_HOME%
        
使用说明
----

##### 查看帮助信息
    
    ~$ nanogo -h
    NanoGo is a CLI library for Go that empowers applications.
    This application is a tool to generate the needed files
    to quickly create a Maven application.

    Usage:
        nanogo [command]

    Available Commands:
        add         Add Source file of flags
        init        Init a maven project
        version     Display the version

    Flags:
        -l, --log-level string   Log level (options "debug", "info", "warn", "error", "fatal", "panic") (default "info")

    Use "nanogo [command] --help" for more information about a command.

##### 查看项目初始化帮助信息
     ~$ nanogo init -h
    Init a maven project

    Usage:
      nanogo init [flags]

    Flags:
      -n, --name string       Maven project name definition, format: "groupId:artifactId:version", version is optional, the default use of 0.0.1
          --parent string     Maven top POM dependency, format: "groupId:artifactId:version" (default "org.nanoframework:super:x.x.x")
          --path string       The project path by default using the current path (default "...")
      -p, --publish uint      Project default port (default 7000)
      -s, --scheduler         Init a scheduler project of nano framework
          --server string     Set up the server (option 'Tomcat' or 'Jetty')
      -t, --template string   The project template file path
      -w, --web               Init a webapp project of nano framework
     
    Global Flags:
      -l, --log-level string   Log level (options 'debug', 'info', 'warn', 'error', 'fatal', 'panic') (default "info")
      
##### 查看代码构建帮助信息
    ~$ nanogo add -h
    Add Source file of flags

    Usage:
      nanogo add [flags]

    Flags:
      -a, --author string   Creation file author
      -n, --name string     The interface or class prefix name
          --path string     The project root path (default "...")

    Global Flags:
      -l, --log-level string   Log level (options 'debug', 'info', 'warn', 'error', 'fatal', 'panic') (default "info")

##### 构建基于NanoFramework的Web项目
    
    ~$ nanogo init -w -n <GroupId>:<ArtifactId>:<Version (Optional)>

##### 构建基于NanoFramework的Scheduler项目

    ~$ nanogo init -s --path <Your creation project path> -n <GroupId>:<ArtifactId>:<Version (Optional)>

##### 构建基于NanoFramework的自定义模板项目

    ~$ nanogo init -n <GroupId>:<ArtifactId>:<Version (Optional)> -t <Your yml template file path>

###### e.g. 
    
    ~$ nanogo init -w --path ____Workspaces____ -n org.nanoframework.nanogo:test
    create file:  ____Workspaces____/test/test-common/pom.xml
    ...
    create file:  ____Workspaces____/test/test-mapper/pom.xml
    ...
    create file:  ____Workspaces____/test/test-core/pom.xml
    ...
    create file:  ____Workspaces____/test/test-webapp-support/pom.xml
    ...
    create file:  ____Workspaces____/test/test-webapp/pom.xml
    ...
    create file:  ____Workspaces____/test/pom.xml
    ...
    
##### 项目结构

    <ArtifactId>
    ├── checkstyle-rules.xml
    ├── checkstyle-suppressions.xml
    ├── findbugs-rules.xml
    ├── .gitignore
    ├── pom.xml
    ├── src
    │   ├── eclipse
    │   │   ├── eclipse-code-style.xml
    │   │   └── eclipse-code-template.xml
    │   ├── mvn
    │   │   └── settings.xml
    │   ├── licensing
    │   │   ├── header-definitions.xml
    │   │   └── header.txt
    │   └── yml
    │       └── nanogo.yml
    │
    ├── <ArtifactId>-common
    │   ├── src/main/java/<GroupID package>/<ArtifactId>/.gitkeep
    │   ├── src/main/resources/.gitkeep
    │   ├── src/test/java/<GroupID package>/<ArtifactId>/.gitkeep
    │   ├── src/test/resources/.gitkeep
    │   └── pom.xml
    │
    ├── <ArtifactId>-mapper
    │   ├── src/main/java/<GroupID package>/<ArtifactId>/.gitkeep
    │   ├── src/main/resources/.gitkeep
    │   ├── src/test/java/<GroupID package>/<ArtifactId>/.gitkeep
    │   ├── src/test/resources/.gitkeep
    │   └── pom.xml
    │
    ├── <ArtifactId>-core
    │   ├── src/main/java/<GroupID package>/<ArtifactId>/.gitkeep
    │   ├── src/main/resources/.gitkeep
    │   ├── src/test/java/<GroupID package>/<ArtifactId>/.gitkeep
    │   ├── src/test/resources/.gitkeep
    │   └── pom.xml
    │
    ├── <ArtifactId>-webapp-support [ <ArtifactId>-scheduler-support ]
    │   ├── src/main/java/<GroupID package>/<ArtifactId>/.gitkeep
    │   ├── src/main/resources/.gitkeep
    │   ├── src/test/java/<GroupID package>/<ArtifactId>/.gitkeep
    │   ├── src/test/resources/.gitkeep
    │   └── pom.xml
    │
    └── <ArtifactId>-webapp [ <ArtifactId>-scheduler ]
        ├── src/main/java/<GroupID package>/<ArtifactId>
        │   ├── .gitkeep
        │   └── XXXBootstrap.java
        ├── src/main/resources
        │   ├── .gitkeep
        │   ├── assembly.xml
        │   └── context.properties
        ├── src/main/webapp
        │   ├── index.jsp
        │   └── WEB-INF
        │       ├── jetty.xml
        │       ├── web.xml
        │       └── webdefault.xml
        ├── src/test/java/<GroupID package>/<ArtifactId>/.gitkeep
        ├── src/test/resources/.gitkeep
        ├── pom.xml
        ├── bin
        │   └── bootstrap.sh
        └── configure
            ├── public/.gitkeep
            ├── sit/.gitkeep
            ├── uat/.gitkeep
            └── release/.gitkeep

项目导入与启动
----

#### Maven仓库配置
    
    IDE导入maven settings.xml的配置, 配置文件位于项目路径: <ArtifactId>/src/mvn/settings.xml
    
#### 项目导入

    使用支持Maven的IDE导入项目
    
#### 项目启动

    运行 <ArtifactId>-webapp/src/main/java/<GroupId package>/<ArtifactId>/XXXBootstrap.java
    访问应用: http://localhost:7000/<ArtifactId>
