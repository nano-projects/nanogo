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

func CheckstyleSuppressions() (*template.Template, error) {
	supperessions := `<?xml version="1.0" encoding="UTF-8"?>
` + license.Xml() + `
<!DOCTYPE suppressions PUBLIC "-//Puppy Crawl//DTD Suppressions 1.1//EN" "http://www.puppycrawl.com/dtds/suppressions_1_1.dtd">

<suppressions>
    <suppress checks="JavadocStyleCheck" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="JavadocMethod" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="JavadocType" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="JavadocVariable" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="MagicNumber" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress checks="DesignForExtension" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress id="stackTraceConsoleLogs" files="(.*Tests*|Mock.*|Test.*)\.java" />
    <suppress id="sysOutConsoleLogs" files="\.xml" />
    <suppress checks="\w+" files="(\.(crt|crl|class|keystore))|rebel.xml" />
</suppressions>
`

	return template.New("CheckstyleSuppressions").Parse(supperessions)
}
