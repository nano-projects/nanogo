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

package cmd

import (
	"github.com/nano-projects/nanogo/initial"
	"github.com/nano-projects/nanogo/initial/conf"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/pom"
	"github.com/spf13/cobra"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"github.com/pkg/errors"
)

// initCmd represents the new command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a maven project",
	RunE: func(cmd *cobra.Command, args []string) error {
		web, err := cmd.Flags().GetBool("web")
		if err != nil {
			return err
		}

		scheduler, err := cmd.Flags().GetBool("scheduler")
		if err != nil {
			return err
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}

		template, err := cmd.Flags().GetString("template")
		if err != nil {
			return err
		}

		parent, err := cmd.Flags().GetString("parent")
		if err != nil {
			return err
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		publish, err := cmd.Flags().GetUint("publish")
		if err != nil {
			return err
		}

		server, err := cmd.Flags().GetString("server")

		parentDep, err := pom.Dependency{}.Parse(parent, false)
		if err != nil {
			return err
		}

		nameDep, err := pom.Dependency{}.Parse(name, true)
		if err != nil {
			return err
		}

		initialConf := conf.InitialConfig{
			Web:       web,
			Scheduler: scheduler,
			Path:      path,
			Template:  template,
			Parent:    parentDep,
			Name:      nameDep,
			Publish:   publish,
		}

		contextRoot := strings.Replace(nameDep.ArtifactId, "-", ".", -1)
		pkg := filepath.Join(strings.Replace(strings.Replace(nameDep.GroupId, ".", "/", -1), "-", "/", -1), strings.Replace(nameDep.ArtifactId, "-", "/", -1))
		srcPkg := strings.Replace(nameDep.GroupId, "-", ".", -1) + "." + strings.Replace(nameDep.ArtifactId, "-", ".", -1)
		componentPkg := srcPkg + ".component"
		schedulerPkg := srcPkg + ".scheduler"

		var bootstrapClassName string
		arts := strings.Split(nameDep.ArtifactId, "-")
		for _, art := range arts {
			bootstrapClassName += strings.ToUpper(art[:1]) + art[1:]
		}

		displayName := bootstrapClassName
		bootstrapClassName += "Bootstrap"
		bootstrap := srcPkg + "." + bootstrapClassName
		tmpConf := conf.TmpConfig{
			GroupId:            nameDep.GroupId,
			ArtifactId:         nameDep.ArtifactId,
			Version:            nameDep.Version,
			ParentGroupId:      parentDep.GroupId,
			ParentArtifactId:   parentDep.ArtifactId,
			ParentVersion:      parentDep.Version,
			Package:            pkg,
			SourcePackage:      srcPkg,
			ComponentPackage:   componentPkg,
			SchedulerPackage:   schedulerPkg,
			BootstrapClassName: bootstrapClassName,
			Bootstrap:          bootstrap,
			BootstrapVersion:   strings.Replace(nameDep.Version, pom.Snapshot, "", -1),
			ContextRoot:        contextRoot,
			Publish:            strconv.Itoa(int(initialConf.Publish)),
			Year:               strconv.Itoa(time.Now().Year()),
			DisplayName:        displayName,
		}

		if server == "Tomcat" || server == "Jetty" {
			tmpConf.Server = server
		} else if server != "" {
			return errors.New("Unknown server type: " + server + " (option 'Tomcat' or 'Jetty')")
		}

		initial := &initial.Initial{
			Conf: initialConf,
			Tmp:  tmpConf,
		}

		return initial.Run()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("web", "w", false, "Init a webapp project of nano framework")
	initCmd.Flags().BoolP("scheduler", "s", false, "Init a scheduler project of nano framework")
	initCmd.Flags().String("path", io.Pwd(), "The project path by default using the current path")
	initCmd.Flags().StringP("template", "t", "", "The project template file path")
	initCmd.Flags().String("parent", "org.nanoframework:super:0.0.14", `Maven top POM dependency, format: "groupId:artifactId:version"`)
	initCmd.Flags().StringP("name", "n", "", `Maven project name definition, format: "groupId:artifactId:version", version is optional, the default use of 0.0.1`)
	initCmd.Flags().UintP("publish", "p", 7000, "Project default port")
	initCmd.Flags().String("server", "", `Set up the server (option 'Tomcat' or 'Jetty')`)
}
