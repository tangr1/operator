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

package main

import (
	"strings"
	"os"
	"github.com/tangr1/hicto/cmd"
	"gopkg.in/readline.v1"
	"github.com/fatih/color"
)

func main() {
	cmd.RootCmd.Help()
	rl, err := readline.NewEx(&readline.Config{
		Prompt: "hicto > ",
		HistoryFile: "/tmp/readline.tmp",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		args := strings.Fields(line)
		os.Args = []string{"./hicto"}
		os.Args = append(os.Args, args...)
		err = cmd.RootCmd.Execute()
		if err != nil {
			color.Red("无效的命令，请运行 help 获得帮助")
		}
	}
}
