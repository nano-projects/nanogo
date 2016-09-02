package template

import (
	"github.com/nano-projects/nanogo/initial/template/license"
	"text/template"
)

func Component() (*template.Template, error) {
	component := license.Class() + `
package {{.Package}}.component;

import org.nanoframework.core.component.stereotype.Component;
import {{.Package}}.component.impl.{{.SourceName}}ComponentImpl;

import com.google.inject.ImplementedBy;

/**
 *
 * @author {{.Author}}
 * @since {{.Version}}
 */
@Component
@ImplementedBy({{.SourceName}}ComponentImpl.class)
public interface {{.SourceName}}Component {

}
`

	return template.New("Component").Parse(component)
}

func ComponentImpl() (*template.Template, error) {
	impl := license.Class() + `
package {{.Package}}.component.impl;

import {{.Package}}.component.{{.SourceName}}Component;
import {{.Package}}.service.{{.SourceName}}Service;

import com.google.inject.Inject;

/**
 *
 * @author {{.Author}}
 * @since {{.Version}}
 */
public class {{.SourceName}}ComponentImpl implements {{.SourceName}}Component {

    @Inject
    private {{.SourceName}}Service {{.AliasSourceName}}Service;

}
`

	return template.New("ComponentImpl").Parse(impl)
}
