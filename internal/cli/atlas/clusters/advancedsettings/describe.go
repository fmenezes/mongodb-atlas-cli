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

package advancedsettings

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

var describeTemplate = `MINIMUM TLS	JAVASCRIPT ENABLED	OPLOG SIZE MB
{{.MinimumEnabledTlsProtocol}}	{{.JavascriptEnabled}}	{{if .GetOplogSizeMB}}{{.GetOplogSizeMB}} {{else}}N/A{{end}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name  string
	store store.AtlasClusterConfigurationOptionsDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.AtlasClusterConfigurationOptions(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas cluster(s) advancedSettings describe <clusterName>  --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe <clusterName>",
		Aliases: []string{"get"},
		Short:   "Retrieve advanced configuration settings for one cluster.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the Atlas cluster for which you want to retrieve configuration settings.",
		},
		Example: fmt.Sprintf(`  %s clusters advancedSettings describe Cluster0`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
