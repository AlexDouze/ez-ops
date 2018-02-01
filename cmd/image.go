// Copyright © 2018 EasyMile <devops@easymile.com>
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
	"path/filepath"

	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:     "image IMAGE_NAME",
	Short:   "Build docker images",
	Long:    `Build a docker image from the definition found in the IMAGE_NAME folder`,
	Aliases: []string{"img", "im", "i"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building image " + args[0])
	},
}

func init() {
	buildCmd.AddCommand(imageCmd)

	var dockerDir string
	currentDir, _ := os.Getwd()
	dockerDirDefault := filepath.Join(currentDir, "dockers")
	imageCmd.Flags().StringVarP(&dockerDir, "docker-dir", "d", dockerDirDefault, "Directory containing image definition folders")
}
