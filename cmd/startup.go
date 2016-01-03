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
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tangr1/hicto/client"
	"github.com/tangr1/hicto/client/startup"
	"github.com/tangr1/hicto/models"
	"strings"
	"github.com/fatih/color"
)

var startupClient = *operations.Default.Startup
var reviewStatus string
var advice string

type Startup struct {
	ID              string `chinese:"ID" list:"true"`
	Certificate     string `chinese:"证书"`
	City            string `chinese:"城市" code:"city" list:"true"`
	ConsumeCtoCoins string `chinese:"消耗的悬赏分"`
	CoverImage      string `chinese:"封面图片"`
	CoverTitle      string `chinese:"封面标题"`
	CtoCoins        string `chinese:"可用的悬赏分"`
	Description     string `chinese:"介绍"`
	Domain          string `chinese:"领域" code:"domain" list:"true"`
	Founders        string `chinese:"创始人"`
	HelpedExperts   string `chinese:"帮助该企业过的专家"`
	Homepage        string `chinese:"主页" list:"true"`
	InvestmentStatus string `chinese:"融资状态" code:"investmentStatus" list:"true"`
	Investors string `chinese:"投资人"`
	Logo string `chinese:"Logo"`
	Name string `chinese:"名称" list:"true"`
	Products string `chinese:"产品"`
	RegistrationName string `chinese:"注册名称" list:"true"`
	ResolvedTopicCount string `chinese:"被解答的问题数"`
	ReviewStatus string `chinese:"审核状态" code:"reviewStatus" list:"true"`
	RunningStatus string `chinese:"运营状态" code:"runningStatus" list:"true"`
	StaffNumber string `chinese:"员工人数" code:"staffNumber"`
	StartYear string `chinese:"成立年份"`
	StartMonth string `chinese:"成立月份"`
	TopicCount string `chinese:"提出的问题总数"`
	ViewCount string `chinese:"被浏览次数"`
}

func ListStartup() {
	result, err := startupClient.GetStartups(nil, writer)
	if err != nil {
		color.Red("获取创业公司列表失败: %v\n", err)
		return
	}
	startups := make([]interface{}, len(result.Payload.Content))
	for i, from := range result.Payload.Content {
		var to Startup
		ModelToRecord(from, &to)
		startups[i] = to
	}
	ListRecords(startups)
}

func GetStartup(id int64) {
	result, err := startupClient.GetStartupsID(&startup.GetStartupsIDParams{ID: id}, writer)
	if err == nil {
		var startup Startup
		ModelToRecord(result.Payload, &startup)
		ShowRecord(startup)
	} else {
		color.Red("获取创业公司失败: %v\n", err)
	}
}

func UpdateStartup(id int64, reviewStatus int32, advice string) {
	body := &models.UpdateStartupRequest{ReviewStatus: reviewStatus, ModifyReason: advice}
	result, err := startupClient.PutStartupsID(&startup.PutStartupsIDParams{ID: id, Body: body}, writer)
	if err == nil {
		var startup Startup
		ModelToRecord(result.Payload, &startup)
		ShowRecord(startup)
	} else {
		color.Red("更新创业公司失败: %v\n", err)
	}
}

// startupCmd represents the startup command
var startupCmd = &cobra.Command{
	Use:   "startup",
	Short: "创业公司相关命令",
	Long:  "创业公司相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var startupListCmd = &cobra.Command{
	Use:   "list",
	Short: "查看创业公司列表",
	Long: "查看创业公司列表",
	Run: func(cmd *cobra.Command, args []string) {
		ListStartup()
	},
}

var startupGetCmd = &cobra.Command{
	Use:   "get ID",
	Short: "查看创业公司",
	Long: "查看创业公司",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			GetStartup(id)
		} else {
			cmd.Usage()
		}
	},
}

var startupUpdateCmd = &cobra.Command{
	Use:   "update ID",
	Short: "更新创业公司",
	Long: "查看创业公司",
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) == 1) {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			if strings.Compare(reviewStatus, "approve") == 0 {
				UpdateStartup(id, 2, advice)
			} else if strings.Compare(reviewStatus, "sendback") == 0 {
				UpdateStartup(id, 3, advice)
			} else {
				color.Red("请指定有效的审核状态: -r approve 或 -r sendback")
			}
		} else {
			cmd.Usage()
		}
	},
}

var startupApproveCmd = &cobra.Command{
	Use:   "approve ID",
	Short: "通过创业公司申请",
	Long: "通过创业公司申请",
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) == 1) {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			UpdateStartup(id, 2, "")
		} else {
			cmd.Usage()
		}
	},
}

var startupSendbackCmd = &cobra.Command{
	Use:   "sendback ID [ADVICE]",
	Short: "打回创业公司申请, 并提出修改意见",
	Long: "打回创业公司申请, 并提出修改意见",
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) == 1) {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			UpdateStartup(id, 3, "")
		} else if (len(args) == 2) {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			UpdateStartup(id, 3, args[1])
		} else {
			cmd.Usage()
		}
	},
}

func init() {
	startupCmd.AddCommand(startupListCmd)
	startupCmd.AddCommand(startupGetCmd)
	startupCmd.AddCommand(startupApproveCmd)
	startupCmd.AddCommand(startupSendbackCmd)
	startupUpdateCmd.Flags().StringVarP(&reviewStatus, "review-status", "r", "", "审核状态, 可为approve或者sendback")
	startupUpdateCmd.Flags().StringVarP(&reviewStatus, "advice", "a", "", "修改建议, 当审核状态为sendback时需要填写")
	startupCmd.AddCommand(startupUpdateCmd)
	RootCmd.AddCommand(startupCmd)
}
