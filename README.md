[![gorelease](https://dn-gorelease.qbox.me/gorelease-download-blue.svg)](https://gobuild.io/nano-projects/nanogo/master)

NanoGo
======

    NanoGo是基于GO语言开发的一款小工具, 主要为NanoFramework框架项目提供快速的项目构建功能.
    用户可以通过NanoGo快速构建基于NanoFramework的Web项目和Scheduler项目.
    
    NanoGo提供了自定义项目构建的方法, 通过编写yaml文件对项目进行定义与编排.(计划中)

下载安装
----

    前往 https://gobuild.io/nano-projects/nanogo/master 下载二进制包
    根据自身操作系统选择对应的版本进行下载, 提供Mac, Linux, Windows等主流操作系统的环境.
    
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
    Usage of nanogo:
      -new
        	新建项目
      -no-chk
        	移除插件: maven-checkstyle-plugin
      -no-doc
        	移除插件: maven-javadoc-plugin
      -no-fb
        	移除插件: findbugs-maven-plugin
      -no-license
        	移除插件: license-maven-plugin
      -no-pmd
        	移除插件: maven-pmd-plugin
      -no-src
        	移除插件: maven-source-plugin
      -parent string
        	Maven顶级POM依赖, 格式: groupId:artifactId:version (default "org.nanoframework:super:6")
      -path string
        	创建项目路径,默认使用当前路径 (default "...")
      -port string
        	项目默认端口 (default "7000")
      -resp string
        	Maven项目资源定义, 格式: groupId:artifactId:version, version为可选项, 默认使用0.0.1
      -scheduler
        	新建基于NanoFramework的任务调度项目
      -web
        	新建基于NanoFramework的Web项目
      -yaml string
        	Yaml配置文件路径
        	
##### 构建基于NanoFramework的Web项目
    
    ~$ nanogo -new -web -path <Your creation project path> -resp <GroupId>:<ArtifactId>:<Version (Optional)>

##### 构建基于NanoFramework的Scheduler项目

    ~$ nanogo -new -scheduler -path <Your creation project path> -resp <GroupId>:<ArtifactId>:<Version (Optional)>

###### e.g. 
    
    ~$ nanogo -new -web -path /Users/yanghe/Works/____Go_Project____/____Workspaces____ -resp org.nanoframework.nanogo:test
    create file:  /Users/yanghe/Works/____Go_Project____/____Workspaces____/test/test-common/pom.xml
    ...
    create file:  /Users/yanghe/Works/____Go_Project____/____Workspaces____/test/test-mapper/pom.xml
    ...
    create file:  /Users/yanghe/Works/____Go_Project____/____Workspaces____/test/test-core/pom.xml
    ...
    create file:  /Users/yanghe/Works/____Go_Project____/____Workspaces____/test/test-webapp-support/pom.xml
    ...
    create file:  /Users/yanghe/Works/____Go_Project____/____Workspaces____/test/test-webapp/pom.xml
    ...
    create file:  /Users/yanghe/Works/____Go_Project____/____Workspaces____/test/pom.xml
    ...
    
##### 项目结构

    <ArtifactId>
    ├── checkstyle-rules.xml
    ├── checkstyle-suppressions.xml
    ├── checkstyle-suppressions.xml
    ├── findbugs-rules.xml
    ├── .gitignore
    ├── pom.xml
    ├── src
    │   ├── eclipse
    │   │   ├── eclipse-code-template.xml
    │   │   └── eclipse-formatter.xml
    │   ├── mvn
    │   │   └── settigs.xml
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
        │   └── Bootstrap.java
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

    运行 <ArtifactId>-web/src/main/java/<GroupId package>/<ArtifactId>/Bootstrap.java
    访问应用: http://localhost:7000/<ArtifactId>
