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
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Build components of a Kubernetes stack (Docker images, Helm charts ...)",
	Aliases: []string{"bld", "bl", "b"},
	Long:    "Build components of a Kubernetes stack (Docker images, Helm charts ...)",
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
