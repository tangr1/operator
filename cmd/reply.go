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
	"github.com/tangr1/hicto/client/reply"
	"github.com/fatih/color"
	"io"
)

type Reply struct {
	ID              string `chinese:"ID" list:"true"`
	TopicID              string `chinese:"问题ID" list:"true"`
	TopicTitle              string `chinese:"问题标题" list:"true"`
	AuthorID string `chinese:"答案作者ID" list:"true"`
	AuthorName string `chinese:"答案作者姓名" list:"true"`
	Content              string `chinese:"内容" list:"true"`
}

func ToReplyRecord(from *models.Reply) interface{} {
	var to Reply
	ModelToRecord(from, &to)
	to.AuthorName = from.Author.Name
	to.AuthorID = strconv.FormatInt(from.Author.ID, 10)
	to.TopicTitle = from.Topic.Title
	return to
}

// replyCmd represents the reply command
var replyCmd = &cobra.Command{
	Use:   "reply",
	Short: "答案相关命令",
	Long:  "答案相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var replyListCmd = &cobra.Command{
	Use:   "list",
	Short: "查看答案列表",
	Long: "查看答案列表",
	Run: func(cmd *cobra.Command, args []string) {
		params := reply.GetRepliesParams{
			Page: IgnoreErrorInt64(cmd.Flags().GetInt64("page")),
			Pagesize: IgnoreErrorInt64(cmd.Flags().GetInt64("page-size")),
			Authorid: IgnoreErrorInt64(cmd.Flags().GetInt64("author-id")),
			Topicid: IgnoreErrorInt64(cmd.Flags().GetInt64("topic-id")),
		}
		result, err := apiClient.Reply.GetReplies(&params, writer)
		if err != nil {
			color.Red("获取答案列表失败: %v\n", err)
			return
		}
		replies := make([]interface{}, len(result.Payload.Content))
		for i, from := range result.Payload.Content {
			replies[i] = ToReplyRecord(&from)
		}
		ListRecords(replies)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var replyGetCmd = &cobra.Command{
	Use:   "get ID",
	Short: "查看答案",
	Long: "查看答案",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			result, err := apiClient.Reply.GetRepliesID(&reply.GetRepliesIDParams{ID: id}, writer)
			if err == nil {
				ShowRecord(ToReplyRecord(result.Payload))
			} else {
				color.Red("获取答案失败: %v\n", err)
			}
		} else {
			cmd.Usage()
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var replyDeleteCmd = &cobra.Command{
	Use:   "delete ID",
	Short: "删除答案",
	Long: "删除答案",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			_, err := apiClient.Reply.DeleteRepliesID(&reply.DeleteRepliesIDParams{ID: id}, writer)
			if err == nil {
				color.Green("删除答案成功\n")
			} else if err == io.EOF {
				color.Green("删除答案成功\n")
			} else {
				color.Red("删除答案失败: %v\n", err)
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
	replyListCmd.Flags().Int64P("page", "p", 0, "页码")
	replyListCmd.Flags().Int64P("page-size", "P", 100, "每页条目数")
	replyListCmd.Flags().Int64P("topic-id", "t", -1, "问题ID")
	replyListCmd.Flags().Int64P("author-id", "a", -1, "专家ID")
	replyCmd.AddCommand(replyListCmd)
	replyCmd.AddCommand(replyGetCmd)
	replyCmd.AddCommand(replyDeleteCmd)
	RootCmd.AddCommand(replyCmd)
}
