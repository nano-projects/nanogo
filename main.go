package main

import (
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/models"
	"fmt"
)

func main() {
	arg := models.Argument{}
	arg.Parse()

	if !*arg.NewWebapp && !*arg.NewScheduler && !*arg.New {
		fmt.Println("必须指定项目类型且只能指定一种项目创建类型, 请使用 -new, -new -web 或 -new -scheduler 创建项目")

	} else if *arg.New && *arg.NewWebapp && *arg.NewScheduler {
		fmt.Println("必须指定项目类型且只能指定一种项目创建类型, 请使用 -new, -new -web 或 -new -scheduler 创建项目")

	} else if *arg.New && !*arg.NewWebapp && !*arg.NewScheduler {
		if arg.ExistYaml(arg.Yaml) {
			New(&arg)
		} else {
			if *arg.Yaml == "" {
				fmt.Println("当前路径下不存在nanogo.yml, 请指定Yaml配置文件的路径或在当前路径下创建文件nanogo.yml")
			} else {
				fmt.Println("不存在" + *arg.Yaml + ", 请指定正确的Yaml配置文件的路径或在当前路径下创建yaml配置文件")
			}
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
		fmt.Println("必须指定GroupId和ArtifactId, 例如: -groupId org.nanoframework -artifactId test")
	}
}

func NewScheduler(arg *models.Argument) {
	if arg.Validation() {

	} else {
		fmt.Println("必须指定GroupId和ArtifactId, 例如: -groupId org.nanoframework -artifactId test")
		return
	}
}



