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
	"github.com/spf13/cobra"
	"mytool/internal/code"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "创建固定的目录结构",
	Long: "创建固定的目录结构,用于存放代码和配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		prepare()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func prepare(){
	json:= `{
  "template": "",
  "output": "",
  "worker_num": 10,
  "separator": "",
  "fields": {
  }
}`
	e:=code.CreateCurrentFolder("template")
	cobra.CheckErr(e)
	e=code.CreateCurrentFolder("output")
	cobra.CheckErr(e)
	e=code.CreateCurrentFile("settings.json",[]byte(json))
	cobra.CheckErr(e)
}
