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
		return &ExecutorWebapp{n}, nil
	}

	if n.Conf.Scheduler {
		return &ExecutorScheduler{&ExecutorWebapp{n}}, nil
	}

	return &ExecutorYml{&ExecutorWebapp{n}}, nil
}
