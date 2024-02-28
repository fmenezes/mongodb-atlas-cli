// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nodes

import (
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/internal/mocks"
	"github.com/fmenezes/mongodb-atlas-cli/internal/test"
	"github.com/golang/mock/gomock"
)

func TestDelete_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSearchNodesDeleter(ctrl)

	deleteOpts := &DeleteOpts{
		store:      mockStore,
		DeleteOpts: cli.NewDeleteOpts(deleteTemplate, "Search node not deleted."),
	}

	const clusterName = "Cluster0"
	deleteOpts.Entry = clusterName
	deleteOpts.ProjectID = "32b6e34b3d91647abb20e111"
	deleteOpts.Confirm = true

	expected := clusterName

	mockStore.
		EXPECT().
		DeleteSearchNodes(deleteOpts.ProjectID, deleteOpts.Entry).
		Return(nil).
		Times(1)

	if err := deleteOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, deleteTemplate, expected)
}

func TestDeleteBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{
			flag.ClusterName,
			flag.ProjectID,
		},
	)
}