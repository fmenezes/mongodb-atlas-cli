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

package events

import (
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/flag"
	mocks "github.com/fmenezes/mongodb-atlas-cli/v2/internal/mocks/atlas"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/test"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func Test_projectListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectEventLister(ctrl)

	expected := &admin.GroupPaginatedEvent{}
	listOpts := &projectListOpts{
		store: mockStore,
	}
	listOpts.ProjectID = "1"
	anyMock := gomock.Any()
	mockStore.
		EXPECT().ProjectEvents(anyMock).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestProjectListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ProjectListBuilder(),
		0,
		[]string{
			flag.Limit,
			flag.Page,
			flag.Output,
			flag.ProjectID,
			flag.TypeFlag,
			flag.MaxDate,
			flag.MinDate,
		},
	)
}

func TestProjectsBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ProjectsBuilder(),
		1,
		[]string{},
	)
}
