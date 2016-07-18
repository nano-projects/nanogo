package resources

import "github.com/nano-projects/nanogo/resources/license"

func JettyXml(port string) string {
	return `<?xml version="1.0" encoding="UTF-8"?>
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
                    <Property name="jetty.port" default="` + port + `" />
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

}
