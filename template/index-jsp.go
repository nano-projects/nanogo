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

package template

import (
	"github.com/nano-projects/nanogo/template/license"
	"text/template"
)

func IndexJsp() (*template.Template, error) {
	jsp := license.Jsp() + `
<html>
<body>
<h2>Hello World!</h2>
</body>
</html>
`

	return template.New("IndexJsp").Parse(jsp)
}
