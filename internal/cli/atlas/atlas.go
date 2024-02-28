// Copyright 2021 MongoDB Inc
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

package atlas

import (
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/accesslists"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/accesslogs"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/cloudproviders"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/clusters"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/customdbroles"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/customdns"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/datalake"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/dbusers"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/integrations"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/livemigrations"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/logs"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/maintenance"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/metrics"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/networking"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/privateendpoints"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/processes"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/security"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/mongocli/alerts"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/mongocli/backup"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/mongocli/events"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/mongocli/performanceadvisor"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/mongocli/serverless"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/config"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/validate"
	"github.com/spf13/cobra"
)

const (
	Use               = "atlas"
	deprecatedMessage = "There's a new, dedicated Atlas CLI available for Atlas users. Install the Atlas CLI to enjoy the same capabilities and keep getting new features: https://dochub.mongodb.org/core/migrate-to-atlas-cli. Atlas commands for MongoCLI are now deprecated, but you can keep using them for 12 months (until April 30, 2023).\n\n"
)

func Builder() *cobra.Command {
	opts := &cli.RefresherOpts{}
	cmd := &cobra.Command{
		Use:        Use,
		Short:      "MongoDB Atlas operations.",
		Deprecated: deprecatedMessage,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := opts.InitFlow(config.Default())(); err != nil {
				return err
			}
			if err := opts.RefreshAccessToken(cmd.Context()); err != nil {
				return err
			}
			if config.Service() == "" {
				config.SetService(config.CloudService)
			}
			return validate.Credentials()
		},
		Annotations: map[string]string{
			"toc": "true",
		},
	}
	cmd.AddCommand(
		clusters.MongoCLIBuilder(),
		dbusers.Builder(),
		customdbroles.Builder(),
		accesslists.Builder(),
		datalake.Builder(),
		alerts.Builder(),
		backup.Builder(),
		events.Builder(),
		metrics.Builder(),
		performanceadvisor.Builder(),
		logs.MongoCLIBuilder(),
		processes.Builder(),
		privateendpoints.Builder(),
		networking.Builder(),
		security.Builder(),
		integrations.Builder(),
		maintenance.Builder(),
		customdns.Builder(),
		cloudproviders.Builder(),
		serverless.Builder(),
		livemigrations.Builder(),
		accesslogs.Builder(),
	)
	return cmd
}
