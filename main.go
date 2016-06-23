/*
	Copyright 2015-2016 the original author or authors.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
 */
package main

import (
	"fmt"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/models"
)

const (
	VERSION = "0.0.1"
)

func main() {
	arg := models.Argument{}
	arg.Parse()

	if !*arg.NewWebapp && !*arg.NewScheduler && !*arg.New {
		fmt.Println("必须指定项目类型且只能指定一种项目构建类型, 请使用 -new, -new -web 或 -new -scheduler 创建项目")

	} else if *arg.New && *arg.NewWebapp && *arg.NewScheduler {
		fmt.Println("必须指定项目类型且只能指定一种项目构建类型, 请使用 -new, -new -web 或 -new -scheduler 创建项目")

	} else if *arg.New && !*arg.NewWebapp && !*arg.NewScheduler {
		if arg.ExistYaml() {
			New(&arg)
		} else {
			if *arg.Yaml == "" {
				fmt.Println("当前路径下不存在nanogo.yml, 请指定Yaml配置文件的路径或在当前路径下创建文件nanogo.yml")
			} else {
				fmt.Println("不存在配置文件: " + *arg.Yaml + ", 请指定正确的Yaml配置文件的路径或在当前路径下创建yaml配置文件")
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
		fmt.Println("必须指定Repository, 例如: -resp org.nanoframework:test:0.0.1-SNAPSHOT")
	}
}

func NewScheduler(arg *models.Argument) {
	if arg.Validation() {
		io.GeneralDefaultScheduler(arg)
	} else {
		fmt.Println("必须指定Repository, 例如: -resp org.nanoframework:test:0.0.1-SNAPSHOT")
		return
	}
}
