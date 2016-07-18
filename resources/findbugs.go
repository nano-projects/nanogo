package resources

import "github.com/nano-projects/nanogo/resources/license"

func Findbugs() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
` + license.Xml() + `
<!-- See http://findbugs.sourceforge.net/manual/filter.html -->
<FindBugsFilter>
    <Match>
        <Confidence value="2" />
        <Rank value="15" />
        <Bug category="SECURITY,PERFORMANCE,MALICIOUS_CODE" />
    </Match>
</FindBugsFilter>
`

}
