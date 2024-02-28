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

package containers

import (
	"context"
	"fmt"

	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/require"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/config"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/store"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.ContainersDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteContainer, opts.ConfigProjectID())
}

// mongocli atlas networking container(s) delete <ID> [--projectId projectId].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Network peering container '%s' deleted\n", "Network peering container not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <containerId>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified network peering container from your project before creating any clusters. Don't run this command if you have clusters in your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"containerIdDesc": "Unique 24-hexadecimal digit string that identifies the network container that you want to remove.",
			"output":          opts.SuccessMessage(),
		},
		Example: fmt.Sprintf(`  # Remove the network peering container with the ID 5e44103f8d614b2f0b6530d8 from the project with the ID 5e2211c17a3e5a48f5497de3:
  %s networking containers delete 5e44103f8d614b2f0b6530d8 --projectId 5e2211c17a3e5a48f5497de3`, cli.ExampleAtlasEntryPoint()),
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

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
