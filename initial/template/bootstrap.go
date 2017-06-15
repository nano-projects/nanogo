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

func Bootstrap() (*template.Template, error) {
	shell := `#!/bin/sh

#check JAVA_HOME & java
#noJavaHome=false
#if [ -z "$JAVA_HOME" ] ; then
#    noJavaHome=true
#fi
#if [ ! -e "$JAVA_HOME/bin/java" ] ; then
#    noJavaHome=true
#fi
#if $noJavaHome ; then
#    echo
#    echo "Error: JAVA_HOME environment variable is not set."
#    echo
#    exit 1
#fi
#==============================================================================
CURR_DIR=` + "`pwd`" + `

# set JAVA_HOME
#cd ` + "`dirname \"$0\"`/.." + `
# cd ..
# JAVA_HOME=` + "`pwd`" + `/jdk
# cd ./bin

#set JAVA_OPTS
JAVA_OPTS="-server -Xms512m -Xmx512m -Xmn64m -Xss256k"
#performance Options
JAVA_OPTS="$JAVA_OPTS -XX:+AggressiveOpts"
JAVA_OPTS="$JAVA_OPTS -XX:+UseBiasedLocking"
JAVA_OPTS="$JAVA_OPTS -XX:+UseFastAccessorMethods"
JAVA_OPTS="$JAVA_OPTS -XX:+DisableExplicitGC"
JAVA_OPTS="$JAVA_OPTS -XX:+UseParNewGC"
JAVA_OPTS="$JAVA_OPTS -XX:+UseConcMarkSweepGC"
JAVA_OPTS="$JAVA_OPTS -XX:+CMSParallelRemarkEnabled"
JAVA_OPTS="$JAVA_OPTS -XX:+UseCMSCompactAtFullCollection"
JAVA_OPTS="$JAVA_OPTS -XX:+UseCMSInitiatingOccupancyOnly"
JAVA_OPTS="$JAVA_OPTS -XX:CMSInitiatingOccupancyFraction=75"
#==============================================================================

#set HOME
# CURR_DIR=` + "`pwd`" + `
cd ` + "`dirname \"$0\"`/.." + `
APP_HOME=` + "`pwd`" + `
cd $CURR_DIR
if [ -z "$APP_HOME" ] ; then
echo
echo "Error: APP_HOME environment variable is not defined correctly."
echo
exit 1
fi
#==============================================================================

#set CLASSPATH
#APP_CLASSPATH="$APP_HOME/app:$APP_HOME/app/lib"

for i in "$APP_HOME"/lib/*.jar
do
APP_CLASSPATH="$APP_CLASSPATH:$i"
done

APP_CLASSPATH="$APP_CLASSPATH:$APP_HOME/conf"

LOGS_HOME="$APP_HOME/bin/logs"
echo $LOGS_HOME
if [ ! -d "$LOGS_HOME" ]; then
mkdir "$LOGS_HOME"
fi

#==============================================================================

#startup Server
RUN_CMD="\"$JAVA_HOME/bin/java\""
RUN_CMD="$RUN_CMD -Dlogic.home=\"$APP_HOME\""
RUN_CMD="$RUN_CMD -classpath \"$APP_CLASSPATH\""
RUN_CMD="$RUN_CMD $JAVA_OPTS"
# replace Bootstrap class name
RUN_CMD="$RUN_CMD {{.Bootstrap}} $@"
RUN_CMD="$RUN_CMD > /dev/null 2>&1"
RUN_CMD="$RUN_CMD &"
echo $RUN_CMD
eval $RUN_CMD
#==============================================================================

`

	return template.New("Bootstrap").Parse(shell)
}

func BootstrapClass() (*template.Template, error) {
	bootstrap := license.Class() + `
package {{.SourcePackage}};

import org.nanoframework.server.GitPull;
import org.nanoframework.server.{{.Server}}CustomServer;

/**
 *
 * @author yanghe
 * @since {{.BootstrapVersion}}
 */
public final class {{.BootstrapClassName}} {
    private {{.BootstrapClassName}}() { }

    /**
     *
     * @param args bootstrap parameters.
     * @throws Throwable Server startup error
     */
    public static void main(final String[] args) throws Throwable {
        GitPull.create().quickPull(args);
        {{.Server}}CustomServer.server().bootstrap(args);
    }
}
`

	return template.New("BootstrapClass").Parse(bootstrap)
}
