package resources

import "github.com/nano-projects/nanogo/resources/license"

func Assembly() string {
	return license.Xml() + `
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
}
