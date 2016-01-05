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
	"reflect"
	"github.com/fatih/structs"
	"github.com/olekukonko/tablewriter"
	"github.com/gosuri/uitable"
	"github.com/labstack/gommon/color"
	fatihcolor "github.com/fatih/color"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/spec"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/spf13/pflag"
)

const (
	templateColorAnsi = "{{ $cmd := . }}" +
	"{{if .HasParent}}" +
	"\x1b[1m用法:\x1b[0m " +
	"{{if .Runnable}}" +
	"{{if not .HasSubCommands}}" +
	"{{if not (eq .Parent.Use \"hicto\")}}" +
	"\n{{ .Parent.Name}} {{ .Use}}{{if .HasFlags}} [FLAGS]{{end}}\n" +
	"{{else}}" +
	"\n{{ .Use}}{{if .HasFlags}} [FLAGS]{{end}}\n" +
	"{{end}}" +
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
var apiClient *operations.CtoFunds

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:           "hicto",
	Short:         "HiCTO 客户端命令行工具",
	Long:          "HiCTO 客户端命令行工具",
	// SilenceErrors: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	PostRun: func(cmd *cobra.Command, args []string) {
		ResetFlags(cmd)
	},
}

func Execute() error {
	cmd, err := RootCmd.ExecuteC()
	cmd.Flag("help").Value.Set("false")
	return err
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().StringVarP(&server, "server", "s", "localhost:8090", "服务器地址")
	swaggerSpec, err := spec.New(operations.SwaggerJSON, "")
	if err != nil {
		// the swagger spec is valid because it was used to generated this code.
		panic(err)
	}
	runtime := httptransport.New(swaggerSpec)
	if len(os.Args) > 1 {
		runtime.Host = os.Args[1]
	} else {
		runtime.Host = "localhost:8090"
	}
	apiClient = operations.New(runtime, strfmt.Default)
	codeResult, err := apiClient.Misc.GetCodes(nil, writer)
	if err != nil {
		fatihcolor.Red("获取码表失败, 退出程序 %s\n", err.Error())
		os.Exit(1)
	}
	codes = make(map[string]*models.Code)
	for _, code := range codeResult.Payload {
		codes[code.Name] = code
	}

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

func ListRecords(records []interface{}) {
	if len(records) == 0 {
		fatihcolor.Magenta("没有符合查询条件的数据")
		return;
	}
	s := structs.New(records[0])
	table := tablewriter.NewWriter(os.Stdout)
	var header []string
	for _, field := range s.Fields() {
		if len(field.Tag("list")) > 0 {
			header = append(header, field.Tag("chinese"));
		}
	}
	table.SetHeader(header)
	for _, record := range records {
		var body []string
		s := structs.New(record)
		for _, field := range s.Fields() {
			if len(field.Tag("list")) > 0 {
				body = append(body, field.Value().(string))
			}
		}
		table.Append(body)
	}
	table.Render()
}

func ShowRecord(record interface{}) {
	table := uitable.New()
	table.MaxColWidth = 120
	table.Wrap = true // wrap columns
	s := structs.New(record)
	for _, field := range s.Fields() {
		table.AddRow(color.Bold(field.Tag("chinese")), field.Value().(string));
	}
	fmt.Println(table)
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if strings.Compare(a, b) == 0 {
			return true
		}
	}
	return false
}

func ModelToRecord(from interface{}, to interface{}) {
	fromStruct := structs.New(from)
	toStruct := structs.New(to)
	for _, fromField := range fromStruct.Fields() {
		if fromField.Kind() == reflect.Struct && StringInSlice(fromField.Name(), []string{"Resource", "User", "Post"}) {
			ModelToRecord(fromField.Value(), to)
		}
		toField, ok := toStruct.FieldOk(fromField.Name())
		if !ok {
			continue
		}
		switch (fromField.Kind()) {
		case reflect.String:
			toField.Set(fromField.Value().(string))
			break;
		case reflect.Int64:
			toField.Set(strconv.FormatInt(fromField.Value().(int64), 10))
			break;
		case reflect.Bool:
			if (fromField.Value().(bool)) {
				toField.Set("是")
			} else {
				toField.Set("否")
			}
			break;
		case reflect.Int32:
			if len(toField.Tag("code")) > 0 {
				toField.Set(GetCode(toField.Tag("code"), fromField.Value().(int32)))
			} else {
				toField.Set(strconv.Itoa(int(fromField.Value().(int32))))
			}
			break;
		default:
			break;
		}
	}
}

func ResetFlags(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		flag.Value.Set(flag.DefValue)
	})
}

func IgnoreErrorInt64(i int64, err error) (int64) {
	return i
}

func IgnoreErrorBool(b bool, err error) (bool) {
	return b
}

func IgnoreErrorString(s string, err error) (string) {
	return s
}
