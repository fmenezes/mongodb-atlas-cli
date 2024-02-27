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

package datalake

import (
	"context"
	"fmt"

	"github.com/fmenezes/mongodb-atlas-cli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/internal/cli/require"
	"github.com/fmenezes/mongodb-atlas-cli/internal/config"
	"github.com/fmenezes/mongodb-atlas-cli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/internal/store"
	"github.com/fmenezes/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.DataLakeDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteDataLake, opts.ConfigProjectID())
}

// mongocli atlas datalake(s) delete <name> --projectId projectId.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Data Lake '%s' deleted\n", "Data Lake not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <name>",
		Aliases: []string{"rm"},
		Short:   "Remove a federated database instance from your project.",
		Long:    `To learn more about Atlas Data Federation (previously named Atlas Data Lake), see https://www.mongodb.com/docs/atlas/data-federation/overview/.`,
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the federated database instance to delete.",
			"output":   opts.SuccessMessage(),
		},
		Example: fmt.Sprintf(`  # Remove the federated database instance named myFDI in the project with the ID 5e2211c17a3e5a48f5497de3:
  %s dataLakes delete myFDI --projectId 5e2211c17a3e5a48f5497de3`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	if config.ToolName == config.AtlasCLI {
		cmd.Deprecated = "Please use 'atlas datafederation delete'"
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
