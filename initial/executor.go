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

type Executor interface {
	Exec() error
}

func WithExecutor(n *New) (Executor, error) {
	if n.Conf.Web {
		return &ExecutorWebapp{n}, nil
	}

	if n.Conf.Scheduler {
		return &ExecutorScheduler{&ExecutorWebapp{n}, n}, nil
	}

	return &ExecutorYml{&ExecutorWebapp{n}, n}, nil
}
