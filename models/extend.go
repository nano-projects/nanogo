package models

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

