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

func JettyXml() (*template.Template, error) {
	xml := `<?xml version="1.0" encoding="UTF-8"?>
` + license.Xml() + `
<!DOCTYPE Configure PUBLIC "-//Jetty//Configure//EN" "http://www.eclipse.org/jetty/configure.dtd">

<Configure id="Server" class="org.eclipse.jetty.server.Server">
    <Set name="ThreadPool">
        <!-- Default queued blocking threadpool -->
        <New class="org.eclipse.jetty.util.thread.QueuedThreadPool">
            <Set name="minThreads">5</Set>
            <Set name="maxThreads">200</Set>
            <Set name="detailedDump">false</Set>
        </New>
    </Set>

    <Call name="addConnector">
        <Arg>
            <New class="org.eclipse.jetty.server.nio.SelectChannelConnector">
                <Set name="host">
                    <Property name="jetty.host" />
                </Set>
                <Set name="port">
                    <Property name="jetty.port" default="{{.Publish}}" />
                </Set>
                <Set name="maxIdleTime">300000</Set>
                <Set name="Acceptors">2</Set>
                <Set name="statsOn">false</Set>
                <Set name="confidentialPort">8443</Set>
                <Set name="lowResourcesConnections">20000</Set>
                <Set name="lowResourcesMaxIdleTime">5000</Set>
            </New>
        </Arg>
    </Call>
</Configure>
`

	return template.New("JettyXml").Parse(xml)
}
