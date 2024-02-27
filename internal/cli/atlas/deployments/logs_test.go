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

package deployments

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"strings"
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/fmenezes/mongodb-atlas-cli/internal/cli/atlas/deployments/test/fixture"
	"github.com/fmenezes/mongodb-atlas-cli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/internal/mocks"
	"github.com/fmenezes/mongodb-atlas-cli/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
)

func TestLogsBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		LogsBuilder(),
		0,
		[]string{
			flag.DeploymentName,
			flag.Output,
		},
	)
}

func TestLogs_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	buf := new(bytes.Buffer)
	expectedLocalDeployment := "localDeployment"
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)
	mockPodman := deploymentTest.MockPodman

	downloadOpts := &DownloadOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymentTest.LocalMockFlow(ctx)

	mockPodman.
		EXPECT().
		ContainerLogs(ctx, options.MongodHostnamePrefix+"-"+expectedLocalDeployment).
		Return([]string{"log1", "log2"}, nil).
		Times(1)

	if err := downloadOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	expectedLogs := "log1\nlog2\n"

	if !strings.Contains(buf.String(), expectedLogs) {
		t.Fatalf("Run() expected output: %s, got: %s", expectedLogs, buf.String())
	}
}

func TestLogs_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	atlasDeployment := "localDeployment1"
	mockStore := mocks.NewMockLogsDownloader(ctrl)
	deploymentTest := fixture.NewMockAtlasDeploymentOpts(ctrl, atlasDeployment)

	downloadOpts := &DownloadOpts{
		GlobalOpts: cli.GlobalOpts{ProjectID: "ProjectID"},
		DownloaderOpts: cli.DownloaderOpts{
			Out: "out",
		},
		DeploymentOpts: *deploymentTest.Opts,
		downloadStore:  mockStore,
		host:           "test",
		name:           "mongodb.gz",
	}

	downloadOpts.Fs = afero.NewMemMapFs()
	deploymentTest.CommonAtlasMocks(downloadOpts.ProjectID)

	b, _ := base64.RawStdEncoding.DecodeString("H4sIAAAAAAAA/8pIzcnJVyjPL8pJAQQAAP//hRFKDQsAAAA") // "hello world" gzipped
	mockStore.
		EXPECT().
		DownloadLog(downloadOpts.newHostLogsParams()).
		Return(io.NopCloser(bytes.NewReader(b)), nil).
		Times(1)

	if err := downloadOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
