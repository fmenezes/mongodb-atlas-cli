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

package policies

import (
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/flag"
	mocks "github.com/fmenezes/mongodb-atlas-cli/v2/internal/mocks/atlas"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/test"
	"github.com/golang/mock/gomock"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{
			flag.ProjectID,
			flag.Output,
		},
	)
}

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyDescriber(ctrl)

	opts := &DescribeOpts{
		store: mockStore,
	}
	expected := &atlasv2.DataProtectionSettings20231001{}

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	test.VerifyOutputTemplate(t, describePoliciesTemplate, expected)
}
