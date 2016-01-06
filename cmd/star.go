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
	"github.com/tangr1/hicto/client/operation"
	"github.com/fatih/color"
	"strings"
)

type Star struct {
	Date              string `chinese:"日期" list:"true"`
	Experts              string `chinese:"创业公司" list:"true"`
	Startups              string `chinese:"专家" list:"true"`
}

func ToStarRecord(from *models.Star) interface{} {
	var to Star
	ModelToRecord(from, &to)
	var experts []string
	for _, expert := range from.Experts {
		experts = append(experts, expert.Name)
	}
	to.Experts = strings.Join(experts, ", ")
	var startups []string
	for _, startup := range from.Startups {
		startups = append(startups, startup.Name)
	}
	to.Startups = strings.Join(startups, ", ")
	return to
}

// starCmd represents the star command
var starCmd = &cobra.Command{
	Use:   "star",
	Short: "明星相关命令",
	Long:  "明星相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var starLatestCmd = &cobra.Command{
	Use:   "latest",
	Short: "查看最新明星列表",
	Long: "查看最新明星列表",
	Run: func(cmd *cobra.Command, args []string) {
		params := operation.GetStarsLatestParams {
		}
		result, err := apiClient.Operation.GetStarsLatest(&params, writer)
		if err == nil {
			ShowRecord(ToStarRecord(result.Payload))
		} else {
			color.Red("获取最新明星列表失败: %v\n", err)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

func init() {
	starCmd.AddCommand(starLatestCmd)
	RootCmd.AddCommand(starCmd)
}
