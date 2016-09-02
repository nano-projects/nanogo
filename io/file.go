// Copyright © 2015-2016 River Yang <comicme_yanghe@nanoframework.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package io

import (
	"bufio"
	"github.com/nano-projects/nanogo/log"
	"os"
	"text/template"
)

const (
	FILE_MODE = 0755
)

func Pwd() (path string) {
	path, err := os.Getwd()
	if err != nil {
		log.Logger.Fatal(err)
	}

	return
}

func IsDirExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return file.IsDir()
	}

}

func IsFileExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return !file.IsDir()
	}
}

func WriteFile(fileName, data string) error {
	log.Logger.Infof("create file: %v", fileName)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, FILE_MODE)
	if err != nil {
		return err
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(data); err != nil {
		return err
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

func WriteTemplate(path string, tmp *template.Template, conf interface{}) error {
	log.Logger.Infof("create file: %v", path)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, FILE_MODE)
	if err != nil {
		return err
	}

	defer file.Close()
	tmp.Execute(file, conf)
	return nil
}
