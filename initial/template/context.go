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
	"github.com/nano-projects/nanogo/initial/template/license"
	"text/template"
)

func WebappContext() (*template.Template, error) {
	context := license.Properties() + `
# 组件服务上下文属性文件列表
context=

# 应用模式, DEV: 开发模式, PROD: 生产模式
context.mode=DEV

# 版本号
context.version={{.Version}}

# 服务根
context.root=/{{.ContextRoot}}

context.component-scan.base-package={{.ComponentPackage}}
`
	return template.New("WebappContext").Parse(context)
}

func SchedulerContext() (*template.Template, error) {
	context := license.Properties() + `
# 组件服务上下文属性文件列表
context=

# 应用模式, DEV: 开发模式, PROD: 生产模式
context.mode=DEV

# 版本号
context.version={{.Version}}

# 服务根
context.root=/{{.ContextRoot}}

context.scheduler-scan.base-package={{.SchedulerPackage}}
`
	return template.New("SchedulerContext").Parse(context)
}

func Context() (*template.Template, error) {
	context := license.Properties() + `
# 组件服务上下文属性文件列表
context=

# 应用模式, DEV: 开发模式, PROD: 生产模式
context.mode=DEV

# 版本号
context.version={{.Version}}

# 服务根
context.root=/{{.ContextRoot}}
`
	return template.New("Context").Parse(context)
}
