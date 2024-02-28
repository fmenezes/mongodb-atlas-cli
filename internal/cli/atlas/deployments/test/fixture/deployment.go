// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package fixture

import (
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/cli/atlas/deployments/options"
	"github.com/fmenezes/mongodb-atlas-cli/v2/internal/mocks"
	"github.com/golang/mock/gomock"
)

type MockDeploymentOpts struct {
	ctrl                      *gomock.Controller
	MockCredentialsGetter     *mocks.MockCredentialsGetter
	MockAtlasClusterListStore *mocks.MockClusterLister
	MockPodman                *mocks.MockClient
	Opts                      *options.DeploymentOpts
}
