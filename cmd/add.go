// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"github.com/nano-projects/nanogo/addition"
	"github.com/nano-projects/nanogo/addition/conf"
	"github.com/nano-projects/nanogo/io"
	"github.com/nano-projects/nanogo/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Source file of flags",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}

		author, err := cmd.Flags().GetString("author")
		if err != nil {
			return err
		}

		if author == "" {
			return errors.New("Author cannot be empty")
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		} else if name == "" {
			return errors.New("Source name cannot be empty")
		}

		if path == "" {
			path = io.Pwd()
		}

		log.Logger.Debugf("Project path: %v", path)
		addConf, err := conf.Make(path, author, name)
		if err != nil {
			return err
		}

		return (&addition.Addition{AddConf: *addConf}).Run()
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().String("path", io.Pwd(), "The project root path")
	addCmd.Flags().StringP("author", "a", "", "Creation file author")
	addCmd.Flags().StringP("name", "n", "", "The interface or class prefix name")

}
