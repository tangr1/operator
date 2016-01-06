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
	"github.com/tangr1/hicto/client/comment"
	"github.com/fatih/color"
	"io"
)

type Comment struct {
	ID              string `chinese:"ID" list:"true"`
	TopicID              string `chinese:"问题ID" list:"true"`
	ReplyID              string `chinese:"答案ID" list:"true"`
	ReferCommentID       string `chinese:"评论ID" list:"true"`
	AuthorID string `chinese:"评论作者ID" list:"true"`
	AuthorName string `chinese:"评论作者姓名" list:"true"`
	Content              string `chinese:"内容" list:"true"`
}

func ToCommentRecord(from *models.Comment) interface{} {
	var to Comment
	ModelToRecord(from, &to)
	to.AuthorName = from.Author.Name
	to.AuthorID = strconv.FormatInt(from.Author.ID, 10)
	if to.TopicID == "0" {
		to.TopicID = ""
	}
	if to.ReplyID == "0" {
		to.ReplyID = ""
	}
	if to.ReferCommentID == "0" {
		to.ReferCommentID = ""
	}
	return to
}

// commentCmd represents the comment command
var commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "评论相关命令",
	Long:  "评论相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var commentListCmd = &cobra.Command{
	Use:   "list",
	Short: "查看评论列表",
	Long: "查看评论列表",
	Run: func(cmd *cobra.Command, args []string) {
		params := comment.GetCommentsParams{
			Page: IgnoreErrorInt64(cmd.Flags().GetInt64("page")),
			Pagesize: IgnoreErrorInt64(cmd.Flags().GetInt64("page-size")),
			Topicid: IgnoreErrorInt64(cmd.Flags().GetInt64("topic-id")),
			Replyid: IgnoreErrorInt64(cmd.Flags().GetInt64("reply-id")),
		}
		result, err := apiClient.Comment.GetComments(&params, writer)
		if err != nil {
			color.Red("获取评论列表失败: %v\n", err)
			return
		}
		comments := make([]interface{}, len(result.Payload.Content))
		for i, from := range result.Payload.Content {
			comments[i] = ToCommentRecord(&from)
		}
		ListRecords(comments)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var commentGetCmd = &cobra.Command{
	Use:   "get ID",
	Short: "查看评论",
	Long: "查看评论",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			result, err := apiClient.Comment.GetCommentsID(&comment.GetCommentsIDParams{ID: id}, writer)
			if err == nil {
				ShowRecord(ToCommentRecord(result.Payload))
			} else {
				color.Red("获取评论失败: %v\n", err)
			}
		} else {
			cmd.Usage()
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var commentDeleteCmd = &cobra.Command{
	Use:   "delete ID",
	Short: "删除评论",
	Long: "删除评论",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			_, err := apiClient.Comment.DeleteCommentsID(&comment.DeleteCommentsIDParams{ID: id}, writer)
			if err == nil {
				color.Green("删除评论成功\n")
			} else if err == io.EOF {
				color.Green("删除评论成功\n")
			} else {
				color.Red("删除评论失败: %v\n", err)
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
	commentListCmd.Flags().Int64P("page", "p", 0, "页码")
	commentListCmd.Flags().Int64P("page-size", "P", 100, "每页条目数")
	commentListCmd.Flags().Int64P("topic-id", "t", -1, "问题ID")
	commentListCmd.Flags().Int64P("reply-id", "r", -1, "问题ID")
	commentCmd.AddCommand(commentListCmd)
	commentCmd.AddCommand(commentGetCmd)
	commentCmd.AddCommand(commentDeleteCmd)
	RootCmd.AddCommand(commentCmd)
}
