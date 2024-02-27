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

package logs

import (
	"context"

	"github.com/fmenezes/mongodb-atlas-cli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/internal/cli/require"
	"github.com/fmenezes/mongodb-atlas-cli/internal/config"
	"github.com/fmenezes/mongodb-atlas-cli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/internal/store"
	"github.com/fmenezes/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

type JobsDownloadOpts struct {
	cli.GlobalOpts
	cli.DownloaderOpts
	id    string
	store store.LogJobsDownloader
}

func (opts *JobsDownloadOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *JobsDownloadOpts) Run() error {
	out, err := opts.NewWriteCloser()
	if err != nil {
		return err
	}

	if err := opts.store.DownloadLogJob(opts.ConfigProjectID(), opts.id, out); err != nil {
		_ = opts.OnError(out)
		return err
	}

	return out.Close()
}

// mongocli om logs jobs download <ID> [--out out] [--projectId projectId].
func JobsDownloadOptsBuilder() *cobra.Command {
	opts := &JobsDownloadOpts{}
	opts.Fs = afero.NewOsFs()
	cmd := &cobra.Command{
		Use:   "download <ID>",
		Short: "Download logs generated by a log collection job.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"IDDesc": "Log job identifier.",
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context()))
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.Out, flag.Out, "", usage.LogOut)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().BoolVar(&opts.Force, flag.Force, false, usage.ForceFile)

	_ = cmd.MarkFlagRequired(flag.Out)
	_ = cmd.MarkFlagFilename(flag.Out)

	return cmd
}
