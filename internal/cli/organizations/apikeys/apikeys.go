// Copyright 2023 MongoDB Inc
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

package apikeys

import (
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli/organizations/apikeys/accesslists"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "apiKeys"
	var cmd = &cobra.Command{
		Use:     use,
		Short:   "Organization API Keys operations.",
		Aliases: cli.GenerateAliases(use),
	}
	cmd.AddCommand(
		ListBuilder(),
		DescribeBuilder(),
		UpdateBuilder(),
		CreateBuilder(),
		DeleteBuilder(),
		accesslists.Builder(),
	)
	return cmd
}
