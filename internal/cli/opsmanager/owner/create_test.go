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

package owner

import (
	"bytes"
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestManagerOwnerCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOwnerCreator(ctrl)

	email := "test@test.com"
	password := "Passw0rd!"
	firstName := "Testy"
	lastName := "McTestyson"
	expected := &opsmngr.CreateUserResponse{}

	t.Run("no access list", func(t *testing.T) {
		createOpts := &CreateOpts{
			email:     email,
			password:  password,
			firstName: firstName,
			lastName:  lastName,
			store:     mockStore,
		}

		mockStore.
			EXPECT().
			CreateOwner(createOpts.newOwner(), createOpts.accessIps).Return(expected, nil).
			Times(1)

		if err := createOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("with access list", func(t *testing.T) {
		createOpts := &CreateOpts{
			email:     email,
			password:  password,
			firstName: firstName,
			lastName:  lastName,
			store:     mockStore,
			accessIps: []string{"192.168.0.1"},
		}

		mockStore.
			EXPECT().
			CreateOwner(createOpts.newOwner(), createOpts.accessIps).Return(expected, nil).
			Times(1)

		if err := createOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestManagerOwnerCreateWithPasswordStdin(t *testing.T) {
	password := "P4ssw0rd!"

	createOpts := &CreateOpts{}

	require.NoError(t, createOpts.InitInput(bytes.NewReader([]byte(password)))())
	require.NoError(t, createOpts.Prompt())

	if u := createOpts.newOwner(); u.Password != password {
		t.Fatalf("failed to set password from input stream")
	}
}
