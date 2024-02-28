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

package copyprotection

import (
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli"
	"github.com/spf13/cobra"
)

const (
	active = "ACTIVE"
)

func Builder() *cobra.Command {
	use := "copyProtection"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage copy protection of the backup compliance policy for your project. Learn more: https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/#-optional--keep-all-snapshots-when-removing-additional-snapshot-regions.",
	}

	cmd.AddCommand(
		EnableBuilder(),
		DisableBuilder(),
	)

	return cmd
}
