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
	"github.com/tangr1/hicto/client/operation"
	"github.com/fatih/color"
	"strings"
)

type Feedback struct {
	ID              string `chinese:"ID" list:"true"`
	AuthorName              string `chinese:"姓名" list:"true"`
	Services              string `chinese:"希望提供的服务" list:"true"`
	Suggestion       string `chinese:"建议" list:"true"`
}

func ToFeedbackRecord(from *models.Feedback) interface{} {
	var to Feedback
	ModelToRecord(from, &to)
	var services []string
	for _, service := range from.Services {
		services = append(services, GetCode("service", service))
	}
	to.Services = strings.Join(services, ", ")
	to.AuthorName = from.Author.Name
	return to
}

// feedbackCmd represents the feedback command
var feedbackCmd = &cobra.Command{
	Use:   "feedback",
	Short: "反馈相关命令",
	Long:  "反馈相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var feedbackListCmd = &cobra.Command{
	Use:   "list",
	Short: "查看反馈列表",
	Long: "查看反馈列表",
	Run: func(cmd *cobra.Command, args []string) {
		params := operation.GetFeedbacksParams{
			Page: IgnoreErrorInt64(cmd.Flags().GetInt64("page")),
			Pagesize: IgnoreErrorInt64(cmd.Flags().GetInt64("page-size")),
		}
		result, err := apiClient.Operation.GetFeedbacks(&params, writer)
		if err != nil {
			color.Red("获取反馈列表失败: %v\n", err)
			return
		}
		feedbacks := make([]interface{}, len(result.Payload.Content))
		for i, from := range result.Payload.Content {
			feedbacks[i] = ToFeedbackRecord(from)
		}
		ListRecords(feedbacks)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var feedbackGetCmd = &cobra.Command{
	Use:   "get ID",
	Short: "查看反馈",
	Long: "查看反馈",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			result, err := apiClient.Operation.GetFeedbacksID(&operation.GetFeedbacksIDParams{ID: id}, writer)
			if err == nil {
				ShowRecord(ToFeedbackRecord(result.Payload))
			} else {
				color.Red("获取反馈失败: %v\n", err)
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
	feedbackListCmd.Flags().Int64P("page", "p", 0, "页码")
	feedbackListCmd.Flags().Int64P("page-size", "P", 100, "每页条目数")
	feedbackCmd.AddCommand(feedbackListCmd)
	feedbackCmd.AddCommand(feedbackGetCmd)
	RootCmd.AddCommand(feedbackCmd)
}
