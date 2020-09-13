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

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new pending task",
	Long: `Create a new pending task.

  todo new <task-title> [timer] [priority]
`,
	Run: func(cmd *cobra.Command, args []string) {
    title, _ := cmd.Flags().GetString("title")
    alarm, _ := cmd.Flags().GetString("alarm")
    priority, _ := cmd.Flags().GetInt("priority")

    if alarm != "" {
      // TODO (miguelaeh): Check time format
    }

    if priority < 1 {
      log.Fatal("'priority'cannot be lower than 1")
    }

    // Build data file path
    path, _ := homedir.Dir()
    path += "/.todo/tasks.json"

    // Get current list of tasks
    f, _ := ioutil.ReadFile(path)
	  currentTasks := types.Tasks{}
	  _ = json.Unmarshal([]byte(f), &currentTasks)

    // Create the new task
    task := types.Task {
      Title: title,
      Alarm: alarm,
      Priority: priority,
    }

    // Check the priority max number
    if currentTasks.Len() > 0 {
      maxPriority := currentTasks[currentTasks.Len() - 1].Priority + 1
      if task.Priority > maxPriority {
        fmt.Printf("WARN: fixing provided priority. Changing from %d to %d\n", task.Priority, maxPriority)
        task.Priority = maxPriority
      }
    } else {
      task.Priority = 1
    }
    // Add the task at the beggining changing all priorities
    for i := 0; i < currentTasks.Len(); i++ {
      if currentTasks[i].Priority >= task.Priority {
        currentTasks[i].Priority += 1
      }
    }
    newTasks :=  append(currentTasks, task)
    sort.Sort(newTasks)

    newContent, _ := json.MarshalIndent(newTasks, "", " ")
	  _ = ioutil.WriteFile(path, newContent, 0644)

		fmt.Println("Task " + task.Title +  " added!")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	newCmd.Flags().String("title", "New task", "Title for the new task")
  newCmd.MarkFlagRequired("title")
  newCmd.Flags().String("alarm", "", "Timer to get a notification")
  newCmd.Flags().Int("priority", 1, "Add the priority number to show it in the tasks' list")
}
