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

//go:build unit

package projects

import (
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/flag"
	mocks "github.com/fmenezes/mongodb-atlas-cli/v2/internal/mocks/atlas"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/pointer"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/test"
	"github.com/golang/mock/gomock"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectCreator(ctrl)

	opts := CreateOpts{}
	expected := &atlasv2.Group{
		Tags: &[]atlasv2.ResourceTag{
			{Key: pointer.Get("environment"), Value: pointer.Get("unit-testing")},
			{Key: pointer.Get("production"), Value: pointer.Get("false")},
		},
	}

	createOpts := &CreateOpts{
		store: mockStore,
		name:  "ProjectBar",
		tag: map[string]string{
			"environment": "unit-testing",
			"production":  "false",
		},
	}
	createOpts.OrgID = "5a0a1e7e0f2912c554080adc"
	params := &atlasv2.CreateProjectApiParams{
		ProjectOwnerId: &opts.projectOwnerID,
		Group:          createOpts.newCreateProjectGroup(),
	}

	mockStore.
		EXPECT().
		CreateProject(params).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateTemplate(t *testing.T) {
	test.VerifyOutputTemplate(t, atlasCreateTemplate, atlasv2.Group{})
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.OrgID, flag.OwnerID, flag.GovCloudRegionsOnly, flag.WithoutDefaultAlertSettings},
	)
}
