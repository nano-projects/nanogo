package template

import (
	"github.com/nano-projects/nanogo/initial/template/license"
	"text/template"
)

func Scheduler() (*template.Template, error) {
	sche := license.Class() + `
package {{.Package}}.scheduler;

import org.nanoframework.concurrent.scheduler.BaseScheduler;
import org.nanoframework.concurrent.scheduler.Scheduler;

import {{.Package}}.service.{{.SourceName}}Service;

import com.google.inject.Inject;

/**
 *
 * @author {{.Author}}
 * @since {{.Version}}
 */
@Scheduler(parallel = 1)
public class {{.SourceName}}Scheduler extends BaseScheduler {

    @Inject
    private {{.SourceName}}Service {{.AliasSourceName}}Service;

    @Override
    public void before() {

    }

    @Override
    public void execute() {

    }

    @Override
    public void after() {

    }

    @Override
    public void destroy() {

    }

}
`

	return template.New("Scheduler").Parse(sche)
}
