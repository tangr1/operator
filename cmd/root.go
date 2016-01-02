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
	"fmt"
	"os"
	"strconv"

	"github.com/go-swagger/go-swagger/client"
	"github.com/spf13/cobra"
	"github.com/tangr1/hicto/client"
	"github.com/tangr1/hicto/models"
	"strings"
	"text/template"
)

const (
	templateColorAnsi = "{{ $cmd := . }}" +
	"{{if .HasParent}}" +
	"\x1b[1m用法:\x1b[0m " +
	"{{if .Runnable}}" +
	"{{if not (eq .Parent.Use \"hicto\")}}" +
	"\n{{ .Parent.Name}} {{ .Use}}{{if .HasFlags}} [FLAGS]{{end}}\n" +
	"{{else}}" +
	"\n{{ .Use}}{{if .HasFlags}} [FLAGS]{{end}}\n" +
	"{{end}}" +
	"{{end}}" +
	"{{if .HasSubCommands}}" +
	"{{if .Runnable}}" +
	"{{ .Use}} COMMAND{{if .HasFlags}} [FLAGS]{{end}}\n" +
	"{{else}}" +
	"\n{{ .Use}} COMMAND{{if .HasFlags}} [FLAGS]{{end}}\n" +
	"{{end}}" +
	"{{end}}" +
	"{{else}}" +
	"{{if .Runnable}}" +
	"\n{{.Use}}{{if .HasFlags}} [FLAGS]{{end}}\n" +
	"{{end}}" +
	"{{if .HasSubCommands}}" +
	"{{if .Runnable}}" +
	"{{ .Use}} COMMAND{{if .HasFlags}} [FLAGS]{{end}}\n" +
	"{{else}}" +
	"{{if (eq .Use \"hicto\")}}" +
	"COMMAND [FLAGS]\n" +
	"{{else}}" +
	"{{ .Use}} COMMAND{{if .HasFlags}} [FLAGS]{{end}}\n" +
	"{{end}}" +
	"{{end}}" +
	"{{end}}" +
	"{{end}}" +
	"{{ if .HasSubCommands}}" +
	"\n\x1b[1m可用COMMAND:\x1b[0m\n" +
	"{{range .Commands}}" +
	"{{if not (eq .Use \"help [command]\") }}" +
	"\x1b[32m{{rpad .Use .UsagePadding }}\x1b[0m {{.Short}}\n" +
	"{{end}}" +
	"{{end}}" +
	"{{end}}" +
	"{{if .HasFlags}}" +
	"\n\x1b[1m可用FLAG:\x1b[0m\n" +
	"{{.Flags.FlagUsages}}" +
	"{{end}}" +
	"{{if .HasSubCommands}}" +
	"{{if .HasParent}}" +
	"\n使用 \x1b[1mhelp {{ .Name}} COMMAND\x1b[0m 来获取每个命令的更多帮助\n\x1b[0m" +
	"{{else}}" +
	"\n使用 \x1b[1mhelp COMMAND\x1b[0m 来获取每个命令的更多帮助\n\x1b[0m" +
	"{{end}}" +
	"{{end}}"
)

var writer client.AuthInfoWriter
var codes map[string]*models.Code

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "hicto",
	Short:         "HiCTO 客户端命令行工具",
	Long:          "HiCTO 客户端命令行工具",
	SilenceErrors: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	miscClient := *operations.Default.Misc
	codeResult, err := miscClient.GetCodes(nil, writer)
	if err != nil {
		fmt.Printf("获取码表失败: %v\n", err)
		os.Exit(1)
	}
	codes = make(map[string]*models.Code)
	for _, code := range codeResult.Payload {
		codes[code.Name] = code
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().StringVarP(&server, "server", "s", "localhost:8090", "服务器地址")
	RootCmd.SetUsageFunc(UsageFunc)
}

func GetCode(name string, key int32) string {
	return codes[name].Mapping[strconv.Itoa(int(key))]
}

func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}

func UsageFunc(cmd *cobra.Command) error {
	t := template.New("top")
	t.Funcs(template.FuncMap{
		"trim": strings.TrimSpace,
		"rpad": rpad,
		"gt":   cobra.Gt,
		"eq":   cobra.Eq,
	})
	usageTmpl := templateColorAnsi
	template.Must(t.Parse(usageTmpl))
	return t.Execute(cmd.Out(), cmd)
}
