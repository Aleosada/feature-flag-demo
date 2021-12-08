/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"github.com/aleosada/feature-flag-demo/pkg"
	"github.com/spf13/cobra"
)

// unleashCmd represents the unleash command
var unleashCmd = &cobra.Command{
	Use:   "unleash",
	Short: "Feature flag demo for unleash",
	Run: func(cmd *cobra.Command, args []string) {
        pkg.InitConfig("unleash")
        pkg.InitUnleash()
        pkg.Run()
	},
}

func init() {
	rootCmd.AddCommand(unleashCmd)
}
