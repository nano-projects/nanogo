package resources

import "github.com/nano-projects/nanogo/resources/license"

func IndexJsp() string {
	return license.Jsp() + `
<html>
<body>
<h2>Hello World!</h2>
</body>
</html>
`

}