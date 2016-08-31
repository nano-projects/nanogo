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

package cmd

import (
	"github.com/nano-projects/nanogo/log"
	"github.com/spf13/cobra"
	"github.com/Sirupsen/logrus"
	"os"
)

var (
	RootCmd = &cobra.Command{
		Use:          "nanogo",
		Short:        "Build a maven project",
		Long: `NanoGo is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Maven application.`,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			logrus.SetOutput(os.Stderr)
			flag, err := cmd.Flags().GetString("log-level")
			if err != nil {
				log.Logger.Fatal(err)
			}

			level, err := logrus.ParseLevel(flag)
			logrus.SetLevel(level)
		},
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Logger.Fatal(err)
	}
}

func init() {
	RootCmd.PersistentFlags().StringP("log-level", "l", "info", "Log level (options \"debug\", \"info\", \"warn\", \"error\", \"fatal\", \"panic\")")

}
