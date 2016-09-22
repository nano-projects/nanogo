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

func CodeTemplate() (*template.Template, error) {
	codeTemplate := `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
` + license.Xml() + `
<templates>
<template autoinsert="true" context="filecomment_context" deleted="false" description="Comment for created Java files" enabled="true" id="org.eclipse.jdt.ui.text.codetemplates.filecomment" name="filecomment">/*
 * Copyright © 2015-{{.Year}} the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */</template>
 <template autoinsert="true" context="typecomment_context" deleted="false" description="Comment for create Java type files" enabled="true" id="org.eclipse.jdt.ui.text.codetemplates.typecomment" name="typecomment">/**
 *
 * @author ${user}
 * @since
 * ${tags}
 */</template>
 </templates>
 `

	return template.New("CodeTemplate").Parse(codeTemplate)
}
