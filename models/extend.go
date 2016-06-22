package models

import (
	"flag"
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
	New *bool
	NewWebapp *bool
	NewScheduler *bool

	Path *string
	Yaml *string

	GroupId *string
	ArtifactId *string
	Version *string

	/* Default Webapp or Scheduler General Only */
	Javadoc *bool
	Source *bool
	Checkstyle *bool
	Findbugs *bool
	License *bool
	Pmd *bool

}

func (this *Argument) Parse() {
	this.New = flag.Bool("new", false, "新建项目")
	this.NewWebapp = flag.Bool("web", false, "新建基于NanoFramework的Web项目")
	this.NewScheduler = flag.Bool("scheduler", false, "新建基于NanoFramework的任务调度项目")

	this.Path = flag.String("path", pwd(), "创建项目路径,默认使用当前路径")
	this.Yaml = flag.String("yaml", "", "Yaml配置文件路径")

	this.GroupId = flag.String("groupId", "", "Maven项目的groupId属性")
	this.ArtifactId = flag.String("artifactId", "", "Maven项目的artifactId属性")
	this.Version = flag.String("version", "0.0.1", "Maven项目的version属性")

	this.Javadoc = flag.Bool("no-doc", false, "移除插件: maven-javadoc-plugin")
	this.Source = flag.Bool("no-src", false, "移除插件: maven-source-plugin")
	this.Checkstyle = flag.Bool("no-chk", false, "移除插件: maven-checkstyle-plugin")
	this.Findbugs = flag.Bool("no-fb", false, "移除插件: findbugs-maven-plugin")
	this.License = flag.Bool("no-license", false, "移除插件: license-maven-plugin")
	this.Pmd = flag.Bool("no-pmd", false, "移除插件: maven-pmd-plugin")

	flag.Parse()

	if !strings.HasSuffix(*this.Path, "/") {
		path := *this.Path
		path += "/"
		this.Path = &path
	}
}

func (this *Argument) Validation() bool {
	if *this.GroupId == "" || *this.ArtifactId == "" {
		return false
	}

	return true
}

func (this *Argument) ExistYaml(yaml *string) bool {
	var yamlPath string
	if *yaml == "" {
		yamlPath = pwd() + "/nanogo.yml"
	} else {
		yamlPath = *yaml
	}

	if file, err := os.Open(yamlPath); err != nil && os.IsNotExist(err) {
		return false
	} else {
		defer file.Close()
		return true
	}
}

func pwd() (path string) {
	if p, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		path = p
	}

	return
}