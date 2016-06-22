package models

type Schema struct {
	Root string `yaml:"root" xml:"-"`
	// The <code>&lt;project&gt;</code> element is the root of the descriptor. The following table lists all of the possible child elements.
	// @version 3.0.0+
	Projects map[string] project `yaml:"projects,flow" xml:"projects"`
}

func (this *Schema) Project(name string) project {
	return this.Projects[name]
}

// The <code>&lt;project&gt;</code> element is the root of the descriptor. The following table lists all of the possible child elements.
// @version 3.0.0+
type project struct {
	Xmlns string `yaml:"-" xml:"xmlns,attr,omitempty"`

	XmlnsXsi string `yaml:"-" xml:"xmlns:xsi,attr,omitempty"`

	XsiSchemaLocation string `yaml:"-" xml:"xsi:schemaLocation,attr,omitempty"`

	ModuleType string `yaml:"moduleType" xml:"-"`

	// Declares to which version of project descriptor this POM conforms.
	// @version 4.0.0+
	ModelVersion string `yaml:"modelVersion" xml:"modelVersion,omitempty"`

	// The location of the parent project, if one exists.
	// Values from the parent project will be the default for this project if they are left unspecified.
	// The location is given as a group ID, artifact ID and version.
	// @version 4.0.0+
	Parent *Parent `yaml:"parent" xml:"parent,omitempty"`

	// A universally unique identifier for a project.
	// It is normal to use a fully-qualified package name to distinguish it from other projects with a similar name (eg. <code>org.apache.maven</code>).
	// @version 3.0.0+
	GroupId string `yaml:"groupId" xml:"groupId,omitempty"`

	// The identifier for this artifact that is unique within the group given by the group ID.
	// An artifact is something that is either produced or used by a project.
	// Examples of artifacts produced by Maven for a project include: JARs, source and binary distributions, and WARs.
	// @version 3.0.0+
	ArtifactId string `yaml:"artifactId" xml:"artifactId,omitempty"`

	// The current version of the artifact produced by this project.
	// @version 4.0.0+
	Version string `yaml:"version" xml:"version,omitempty"`

	// The type of artifact this project produces,
	// for example
	// 		<code>jar</code>
	// 		<code>war</code>
	// 		<code>ear</code>
	// 		<code>pom</code>.
	// Plugins can create their own packaging, and therefore their own packaging types, so this list does not contain all possible types.
	// @version 4.0.0+
	Packaging string `yaml:"packaging" xml:"packaging,omitempty"`

	// The full name of the project.
	// @version 3.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// A detailed description of the project, used by Maven whenever it needs to describe the project, such as on the web site.
	// While this element can be specified as CDATA to enable the use of HTML tags within the description, it is discouraged to allow plain text representation.
	// If you need to modify the index page of the generated web site, you are able to specify your own instead of adjusting this text.
	// @version 3.0.0+
	Description string `yaml:"description" xml:"description,omitempty"`

	// The URL to the project's homepage. <br /><b>Default value is</b>: parent value [+ path adjustment] + artifactId
	// @version 3.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`

	// The year of the project's inception, specified with 4 digits. This value is used when generating copyright notices as well as being informational.
	// @version 3.0.0+
	InceptionYear string `yaml:"inceptionYear" xml:"inceptionYear,omitempty"`

	// This element describes various attributes of the organization to which the project belongs.
	// These attributes are utilized when documentation is created (for copyright notices and links).
	// @version 3.0.0+
	Organization *Organization `yaml:"organization" xml:"organization,omitempty"`

	// This element describes all of the licenses for this project.
	// Each license is described by a <code>license</code> element, which is then described by additional elements.
	// Projects should only list the license(s) that applies to the project and not the licenses that apply to dependencies.
	// If multiple licenses are listed, it is assumed that the user can select any of them, not that they must accept all.
	// @version 3.0.0+
	Licenses *Licenses `yaml:"licenses" xml:"licenses,omitempty"`

	// Describes the committers of a project.
	// @version 3.0.0+
	Developers *Developers `yaml:"developers" xml:"developers,omitempty"`

	// Describes the contributors to a project that are not yet committers.
	// @version 3.0.0+
	Contributors *Contributors `yaml:"contributors,flow" xml:"contributors,omitempty"`

	// Contains information about a project's mailing lists.
	// @version 3.0.0+
	MailingLists *MailingLists `yaml:"mailingLists,flow" xml:"mailingLists,omitempty"`

	// Describes the prerequisites in the build environment for this project.
	// @version 4.0.0+
	Prerequisites *Prerequisites `yaml:"prerequisites" xml:"prerequisites,omitempty"`

	// The modules (sometimes called subprojects) to build as a part of this project. Each module listed is a relative path to the directory containing the module.
	// To be consistent with the way default urls are calculated from parent, it is recommended to have module names match artifact ids.
	// @version 4.0.0+
	Modules *Modules `yaml:"modules,flow" xml:"modules,omitempty"`

	// Specification for the SCM used by the project, such as CVS, Subversion, etc.
	// @version 4.0.0+
	Scm *Scm `yaml:"scm" xml:"scm,omitempty"`

	// The project's issue management system information.
	// @version 4.0.0+
	IssueManagement *IssueManagement `yaml:"issueManagement" xml:"issueManagement,omitempty"`

	// The project's continuous integration information.
	// @version 4.0.0+
	CiManagement *CiManagement `yaml:"ciManagement" xml:"ciManagement,omitempty"`

	// Distribution information for a project that enables deployment of the site and artifacts to remote web servers and repositories respectively.
	// @version 4.0.0+
	DistributionManagement *DistributionManagement `yaml:"distributionManagement" xml:"distributionManagement,omitempty"`

	// Properties that can be used throughout the POM as a substitution, and are used as filters in resources if enabled.
	// The format is <code>&lt;name&gt;value&lt;/name&gt;</code>.
	// @version 4.0.0+
	Properties             InterfaceMap `yaml:"properties,flow" xml:"properties,omitempty"`

	// Default dependency information for projects that inherit from this one.
	// The dependencies in this section are not immediately resolved.
	// Instead, when a POM derived from this one declares a dependency described by a matching groupId and artifactId,
	// the version and other values from this section are used for that dependency if they were not already specified.
	// @version 4.0.0+
	DependencyManagement   *DependencyManagement `yaml:"dependencyManagement" xml:"dependencyManagement,omitempty"`

	// This element describes all of the dependencies associated with a project.
	// These dependencies are used to construct a classpath for your project during the build process.
	// They are automatically downloaded from the repositories defined in this project.
	// See <a href="http://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html">the dependency mechanism</a> for more information.
	// @version 3.0.0+
	Dependencies *Dependencies `yaml:"dependencies" xml:"dependencies,omitempty"`

	// The lists of the remote repositories for discovering dependencies and extensions.
	// @version 4.0.0+
	Repositories *Repositories `yaml:"repositories" xml:"repositories,omitempty"`

	// The lists of the remote repositories for discovering plugins for builds and reports.
	// @version 4.0.0+
	PluginRepositories *PluginRepositories `yaml:"pluginRepositories" xml:"pluginRepositories,omitempty"`

	// Information required to build the project.
	// @version 3.0.0+
	Build *Build `yaml:"build" xml:"build,omitempty"`

	// <b>Deprecated</b>.
	// Now ignored by Maven.
	// @version 4.0.0+
	Reports *Reports `yaml:"reports" xml:"reports,omitempty"`

	// This element includes the specification of report plugins to use to generate the reports on the Maven-generated site.
	// These reports will be run when a user executes <code>mvn site</code>.
	// All of the reports will be included in the navigation bar for browsing.
	// @version 4.0.0+
	Reporting *Reporting `yaml:"reporting" xml:"reporting,omitempty"`

	// A listing of project-local build profiles which will modify the build process when activated.
	// @version 4.0.0+
	Profiles *Profiles `yaml:"profiles" xml:"profiles,omitempty"`
}

// The <code>&lt;parent&gt;</code> element contains information required to locate the parent project from which this project will inherit from.
// <strong>Note:</strong> The children of this element are not interpolated and must be given as literal values.
// @version 4.0.0+
type Parent struct {
	// The group id of the parent project to inherit from.
	// @version 4.0.0+
	GroupId string `yaml:"groupId" xml:"groupId,omitempty"`

	// The artifact id of the parent project to inherit from.
	// @version 4.0.0+
	ArtifactId string `yaml:"artifactId" xml:"artifactId,omitempty"`

	// The version of the parent project to inherit.
	// @version 4.0.0+
	Version string `yaml:"version" xml:"version,omitempty"`

	// The relative path of the parent <code>pom.xml</code> file within the check out. If not specified, it defaults to <code>../pom.xml</code>.
	// Maven looks for the parent POM first in this location on the filesystem, then the local repository, and lastly in the remote repo.
	// <code>relativePath</code> allows you to select a different location, for example when your structure is flat, or deeper without an intermediate parent POM.
	// However, the group ID, artifact ID and version are still required, and must match the file in the location given or it will revert to the repository for the POM.
	// This feature is only for enhancing the development in a local checkout of that project.
	// Set the value to an empty string in case you want to disable the feature and always resolve the parent POM from the repositories.
	// @version 4.0.0+
	RelativePath string `yaml:"relativePath" xml:"relativePath,omitempty"`
}

// Specifies the organization that produces this project.
// @version 3.0.0+
type Organization struct {
	// The full name of the organization.
	// @version 3.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// The URL to the organization's home page.
	// @version 3.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`
}

// This elements describes all that pertains to distribution for a project.
// It is primarily used for deployment of artifacts and the site produced by the build.
// @version 4.0.0+
type DistributionManagement struct {
	// Information needed to deploy the artifacts generated by the project to a remote repository.
	// @version 4.0.0+
	Repository *DeploymentRepository `yaml:"repository" xml:"repository,omitempty"`

	// Where to deploy snapshots of artifacts to. If not given, it defaults to the <code>repository</code> element.
	// @version 4.0.0+
	SnapshotRepository *DeploymentRepository `yaml:"snapshotRepository" xml:"snapshotRepository,omitempty"`

	// Information needed for deploying the web site of the project.
	// @version 4.0.0+
	Site *Site `yaml:"site" xml:"site,omitempty"`

	// The URL of the project's download page. If not given users will be referred to the homepage given by <code>url</code>.
	// This is given to assist in locating artifacts that are not in the repository due to licensing restrictions.
	// @version 4.0.0+
	DownloadUrl string `yaml:"downloadUrl" xml:"downloadUrl,omitempty"`

	// Relocation information of the artifact if it has been moved to a new group ID and/or artifact ID.
	// @version 4.0.0+
	Relocation *Relocation `yaml:"relocation" xml:"relocation,omitempty"`

	// Gives the status of this artifact in the remote repository.
	// This must not be set in your local project, as it is updated by tools placing it in the reposiory.
	// Valid values are:
	// 		<code>none</code> (default),
	// 		<code>converted</code> (repository manager converted this from an Maven 1 POM),
	// 		<code>partner</code> (directly synced from a partner Maven 2 repository),
	// 		<code>deployed</code> (was deployed from a Maven 2 instance),
	// 		<code>verified</code> (has been hand verified as correct and final).
	// @version 4.0.0+
	Status string `yaml:"status" xml:"status,omitempty"`
}

// Describes where an artifact has moved to. If any of the values are omitted, it is assumed to be the same as it was before.
// @version 4.0.0+
type Relocation struct {
	// The group ID the artifact has moved to.
	// @version 4.0.0+
	GroupId string `yaml:"groupId" xml:"groupId,omitempty"`

	// The new artifact ID of the artifact.
	// @version 4.0.0+
	ArtifactId string `yaml:"artifactId" xml:"artifactId,omitempty"`

	// The new version of the artifact.
	// @version 4.0.0+
	Version string `yaml:"version" xml:"version,omitempty"`

	// An additional message to show the user about the move, such as the reason.
	// @version 4.0.0+
	Message string `yaml:"message" xml:"message,omitempty"`
}

// Contains the information needed for deploying websites.
// @version 4.0.0+
type Site struct {
	// A unique identifier for a deployment location.
	// This is used to match the site to configuration in the <code>settings.xml</code> file, for example.
	// @version 4.0.0+
	Id string `yaml:"id" xml:"id,omitempty"`

	// Human readable name of the deployment location.
	// @version 4.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// The url of the location where website is deployed, in the form <code>protocol://hostname/path</code>.
	// <br />
	// <b>Default value is</b>:
	// 		parent value [+ path adjustment] + artifactId
	// @version 4.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`
}

// Repository contains the information needed for deploying to the remote repository.
// @version 4.0.0+
type DeploymentRepository struct {
	// Whether to assign snapshots a unique version comprised of the timestamp and build number, or to use the same version each time
	// @version 4.0.0+
	UniqueVersion string /*bool*/ `yaml:"uniqueVersion" xml:"uniqueVersion,omitempty"`

	// How to handle downloading of releases from this repository.
	// @version 4.0.0+
	Releases *RepositoryPolicy `yaml:"releases" xml:"releases,omitempty"`

	// How to handle downloading of snapshots from this repository.
	// @version 4.0.0+
	Snapshots *RepositoryPolicy `yaml:"snapshots" xml:"snapshots,omitempty"`

	// A unique identifier for a repository.
	// This is used to match the repository to configuration in the <code>settings.xml</code> file, for example.
	// Furthermore, the identifier is used during POM inheritance and profile injection to detect repositories that should be merged.
	// @version 4.0.0+
	Id string `yaml:"id" xml:"id,omitempty"`

	// Human readable name of the repository.
	// @version 4.0.0 +
	Name string `yaml:"name" xml:"name,omitempty"`

	// The url of the repository, in the form <code>protocol://hostname/path</code>.
	// @version 4.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`

	// The type of layout this repository uses for locating and storing artifacts - can be <code>legacy</code> or <code>default</code>.
	// @version 4.0.0+
	Layout string `yaml:"layout" xml:"layout,omitempty"`
}

// Download policy.
// @version 4.0.0+
type RepositoryPolicy struct {
	// Whether to use this repository for downloading this type of artifact.
	// Note: While the type of this field is <code>String</code> for technical reasons, the semantic type is actually <code>Boolean</code>.
	// Default value is <code>true</code>.
	// @version 4.0.0+
	Enabled string `yaml:"enabled" xml:"enabled,omitempty"`

	// The frequency for downloading updates -
	// can be
	// 		<code>always,</code>
	// 		<code>daily</code> (default),
	// 		<code>interval:XXX</code> (in minutes) or
	// 		<code>never</code> (only if it doesn't exist locally).
	// @version 4.0.0+
	UpdatePolicy string `yaml:"updatePolicy" xml:"updatePolicy,omitempty"`

	// What to do when verification of an artifact checksum fails.
	// Valid values are
	// 		<code>ignore</code>,
	// 		<code>fail</code> or
	// 		<code>warn</code> (the default).
	// @version 4.0.0+
	ChecksumPolicy string `yaml:"checksumPolicy" xml:"checksumPolicy,omitempty"`
}

// Describes the prerequisites a project can have.
// @version 4.0.0+
type Prerequisites struct {
	// For a plugin project, the minimum version of Maven required to use the resulting plugin.<br />
	// For specifying the minimum version of Maven required to build a project, this element is <b>deprecated</b>.
	// Use the Maven Enforcer Plugin's <a href="https://maven.apache.org/enforcer/enforcer-rules/requireMavenVersion.html"><code>requireMavenVersion</code></a> rule instead.
	// @version 4.0.0+
	Maven string `yaml:"maven" xml:"maven,omitempty"`
}

// Description of a person who has contributed to the project, but who does not have commit privileges. Usually, these contributions come in the form of patches submitted.
// @version 3.0.0+
type Contributor struct {
	// The full name of the contributor.
	// @version 3.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// The email address of the contributor.
	// @version 3.0.0+
	Email string `yaml:"email" xml:"email,omitempty"`

	// The URL for the homepage of the contributor.
	// @version 3.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`

	// The organization to which the contributor belongs.
	// @version 3.0.0+
	Organization string `yaml:"organization" xml:"organization,omitempty"`

	// The URL of the organization.
	// @version 3.0.0+
	OrganizationUrl string `yaml:"organizationUrl" xml:"organizationUrl,omitempty"`

	// The roles the contributor plays in the project.
	// Each role is described by a <code>role</code> element, the body of which is a role name.
	// This can also be used to describe the contribution.
	// @version 3.0.0+
	Roles      *Roles `yaml:"roles" xml:"roles,omitempty"`

	// The timezone the contributor is in.
	// Typically, this is a number in the range <a href="http://en.wikipedia.org/wiki/UTC%E2%88%9212:00">-12</a> to <a href="http://en.wikipedia.org/wiki/UTC%2B14:00">+14</a>
	// or a valid time zone id like "America/Montreal" (UTC-05:00) or "Europe/Paris" (UTC+01:00).
	// @version 3.0.0+
	Timezone   string `yaml:"timezone" xml:"timezone,omitempty"`

	// Properties about the contributor, such as an instant messenger handle.
	// @version 3.0.0+
	Properties InterfaceMap `yaml:"properties,flow" xml:"properties,omitempty"`
}

// The <code>&lt;scm&gt;</code> element contains informations required to the SCM (Source Control Management) of the project.
// @version 4.0.0+
type Scm struct {
	// The source control management system URL that describes the repository and how to connect to the repository.
	// For more information, see the <a href="http://maven.apache.org/scm/scm-url-format.html">URL format</a> and <a href="http://maven.apache.org/scm/scms-overview.html">list of supported SCMs</a>.
	// This connection is read-only. <br />
	// <b>Default value is</b>: parent value [+ path adjustment] + artifactId
	// @version 4.0.0+
	Connection string `yaml:"connection" xml:"connection,omitempty"`

	// Just like <code>connection</code>, but for developers, i.e. this scm connection will not be read only. <br />
	// <b>Default value is</b>: parent value [+ path adjustment] + artifactId
	// @version 4.0.0+
	DeveloperConnection string `yaml:"developerConnection" xml:"developerConnection,omitempty"`

	// The tag of current code. By default, it's set to HEAD during development.
	// @version 4.0.0+
	Tag string `yaml:"tag" xml:"tag,omitempty"`

	// The URL to the project's browsable SCM repository, such as ViewVC or Fisheye. <br />
	// <b>Default value is</b>: parent value [+ path adjustment] + artifactId
	// @version 4.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`
}

// A repository contains the information needed for establishing connections with remote repository.
// @version 4.0.0+
type Repository struct {
	// How to handle downloading of releases from this repository.
	// @version 4.0.0+
	Releases *RepositoryPolicy `yaml:"releases" xml:"releases,omitempty"`

	// How to handle downloading of snapshots from this repository.
	// @version 4.0.0+
	Snapshots *RepositoryPolicy `yaml:"snapshots" xml:"snapshots,omitempty"`

	// A unique identifier for a repository.
	// This is used to match the repository to configuration in the <code>settings.xml</code> file, for example.
	// Furthermore, the identifier is used during POM inheritance and profile injection to detect repositories that should be merged.
	// @version 4.0.0+
	Id string `yaml:"id" xml:"id,omitempty"`

	// Human readable name of the repository.
	// @version 4.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// The url of the repository, in the form <code>protocol://hostname/path</code>.
	// @version 4.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`

	// The type of layout this repository uses for locating and storing artifacts - can be <code>legacy</code> or <code>default</code>.
	// @version 4.0.0+
	Layout string `yaml:"layout" xml:"layout,omitempty"`
}

// Information about the issue tracking (or bug tracking) system used to manage this project.
// @version 4.0.0+
type IssueManagement struct {
	// The name of the issue management system, e.g. Bugzilla
	// @version 4.0.0+
	System string `yaml:"system" xml:"system,omitempty"`

	// URL for the issue management system used by the project.
	// @version 4.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`
}

// The <code>&lt;CiManagement&gt;</code> element contains informations required to the continuous integration system of the project.
// @version 4.0.0+
type CiManagement struct {
	// The name of the continuous integration system, e.g. <code>continuum</code>.
	// @version 4.0.0+
	System string `yaml:"system" xml:"system,omitempty"`

	// URL for the continuous integration system used by the project if it has a web interface.
	// @version 4.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`

	// Configuration for notifying developers/users when a build is unsuccessful, including user information and notification mode.
	// @version 4.0.0+
	Notifiers *Notifiers `yaml:"notifiers" xml:"notifiers,omitempty"`
}

// Configures one method for notifying users/developers when a build breaks.
// @version 4.0.0+
type Notifier struct {
	// The mechanism used to deliver notifications.
	// @version 4.0.0+
	Type string `yaml:"type" xml:"type,omitempty"`

	// Whether to send notifications on error.
	// @version 4.0.0+
	SendOnError string /*bool*/ `yaml:"sendOnError" xml:"sendOnError,omitempty"`

	// Whether to send notifications on failure.
	// @version 4.0.0+
	SendOnFailure string /*bool*/ `yaml:"sendOnFailure" xml:"sendOnFailure,omitempty"`

	// Whether to send notifications on success.
	// @version 4.0.0+
	SendOnSuccess string /*bool*/ `yaml:"sendOnSuccess" xml:"sendOnSuccess,omitempty"`

	// Whether to send notifications on warning.
	// @version 4.0.0+
	SendOnWarning string /*bool*/ `yaml:"sendOnWarning" xml:"sendOnWarning,omitempty"`

	// <b>Deprecated</b>. Where to send the notification to - eg email address.
	// @version 4.0.0+
	Address       string `yaml:"address" xml:"address,omitempty"`

	// Extended configuration specific to this notifier goes here.
	// @version 4.0.0+
	Configuration InterfaceMap /* *interface{} */ `yaml:"configuration,flow" xml:"configuration,omitempty"`
}

// Modifications to the build process which is activated based on environmental parameters or command line arguments.
// @version 4.0.0+
type Profile struct {
	// The identifier of this build profile.
	// This is used for command line activation, and identifies profiles to be merged.
	// @version 4.0.0+
	Id                     string `yaml:"id" xml:"id,omitempty"`

	// The conditional logic which will automatically trigger the inclusion of this profile.
	// @version 4.0.0+
	Activation *Activation `yaml:"activation" xml:"activation,omitempty"`

	// Information required to build the project.
	// @version 4.0.0+
	Build *BuildBase `yaml:"build" xml:"build,omitempty"`

	// The modules (sometimes called subprojects) to build as a part of this project.
	// Each module listed is a relative path to the directory containing the module.
	// To be consistent with the way default urls are calculated from parent, it is recommended to have module names match artifact ids.
	// @version 4.0.0+
	Modules *Modules `yaml:"modules" xml:"modules,omitempty"`

	// Distribution information for a project that enables deployment of the site and artifacts to remote web servers and repositories respectively.
	// @version 4.0.0+
	DistributionManagement *DistributionManagement `yaml:"distributionManagement" xml:"distributionManagement,omitempty"`

	// Properties that can be used throughout the POM as a substitution, and are used as filters in resources if enabled.
	// The format is <code>&lt;name&gt;value&lt;/name&gt;</code>.
	// @version 4.0.0+
	Properties             InterfaceMap `yaml:"properties,flow" xml:"properties,omitempty"`

	// Default dependency information for projects that inherit from this one.
	// The dependencies in this section are not immediately resolved.
	// Instead, when a POM derived from this one declares a dependency described by a matching groupId and artifactId, the version and other values from this section are used for that dependency if they were not already specified.
	// @version 4.0.0+
	DependencyManagement   *DependencyManagement `yaml:"dependencyManagement" xml:"dependencyManagement,omitempty"`

	// This element describes all of the dependencies associated with a project.
	// These dependencies are used to construct a classpath for your project during the build process.
	// They are automatically downloaded from the repositories defined in this project.
	// See <a href="http://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html">the dependency mechanism</a> for more information.
	// @version 3.0.0+
	Dependencies *Dependencies `yaml:"dependencies" xml:"dependencies,omitempty"`

	// The lists of the remote repositories for discovering dependencies and extensions.
	// @version 4.0.0+
	Repositories *Repositories `yaml:"repositories" xml:"repositories,omitempty"`

	// The lists of the remote repositories for discovering plugins for builds and reports.
	// @version 4.0.0+
	PluginRepositories *PluginRepositories `yaml:"pluginRepositories" xml:"pluginRepositories,omitempty"`

	// <b>Deprecated</b>. Now ignored by Maven.
	// @version 4.0.0+
	Reports *Reports `yaml:"reports" xml:"reports,omitempty"`

	// This element includes the specification of report plugins to use to generate the reports on the Maven-generated site.
	// These reports will be run when a user executes <code>mvn site</code>.
	// All of the reports will be included in the navigation bar for browsing.
	// @version 4.0.0+
	Reporting *Reporting `yaml:"reporting" xml:"reporting,omitempty"`
}

// Generic informations for a build.
// @version 3.0.0+
type BuildBase struct {
	// The default goal (or phase in Maven 2) to execute when none is specified for the project.
	// Note that in case of a multi-module build, only the default goal of the top-level project is relevant, i.e. the default goals of child modules are ignored.
	// Since Maven 3, multiple goals/phases can be separated by whitespace.
	// @version 3.0.0+
	DefaultGoal string `yaml:"defaultGoal" xml:"defaultGoal,omitempty"`

	// This element describes all of the classpath resources such as properties files associated with a project.
	// These resources are often included in the final package.
	// The default value is <code>src/main/resources</code>.
	// @version 3.0.0+
	Resources *Resources `yaml:"resources" xml:"resources,omitempty"`

	// This element describes all of the classpath resources such as properties files associated with a project's unit tests.
	// The default value is <code>src/test/resources</code>.
	// @version 4.0.0+
	TestResources *TestResources `yaml:"testResources" xml:"testResources,omitempty"`

	// The directory where all files generated by the build are placed. The default value is <code>target</code>.
	// @version 4.0.0+
	Directory string `yaml:"directory" xml:"directory,omitempty"`

	// The filename (excluding the extension, and with no path information) that the produced artifact will be called. The default value is <code>${artifactId}-${version}</code>.
	// @version 4.0.0+
	FinalName string `yaml:"finalName" xml:"finalName,omitempty"`

	// The list of filter properties files that are used when filtering is enabled.
	// @version 4.0.0+
	Filters *Filters `yaml:"filters" xml:"filters,omitempty"`

	// Default plugin information to be made available for reference by projects derived from this one.
	// This plugin configuration will not be resolved or bound to the lifecycle unless referenced.
	// Any local configuration for a given plugin will override the plugin's entire definition here.
	// @version 4.0.0+
	PluginManagement *PluginManagement `yaml:"pluginManagement" xml:"pluginManagement,omitempty"`

	// The list of plugins to use.
	// @version 4.0.0+
	Plugins *Plugins `yaml:"plugins" xml:"plugins,omitempty"`
}

// The <code>&lt;plugin&gt;</code> element contains informations required for a plugin.
// @version 4.0.0+
type Plugin struct {
	// The group ID of the plugin in the repository.
	// @version 4.0.0+
	GroupId string `yaml:"groupId" xml:"groupId,omitempty"`

	// The artifact ID of the plugin in the repository.
	// @version 4.0.0+
	ArtifactId string `yaml:"artifactId" xml:"artifactId,omitempty"`

	// The version (or valid range of versions) of the plugin to be used.
	// @version 4.0.0+
	Version string `yaml:"version" xml:"version,omitempty"`

	// Whether to load Maven extensions (such as packaging and type handlers) from this plugin.
	// For performance reasons, this should only be enabled when necessary.
	// Note: While the type of this field is <code>String</code> for technical reasons, the semantic type is actually <code>Boolean</code>.
	// Default value is <code>false</code>.
	// @version 4.0.0+
	Extensions string `yaml:"extensions" xml:"extensions,omitempty"`

	// Multiple specifications of a set of goals to execute during the build lifecycle, each having (possibly) a different configuration.
	// @version 4.0.0+
	Executions *Executions `yaml:"executions,flow" xml:"executions,omitempty"`

	// Additional dependencies that this project needs to introduce to the plugin's classloader.
	// @version 4.0.0+
	Dependencies *Dependencies `yaml:"dependencies" xml:"dependencies,omitempty"`

	// <b>Deprecated</b>. Unused by Maven.
	// @version 4.0.0+
	Goals *Goals `yaml:"goals" xml:"goals,omitempty"`

	// Whether any configuration should be propagated to child POMs.
	// Note: While the type of this field is <code>String</code> for technical reasons, the semantic type is actually <code>Boolean</code>. Default value is <code>true</code>.
	// @version 4.0.0+
	Inherited string `yaml:"inherited" xml:"inherited,omitempty"`

	// <p>The configuration as DOM object.</p> <p>By default, every element content is trimmed, but starting with Maven 3.1.0, you can add <code>xml:space="preserve"</code> to elements you want to preserve whitespace.
	// </p> <p>You can control how child POMs inherit configuration from parent POMs by adding <code>combine.children</code> or <code>combine.self</code> attributes to the children
	// of the configuration element:
	// </p> <ul> <li><code>combine.children</code>:
	// 		available values are <code>merge</code> (default) and <code>append</code>,
	// </li> <li><code>combine.self</code>:
	// 		available values are <code>merge</code> (default) and <code>override</code>.
	// </li> </ul> <p>See <a href="http://maven.apache.org/pom.html#Plugins">POM Reference documentation</a> and <a href="http://plexus.codehaus.org/plexus-utils/apidocs/org/codehaus/plexus/util/xml/Xpp3DomUtils.html">Xpp3DomUtils</a> for more information.</p>
	// @version 4.0.0+
	Configuration InterfaceMap /* *interface{} */ `yaml:"configuration,flow" xml:"configuration,omitempty"`
}

// The <code>&lt;dependency&gt;</code> element contains information about a dependency of the project.
// @version 3.0.0+
type Dependency struct {
	// The project group that produced the dependency, e.g. <code>org.apache.maven</code>.
	// @version 3.0.0+
	GroupId string `yaml:"groupId" xml:"groupId,omitempty"`

	// The unique id for an artifact produced by the project group, e.g. <code>maven-artifact</code>.
	// @version 3.0.0+
	ArtifactId string `yaml:"artifactId" xml:"artifactId,omitempty"`

	// The version of the dependency, e.g. <code>3.2.1</code>. In Maven 2, this can also be specified as a range of versions.
	// @version 3.0.0+
	Version string `yaml:"version" xml:"version,omitempty"`

	// The type of dependency. While it usually represents the extension on the filename of the dependency, that is not always the case.
	// A type can be mapped to a different extension and a classifier.
	// The type often corresponds to the packaging used, though this is also not always the case.
	// Some examples are
	// 		<code>jar</code>,
	// 		<code>war</code>,
	// 		<code>ejb-client</code> and
	// 		<code>test-jar</code>:
	// see <a href="../maven-core/artifact-handlers.html">default artifact handlers</a> for a list.
	// New types can be defined by plugins that set <code>extensions</code> to <code>true</code>, so this is not a complete list.
	// @version 4.0.0+
	Type string `yaml:"type" xml:"type,omitempty"`

	// The classifier of the dependency.
	// It is appended to the filename after the version.
	// This allows: <ul> <li>refering to attached artifact,
	// for example
	// 		<code>sources</code> and
	// 		<code>javadoc</code>:
	// see <a href="../maven-core/artifact-handlers.html">default artifact handlers</a> for a list,</li> <li>distinguishing two artifacts that belong to the same POM but were built differently.
	// For example,
	// 		<code>jdk14</code> and
	// 		<code>jdk15</code>.</li> </ul>
	// @version 4.0.0+
	Classifier string `yaml:"classifier" xml:"classifier,omitempty"`

	// The scope of the dependency -
	// 		<code>compile</code>,
	// 		<code>runtime</code>,
	// 		<code>test</code>,
	// 		<code>system</code>, and
	// 		<code>provided</code>.
	// Used to calculate the various classpaths used for compilation, testing, and so on.
	// It also assists in determining which artifacts to include in a distribution of this project.
	// For more information, see <a href="http://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html">the dependency mechanism</a>.
	// @version 4.0.0+
	Scope string `yaml:"scope" xml:"scope,omitempty"`

	// FOR SYSTEM SCOPE ONLY.
	// Note that use of this property is <b>discouraged</b> and may be replaced in later versions.
	// This specifies the path on the filesystem for this dependency. Requires an absolute path for the value, not relative.
	// Use a property that gives the machine specific absolute path, e.g. <code>${java.home}</code>.
	// @version 4.0.0+
	SystemPath string `yaml:"systemPath" xml:"systemPath,omitempty"`

	// Lists a set of artifacts that should be excluded from this dependency's artifact list when it comes to calculating transitive dependencies.
	// @version 4.0.0+
	Exclusions *Exclusions `yaml:"exclusions" xml:"exclusions,omitempty"`

	// Indicates the dependency is optional for use of this library.
	// While the version of the dependency will be taken into account for dependency calculation if the library is used elsewhere, it will not be passed on transitively.
	// Note: While the type of this field is <code>String</code> for technical reasons, the semantic type is actually <code>Boolean</code>.
	// Default value is <code>false</code>.
	// @version 4.0.0+
	Optional string `yaml:"optional" xml:"optional,omitempty"`
}

// The <code>&lt;exclusion&gt;</code> element contains informations required to exclude an artifact to the project.
// @version 4.0.0+
type Exclusion struct {
	// The artifact ID of the project to exclude.
	// @version 4.0.0+
	ArtifactId string `yaml:"artifactId" xml:"artifactId,omitempty"`

	// The group ID of the project to exclude.
	// @version 4.0.0+
	GroupId string `yaml:"groupId" xml:"groupId,omitempty"`
}

// The <code>&lt;execution&gt;</code> element contains informations required for the execution of a plugin.
// @version 4.0.0+
type PluginExecution struct {
	// The identifier of this execution for labelling the goals during the build, and for matching executions to merge during inheritance and profile injection.
	// @version 4.0.0+
	Id string `yaml:"id" xml:"id,omitempty"`

	// The build lifecycle phase to bind the goals in this execution to. If omitted, the goals will be bound to the default phase specified by the plugin.
	// @version 4.0.0+
	Phase string `yaml:"phase" xml:"phase,omitempty"`

	// The goals to execute with the given configuration.
	// @version 4.0.0+
	Goals *Goals `yaml:"goals" xml:"goals,omitempty"`

	// Whether any configuration should be propagated to child POMs.
	// Note: While the type of this field is <code>String</code> for technical reasons, the semantic type is actually <code>Boolean</code>.
	// Default value is <code>true</code>.
	// @version 4.0.0+
	Inherited string `yaml:"inherited" xml:"inherited,omitempty"`

	// <p>The configuration as DOM object.</p> <p>By default, every element content is trimmed, but starting with Maven 3.1.0, you can add <code>xml:space="preserve"</code> to elements you want to preserve whitespace.</p>
	// <p>You can control how child POMs inherit configuration from parent POMs by adding <code>combine.children</code> or <code>combine.self</code> attributes to the children of the configuration element:</p>
	// <ul>
	// 	<li> <code>combine.children</code>: available values are <code>merge</code> (default) and <code>append</code>,</li>
	// 	<li><code>combine.self</code>: available values are <code>merge</code> (default) and <code>override</code>.</li>
	// </ul>
	// <p>See
	// 	<a href="http://maven.apache.org/pom.html#Plugins">POM Reference documentation</a> and
	// 	<a href="http://plexus.codehaus.org/plexus-utils/apidocs/org/codehaus/plexus/util/xml/Xpp3DomUtils.html">Xpp3DomUtils</a>
	// 	for more information.
	// </p>
	// @version 0.0.0+
	Configuration InterfaceMap /* *interface{} */ `yaml:"configuration,flow" xml:"configuration,omitempty"`
}

// This element describes all of the classpath resources associated with a project or unit tests.
// @version 3.0.0+
type Resource struct {
	// Describe the resource target path.
	// The path is relative to the target/classes directory (i.e. <code>${project.build.outputDirectory}</code>).
	// For example, if you want that resource to appear in a specific package (<code>org.apache.maven.messages</code>),
	// you must specify this element with this value: <code>org/apache/maven/messages</code>.
	// This is not required if you simply put the resources in that directory structure at the source, however.
	// @version 3.0.0+
	TargetPath string `yaml:"targetPath" xml:"targetPath,omitempty"`

	// Whether resources are filtered to replace tokens with parameterised values or not.
	// The values are taken from the <code>properties</code> element and from the properties in the files listed in the <code>filters</code> element.
	// Note: While the type of this field is <code>String</code> for technical reasons, the semantic type is actually <code>Boolean</code>.
	// Default value is <code>false</code>.
	// @version 3.0.0+
	Filtering string `yaml:"filtering" xml:"filtering,omitempty"`

	// Describe the directory where the resources are stored. The path is relative to the POM.
	// @version 3.0.0+
	Directory string `yaml:"directory" xml:"directory,omitempty"`

	// A list of patterns to include, e.g. <code>**&#47;*.xml</code>.
	// @version 3.0.0+
	Includes *Includes `yaml:"includes" xml:"includes,omitempty"`

	// A list of patterns to exclude, e.g. <code>**&#47;*.xml</code>
	// @version 3.0.0+
	Excludes *Excludes `yaml:"excludes" xml:"excludes,omitempty"`
}

// Section for management of default plugin information for use in a group of POMs.
// @version 4.0.0+
type PluginManagement struct {
	// The list of plugins to use.
	// @version 4.0.0+
	Plugins *Plugins `yaml:"plugins" xml:"plugins,omitempty"`
}

// Section for management of reports and their configuration.
// @version 4.0.0+
type Reporting struct {
	// If true, then the default reports are not included in the site generation.
	// This includes the reports in the "Project Info" menu.
	// Note: While the type of this field is <code>String</code> for technical reasons, the semantic type is actually <code>Boolean</code>.
	// Default value is <code>false</code>.
	// @version 4.0.0+
	ExcludeDefaults string `yaml:"excludeDefaults" xml:"excludeDefaults,omitempty"`

	// Where to store all of the generated reports. The default is <code>${project.build.directory}/site</code>.
	// @version 4.0.0+
	OutputDirectory string `yaml:"outputDirectory" xml:"outputDirectory,omitempty"`

	// The reporting plugins to use and their configuration.
	// @version 4.0.0+
	Plugins *ReportPlugins `yaml:"plugins" xml:"plugins,omitempty"`
}

// The <code>&lt;plugin&gt;</code> element contains informations required for a report plugin.
// @version 4.0.0+
type ReportPlugin struct {
	// The group ID of the reporting plugin in the repository.
	// @version 4.0.0+
	GroupId string `yaml:"groupId" xml:"groupId,omitempty"`

	// The artifact ID of the reporting plugin in the repository.
	// @version 4.0.0+
	ArtifactId string `yaml:"artifactId" xml:"artifactId,omitempty"`

	// The version of the reporting plugin to be used.
	// @version 4.0.0+
	Version string `yaml:"version" xml:"version,omitempty"`

	// Multiple specifications of a set of reports, each having (possibly) different configuration.
	// This is the reporting parallel to an <code>execution</code> in the build.
	// @version 4.0.0+
	ReportSets *ReportSets `yaml:"reportSets" xml:"reportSets,omitempty"`

	// Whether any configuration should be propagated to child POMs.
	// Note: While the type of this field is <code>String</code> for technical reasons, the semantic type is actually <code>Boolean</code>.
	// Default value is <code>true</code>.
	// @version 4.0.0+
	Inherited string `yaml:"inherited" xml:"inherited,omitempty"`

	// <p>The configuration as DOM object.</p>
	// <p>By default, every element content is trimmed, but starting with Maven 3.1.0, you can add <code>xml:space="preserve"</code> to elements you want to preserve whitespace.</p>
	// <p>You can control how child POMs inherit configuration from parent POMs by adding <code>combine.children</code> or <code>combine.self</code> attributes to the children of the configuration element:</p>
	// <ul>
	// 	<li><code>combine.children</code>: available values are <code>merge</code> (default) and <code>append</code>,</li>
	// 	<li><code>combine.self</code>: available values are <code>merge</code> (default) and <code>override</code>.</li>
	// </ul>
	// <p>See
	// 	<a href="http://maven.apache.org/pom.html#Plugins">POM Reference documentation</a> and
	// 	<a href="http://plexus.codehaus.org/plexus-utils/apidocs/org/codehaus/plexus/util/xml/Xpp3DomUtils.html">Xpp3DomUtils</a>
	// 	for more information.
	// </p>
	// @version 0.0.0+
	Configuration InterfaceMap /* *interface{} */ `yaml:"configuration,flow" xml:"configuration,omitempty"`
}

// Represents a set of reports and configuration to be used to generate them.
// @version 4.0.0+
type ReportSet struct {
	// The unique id for this report set, to be used during POM inheritance and profile injection for merging of report sets.
	// @version 0.0.0+
	Id            string `yaml:"id" xml:"id,omitempty"`

	// The list of reports from this plugin which should be generated from this set.
	// @version 4.0.0+
	Reports       *Reports `yaml:"reports" xml:"reports,omitempty"`

	// Whether any configuration should be propagated to child POMs.
	// Note: While the type of this field is <code>String</code> for technical reasons, the semantic type is actually <code>Boolean</code>.
	// Default value is <code>true</code>.
	// @version 4.0.0+
	Inherited string `yaml:"inherited" xml:"inherited,omitempty"`

	// <p>The configuration as DOM object.</p>
	// <p>By default, every element content is trimmed, but starting with Maven 3.1.0, you can add <code>xml:space="preserve"</code> to elements you want to preserve whitespace.</p>
	// <p>You can control how child POMs inherit configuration from parent POMs by adding <code>combine.children</code> or <code>combine.self</code> attributes to the children of the configuration element:</p>
	// <ul>
	// 	<li><code>combine.children</code>: available values are <code>merge</code> (default) and <code>append</code>,</li>
	// 	<li><code>combine.self</code>: available values are <code>merge</code> (default) and <code>override</code>.</li>
	// </ul>
	// <p>See
	// 	<a href="http://maven.apache.org/pom.html#Plugins">POM Reference documentation</a> and
	// 	<a href="http://plexus.codehaus.org/plexus-utils/apidocs/org/codehaus/plexus/util/xml/Xpp3DomUtils.html">Xpp3DomUtils</a>
	// 	for more information.
	// </p>
	// @version 0.0.0+
	Configuration InterfaceMap /* *interface{} */ `yaml:"configuration,flow" xml:"configuration,omitempty"`
}

// The conditions within the build runtime environment which will trigger the automatic inclusion of the build profile.
// Multiple conditions can be defined, which must be all satisfied to activate the profile.
// @version 4.0.0+
type Activation struct {
	// If set to true, this profile will be active unless another profile in this pom is activated using the command line -P option or by one of that profile's activators.
	// @version 4.0.0+
	ActiveByDefault string /*bool*/ `yaml:"activeByDefault" xml:"activeByDefault,omitempty"`

	// Specifies that this profile will be activated when a matching JDK is detected.
	// For example, <code>1.4</code> only activates on JDKs versioned 1.4, while <code>!1.4</code> matches any JDK that is not version 1.4.
	// Ranges are supported too: <code>[1.5,)</code> activates when the JDK is 1.5 minimum.
	// @version 4.0.0+
	Jdk string `yaml:"jdk" xml:"jdk,omitempty"`

	// Specifies that this profile will be activated when matching operating system attributes are detected.
	// @version 4.0.0+
	Os *ActivationOS `yaml:"os" xml:"os,omitempty"`

	// Specifies that this profile will be activated when this system property is specified.
	// @version 4.0.0+
	Property *ActivationProperty `yaml:"property" xml:"property,omitempty"`

	// Specifies that this profile will be activated based on existence of a file.
	// @version 4.0.0+
	File *ActivationFile `yaml:"file" xml:"file,omitempty"`
}

// This is the property specification used to activate a profile.
// If the value field is empty, then the existence of the named property will activate the profile, otherwise it does a case-sensitive match against the property value as well.
// @version 4.0.0+
type ActivationProperty struct {
	// The name of the property to be used to activate a profile.
	// @version 4.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// The value of the property required to activate a profile.
	// @version 4.0.0+
	Value string `yaml:"value" xml:"value,omitempty"`
}

// This is an activator which will detect an operating system's attributes in order to activate its profile.
// @version 4.0.0+
type ActivationOS struct {
	// The name of the operating system to be used to activate the profile.
	// This must be an exact match of the <code>${os.name}</code> Java property, such as <code>Windows XP</code>.
	// @version 4.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// The general family of the OS to be used to activate the profile, such as <code>windows</code> or <code>unix</code>.
	// @version 4.0.0+
	Family string `yaml:"family" xml:"family,omitempty"`

	// The architecture of the operating system to be used to activate the profile.
	// @version 4.0.0+
	Arch string `yaml:"arch" xml:"arch,omitempty"`

	// The version of the operating system to be used to activate the profile.
	// @version 4.0.0+
	Version string `yaml:"version" xml:"version,omitempty"`
}

// This is the file specification used to activate the profile.
// The <code>missing</code> value is the location of a file that needs to exist, and if it doesn't, the profile will be activated.
// On the other hand, <code>exists</code> will test for the existence of the file and if it is there, the profile will be activated.<br/>
// Variable interpolation for these file specifications is limited to <code>${basedir}</code>, System properties and request properties.
// @version 4.0.0+
type ActivationFile struct {
	// The name of the file that must be missing to activate the profile.
	// @version 4.0.0+
	Missing string `yaml:"missing" xml:"missing,omitempty"`

	// The name of the file that must exist to activate the profile.
	// @version 4.0.0+
	Exists string `yaml:"exists" xml:"exists,omitempty"`
}

// Section for management of default dependency information for use in a group of POMs.
// @version 4.0.0+
type DependencyManagement struct {
	// The dependencies specified here are not used until they are referenced in a POM within the group.
	// This allows the specification of a "standard" version for a particular dependency.
	// @version 4.0.0+
	Dependencies *Dependencies `yaml:"dependencies" xml:"dependencies,omitempty"`
}

// The <code>&lt;build&gt;</code> element contains informations required to build the project. Default values are defined in Super POM.
// @version 3.0.0+
type Build struct {
	// This element specifies a directory containing the source of the project.
	// The generated build system will compile the sources from this directory when the project is built.
	// The path given is relative to the project descriptor.
	// The default value is <code>src/main/java</code>.
	// @version 3.0.0+
	SourceDirectory string `yaml:"sourceDirectory" xml:"sourceDirectory,omitempty"`

	// This element specifies a directory containing the script sources of the project.
	// This directory is meant to be different from the sourceDirectory, in that its contents will be copied to the output directory in most cases (since scripts are interpreted rather than compiled).
	// The default value is <code>src/main/scripts</code>.
	// @version 4.0.0+
	ScriptSourceDirectory string `yaml:"scriptSourceDirectory" xml:"scriptSourceDirectory,omitempty"`

	// This element specifies a directory containing the unit test source of the project.
	// The generated build system will compile these directories when the project is being tested. The path given is relative to the project descriptor.
	// The default value is <code>src/test/java</code>.
	// @version 4.0.0+
	TestSourceDirectory string `yaml:"testSourceDirectory" xml:"testSourceDirectory,omitempty"`

	// The directory where compiled application classes are placed.
	// The default value is <code>target/classes</code>.
	// @version 4.0.0+
	OutputDirectory string `yaml:"outputDirectory" xml:"outputDirectory,omitempty"`

	// The directory where compiled test classes are placed.
	// The default value is <code>target/test-classes</code>.
	// @version 4.0.0+
	TestOutputDirectory string `yaml:"testOutputDirectory" xml:"testOutputDirectory,omitempty"`

	// A set of build extensions to use from this project.
	// @version 4.0.0+
	Extensions *Extensions `yaml:"extensions" xml:"extensions,omitempty"`

	// The default goal (or phase in Maven 2) to execute when none is specified for the project.
	// Note that in case of a multi-module build, only the default goal of the top-level project is relevant, i.e. the default goals of child modules are ignored.
	// Since Maven 3, multiple goals/phases can be separated by whitespace.
	// @version 3.0.0+
	DefaultGoal string `yaml:"defaultGoal" xml:"defaultGoal,omitempty"`

	// This element describes all of the classpath resources such as properties files associated with a project.
	// These resources are often included in the final package.
	// The default value is <code>src/main/resources</code>.
	// @version 3.0.0+
	Resources *Resources `yaml:"resources" xml:"resources,omitempty"`

	// This element describes all of the classpath resources such as properties files associated with a project's unit tests.
	// The default value is <code>src/test/resources</code>.
	// @version 4.0.0+
	TestResources *TestResources `yaml:"testResources" xml:"testResources,omitempty"`

	// The directory where all files generated by the build are placed.
	// The default value is <code>target</code>.
	// @version 4.0.0+
	Directory string `yaml:"directory" xml:"directory,omitempty"`

	// The filename (excluding the extension, and with no path information) that the produced artifact will be called.
	// The default value is <code>${artifactId}-${version}</code>.
	// @version 4.0.0+
	FinalName string `yaml:"finalName" xml:"finalName,omitempty"`

	// The list of filter properties files that are used when filtering is enabled.
	// @version 4.0.0+
	Filters *Filters `yaml:"filters" xml:"filters,omitempty"`

	// Default plugin information to be made available for reference by projects derived from this one.
	// This plugin configuration will not be resolved or bound to the lifecycle unless referenced.
	// Any local configuration for a given plugin will override the plugin's entire definition here.
	// @version 4.0.0+
	PluginManagement *PluginManagement `yaml:"pluginManagement" xml:"pluginManagement,omitempty"`

	// The list of plugins to use.
	// @version 4.0.0+
	Plugins *Plugins `yaml:"plugins" xml:"plugins,omitempty"`
}

// Describes a build extension to utilise.
// @version 4.0.0+
type Extension struct {
	// The group ID of the extension's artifact.
	// @version 4.0.0+
	GroupId string `yaml:"groupId" xml:"groupId,omitempty"`

	// The artifact ID of the extension.
	// @version 4.0.0+
	ArtifactId string `yaml:"artifactId" xml:"artifactId,omitempty"`

	// The version of the extension.
	// @version 4.0.0+
	Version string `yaml:"version" xml:"version,omitempty"`
}

// Describes the licenses for this project.
// This is used to generate the license page of the project's web site, as well as being taken into consideration in other reporting and validation.
// The licenses listed for the project are that of the project itself, and not of dependencies.
// @version 3.0.0+
type License struct {
	// The full legal name of the license.
	// @version 3.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// The official url for the license text.
	// @version 3.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`

	// The primary method by which this project may be distributed.
	// <dl>
	// 	<dt>repo</dt>
	// 	<dd>may be downloaded from the Maven repository</dd>
	// 	<dt>manual</dt>
	// 	<dd>user must manually download and install the dependency.</dd>
	// </dl>
	// @version 3.0.0+
	Distribution string `yaml:"distribution" xml:"distribution,omitempty"`

	// Addendum information pertaining to this license.
	// @version 3.0.0+
	Comments string `yaml:"comments" xml:"comments,omitempty"`
}

// This element describes all of the mailing lists associated with a project.
// The auto-generated site references this information.
// @version 3.0.0+
type MailingList struct {
	// The name of the mailing list.
	// @version 3.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// The email address or link that can be used to subscribe to the mailing list.
	// If this is an email address, a <code>mailto:</code> link will automatically be created when the documentation is created.
	// @version 3.0.0+
	Subscribe string `yaml:"subscribe" xml:"subscribe,omitempty"`

	// The email address or link that can be used to unsubscribe to the mailing list.
	// If this is an email address, a <code>mailto:</code> link will automatically be created when the documentation is created.
	// @version 3.0.0+
	Unsubscribe string `yaml:"unsubscribe" xml:"unsubscribe,omitempty"`

	// The email address or link that can be used to post to the mailing list.
	// If this is an email address, a <code>mailto:</code> link will automatically be created when the documentation is created.
	// @version 3.0.0+
	Post string `yaml:"post" xml:"post,omitempty"`

	// The link to a URL where you can browse the mailing list archive.
	// @version 3.0.0+
	Archive string `yaml:"archive" xml:"archive,omitempty"`

	// The link to alternate URLs where you can browse the list archive.
	// @version 3.0.0+
	OtherArchives *OtherArchives `yaml:"otherArchives" xml:"otherArchives,omitempty"`
}

// Information about one of the committers on this project.
// @version 3.0.0+
type Developer struct {
	// The unique ID of the developer in the SCM.
	// @version 3.0.0+
	Id string `yaml:"id" xml:"id,omitempty"`

	// The full name of the contributor.
	// @version 3.0.0+
	Name string `yaml:"name" xml:"name,omitempty"`

	// The email address of the contributor.
	// @version 3.0.0+
	Email string `yaml:"email" xml:"email,omitempty"`

	// The URL for the homepage of the contributor.
	// @version 3.0.0+
	Url string `yaml:"url" xml:"url,omitempty"`

	// The organization to which the contributor belongs.
	// @version 3.0.0+
	Organization string `yaml:"organization" xml:"organization,omitempty"`

	// The URL of the organization.
	// @version 3.0.0+
	OrganizationUrl string `yaml:"organizationUrl" xml:"organizationUrl,omitempty"`

	// The roles the contributor plays in the project.
	// Each role is described by a <code>role</code> element, the body of which is a role name.
	// This can also be used to describe the contribution.
	// @version 3.0.0+
	Roles      *Roles `yaml:"roles" xml:"roles,omitempty"`

	// The timezone the contributor is in.
	// Typically, this is a number in the range <a href="http://en.wikipedia.org/wiki/UTC%E2%88%9212:00">-12</a> to <a href="http://en.wikipedia.org/wiki/UTC%2B14:00">+14</a>
	// or a valid time zone id like "America/Montreal" (UTC-05:00) or "Europe/Paris" (UTC+01:00).
	// @version 3.0.0+
	Timezone   string `yaml:"timezone" xml:"timezone,omitempty"`

	// Properties about the contributor, such as an instant messenger handle.
	// @version 3.0.0+
	Properties InterfaceMap `yaml:"properties,flow" xml:"properties,omitempty"`
}
