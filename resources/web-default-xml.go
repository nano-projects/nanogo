package resources

import "github.com/nano-projects/nanogo/resources/license"

func WebDefaultXml() string {
	return `<?xml version="1.0" encoding="ISO-8859-1"?>
` + license.Xml() + `

<!-- ===================================================================== -->
<!-- This file contains the default descriptor for web applications.       -->
<!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -->
<!-- The intent of this descriptor is to include jetty specific or common  -->
<!-- configuration for all webapps.   If a context has a webdefault.xml    -->
<!-- descriptor, it is applied before the contexts own web.xml file        -->
<!--                                                                       -->
<!-- A context may be assigned a default descriptor by:                    -->
<!--  + Calling WebApplicationContext.setDefaultsDescriptor                -->
<!--  + Passed an arg to addWebApplications                                -->
<!--                                                                       -->
<!-- This file is used both as the resource within the jetty.jar (which is -->
<!-- used as the default if no explicit defaults descriptor is set) and it -->
<!-- is copied to the etc directory of the Jetty distro and explicitly     -->
<!-- by the jetty.xml file.                                                -->
<!--                                                                       -->
<!-- ===================================================================== -->
<web-app
   xmlns="http://java.sun.com/xml/ns/javaee"
   xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
   xsi:schemaLocation="http://java.sun.com/xml/ns/javaee http://java.sun.com/xml/ns/javaee/web-app_2_5.xsd"
   metadata-complete="true"
   version="2.5">

  <description>
    Default web.xml file.
    This file is applied to a Web application before it's own WEB_INF/web.xml file
  </description>


  <!-- ==================================================================== -->
  <!-- Context params to control Session Cookies                            -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <!-- UNCOMMENT TO ACTIVATE
  <context-param>
    <param-name>org.mortbay.jetty.servlet.SessionDomain</param-name>
    <param-value>127.0.0.1</param-value>
  </context-param>

  <context-param>
    <param-name>org.mortbay.jetty.servlet.SessionPath</param-name>
    <param-value>/</param-value>
  </context-param>

  <context-param>
    <param-name>org.mortbay.jetty.servlet.MaxAge</param-name>
    <param-value>-1</param-value>
  </context-param>
  -->

  <context-param>
    <param-name>org.mortbay.jetty.webapp.NoTLDJarPattern</param-name>
    <param-value>start.jar|ant-.*\.jar|dojo-.*\.jar|jetty-.*\.jar|jsp-api-.*\.jar|junit-.*\.jar|servlet-api-.*\.jar|dnsns\.jar|rt\.jar|jsse\.jar|tools\.jar|sunpkcs11\.jar|sunjce_provider\.jar|xerces.*\.jar</param-value>
  </context-param>



  <!-- ==================================================================== -->
  <!-- The default servlet.                                                 -->
  <!-- This servlet, normally mapped to /, provides the handling for static -->
  <!-- content, OPTIONS and TRACE methods for the context.                  -->
  <!-- The following initParameters are supported:                          -->
  <!--                                                                      -->
  <!--   acceptRanges     If true, range requests and responses are         -->
  <!--                    supported                                         -->
  <!--                                                                      -->
  <!--   dirAllowed       If true, directory listings are returned if no    -->
  <!--                    welcome file is found. Else 403 Forbidden.        -->
  <!--                                                                      -->
  <!--   welcomeServlets  If true, attempt to dispatch to welcome files     -->
  <!--                    that are servlets, if no matching static          -->
  <!--                    resources can be found.                           -->
  <!--                                                                      -->
  <!--   redirectWelcome  If true, redirect welcome file requests           -->
  <!--                    else use request dispatcher forwards              -->
  <!--                                                                      -->
  <!--   gzip             If set to true, then static content will be served-->
  <!--                    as gzip content encoded if a matching resource is -->
  <!--                    found ending with ".gz"                           -->
  <!--                                                                      -->
  <!--   resoureBase      Can be set to replace the context resource base   -->
  <!--                                                                      -->
  <!--   relativeResourceBase                                               -->
  <!--                    Set with a pathname relative to the base of the   -->
  <!--                    servlet context root. Useful for only serving     -->
  <!--                    static content from only specific subdirectories. -->
  <!--                                                                      -->
  <!--   useFileMappedBuffer                                                -->
  <!--                    If set to true (the default), a  memory mapped    -->
  <!--                    file buffer will be used to serve static content  -->
  <!--                    when using an NIO connector. Setting this value   -->
  <!--                    to false means that a direct buffer will be used  -->
  <!--                    instead. If you are having trouble with Windows   -->
  <!--                    file locking, set this to false.                  -->
  <!--                                                                      -->
  <!--  cacheControl      If set, all static content will have this value   -->
  <!--                    set as the cache-control header.                  -->
  <!--                                                                      -->
  <!--  maxCacheSize      Maximum size of the static resource cache         -->
  <!--                                                                      -->
  <!--  maxCachedFileSize Maximum size of any single file in the cache      -->
  <!--                                                                      -->
  <!--  maxCachedFiles    Maximum number of files in the cache              -->
  <!--                                                                      -->
  <!--  cacheType         "nio", "bio" or "both" to determine the type(s)   -->
  <!--                    of resource cache. A bio cached buffer may be used-->
  <!--                    by nio but is not as efficient as a nio buffer.   -->
  <!--                    An nio cached buffer may not be used by bio.      -->
  <!--                                                                      -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <servlet>
    <servlet-name>default</servlet-name>
    <servlet-class>org.eclipse.jetty.servlet.DefaultServlet</servlet-class>
    <init-param>
      <param-name>acceptRanges</param-name>
      <param-value>true</param-value>
    </init-param>
    <init-param>
      <param-name>dirAllowed</param-name>
      <param-value>true</param-value>
    </init-param>
    <init-param>
      <param-name>welcomeServlets</param-name>
      <param-value>false</param-value>
    </init-param>
    <init-param>
      <param-name>redirectWelcome</param-name>
      <param-value>false</param-value>
    </init-param>
    <init-param>
      <param-name>maxCacheSize</param-name>
      <param-value>256000000</param-value>
    </init-param>
    <init-param>
      <param-name>maxCachedFileSize</param-name>
      <param-value>10000000</param-value>
    </init-param>
    <init-param>
      <param-name>maxCachedFiles</param-name>
      <param-value>1000</param-value>
    </init-param>
    <init-param>
      <param-name>cacheType</param-name>
      <param-value>both</param-value>
    </init-param>
    <init-param>
      <param-name>gzip</param-name>
      <param-value>true</param-value>
    </init-param>
    <init-param>
      <param-name>useFileMappedBuffer</param-name>
      <param-value>false</param-value>
    </init-param>
    <!--
    <init-param>
      <param-name>cacheControl</param-name>
      <param-value>max-age=3600,public</param-value>
    </init-param>
    -->
    <load-on-startup>0</load-on-startup>
  </servlet>

  <servlet-mapping> <servlet-name>default</servlet-name> <url-pattern>/</url-pattern> </servlet-mapping>


  <!-- ==================================================================== -->
  <!-- JSP Servlet                                                          -->
  <!-- This is the jasper JSP servlet from the jakarta project              -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <!-- The JSP page compiler and execution servlet, which is the mechanism  -->
  <!-- used by Glassfish to support JSP pages.  Traditionally, this servlet -->
  <!-- is mapped to URL patterh "*.jsp".  This servlet supports the         -->
  <!-- following initialization parameters (default values are in square    -->
  <!-- brackets):                                                           -->
  <!--                                                                      -->
  <!--   checkInterval       If development is false and reloading is true, -->
  <!--                       background compiles are enabled. checkInterval -->
  <!--                       is the time in seconds between checks to see   -->
  <!--                       if a JSP page needs to be recompiled. [300]    -->
  <!--                                                                      -->
  <!--   compiler            Which compiler Ant should use to compile JSP   -->
  <!--                       pages.  See the Ant documenation for more      -->
  <!--                       information. [javac]                           -->
  <!--                                                                      -->
  <!--   classdebuginfo      Should the class file be compiled with         -->
  <!--                       debugging information?  [true]                 -->
  <!--                                                                      -->
  <!--   classpath           What class path should I use while compiling   -->
  <!--                       generated servlets?  [Created dynamically      -->
  <!--                       based on the current web application]          -->
  <!--                       Set to ? to make the container explicitly set  -->
  <!--                       this parameter.                                -->
  <!--                                                                      -->
  <!--   development         Is Jasper used in development mode (will check -->
  <!--                       for JSP modification on every access)?  [true] -->
  <!--                                                                      -->
  <!--   enablePooling       Determines whether tag handler pooling is      -->
  <!--                       enabled  [true]                                -->
  <!--                                                                      -->
  <!--   fork                Tell Ant to fork compiles of JSP pages so that -->
  <!--                       a separate JVM is used for JSP page compiles   -->
  <!--                       from the one Tomcat is running in. [true]      -->
  <!--                                                                      -->
  <!--   ieClassId           The class-id value to be sent to Internet      -->
  <!--                       Explorer when using <jsp:plugin> tags.         -->
  <!--                       [clsid:8AD9C840-044E-11D1-B3E9-00805F499D93]   -->
  <!--                                                                      -->
  <!--   javaEncoding        Java file encoding to use for generating java  -->
  <!--                       source files. [UTF-8]                          -->
  <!--                                                                      -->
  <!--   keepgenerated       Should we keep the generated Java source code  -->
  <!--                       for each page instead of deleting it? [true]   -->
  <!--                                                                      -->
  <!--   logVerbosityLevel   The level of detailed messages to be produced  -->
  <!--                       by this servlet.  Increasing levels cause the  -->
  <!--                       generation of more messages.  Valid values are -->
  <!--                       FATAL, ERROR, WARNING, INFORMATION, and DEBUG. -->
  <!--                       [WARNING]                                      -->
  <!--                                                                      -->
  <!--   mappedfile          Should we generate static content with one     -->
  <!--                       print statement per input line, to ease        -->
  <!--                       debugging?  [false]                            -->
  <!--                                                                      -->
  <!--                                                                      -->
  <!--   reloading           Should Jasper check for modified JSPs?  [true] -->
  <!--                                                                      -->
  <!--   suppressSmap        Should the generation of SMAP info for JSR45   -->
  <!--                       debugging be suppressed?  [false]              -->
  <!--                                                                      -->
  <!--   dumpSmap            Should the SMAP info for JSR45 debugging be    -->
  <!--                       dumped to a file? [false]                      -->
  <!--                       False if suppressSmap is true                  -->
  <!--                                                                      -->
  <!--   scratchdir          What scratch directory should we use when      -->
  <!--                       compiling JSP pages?  [default work directory  -->
  <!--                       for the current web application]               -->
  <!--                                                                      -->
  <!--   tagpoolMaxSize      The maximum tag handler pool size  [5]         -->
  <!--                                                                      -->
  <!--   xpoweredBy          Determines whether X-Powered-By response       -->
  <!--                       header is added by generated servlet  [false]  -->
  <!--                                                                      -->
  <!-- If you wish to use Jikes to compile JSP pages:                       -->
  <!--   Set the init parameter "compiler" to "jikes".  Define              -->
  <!--   the property "-Dbuild.compiler.emacs=true" when starting Jetty     -->
  <!--   to cause Jikes to emit error messages in a format compatible with  -->
  <!--   Jasper.                                                            -->
  <!--   If you get an error reporting that jikes can't use UTF-8 encoding, -->
  <!--   try setting the init parameter "javaEncoding" to "ISO-8859-1".     -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <servlet id="jsp">
    <servlet-name>jsp</servlet-name>
    <servlet-class>org.apache.jasper.servlet.JspServlet</servlet-class>
    <init-param>
        <param-name>logVerbosityLevel</param-name>
        <param-value>DEBUG</param-value>
    </init-param>
    <init-param>
        <param-name>fork</param-name>
        <param-value>false</param-value>
    </init-param>
    <init-param>
        <param-name>xpoweredBy</param-name>
        <param-value>false</param-value>
    </init-param>
    <!--
    <init-param>
        <param-name>classpath</param-name>
        <param-value>?</param-value>
    </init-param>
    -->
    <load-on-startup>0</load-on-startup>
  </servlet>

  <servlet-mapping>
    <servlet-name>jsp</servlet-name>
    <url-pattern>*.jsp</url-pattern>
    <url-pattern>*.jspf</url-pattern>
    <url-pattern>*.jspx</url-pattern>
    <url-pattern>*.xsp</url-pattern>
    <url-pattern>*.JSP</url-pattern>
    <url-pattern>*.JSPF</url-pattern>
    <url-pattern>*.JSPX</url-pattern>
    <url-pattern>*.XSP</url-pattern>
  </servlet-mapping>

  <!-- ==================================================================== -->
  <!-- Dynamic Servlet Invoker.                                             -->
  <!-- This servlet invokes anonymous servlets that have not been defined   -->
  <!-- in the web.xml or by other means. The first element of the pathInfo  -->
  <!-- of a request passed to the envoker is treated as a servlet name for  -->
  <!-- an existing servlet, or as a class name of a new servlet.            -->
  <!-- This servlet is normally mapped to /servlet/*                        -->
  <!-- This servlet support the following initParams:                       -->
  <!--                                                                      -->
  <!--  nonContextServlets       If false, the invoker can only load        -->
  <!--                           servlets from the contexts classloader.    -->
  <!--                           This is false by default and setting this  -->
  <!--                           to true may have security implications.    -->
  <!--                                                                      -->
  <!--  verbose                  If true, log dynamic loads                 -->
  <!--                                                                      -->
  <!--  *                        All other parameters are copied to the     -->
  <!--                           each dynamic servlet as init parameters    -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <!-- Uncomment for dynamic invocation
  <servlet>
    <servlet-name>invoker</servlet-name>
    <servlet-class>org.mortbay.jetty.servlet.Invoker</servlet-class>
    <init-param>
      <param-name>verbose</param-name>
      <param-value>false</param-value>
    </init-param>
    <init-param>
      <param-name>nonContextServlets</param-name>
      <param-value>false</param-value>
    </init-param>
    <init-param>
      <param-name>dynamicParam</param-name>
      <param-value>anyValue</param-value>
    </init-param>
    <load-on-startup>0</load-on-startup>
  </servlet>

  <servlet-mapping> <servlet-name>invoker</servlet-name> <url-pattern>/servlet/*</url-pattern> </servlet-mapping>
  -->



  <!-- ==================================================================== -->
  <session-config>
    <session-timeout>30</session-timeout>
  </session-config>

  <!-- ==================================================================== -->
  <!-- Default MIME mappings                                                -->
  <!-- The default MIME mappings are provided by the mime.properties        -->
  <!-- resource in the org.mortbay.jetty.jar file.  Additional or modified  -->
  <!-- mappings may be specified here                                       -->
  <!-- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  -->
  <!-- UNCOMMENT TO ACTIVATE
  <mime-mapping>
    <extension>mysuffix</extension>
    <mime-type>mymime/type</mime-type>
  </mime-mapping>
  -->

  <!-- ==================================================================== -->
  <welcome-file-list>
    <welcome-file>index.html</welcome-file>
    <welcome-file>index.htm</welcome-file>
    <welcome-file>index.jsp</welcome-file>
  </welcome-file-list>

  <!-- ==================================================================== -->
  <locale-encoding-mapping-list>
    <locale-encoding-mapping><locale>ar</locale><encoding>ISO-8859-6</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>be</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>bg</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ca</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>cs</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>da</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>de</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>el</locale><encoding>ISO-8859-7</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>en</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>es</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>et</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>fi</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>fr</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>hr</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>hu</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>is</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>it</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>iw</locale><encoding>ISO-8859-8</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ja</locale><encoding>Shift_JIS</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ko</locale><encoding>EUC-KR</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>lt</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>lv</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>mk</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>nl</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>no</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>pl</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>pt</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ro</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>ru</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sh</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sk</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sl</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sq</locale><encoding>ISO-8859-2</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sr</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>sv</locale><encoding>ISO-8859-1</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>tr</locale><encoding>ISO-8859-9</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>uk</locale><encoding>ISO-8859-5</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>zh</locale><encoding>GB2312</encoding></locale-encoding-mapping>
    <locale-encoding-mapping><locale>zh_TW</locale><encoding>Big5</encoding></locale-encoding-mapping>
  </locale-encoding-mapping-list>

  <security-constraint>
    <web-resource-collection>
      <web-resource-name>Disable TRACE</web-resource-name>
      <url-pattern>/</url-pattern>
      <http-method>TRACE</http-method>
    </web-resource-collection>
    <auth-constraint/>
  </security-constraint>

</web-app>

`

}