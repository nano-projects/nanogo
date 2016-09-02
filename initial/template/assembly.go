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

func Assembly() (*template.Template, error) {
	xml := license.Xml() + `
<assembly xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSch ema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/assembly-1.0.0.xsd">

  <formats>
    <format>tar.gz</format>
  </formats>
  <includeBaseDirectory>true</includeBaseDirectory>
  <files>
  </files>
  <fileSets>
    <fileSet>
      <directory>bin</directory>
      <outputDirectory>bin</outputDirectory>
      <fileMode>0755</fileMode>
      <lineEnding>keep</lineEnding>
    </fileSet>
    <fileSet>
      <directory>webRoot</directory>
      <outputDirectory>bin/webRoot</outputDirectory>
    </fileSet>
    <fileSet>
      <directory>lib</directory>
      <outputDirectory>lib</outputDirectory>
    </fileSet>
    <fileSet>
        <directory>target/conf</directory>
        <outputDirectory>conf</outputDirectory>
    </fileSet>
  </fileSets>
  <dependencySets>
    <dependencySet>
      <outputDirectory>lib</outputDirectory>
      <scope>runtime</scope>
    </dependencySet>
  </dependencySets>

</assembly>
`

	return template.New("Assembly").Parse(xml)
}
