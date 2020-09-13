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
  "fmt"
  "encoding/json"
  "io/ioutil"
  "sort"
  "log"

  "github.com/miguelaeh/todo/types"
  "github.com/spf13/cobra"
  homedir "github.com/mitchellh/go-homedir"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done and remove it from the list",
	Long: "Mark a task as done and remove it from the list",
	Run: func(cmd *cobra.Command, args []string) {
    id, _ := cmd.Flags().GetInt("id")
    if id < 0 {
      log.Fatal("'id' cannot lower than 1")
    }

    // Build data file path
    path, _ := homedir.Dir()
    path += "/.todo/tasks.json"

    // Get current list of tasks
    f, _ := ioutil.ReadFile(path)
	  tasks := types.Tasks{}
	  _ = json.Unmarshal([]byte(f), &tasks)

    if id > tasks.Len() {
      log.Fatal("'id' cannot be greater than the number of tasks")
    }

    title := tasks[id-1].Title

    // Dicrease the priority of the next elements
    for i := 0; i < tasks.Len(); i++ {
      if tasks[i].Priority >= id {
        tasks[i].Priority -= 1
      }
    }

    // Remove the element
    newTasks := append(tasks[:id-1], tasks[id:]...)
    sort.Sort(newTasks)

    newContent, _ := json.MarshalIndent(newTasks, "", " ")
	  _ = ioutil.WriteFile(path, newContent, 0644)

		fmt.Println("Task " + title +  " deleted!!")
  },
}

func init() {
	rootCmd.AddCommand(doneCmd)

  doneCmd.Flags().Int("id", 0, "Id of task to remove")
  doneCmd.MarkFlagRequired("id")
}
