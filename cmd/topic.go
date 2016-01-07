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
	"strings"
	"strconv"
	"github.com/tangr1/hicto/client/topic"
	"github.com/fatih/color"
	"io"
)

type Topic struct {
	ID              string `chinese:"ID" list:"true"`
	Title              string `chinese:"标题" list:"true"`
	Summary              string `chinese:"摘要" list:"true"`
	Content              string `chinese:"内容"`
	Category string `chinese:"类别" list:"true" code:"category"`
	CtoCoins string `chinese:"悬赏分" list:"true"`
	Tags string `chinese:"标签"`
	StartupName string `chinese:"创业公司" list:"true"`
	AuthorName string `chinese:"创业公司员工" list:"true"`
	ReplyCount string `chinese:"回复数" list:"true"`
	ResolvedComment string `chinese:"采纳评论"`
	ResolvedReplyID string `chinese:"采纳回复ID" list:"true"`
	ReadableCreateTime string `chinese:"创建时间" list:"true"`
	ReadableUpdateTime string `chinese:"更新时间" list:"true"`
	Status string `chinese:"问题状态" list:"true" code:"postStatus"`
}

func ToTopicRecord(from *models.Topic) interface{} {
	var to Topic
	ModelToRecord(from, &to)
	to.Tags = strings.Join(from.Tags, ", ")
	to.AuthorName = from.Author.Name
	to.StartupName = from.Startup.Name
	if to.ResolvedReplyID == "0" {
		to.ResolvedReplyID = ""
	}
	return to
}

// topicCmd represents the topic command
var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "问题相关命令",
	Long:  "问题相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var topicListCmd = &cobra.Command{
	Use:   "list",
	Short: "查看问题列表",
	Long: "查看问题列表",
	Run: func(cmd *cobra.Command, args []string) {
		params := topic.GetTopicsParams{
			Page: IgnoreErrorInt64(cmd.Flags().GetInt64("page")),
			Pagesize: IgnoreErrorInt64(cmd.Flags().GetInt64("page-size")),
			Category: IgnoreErrorInt64(cmd.Flags().GetInt64("category")),
			Resolved: IgnoreErrorBool(cmd.Flags().GetBool("resolved")),
			Unresolved: IgnoreErrorBool(cmd.Flags().GetBool("unresolved")),
			Startupid: IgnoreErrorInt64(cmd.Flags().GetInt64("startup-id")),
			Authorid: IgnoreErrorInt64(cmd.Flags().GetInt64("author-id")),
			Status: IgnoreErrorInt64(cmd.Flags().GetInt64("status")),
			Wonderful: IgnoreErrorInt64(cmd.Flags().GetInt64("wonderful")),
		}
		result, err := apiClient.Topic.GetTopics(&params, writer)
		if err != nil {
			color.Red("获取问题列表失败: %v\n", err)
			return
		}
		topics := make([]interface{}, len(result.Payload.Content))
		for i, from := range result.Payload.Content {
			topics[i] = ToTopicRecord(&from)
		}
		ListRecords(topics)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var topicGetCmd = &cobra.Command{
	Use:   "get ID",
	Short: "查看问题",
	Long: "查看问题",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			result, err := apiClient.Topic.GetTopicsID(&topic.GetTopicsIDParams{ID: id}, writer)
			if err == nil {
				ShowRecord(ToTopicRecord(result.Payload))
			} else {
				color.Red("获取问题失败: %v\n", err)
			}
		} else {
			cmd.Usage()
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var topicDeleteCmd = &cobra.Command{
	Use:   "delete ID",
	Short: "删除问题",
	Long: "删除问题",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			_, err := apiClient.Topic.DeleteTopicsID(&topic.DeleteTopicsIDParams{ID: id}, writer)
			if err == nil {
				color.Green("删除问题成功\n")
			} else if err == io.EOF {
				color.Green("删除问题成功\n")
			} else {
				color.Red("删除问题失败: %v\n", err)
			}
		} else {
			cmd.Usage()
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var topicDeleteApproveCmd = &cobra.Command{
	Use:   "approve ID",
	Short: "通过问题删除申请",
	Long: "通过问题删除申请",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			body := models.UpdateTopicRequest {
				Status: 3,
			}
			result, err := apiClient.Topic.PutTopicsID(&topic.PutTopicsIDParams{ID: id, Body: &body}, writer)
			if err == nil {
				ShowRecord(ToTopicRecord(result.Payload))
			} else {
				color.Red("更新问题失败: %v\n", err)
			}
		} else {
			cmd.Usage()
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var topicDeleteRejectCmd = &cobra.Command{
	Use:   "reject ID [REJECT-REASON]",
	Short: "拒绝问题删除申请",
	Long: "拒绝问题删除申请",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			body := models.UpdateTopicRequest {
				Status: 1,
			}
			result, err := apiClient.Topic.PutTopicsID(&topic.PutTopicsIDParams{ID: id, Body: &body}, writer)
			if err == nil {
				ShowRecord(ToTopicRecord(result.Payload))
			} else {
				color.Red("更新问题失败: %v\n", err)
			}
		} else if len(args) == 2 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			body := models.UpdateTopicRequest {
				Status: 1,
				RejectReason: args[1],
			}
			result, err := apiClient.Topic.PutTopicsID(&topic.PutTopicsIDParams{ID: id, Body: &body}, writer)
			if err == nil {
				ShowRecord(ToTopicRecord(result.Payload))
			} else {
				color.Red("更新问题失败: %v\n", err)
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
	topicListCmd.Flags().Int64P("page", "p", 0, "页码")
	topicListCmd.Flags().Int64P("page-size", "P", 100, "每页条目数")
	topicListCmd.Flags().Int64P("author-id", "a", -1, "作者ID")
	topicListCmd.Flags().Int64P("category", "c", -1, "类别")
	topicListCmd.Flags().BoolP("resolved", "r", false, "已解决优先")
	topicListCmd.Flags().BoolP("unresolved", "u", false, "未解决优先")
	topicListCmd.Flags().Int64P("status", "s", 1, "问题状态")
	topicListCmd.Flags().Int64P("startup-id", "S", -1, "创业公司ID")
	topicListCmd.Flags().Int64P("wonderful", "w", -1, "精彩话题数")
	topicCmd.AddCommand(topicListCmd)
	topicCmd.AddCommand(topicGetCmd)
	topicCmd.AddCommand(topicDeleteCmd)
	topicCmd.AddCommand(topicDeleteApproveCmd)
	topicCmd.AddCommand(topicDeleteRejectCmd)
	RootCmd.AddCommand(topicCmd)
}
