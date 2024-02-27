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

package blockstore

import (
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/internal/mocks"
	"github.com/fmenezes/mongodb-atlas-cli/internal/test"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockBlockstoresCreator(ctrl)

	expected := &opsmngr.BackupStore{}

	createOpts := &CreateOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().CreateBlockstore(createOpts.NewBackupStore()).
		Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.Output, flag.Name, flag.SSL, flag.EncryptedCredentials, flag.LoadFactor,
			flag.MaxCapacityGB, flag.Assignment, flag.Label, flag.URI, flag.WriteConcern},
	)
}
