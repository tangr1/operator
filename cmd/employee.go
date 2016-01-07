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
	"github.com/tangr1/hicto/client/employee"
	"github.com/fatih/color"
)

type Employee struct {
	ID              string `chinese:"ID" list:"true"`
	Name string `chinese:"姓名" list:"true"`
	Email string `chinese:"邮箱" list:"true"`
	Phone string `chinese:"电话"`
	StartupID string `chinese:"创业公司ID"`
	Position string `chinese:"职位" list:"true"`
	Admin string `chinese:"是否为管理员" list:"true"`
	ConsumeCtoCoins string `chinese:"消费的悬赏分总额" list:"true"`
	TopicCount string `chinese:"问题总数" list:"true"`
	Avatar string `chinese:"头像"`
	Description string `chinese:"个人简介"`
	DeviceToken string `chinese:"设备token"`
	NotifyOnlyMyTopic string `chinese:"是否只接收自己问题的通知"`
	NotifyCommentNewCommentByEmail string `chinese:"是否接收评论的新评论邮件通知"`
	NotifyCommentNewCommentByPush string `chinese:"是否接收评论的新评论推送通知"`
	NotifyNewReplyByEmail string `chinese:"是否接收问题的新答案邮件通知"`
	NotifyNewReplyByPush string `chinese:"是否接收问题的新答案推送通知"`
	NotifyTopicNewCommentByEmail string `chinese:"是否接收问题的新评论邮件通知"`
	NotifyTopicNewCommentByPush string `chinese:"是否接收问题的新评论推送通知"`
	ReadableCreateTime string `chinese:"创建时间" list:"true"`
	ReadableUpdateTime string `chinese:"更新时间" list:"true"`
}

func ToEmployeeRecord(from *models.Employee) interface{} {
	var to Employee
	ModelToRecord(from, &to)
	return to
}

var employeeCmd = &cobra.Command{
	Use:   "employee",
	Short: "员工相关命令",
	Long:  "员工相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var employeeListCmd = &cobra.Command{
	Use:   "list STARTUP-ID",
	Short: "查看员工列表",
	Long: "查看员工列表",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			return
		}
		params := employee.GetEmployeesParams {
			Page: IgnoreErrorInt64(cmd.Flags().GetInt64("page")),
			Pagesize: IgnoreErrorInt64(cmd.Flags().GetInt64("page-size")),
			Startupid: IgnoreErrorInt64(strconv.ParseInt(args[0], 10, 64)),
		}
		result, err := apiClient.Employee.GetEmployees(&params, writer)
		if err != nil {
			color.Red("获取员工列表失败: %v\n", err)
			return
		}
		employees := make([]interface{}, len(result.Payload.Content))
		for i, from := range result.Payload.Content {
			employees[i] = ToEmployeeRecord(&from)
		}
		ListRecords(employees)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

var employeeGetCmd = &cobra.Command{
	Use:   "get ID",
	Short: "查看员工",
	Long: "查看员工",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			result, err := apiClient.Employee.GetEmployeesID(&employee.GetEmployeesIDParams{ID: id}, writer)
			if err == nil {
				ShowRecord(ToEmployeeRecord(result.Payload))
			} else {
				color.Red("获取员工失败: %v\n", err)
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
	employeeListCmd.Flags().Int64P("page", "p", 0, "页码")
	employeeListCmd.Flags().Int64P("page-size", "P", 100, "每页条目数")
	employeeCmd.AddCommand(employeeListCmd)
	employeeCmd.AddCommand(employeeGetCmd)
	RootCmd.AddCommand(employeeCmd)
}
