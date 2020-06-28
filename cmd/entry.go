/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "create a new entry",
	Long:  `Create a new entry in your journalp directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("journalpDir")
		if _, err := os.Stat(path); err == nil {
			fmt.Fprintf(os.Stderr, "directory %s does exist. \\o/\n", path)
			year, month, day := time.Now().Date()
			md := filepath.Join(path, fmt.Sprintf("%d", year), strings.ToLower(month.String()))
			df := filepath.Join(md, fmt.Sprintf("%d.md", day))

			if err = os.MkdirAll(md, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "directory %s can not be created.\n", md)
			}
			if err := ioutil.WriteFile(df, []byte(fmt.Sprintf("# %d-%s-%d\n", year, month, day)), 0644); err != nil {
				fmt.Fprintf(os.Stderr, "file %s can not be written.\n", df)
			}
			fmt.Fprintf(os.Stdout, "%s\n", df)

		} else {
			fmt.Fprintf(os.Stderr, "directory %s does not exist.\n", path)
		}
	},
}

func init() {
	createCmd.AddCommand(entryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// entryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// entryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
