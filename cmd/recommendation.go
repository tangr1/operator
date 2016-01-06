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
)

type Recommendation struct {
	ID              string `chinese:"ID" list:"true"`
	Name              string `chinese:"姓名" list:"true"`
	Position              string `chinese:"职位" list:"true"`
	Photo              string `chinese:"照片" list:"true"`
	Words       string `chinese:"推荐语" list:"true"`
}

func ToRecommendationRecord(from *models.Recommendation) interface{} {
	var to Recommendation
	ModelToRecord(from, &to)
	return to
}

// recommendationCmd represents the recommendation command
var recommendationCmd = &cobra.Command{
	Use:   "recommendation",
	Short: "推荐相关命令",
	Long:  "推荐相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var recommendationListCmd = &cobra.Command{
	Use:   "list",
	Short: "查看推荐列表",
	Long: "查看推荐列表",
	Run: func(cmd *cobra.Command, args []string) {
		params := operation.GetRecommendationsParams{
			Page: IgnoreErrorInt64(cmd.Flags().GetInt64("page")),
			Pagesize: IgnoreErrorInt64(cmd.Flags().GetInt64("page-size")),
		}
		result, err := apiClient.Operation.GetRecommendations(&params, writer)
		if err != nil {
			color.Red("获取推荐列表失败: %v\n", err)
			return
		}
		recommendations := make([]interface{}, len(result.Payload.Content))
		for i, from := range result.Payload.Content {
			recommendations[i] = ToRecommendationRecord(from)
		}
		ListRecords(recommendations)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var recommendationGetCmd = &cobra.Command{
	Use:   "get ID",
	Short: "查看推荐",
	Long: "查看推荐",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			result, err := apiClient.Operation.GetRecommendationsID(&operation.GetRecommendationsIDParams{ID: id}, writer)
			if err == nil {
				ShowRecord(ToRecommendationRecord(result.Payload))
			} else {
				color.Red("获取推荐失败: %v\n", err)
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
	recommendationListCmd.Flags().Int64P("page", "p", 0, "页码")
	recommendationListCmd.Flags().Int64P("page-size", "P", 100, "每页条目数")
	recommendationCmd.AddCommand(recommendationListCmd)
	recommendationCmd.AddCommand(recommendationGetCmd)
	RootCmd.AddCommand(recommendationCmd)
}
