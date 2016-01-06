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
	"github.com/spf13/cobra"
	"github.com/tangr1/hicto/models"
	"strconv"
	"github.com/tangr1/hicto/client/expert"
	"github.com/fatih/color"
	"strings"
)

type Expert struct {
	ID              string `chinese:"ID" list:"true"`
	Name string `chinese:"姓名" list:"true"`
	Email string `chinese:"邮箱" list:"true"`
	Phone string `chinese:"电话"`
	City string `chinese:"城市" code:"city"`
	Company string `chinese:"公司"`
	Position string `chinese:"职位"`
	Expertise string `chinese:"专长" list:"true"`
	CtoCoins string `chinese:"悬赏分" list:"true"`
	Reputation string `chinese:"威望分"`
	ReplyCount string `chinese:"回答总数"`
	AcceptedReplyCount string `chinese:"被采纳的答案数" list:"true"`
	Avatar string `chinese:"头像"`
	Description string `chinese:"个人简介"`
	DeviceToken string `chinese:"设备token"`
	ViewCount string `chinese:"浏览次数"`
	CoverImage string `chinese:"封面图片"`
	HelpedStartups string `chinese:"帮助过的创业公司"`
	Internal string `chinese:"是否为内部专家"`
	InviteQuota string `chinese:"可邀请专家数"`
	InviteUsed string `chinese:"已邀请专家数"`
	ManagementSkill string `chinese:"管理技能" list:"true" code:"managementSkill"`
	NotifyNewTopicByPush string `chinese:"是否接收新主题推送通知"`
	NotifyOnlyFreeTime string `chinese:"是否只在空闲时间发通知"`
	NotifyNewTopicByEmail string `chinese:"是否接收新主题邮件通知"`
	NotifyCommentNewCommentByEmail string `chinese:"是否接收评论的新评论邮件通知"`
	NotifyCommentNewCommentByPush string `chinese:"是否接收评论的新评论推送通知"`
	NotifyReplyAcceptedByEmail string `chinese:"是否接收答案被采纳邮件通知"`
	NotifyReplyAcceptedByPush string `chinese:"是否接收答案被采纳推送通知"`
	NotifyReplyNewCommentByEmail string `chinese:"是否接收答案的新评论邮件通知"`
	NotifyReplyNewCommentByPush string `chinese:"是否接收答案的新评论推送通知"`
	ReviewStatus string `chinese:"审核状态" list:"true" code:"reviewStatus" list:"true"`
}

func ToExpertRecord(from *models.Expert) interface{} {
	var to Expert
	ModelToRecord(from, &to)
	startups := make([]string, len(from.HelpedStartups))
	for i, startup := range from.HelpedStartups {
		startups[i] = startup.Name
	}
	to.HelpedStartups = strings.Join(startups, ", ")
	expertises := make([]string, len(from.Expertise))
	for i, expertise := range from.Expertise {
		expertises[i] = GetCode("category", expertise)
	}
	to.Expertise = strings.Join(expertises, ", ")
	return to
}

var expertCmd = &cobra.Command{
	Use:   "expert",
	Short: "专家相关命令",
	Long:  "专家相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var expertListCmd = &cobra.Command{
	Use:   "list",
	Short: "查看员工列表",
	Long: "查看员工列表",
	Run: func(cmd *cobra.Command, args []string) {
		params := expert.GetExpertsParams {
			Page: IgnoreErrorInt64(cmd.Flags().GetInt64("page")),
			Pagesize: IgnoreErrorInt64(cmd.Flags().GetInt64("page-size")),
			Category: IgnoreErrorInt64(cmd.Flags().GetInt64("category")),
		}
		result, err := apiClient.Expert.GetExperts(&params, writer)
		if err != nil {
			color.Red("获取专家列表失败: %v\n", err)
			return
		}
		experts := make([]interface{}, len(result.Payload.Content))
		for i, from := range result.Payload.Content {
			experts[i] = ToExpertRecord(&from)
		}
		ListRecords(experts)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var expertGetCmd = &cobra.Command{
	Use:   "get ID",
	Short: "查看专家",
	Long: "查看专家",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			result, err := apiClient.Expert.GetExpertsID(&expert.GetExpertsIDParams{ID: id}, writer)
			if err == nil {
				ShowRecord(ToExpertRecord(result.Payload))
			} else {
				color.Red("获取专家失败: %v\n", err)
			}
		} else {
			cmd.Usage()
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var expertApproveCmd = &cobra.Command{
	Use:   "approve ID",
	Short: "通过专家申请",
	Long: "通过专家申请",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			body := models.UpdateExpertRequest {
				ReviewStatus: 2,
			}
			result, err := apiClient.Expert.PutExpertsID(&expert.PutExpertsIDParams{ID: id, Body: &body}, writer)
			if err == nil {
				ShowRecord(ToExpertRecord(result.Payload))
			} else {
				color.Red("更新专家失败: %v\n", err)
			}
		} else {
			cmd.Usage()
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var expertSendbackCmd = &cobra.Command{
	Use:   "sendback ID [ADVICE]",
	Short: "打回专家申请, 并提出修改意见",
	Long: "打回专家申请, 并提出修改意见",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			body := models.UpdateExpertRequest {
				ReviewStatus: 2,
			}
			result, err := apiClient.Expert.PutExpertsID(&expert.PutExpertsIDParams{ID: id, Body: &body}, writer)
			if err == nil {
				ShowRecord(ToExpertRecord(result.Payload))
			} else {
				color.Red("更新专家失败: %v\n", err)
			}
		} else if len(args) == 2 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			body := models.UpdateExpertRequest {
				ReviewStatus: 1,
				ModifyReason: args[1],
			}
			result, err := apiClient.Expert.PutExpertsID(&expert.PutExpertsIDParams{ID: id, Body: &body}, writer)
			if err == nil {
				ShowRecord(ToExpertRecord(result.Payload))
			} else {
				color.Red("更新专家失败: %v\n", err)
			}
		} else {
			cmd.Usage()
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

func init() {
	expertListCmd.Flags().Int64P("page", "p", 0, "页码")
	expertListCmd.Flags().Int64P("page-size", "P", 100, "每页条目数")
	expertListCmd.Flags().Int64P("category", "c", -1, "类别")
	expertCmd.AddCommand(expertListCmd)
	expertCmd.AddCommand(expertGetCmd)
	expertCmd.AddCommand(expertApproveCmd)
	expertCmd.AddCommand(expertSendbackCmd)
	RootCmd.AddCommand(expertCmd)
}
