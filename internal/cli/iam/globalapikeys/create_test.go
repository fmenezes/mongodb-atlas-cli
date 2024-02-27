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

//go:build unit

package globalapikeys

import (
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/internal/mocks"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockGlobalAPIKeyCreator(ctrl)

	expected := &opsmngr.APIKey{
		ID: "1",
	}

	createOpts := &CreateOpts{
		store: mockStore,
		roles: []string{"ORG_OWNER"},
	}

	mockStore.
		EXPECT().
		CreateGlobalAPIKey(createOpts.newAPIKeyInput()).
		Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
