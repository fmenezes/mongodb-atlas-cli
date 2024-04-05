// Copyright 2022 MongoDB Inc
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

package jobs

import (
	"context"
	"fmt"

	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	clusterName string
	exportID    string
	store       store.ExportJobsDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `ID	EXPORT BUCKET ID	STATE	SNAPSHOT ID
{{.Id}}	{{.ExportBucketId}}	{{.State}}	{{.SnapshotId}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.ExportJob(opts.ConfigProjectID(), opts.clusterName, opts.exportID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas backup(s) export(s) job(s) describe --clusterName <clusterName> --exportID <exportID> [--projectID <projectID>].
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"get"},
		Short:   "Return one cloud backup export job for your project, cluster and job.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output": describeTemplate,
		},
		Example: `  # Return the details for the continuous backup export job with the ID 5df90590f10fab5e33de2305 for the cluster named Cluster0:
  atlas backup exports jobs describe --clusterName Cluster0 --exportID 5df90590f10fab5e33de2305`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.exportID, flag.ExportID, "", usage.ExportID)
	cmd.Flags().StringVar(&opts.exportID, flag.BucketID, "", usage.ExportID)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	cmd.MarkFlagsMutuallyExclusive(flag.BucketID, flag.ExportID)
	_ = cmd.Flags().MarkDeprecated(flag.BucketID, fmt.Sprintf("please use --%s instead", flag.ExportID))

	return cmd
}
