package template

import (
	"github.com/nano-projects/nanogo/initial/template/license"
	"text/template"
)

func Service() (*template.Template, error) {
	service := license.Class() + `
package {{.Package}}.service;

import {{.Package}}.service.impl.{{.SourceName}}ServiceImpl;

import com.google.inject.ImplementedBy;

/**
 *
 * @author {{.Author}}
 * @since {{.Version}}
 */
@ImplementedBy({{.SourceName}}ServiceImpl.class)
public interface {{.SourceName}}Service {

}
`

	return template.New("Service").Parse(service)
}

func ServiceImpl() (*template.Template, error) {
	impl := license.Class() + `
package {{.Package}}.service.impl;

import {{.Package}}.service.{{.SourceName}}Service;

/**
 *
 * @author {{.Author}}
 * @since {{.Version}}
 */
public class {{.SourceName}}ServiceImpl implements {{.SourceName}}Service {

}
`

	return template.New("ServiceImpl").Parse(impl)
}
