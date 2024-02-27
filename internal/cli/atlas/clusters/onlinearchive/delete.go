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

package onlinearchive

import (
	"context"
	"fmt"

	"github.com/fmenezes/mongodb-atlas-cli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/internal/cli/require"
	"github.com/fmenezes/mongodb-atlas-cli/internal/config"
	"github.com/fmenezes/mongodb-atlas-cli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/internal/store"
	"github.com/fmenezes/mongodb-atlas-cli/internal/usage"
	"github.com/fmenezes/mongodb-atlas-cli/internal/validate"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	clusterName string
	store       store.OnlineArchiveDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteOnlineArchive, opts.ConfigProjectID(), opts.clusterName)
}

// mongocli atlas cluster(s) onlineArchive(s) delete <archiveId> [--clusterName name][--projectId projectId][--force].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Archive '%s' deleted\n", "Archive not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <archiveId>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified online archive from your cluster.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Data Access Admin"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"archiveIdDesc": "Unique identifier of the online archive to delete.",
			"output":        opts.SuccessMessage(),
		},
		Example: fmt.Sprintf(`  # Remove an online archive with the ID 5f189832e26ec075e10c32d3 for the cluster named myCluster:
  %s clusters onlineArchives delete 5f189832e26ec075e10c32d3 --clusterName myCluster`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context())); err != nil {
				return err
			}
			if err := validate.ObjectID(args[0]); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}
