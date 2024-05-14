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

// This code was autogenerated at 2023-07-05T01:21:22Z. Note: Manual updates are allowed, but may be overwritten.

package connection

import (
	"context"
	"fmt"

	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store           store.ConnectionUpdater
	name            string
	filename        string
	streamsInstance string
	fs              afero.Fs
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = "Connection {{.Name}} updated.\n"

func (opts *UpdateOpts) Run() error {
	updateRequest, err := opts.newUpdateRequest()
	if err != nil {
		return err
	}

	r, err := opts.store.UpdateConnection(opts.ConfigProjectID(), opts.streamsInstance, opts.name, updateRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newUpdateRequest() (*atlasv2.StreamsConnection, error) {
	out := atlasv2.NewStreamsConnectionWithDefaults()
	if err := file.Load(opts.fs, opts.filename, out); err != nil {
		return nil, err
	}
	out.Name = &opts.name

	return out, nil
}

// atlas streams connection update <connectionName> [--projectId projectId].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update <connectionName>",
		Short: "Modify the details of the specified connection within your Atlas Stream Processing instance.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"connectionNameDesc": "Name of the connection.",
			"output":             updateTemplate,
		},
		Example: `# update an Atlas Stream Processing connection:
  atlas streams connection update kafkaprod --instance test01 -f kafkaConfig.json
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.StreamsConnectionFilename)
	_ = cmd.MarkFlagFilename(flag.File)

	_ = cmd.MarkFlagRequired(flag.Instance)
	_ = cmd.MarkFlagRequired(flag.File)

	return cmd
}
