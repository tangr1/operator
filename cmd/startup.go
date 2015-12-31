// Copyright © 2015 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"
	"strconv"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/tangr1/hicto/client"
	"github.com/tangr1/hicto/client/startup"
	"github.com/tangr1/hicto/models"
)

var startupClient = *operations.Default.Startup

func ListStartup() {
	result, err := startupClient.GetStartups(nil, writer)
	if err != nil {
		fmt.Printf("获取创业公司列表失败: %v\n", err)
		os.Exit(1)
	}
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("ID", "名称", "工商注册名称", "城市", "领域", "审核状态")
	for _, startup := range result.Payload.Content {
		table.AddRow(startup.ID, startup.Name, startup.RegistrationName, GetCode("city", startup.City),
			GetCode("domain", startup.Domain), GetCode("reviewStatus", startup.ReviewStatus))
	}
	fmt.Println(table)
}

func ShowStartup(startup *models.Startup) {
	table := uitable.New()
	table.MaxColWidth = 120
	table.Wrap = true // wrap columns
	table.AddRow("ID", startup.ID)
	table.AddRow("名称", startup.Name)
	table.AddRow("注册名称", startup.RegistrationName)
	table.AddRow("城市", GetCode("city", startup.City))
	table.AddRow("领域", GetCode("domain", startup.Domain))
	table.AddRow("主页", startup.Homepage)
	table.AddRow("描述", startup.Description)
	table.AddRow("审核状态", GetCode("reviewStatus", startup.ReviewStatus))
	fmt.Println(table)
}

func GetStartup(id int64) {
	result, err := startupClient.GetStartupsID(&startup.GetStartupsIDParams{ID: id}, writer)
	if err != nil {
		fmt.Printf("获取创业公司失败: %v\n", err)
		os.Exit(1)
	}
	ShowStartup(result.Payload)
}

func UpdateStartup(id int64, reviewStatus int32, advice string) {
	body := &models.UpdateStartupRequest{ReviewStatus: reviewStatus, ModifyReason: advice}
	result, err := startupClient.PutStartupsID(&startup.PutStartupsIDParams{ID: id, Body: body}, writer)
	if err != nil {
		fmt.Printf("更新创业公司失败: %v\n", err)
		os.Exit(1)
	}
	ShowStartup(result.Payload)
}

// startupCmd represents the startup command
var startupCmd = &cobra.Command{
	Use:   "startup [STARTUP-ID] [approve|sendback] [ADVICE]",
	Short: "创业公司相关命令",
	Long:  `创业公司相关命令，可以查看审核创业公司列表并进行审核`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			ListStartup()
		} else if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			GetStartup(id)
		} else {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			switch args[1] {
			case "approve":
				UpdateStartup(id, 2, "")
				break
			case "sendback":
				if len(args) < 3 {
					fmt.Println("请给出修改意见")
					os.Exit(1)
				}
				UpdateStartup(id, 3, args[2])
				break
			default:
				fmt.Printf("不允许的动作: %s\n", args[1])
				os.Exit(1)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(startupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
