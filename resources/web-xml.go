package resources

import "github.com/nano-projects/nanogo/resources/license"

func WebXml(displayName string) string {
	return license.Xml() + `
<web-app xmlns="http://java.sun.com/xml/ns/javaee" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://java.sun.com/xml/ns/javaee  http://java.sun.com/xml/ns/javaee/web-app_3_0.xsd"
    version="3.0">
    <display-name>` + displayName + `</display-name>

    <filter>
        <filter-name>httpRequestFilter</filter-name>
        <filter-class>org.nanoframework.web.server.filter.HttpRequestFilter</filter-class>
    </filter>

    <filter-mapping>
        <filter-name>httpRequestFilter</filter-name>
        <url-pattern>/*</url-pattern>
    </filter-mapping>

    <servlet>
        <servlet-name>Dispatcher-Servlet</servlet-name>
        <servlet-class>org.nanoframework.web.server.servlet.DispatcherServlet</servlet-class>
        <init-param>
            <param-name>context</param-name>
            <param-value>/context.properties</param-value>
        </init-param>
        <load-on-startup>1</load-on-startup>
    </servlet>

    <servlet-mapping>
        <servlet-name>Dispatcher-Servlet</servlet-name>
        <url-pattern>/dispatcher/*</url-pattern>
    </servlet-mapping>

    <welcome-file-list>
        <welcome-file>index.jsp</welcome-file>
    </welcome-file-list>
</web-app>
`
}
