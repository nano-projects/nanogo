package models

import (
	"bufio"
	"flag"
	"io"
	"os"
	"strings"
)

type Licenses struct {
	License *[]License `yaml:"license,flow" xml:"license,omitempty"`
}

type Developers struct {
	Developer *[]Developer `yaml:"developer,flow" xml:"developer,omitempty"`
}

type Contributors struct {
	Contributor *[]Contributor `yaml:"contributor,flow" xml:"contributor,omitempty"`
}

type MailingLists struct {
	MailingList []MailingList `yaml:"mailingList,flow" xml:"mailingList,omitempty"`
}

type Modules struct {
	Module *[]string `yaml:"module,flow" xml:"module,omitempty"`
}

type Dependencies struct {
	Dependency *[]Dependency `yaml:"dependency,flow" xml:"dependency,omitempty"`
}

type Repositories struct {
	Repository *[]Repository `yaml:"repository,flow" xml:"repository,omitempty"`
}

type PluginRepositories struct {
	PluginRepository *[]Repository `yaml:"pluginRepository,flow" xml:"pluginRepository,omitempty"`
}

type Reports struct {
	Project *[]project `yaml:"project,flow" xml:"project,omitempty"`
	Report  *[]string  `yaml:"report,flow" xml:"report,omitempty"`
}

type Profiles struct {
	Profile *[]Profile `yaml:"profile,flow" xml:"profile,omitempty"`
}

type Roles struct {
	Role *[]string `yaml:"role,flow" xml:"role,omitempty"`
}

type Notifiers struct {
	Notifier *[]Notifier `yaml:"notifier,flow" xml:"notifier,omitempty"`
}

type Resources struct {
	Resource *[]Resource `yaml:"resource,flow" xml:"resource,omitempty"`
}

type TestResources struct {
	TestResource *[]Resource `yaml:"testResource,flow" xml:"testResource,omitempty"`
}

type Filters struct {
	Filter *[]string `yaml:"filter,flow" xml:"filter,omitempty"`
}

type Plugins struct {
	Plugin *[]Plugin `yaml:"plugin,flow" xml:"plugin,omitempty"`
}

type Executions struct {
	Execution *[]PluginExecution `yaml:"execution,flow" xml:"execution,omitempty"`
}

type Goals struct {
	Project *[]project `yaml:"project,flow" xml:"project,omitempty"`
	Goal    *[]string  `yaml:"goal,flow" xml:"goal,omitempty"`
}

type Exclusions struct {
	Exclusion *[]Exclusion `yaml:"exclusion,flow" xml:"exclusion,omitempty"`
}

type Includes struct {
	Include *[]string `yaml:"include,flow" xml:"include,omitempty"`
}

type Excludes struct {
	Exclude *[]string `yaml:"exclude,flow" xml:"exclude,omitempty"`
}

type ReportPlugins struct {
	Plugin *[]ReportPlugin `yaml:"plugin,flow" xml:"plugin,omitempty"`
}

type ReportSets struct {
	ReportSet *[]ReportSet `yaml:"reportSet,flow" xml:"reportSet,omitempty"`
}

type Extensions struct {
	Extension *[]Extension `yaml:"extension,flow" xml:"extension,omitempty"`
}

type OtherArchives struct {
	OtherArchive *[]string `yaml:"otherArchive,flow" xml:"otherArchive,omitempty"`
}

type Argument struct {
	New          *bool
	NewWebapp    *bool
	NewScheduler *bool

	Path *string
	Yaml *string

	Parent *Parent

	GroupId    *string
	ArtifactId *string
	Version    *string

	/* Default Webapp or Scheduler General Only */
	Javadoc    *bool
	Source     *bool
	Checkstyle *bool
	Findbugs   *bool
	License    *bool
	Pmd        *bool

	Port *string
}

func (this *Argument) Parse() {
	this.New = flag.Bool("new", false, "新建项目")
	this.NewWebapp = flag.Bool("web", false, "新建基于NanoFramework的Web项目")
	this.NewScheduler = flag.Bool("scheduler", false, "新建基于NanoFramework的任务调度项目")

	this.Path = flag.String("path", pwd(), "创建项目路径,默认使用当前路径")
	this.Yaml = flag.String("yaml", "", "Yaml配置文件路径")

	parent := flag.String("parent", "org.nanoframework:super:6", "Maven顶级POM依赖, 格式: groupId:artifactId:version")
	resp := flag.String("resp", "", "Maven项目资源定义, 格式: groupId:artifactId:version, version为可选项, 默认使用0.0.1")

	this.Javadoc = flag.Bool("no-doc", false, "移除插件: maven-javadoc-plugin")
	this.Source = flag.Bool("no-src", false, "移除插件: maven-source-plugin")
	this.Checkstyle = flag.Bool("no-chk", false, "移除插件: maven-checkstyle-plugin")
	this.Findbugs = flag.Bool("no-fb", false, "移除插件: findbugs-maven-plugin")
	this.License = flag.Bool("no-license", false, "移除插件: license-maven-plugin")
	this.Pmd = flag.Bool("no-pmd", false, "移除插件: maven-pmd-plugin")

	this.Port = flag.String("port", "7000", "项目默认端口")

	flag.Parse()

	if !strings.HasSuffix(*this.Path, "/") {
		path := *this.Path
		path += "/"
		this.Path = &path
	}

	if *parent != "" {
		parents := strings.Split(*parent, ":")
		parentsLen := len(parents)
		if parentsLen < 3 {
			panic("无效的Maven顶级POM资源定义(-parent), 格式: groupId:artifactId:version")
		}

		p := Parent{}
		p.GroupId = parents[0]
		p.ArtifactId = parents[1]
		p.Version = parents[2]
		this.Parent = &p
	}

	defaultVersion := "0.0.1-SNAPSHOT"
	if *resp != "" {
		resps := strings.Split(*resp, ":")
		respsLen := len(resps)
		if respsLen < 2 {
			panic("无效的Maven项目资源定义(-resp), 格式: groupId:artifactId:version")
		} else {
			this.GroupId = &resps[0]
			this.ArtifactId = &resps[1]
			if respsLen > 2 {
				version := resps[2]
				if !strings.HasSuffix(version, "-SNAPSHOT") {
					version += "-SNAPSHOT"
				}

				this.Version = &version
			} else {
				this.Version = &defaultVersion
			}
		}
	} else {
		empty := ""
		this.GroupId = &empty
		this.ArtifactId = &empty
		this.Version = &defaultVersion
	}
}

func (this *Argument) Validation() bool {
	if *this.GroupId == "" || *this.ArtifactId == "" {
		return false
	}

	return true
}

func (this *Argument) ExistYaml() bool {
	var yamlPath string
	if *this.Yaml == "" {
		yamlPath = pwd() + "/nanogo.yml"
	} else {
		yamlPath = *this.Yaml
	}

	if file, err := os.Open(yamlPath); err != nil && os.IsNotExist(err) {
		return false
	} else {
		defer file.Close()
		return true
	}
}

func (this *Argument) LoadYaml() (pom *string) {
	var yamlPath string
	if *this.Yaml == "" {
		yamlPath = pwd() + "/nanogo.yml"
	} else {
		yamlPath = *this.Yaml
	}

	if file, err := os.Open(yamlPath); err != nil && os.IsNotExist(err) {
		panic(err)
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)

		data := ""
		for {
			if d, err := reader.ReadString('\n'); err == io.EOF {
				break
			} else {
				data += d
			}
		}

		pom = &data
	}

	return
}

func pwd() (path string) {
	if p, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		path = p
	}

	return
}
