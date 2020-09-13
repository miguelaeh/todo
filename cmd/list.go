/*
Copyright © 2020 Miguel Angel Cabrera Minagorri <devgorri@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
  "encoding/json"
  "io/ioutil"

  "github.com/miguelaeh/todo/types"

  homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the collection of pending tasks ordered by priority",
	Long: "List the collection of tasks ordered by priority",
	Run: func(cmd *cobra.Command, args []string) {
	  // Build data file path
    path, _ := homedir.Dir()
    path += "/.todo/tasks.json"

    // Get current list of tasks
    f, _ := ioutil.ReadFile(path)
	  tasks := types.Tasks{}
	  _ = json.Unmarshal([]byte(f), &tasks)
    tasks.Print()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
