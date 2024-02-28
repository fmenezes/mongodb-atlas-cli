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

//go:build e2e || (atlas && generic)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/v2/test/e2e"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func TestAuditing(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("auditing")
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			auditingEntity,
			"describe",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var setting *atlasv2.AuditLog
		require.NoError(t, json.Unmarshal(resp, &setting), string(resp))
	})
}
