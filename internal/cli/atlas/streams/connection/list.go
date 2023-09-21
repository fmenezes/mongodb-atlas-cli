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

// This code was autogenerated at 2023-07-03T04:46:41Z. Note: Manual updates are allowed, but may be overwritten.

package connection

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

var listTemplate = `NAME	TYPE	SERVERS{{range .Results}}
{{.Name}}	{{.Type}}	{{if .ClusterName}}{{.ClusterName}}{{else if .BootstrapServers}}{{.BootstrapServers}}{{else}}nil{{end}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	streamsInstance string
	store           store.StreamsConnectionLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.StreamsConnections(opts.ConfigProjectID(), opts.streamsInstance)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas streams connection list --instance <instanceName> [--projectId projectId].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Returns all Atlas Stream Processing connections from your Atlas Stream Processing instance.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.ExactArgs(0),
		Example: fmt.Sprintf(`# list all connections within ExampleInstance:
%[1]s streams connection list --instance ExampleInstance
`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	_ = cmd.MarkFlagRequired(flag.Instance)

	return cmd
}
