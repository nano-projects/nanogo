package resources

import "github.com/nano-projects/nanogo/resources/license"

func CheckstyleSuppressions() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
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

}