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

//go:build unit

package teams

import (
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func TestRenameBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		RenameBuilder(),
		0,
		[]string{flag.TeamID, flag.OrgID},
	)
}

func Test_renameOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockTeamRenamer(ctrl)

	expected := &atlasv2.TeamResponse{
		Name: pointer.Get("test"),
	}

	opts := &renameOpts{
		store: mockStore,
		name:  "test",
	}

	mockStore.
		EXPECT().
		RenameTeam(opts.OrgID, opts.teamID, opts.newTeam()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, opts.Run())
}
