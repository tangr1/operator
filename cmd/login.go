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
	httpclient "github.com/go-swagger/go-swagger/httpkit/client"

	"github.com/spf13/cobra"
	"github.com/tangr1/hicto/client"
	"github.com/tangr1/hicto/models"
	"github.com/tangr1/hicto/client/security"
	"github.com/fatih/color"
)

var password string
var operator bool

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login [EMAIL] [FLAGS]",
	Short: "登录相关命令",
	Long: "登录相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
		} else {
			securityClient := operations.Default.Security
			if operator {
				body := &models.AuthenticationRequest{Email: args[0], Password: password}
				loginResult, err := securityClient.PostOperatorLogin(&security.PostOperatorLoginParams{Body: body}, writer)
				if err == nil {
					color.Green("登录成功")
					writer = httpclient.APIKeyAuth("X-AUTH-TOKEN", "header", loginResult.Payload.Token)
				} else {
					color.Red("登录失败: %v\n", err)
				}
			} else {
				body := &models.AuthenticationRequest{Email: args[0], Password: password}
				loginResult, err := securityClient.PostAuthLogin(&security.PostAuthLoginParams{Body: body}, writer)
				if err == nil {
					color.Green("登录成功")
					writer = httpclient.APIKeyAuth("X-AUTH-TOKEN", "header", loginResult.Payload.Token)
				} else {
					color.Red("登录失败: %v\n", err)
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringVarP(&password, "password", "p", "password", "用户密码")
	loginCmd.Flags().BoolVarP(&operator, "operator", "o", false, "用户为运营人员")
}
