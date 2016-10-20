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

package initial

import (
	"github.com/nano-projects/nanogo/exec"
	"github.com/nano-projects/nanogo/initial/conf"
)

type Initial struct {
	Conf conf.InitialConfig
	Tmp  conf.TmpConfig
}

func (n *Initial) Run() error {
	if err := n.Conf.Valid(); err != nil {
		return err
	}

	exec, err := n.withExecutor()
	if err != nil {
		return err
	}

	return exec.Exec()
}

func (n *Initial) withExecutor() (exec.Executor, error) {
	if n.Conf.Web {
		n.serverDependencies("Tomcat")
		return &ExecutorWebapp{n}, nil
	}

	if n.Conf.Scheduler {
		n.serverDependencies("Jetty")
		return &ExecutorScheduler{&ExecutorWebapp{n}}, nil
	}

	n.serverDependencies("Tomcat")
	return &ExecutorYml{&ExecutorWebapp{n}}, nil
}

func (n *Initial) serverDependencies(defServer string) {
	if n.Tmp.Server == "" {
		n.Tmp.Server = defServer
	}

	switch n.Tmp.Server {
	case "Tomcat":
		n.Tmp.ServerDependencies = `- groupId: org.nanoframework
          artifactId: nano-tomcat-server`
	case "Jetty":
		n.Tmp.ServerDependencies = `- groupId: org.eclipse.jetty.orbit
          artifactId: javax.servlet.jsp
        - groupId: org.nanoframework
          artifactId: nano-jetty-server
          exclusions:
            exclusion:
              - groupId: javax.servlet.jsp
                artifactId: javax.servlet.jsp-api`
	default:
		n.Tmp.ServerDependencies = `- groupId: org.nanoframework
          artifactId: nano-tomcat-server`
	}
}