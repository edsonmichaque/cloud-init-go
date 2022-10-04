// Copyright 2022 Edson Michaque
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloudinit

import (
	"github.com/spf13/cobra"
)

func CmdStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use: "status",
	}

	cmd.Flags().Bool("wait", false, "Wait")
	cmd.Flags().Bool("long", false, "Long")
	cmd.Flags().String("format", "", "Format")

	return cmd
}
