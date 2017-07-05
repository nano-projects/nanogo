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

func CheckstyleRules() (*template.Template, error) {
	rules := `<?xml version="1.0" encoding="UTF-8"?>
` + license.Xml() + `
<!DOCTYPE module PUBLIC "-//Puppy Crawl//DTD Check Configuration 1.3//EN" "http://www.puppycrawl.com/dtds/configuration_1_3.dtd">
<module name="Checker">
    <property name="severity" value="warning"/>
    <property name="charset" value="UTF-8"/>

    <!-- TreeWalker module checks -->
    <module name="TreeWalker">
        <property name="tabWidth" value="4"/>
        <module name="JavadocMethod">
            <property name="scope" value="protected"/>
            <property name="severity" value="error"/>
            <property name="allowUndeclaredRTE" value="true"/>
            <property name="allowMissingPropertyJavadoc" value="true"/>
            <property name="allowThrowsTagsForSubclasses" value="true"/>
        </module>
        <module name="JavadocType">
            <property name="severity" value="error"/>
            <property name="scope" value="public"/>
            <property name="authorFormat" value="\w+"/>
            <property name="allowMissingParamTags" value="true"/>
        </module>
        <module name="JavadocVariable">
            <property name="severity" value="error"/>
            <property name="scope" value="protected"/>
            <property name="ignoreNamePattern" value="log|logger|LOG|LOGGER"/>
        </module>
        <module name="JavadocStyle">
            <property name="severity" value="error"/>
        </module>
        <module name="ClassFanOutComplexity">
            <property name="max" value="50"/>
            <property name="severity" value="error"/>
        </module>
        <!-- <module name="CommentsIndentation">
            <property name="severity" value="error"/>
        </module> -->
        <!-- <module name="Indentation">
            <property name="severity" value="error"/>
        </module> -->
        <module name="CyclomaticComplexity">
            <property name="max" value="15"/>
            <property name="severity" value="error"/>
        </module>
        <module name="DefaultComesLast">
            <property name="severity" value="error"/>
        </module>

        <module name="AnnotationLocation">
            <property name="severity" value="error"/>
            <property name="allowSamelineSingleParameterlessAnnotation" value="false"/>
        </module>
        <module name="ConstantName">
            <property name="severity" value="error"/>
        </module>
        <module name="GenericWhitespaceCheck">
            <property name="severity" value="error"/>
        </module>
        <module name="ModifiedControlVariable">
            <property name="severity" value="error"/>
        </module>
        <module name="MagicNumber">
            <property name="ignoreAnnotation" value="true"/>
            <property name="ignoreHashCodeMethod" value="true"/>
            <property name="ignoreFieldDeclaration" value="true"/>
            <property name="severity" value="error"/>
        </module>
        <module name="LocalFinalVariableName">
            <property name="severity" value="error"/>
        </module>
        <module name="LocalVariableName">
            <property name="severity" value="error"/>
        </module>
        <module name="AbstractClassName">
            <property name="ignoreModifier" value="true"/>
            <property name="severity" value="error"/>
            <property name="format" value="^Abstract.*$|^.*Factory$|^Base.*$|^Root.*$"/>
        </module>
        <module name="MemberName">
            <property name="severity" value="error"/>
        </module>
        <module name="MethodName">
            <property name="severity" value="error"/>
        </module>
        <module name="GenericWhitespace">
            <property name="severity" value="error"/>
        </module>
        <module name="PackageName">
            <property name="severity" value="error"/>
        </module>
        <module name="ParameterName">
            <property name="severity" value="error"/>
            <!-- <property name="ignoreOverridden" value="true" /> -->
        </module>
        <module name="StaticVariableName">
            <property name="severity" value="error"/>
            <property name="format" value="(^[A-Z][A-Z0-9]*(_[A-Z0-9]+)*$)"/>
        </module>
        <module name="AbbreviationAsWordInName"/>
        <module name="TypeName">
            <property name="severity" value="error"/>
        </module>
        <module name="AvoidStarImport">
            <property name="severity" value="error"/>
            <property name="excludes"
                      value="java.io,java.net,java.lang.Math,org.junit.Assert,org.mockito.Mockito,
                      org.mockito.Matchers,org.springframework.test.web.servlet.request.MockMvcRequestBuilders,
                      org.springframework.test.web.servlet.result.MockMvcResultMatchers,
                      java.nio.file.StandardWatchEventKinds"/>
            <property name="allowStaticMemberImports" value="false"/>
        </module>
        <module name="SingleLineJavadoc">
            <property name="severity" value="warning"/>
        </module>
        <module name="IllegalImport">
            <property name="severity" value="error"/>
        </module>
        <module name="RedundantImport">
            <property name="severity" value="error"/>
        </module>
        <module name="UnusedImports">
            <property name="severity" value="error"/>
        </module>
        <module name="SuperClone">
            <property name="severity" value="info"/>
        </module>
        <module name="SuperFinalize">
            <property name="severity" value="error"/>
        </module>
        <module name="MethodLength"/>
        <module name="ParameterNumber">
            <property name="max" value="10"/>
            <property name="severity" value="error"/>
            <property name="tokens" value="METHOD_DEF"/>
        </module>
        <module name="LineLength">
            <property name="max" value="200"/>
            <property name="tabWidth" value="4"/>
            <property name="severity" value="error"/>
        </module>
        <module name="EmptyForIteratorPad"/>
        <module name="MethodParamPad">
            <property name="severity" value="error"/>
        </module>
        <module name="NoWhitespaceAfter">
            <property name="severity" value="error"/>
        </module>
        <module name="NoWhitespaceBefore">
            <property name="severity" value="error"/>
        </module>
        <module name="OperatorWrap">
            <property name="severity" value="error"/>
        </module>
        <module name="ParenPad"/>
        <module name="TypecastParenPad"/>
        <module name="WhitespaceAfter">
            <property name="severity" value="error"/>
        </module>
        <module name="ModifierOrder">
            <property name="severity" value="error"/>
        </module>
        <module name="RedundantModifier">
            <property name="severity" value="error"/>
        </module>
        <module name="AvoidNestedBlocks"/>
        <module name="EmptyBlock">
            <property name="severity" value="error"/>
        </module>
        <module name="FallThrough">
            <property name="severity" value="error"/>
        </module>
        <module name="DeclarationOrder">
            <property name="severity" value="error"/>
        </module>
        <module name="CovariantEquals">
            <property name="severity" value="error"/>
        </module>
        <module name="ExplicitInitialization">
            <property name="severity" value="error"/>
        </module>
        <module name="LeftCurly">
            <property name="severity" value="error"/>
        </module>
        <module name="NeedBraces">
            <property name="severity" value="error"/>
        </module>
        <module name="EqualsAvoidNull">
            <property name="severity" value="error"/>
        </module>
        <module name="RightCurly">
            <property name="severity" value="error"/>
        </module>
        <module name="NoFinalizer">
            <property name="severity" value="error"/>
        </module>
        <module name="EmptyStatement"/>
        <module name="EqualsHashCode"/>
        <module name="IllegalInstantiation"/>
        <module name="InnerAssignment"/>
        <module name="MissingSwitchDefault">
            <property name="severity" value="info"/>
        </module>
        <module name="SimplifyBooleanExpression">
            <property name="severity" value="error"/>
        </module>
        <module name="SimplifyBooleanReturn">
            <property name="severity" value="error"/>
        </module>
        <module name="FinalClass">
            <property name="severity" value="error"/>
        </module>
        <module name="HideUtilityClassConstructor">
            <property name="severity" value="error"/>
        </module>
        <module name="InterfaceIsType">
            <property name="severity" value="info"/>
        </module>
        <module name="VisibilityModifier">
            <property name="severity" value="error"/>
            <property name="protectedAllowed" value="true"/>
        </module>
        <module name="AtclauseOrder">
            <property name="severity" value="error"/>
        </module>
        <module name="BooleanExpressionComplexity">
            <property name="severity" value="error"/>
            <property name="max" value="4"/>
        </module>
        <module name="ArrayTypeStyle">
            <property name="severity" value="error"/>
        </module>
        <module name="FinalParameters">
            <property name="severity" value="error"/>
            <property name="tokens" value="METHOD_DEF, CTOR_DEF, LITERAL_CATCH"/>
        </module>
        <module name="FinalLocalVariable">
            <property name="severity" value="ignore"/>
            <property name="tokens" value="PARAMETER_DEF, VARIABLE_DEF"/>
            <property name="validateEnhancedForLoopVariable" value="true"/>
        </module>
        <module name="TodoComment">
            <property name="severity" value="error"/>
        </module>
        <module name="MutableException">
            <property name="severity" value="error"/>
        </module>
        <module name="UpperEll"/>
        <module name="MissingOverride">
            <property name="severity" value="error"/>
        </module>
        <module name="MissingDeprecated">
            <property name="severity" value="error"/>
        </module>

        <!-- Custom checks based on regular expressions -->
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Parameteric log messages"/>
            <property name="severity" value="error"/>
            <property name="format" value="log\.\w+\(((\&quot;.+\&quot;\s*\+)|(.*\s*\+\s*\&quot;))"/>
            <property name="message"
                      value="Avoid string concatenation for constructing log messages. Use parametric messages instead"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="JavaDoc @version tag"/>
            <property name="severity" value="error"/>
            <property name="format" value="@version\s+(.+)*(\$Revision|\$Date)"/>
            <property name="message" value="Invalid JavaDoc @version tag."/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Logging framework selection"/>
            <property name="severity" value="error"/>
            <property name="format" value="LogFactory\.getLog"/>
            <property name="message" value="CAS uses the slf4j logging framework."/>
        </module>
        <module name="RegexpSinglelineJava">
            <property name="id" value="sysOutConsoleLogs"/>
            <metadata name="net.sf.eclipsecs.core.comment" value="Console output messages"/>
            <property name="severity" value="error"/>
            <property name="format" value="System\.(out|err)"/>
            <property name="message"
                      value="Avoid sending messages to the console directly. Use a logger object instead"/>
        </module>
        <module name="RegexpSinglelineJava">
            <property name="id" value="stackTraceConsoleLogs"/>
            <metadata name="net.sf.eclipsecs.core.comment" value="Printing stack traces to the console"/>
            <property name="severity" value="error"/>
            <property name="format" value="\.printStackTrace\(\)"/>
            <property name="message"
                      value="Avoid sending stack traces to the console directly. Use a logger object instead"/>
        </module>
        <module name="RegexpSinglelineJava">
            <property name="id" value="junitTestMethodName"/>
            <metadata name="net.sf.eclipsecs.core.comment" value="Using 'test' prefix for JUnit Tests"/>
            <property name="severity" value="error"/>
            <property name="format" value="(public|protected)\s+void\s+test\w+\(.+\{$"/>
            <property name="message"
                      value="JUnit test methods should not begin with the 'test' prefix. Use annotations instead and/or rename the method"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="The logger object must be named &quot;logger&quot;"/>
            <property name="severity" value="error"/>
            <property name="format" value="\s+(static\s)*(final\s)*(static\s)*Logger\s+(log|LOG)\b"/>
            <property name="message"
                      value="The Logger object must only be called &quot;logger&quot; or &quot;LOGGER&quot;"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Qualifying logger with the &quot;this&quot; keyword"/>
            <property name="severity" value="error"/>
            <property name="format" value="((this\.logger)|(super\.logger))\.\w+\("/>
            <property name="message" value="The Logger object need not be qualified with the &quot;this&quot; keyword"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Using @Test instead of extending TestCase"/>
            <property name="severity" value="error"/>
            <property name="format" value="class\s+\w+\s+extends\s+(junit\.framework\.)*TestCase"/>
            <property name="message"
                      value="All testcase must use annotations (@Test) instead of extending junit.framework.TestCase"/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Using the junit.framework package"/>
            <property name="severity" value="error"/>
            <property name="format" value="junit.framework"/>
            <property name="message" value="The package junit.framework belongs to JUnit v3. Use org.junit instead."/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Checking for logging level"/>
            <property name="severity" value="warning"/>
            <property name="format" value="log\.is\w+Enabled\("/>
            <property name="message"
                      value="If the construction of the log message is cheap, consider not checking for logging levels."/>
        </module>
        <module name="RegexpSinglelineJava">
            <metadata name="net.sf.eclipsecs.core.comment" value="Non-static inner class"/>
            <property name="severity" value="error"/>
            <property name="format" value="\s+(private|public|protected)*\s+(abstract\s)*class\s+\w+"/>
            <property name="message"
                      value="Non-static nested classes are a security compromise. Consider using a static class instead"/>
        </module>
    </module>

    <!-- Checker module checks -->
    <module name="UniqueProperties">
        <property name="severity" value="error"/>
    </module>
    <module name="NewlineAtEndOfFile">
        <property name="fileExtensions" value="java, xml, properties, txt"/>
        <property name="lineSeparator" value="lf"/>
        <property name="severity" value="error"/>
    </module>
    <module name="Translation">
        <property name="severity" value="ignore"/>
    </module>
    <module name="FileLength"/>
    <module name="FileTabCharacter">
        <property name="severity" value="error"/>
    </module>

    <!-- Custom checks based on regular expressions -->
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Trailing spaces"/>
        <property name="severity" value="error"/>
        <property name="format" value="\s+$"/>
        <property name="message" value="Line has trailing spaces."/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Space instead of tabs"/>
        <property name="severity" value="error"/>
        <property name="format" value="^\t+"/>
        <property name="message" value="Tabs should never be used for indentation. Use spaces instead"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Space after cast"/>
        <property name="severity" value="error"/>
        <property name="format" value="\(\w+\)\w+"/>
        <property name="message" value="There are no spaces after cast."/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Usage of java.util.Random"/>
        <property name="severity" value="error"/>
        <property name="format" value="(java.util.Random)|(new Random\()"/>
        <property name="message" value="For security purposes, use 'java.security.SecureRandom' instead"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="final clone() method"/>
        <property name="severity" value="error"/>
        <property name="format" value="public\s+\w+\s+clone\(\)"/>
        <property name="message"
                  value="Consider marking the clone() method as final to reduce chances of data corruption"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="JUnit Assert static import"/>
        <property name="severity" value="error"/>
        <property name="format" value="import\s+static\s+org\.junit\.Assert\.\w+"/>
        <property name="message" value="JUnit Assert methods MUST be imported statically with a *"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Mockito static import"/>
        <property name="severity" value="error"/>
        <property name="format" value="import\s+static\s+org\.mockito\.Mockito\.\w+"/>
        <property name="message" value="Mockito methods MUST be imported statically with a *"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Missing @since tag"/>
        <property name="severity" value="error"/>
        <property name="format" value="\s+\*+\s+@since\s+\d+\.\d+.*"/>
        <property name="maximum" value="200"/>
        <property name="minimum" value="1"/>
        <property name="fileExtensions" value="java,groovy"/>
        <property name="message" value="There are no @since tags defined for this component's Javadocs."/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Before space"/>
        <property name="severity" value="error"/>
        <property name="format" value="\w+\{|\)\{"/>
        <property name="message" value="'{' 前应有空格"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="Before multi space"/>
        <property name="severity" value="error"/>
        <property name="format" value="\w+\s{2,}\{|\)\s{2,}\{"/>
        <property name="message" value="'{' 前仅能有一个空格"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="if after multi space"/>
        <property name="severity" value="error"/>
        <property name="format" value="if\s{2,}\("/>
        <property name="message" value="'if' 后仅能有一个空格"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="For after space"/>
        <property name="severity" value="error"/>
        <property name="format" value="for\("/>
        <property name="message" value="'for' 后应有空格"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="For after multi space"/>
        <property name="severity" value="error"/>
        <property name="format" value="for\s{2,}\("/>
        <property name="message" value="'for' 后仅能有一个空格"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="While after space"/>
        <property name="severity" value="error"/>
        <property name="format" value="while\("/>
        <property name="message" value="'while' 后应有空格"/>
    </module>
    <module name="RegexpSingleline">
        <metadata name="net.sf.eclipsecs.core.comment" value="While after multi space"/>
        <property name="severity" value="error"/>
        <property name="format" value="while\s{2,}\("/>
        <property name="message" value="'while' 后仅能有一个空格"/>
    </module>
    <module name="SuppressionFilter">
        <property name="file" value="${checkstyle.suppressions.file}"/>
    </module>
</module>
`

	return template.New("CheckstyleRules").Parse(rules)
}
