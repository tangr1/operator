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

var topicsParams = topic.GetTopicsParams {
	Authorid: -1,
	Category: "",
	Page: 0,
	Pagesize: 100,
	Resolved: false,
	Startupid: -1,
	Status: 1,
	Unresolved: false,
	Wonderful: -1,
}

type Topic struct {
	ID              string `chinese:"ID" list:"true"`
	Title              string `chinese:"标题" list:"true"`
	Summary              string `chinese:"摘要"`
	Content              string `chinese:"内容"`
	Category string `chinese:"类别" list:"true" code:"category"`
	CtoCoins string `chinese:"悬赏分" list:"true"`
	Tags string `chinese:"标签" list:"true"`
	StartupName string `chinese:"问题" list:"true"`
	AuthorName string `chinese:"问题员工" list:"true"`
	ReplyCount string `chinese:"回复数" list:"true"`
	ResolvedComment string `chinese:"采纳评论"`
	ResolvedReplyID string `chinese:"采纳回复ID" list:"true"`
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
}

var topicListCmd = &cobra.Command{
	Use:   "list",
	Short: "查看问题列表",
	Long: "查看问题列表",
	Run: func(cmd *cobra.Command, args []string) {
		var params = topicsParams
		result, err := apiClient.Topic.GetTopics(&params, writer)
		if err != nil {
			color.Red("获取问题列表失败: %v\n", err)
			return
		}
		topics := make([]interface{}, len(result.Payload.Content))
		for i, from := range result.Payload.Content {
			topics[i] = ToTopicRecord(&from)
		}
		if operator {
			params.Status = 2
			result, err = apiClient.Topic.GetTopics(&params, writer)
			if err != nil {
				color.Red("获取问题列表失败: %v\n", err)
				return
			}
			for _, from := range result.Payload.Content {
				topics = append(topics, from)
			}
			params.Status = 3
			result, err = apiClient.Topic.GetTopics(&params, writer)
			if err != nil {
				color.Red("获取问题列表失败: %v\n", err)
				return
			}
			for _, from := range result.Payload.Content {
				topics = append(topics, from)
			}
		}
		ListRecords(topics)
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
}

func init() {
	topicListCmd.Flags().Int64VarP(&topicsParams.Page, "page", "p", 0, "页码")
	topicListCmd.Flags().Int64VarP(&topicsParams.Pagesize, "page-size", "P", 100, "每页条目数")
	topicCmd.AddCommand(topicListCmd)
	topicCmd.AddCommand(topicGetCmd)
	topicCmd.AddCommand(topicDeleteCmd)
	RootCmd.AddCommand(topicCmd)
}
