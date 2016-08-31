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

package pom

import (
	"github.com/pkg/errors"
	"strings"
)

const (
	DefaultDependencyVersion = "0.0.1-SNAPSHOT"
	Snapshot                 = "-SNAPSHOT"
)

func (d Dependency) Parse(dep string) (Dependency, error) {
	if dep == "" {
		return d, errors.New("Unknown maven project dependency definition")
	}

	deps := strings.Split(dep, ":")
	if len(deps) < 2 || len(deps) > 3 {
		return d, errors.New("Unknown format of Dependency: " + dep)
	}

	d.GroupId = deps[0]
	d.ArtifactId = deps[1]
	if len(deps) == 3 {
		if !strings.HasSuffix(deps[2], Snapshot) {
			d.Version = deps[2] + Snapshot
		} else {
			d.Version = deps[2]
		}
	} else {
		d.Version = DefaultDependencyVersion
	}

	return d, nil
}
