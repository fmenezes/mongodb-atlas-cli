// Copyright 2020 MongoDB Inc
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

package backup

import (
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/backup/compliancepolicy"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/backup/exports"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/backup/restores"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/backup/schedule"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/backup/snapshots"
	"github.com/spf13/cobra"
)

func baseCommand() *cobra.Command {
	const use = "backups"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage cloud backups for your project.",
	}

	return cmd
}

func Builder() *cobra.Command {
	cmd := baseCommand()

	cmd.AddCommand(
		snapshots.Builder(),
		restores.Builder(),
		exports.Builder(),
		schedule.Builder(),
		compliancepolicy.Builder(),
	)

	return cmd
}
