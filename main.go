package main

import (
	"log"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/models"
)

func main() {
	//path := "/Users/yanghe/Works/____Go_Project____/____Workspaces____"
	//if !strings.HasSuffix(path, "/") {
	//	path += "/"
	//}
	//
	//groupId := "cn.net.yto.test"
	//artifactId := "nanogo"
	//version := "0.0.1"
	//
	//

	arg := models.Argument{}
	arg.Parse()

	if !*arg.NewWebapp && !*arg.NewScheduler && !*arg.New {
		log.Fatalf("必须指定项目类型且只能指定一种项目创建类型, 请使用 -new, -new -web 或 -new -scheduler 创建项目")

	} else if *arg.New && *arg.NewWebapp && *arg.NewScheduler {
		log.Fatalf("必须指定项目类型且只能指定一种项目创建类型, 请使用 -new, -new -web 或 -new -scheduler 创建项目")

	} else if *arg.New && !*arg.NewWebapp && !*arg.NewScheduler {
		if *arg.Yaml == "" {
			if !arg.ExistYaml() {
				return
			}

			New(&arg)
		} else {
			log.Fatalf("必须指定Yaml配置文件的路径")
		}

	} else if *arg.New && *arg.NewWebapp && !*arg.NewScheduler {
		NewWebapp(&arg)

	} else if *arg.New && !*arg.NewWebapp && *arg.NewScheduler {
		NewScheduler(&arg)

	}
}

func New(_arg *models.Argument) {

}

func NewWebapp(arg *models.Argument) {
	if arg.Validation() {
		io.GeneralDefaultWebapp(arg)
	} else {
		log.Fatalf("必须指定GroupId和ArtifactId, 例如: -groupId org.nanoframework -artifactId test")
	}
}

func NewScheduler(arg *models.Argument) {
	if arg.Validation() {

	} else {
		log.Fatalf("必须指定GroupId和ArtifactId, 例如: -groupId org.nanoframework -artifactId test")
		return
	}
}



