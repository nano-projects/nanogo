package resources

func Definitions() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<additionalHeaders>
  <JAVADOC_STYLE>
    <firstLine>/*</firstLine>
    <beforeEachLine> * </beforeEachLine>
    <endLine> */</endLine>
    <firstLineDetectionPattern>( |\t)*/\*( |\t)*$</firstLineDetectionPattern>
    <lastLineDetectionPattern>( |\t)*\*/( |\t)*$</lastLineDetectionPattern>
    <allowBlankLines>true</allowBlankLines>
    <isMultiline>true</isMultiline>
  </JAVADOC_STYLE>
</additionalHeaders>
`

}