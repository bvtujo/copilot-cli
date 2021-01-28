// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package manifest

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/copilot-cli/internal/pkg/template"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestBuildArgs_UnmarshalYAML(t *testing.T) {
	testCases := map[string]struct {
		inContent []byte

		wantedStruct BuildArgsOrString
		wantedError  error
	}{
		"legacy case: simple build string": {
			inContent: []byte(`build: ./Dockerfile`),

			wantedStruct: BuildArgsOrString{
				BuildString: aws.String("./Dockerfile"),
			},
		},
		"Dockerfile specified in build opts": {
			inContent: []byte(`build:
  dockerfile: path/to/Dockerfile
`),
			wantedStruct: BuildArgsOrString{
				BuildArgs: DockerBuildArgs{
					Dockerfile: aws.String("path/to/Dockerfile"),
				},
				BuildString: nil,
			},
		},
		"Dockerfile context, and args specified in build opts": {
			inContent: []byte(`build:
  dockerfile: path/to/Dockerfile
  args:
    arg1: value1
    bestdog: bowie
  context: path/to/source`),
			wantedStruct: BuildArgsOrString{
				BuildArgs: DockerBuildArgs{
					Dockerfile: aws.String("path/to/Dockerfile"),
					Context:    aws.String("path/to/source"),
					Args: map[string]string{
						"arg1":    "value1",
						"bestdog": "bowie",
					},
				},
				BuildString: nil,
			},
		},
		"Dockerfile with cache from and target build opts": {
			inContent: []byte(`build:
  cache_from:
    - foo/bar:latest
    - foo/bar/baz:1.2.3
  target: foobar`),
			wantedStruct: BuildArgsOrString{
				BuildArgs: DockerBuildArgs{
					Target: aws.String("foobar"),
					CacheFrom: []string{
						"foo/bar:latest",
						"foo/bar/baz:1.2.3",
					},
				},
				BuildString: nil,
			},
		},
		"Error if unmarshalable": {
			inContent: []byte(`build:
  badfield: OH NOES
  otherbadfield: DOUBLE BAD`),
			wantedError: errUnmarshalBuildOpts,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			b := Image{
				Build: BuildArgsOrString{
					BuildString: aws.String("./default"),
				},
			}
			err := yaml.Unmarshal(tc.inContent, &b)
			if tc.wantedError != nil {
				require.EqualError(t, err, tc.wantedError.Error())
			} else {
				require.NoError(t, err)
				// check memberwise dereferenced pointer equality
				require.Equal(t, tc.wantedStruct.BuildString, b.Build.BuildString)
				require.Equal(t, tc.wantedStruct.BuildArgs.Context, b.Build.BuildArgs.Context)
				require.Equal(t, tc.wantedStruct.BuildArgs.Dockerfile, b.Build.BuildArgs.Dockerfile)
				require.Equal(t, tc.wantedStruct.BuildArgs.Args, b.Build.BuildArgs.Args)
				require.Equal(t, tc.wantedStruct.BuildArgs.Target, b.Build.BuildArgs.Target)
				require.Equal(t, tc.wantedStruct.BuildArgs.CacheFrom, b.Build.BuildArgs.CacheFrom)
			}
		})
	}
}

func TestExec_UnmarshalYAML(t *testing.T) {
	testCases := map[string]struct {
		inContent []byte

		wantedStruct ExecuteCommand
		wantedError  error
	}{
		"use default with empty value": {
			inContent: []byte(`exec:
count: 1`),

			wantedStruct: ExecuteCommand{
				Enable: aws.Bool(false),
			},
		},
		"use default without any input": {
			inContent: []byte(`count: 1`),

			wantedStruct: ExecuteCommand{
				Enable: aws.Bool(false),
			},
		},
		"simple enable": {
			inContent: []byte(`exec: true`),

			wantedStruct: ExecuteCommand{
				Enable: aws.Bool(true),
			},
		},
		"with config": {
			inContent: []byte(`exec:
  enable: true`),
			wantedStruct: ExecuteCommand{
				Enable: aws.Bool(false),
				Config: ExecuteCommandConfig{
					Enable: aws.Bool(true),
				},
			},
		},
		"Error if unmarshalable": {
			inContent: []byte(`exec:
  badfield: OH NOES
  otherbadfield: DOUBLE BAD`),
			wantedError: errUnmarshalExec,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			b := TaskConfig{
				ExecuteCommand: ExecuteCommand{
					Enable: aws.Bool(false),
				},
			}
			err := yaml.Unmarshal(tc.inContent, &b)
			if tc.wantedError != nil {
				require.EqualError(t, err, tc.wantedError.Error())
			} else {
				require.NoError(t, err)
				// check memberwise dereferenced pointer equality
				require.Equal(t, tc.wantedStruct.Enable, b.ExecuteCommand.Enable)
				require.Equal(t, tc.wantedStruct.Config, b.ExecuteCommand.Config)
			}
		})
	}
}

func TestBuildConfig(t *testing.T) {
	mockWsRoot := "/root/dir"
	testCases := map[string]struct {
		inBuild     BuildArgsOrString
		wantedBuild DockerBuildArgs
	}{
		"simple case: BuildString path to dockerfile": {
			inBuild: BuildArgsOrString{
				BuildString: aws.String("my/Dockerfile"),
			},
			wantedBuild: DockerBuildArgs{
				Dockerfile: aws.String(filepath.Join(mockWsRoot, "my/Dockerfile")),
				Context:    aws.String(filepath.Join(mockWsRoot, "my")),
			},
		},
		"Different context than dockerfile": {
			inBuild: BuildArgsOrString{
				BuildArgs: DockerBuildArgs{
					Dockerfile: aws.String("build/dockerfile"),
					Context:    aws.String("cmd/main"),
				},
			},
			wantedBuild: DockerBuildArgs{
				Dockerfile: aws.String(filepath.Join(mockWsRoot, "build/dockerfile")),
				Context:    aws.String(filepath.Join(mockWsRoot, "cmd/main")),
			},
		},
		"no dockerfile specified": {
			inBuild: BuildArgsOrString{
				BuildArgs: DockerBuildArgs{
					Context: aws.String("cmd/main"),
				},
			},
			wantedBuild: DockerBuildArgs{
				Dockerfile: aws.String(filepath.Join(mockWsRoot, "cmd", "main", "Dockerfile")),
				Context:    aws.String(filepath.Join(mockWsRoot, "cmd", "main")),
			},
		},
		"no dockerfile or context specified": {
			inBuild: BuildArgsOrString{
				BuildArgs: DockerBuildArgs{
					Args: map[string]string{
						"goodDog": "bowie",
					},
				},
			},
			wantedBuild: DockerBuildArgs{
				Dockerfile: aws.String(filepath.Join(mockWsRoot, "Dockerfile")),
				Context:    aws.String(mockWsRoot),
				Args: map[string]string{
					"goodDog": "bowie",
				},
			},
		},
		"including args": {
			inBuild: BuildArgsOrString{
				BuildArgs: DockerBuildArgs{
					Dockerfile: aws.String("my/Dockerfile"),
					Args: map[string]string{
						"goodDog":  "bowie",
						"badGoose": "HONK",
					},
				},
			},
			wantedBuild: DockerBuildArgs{
				Dockerfile: aws.String(filepath.Join(mockWsRoot, "my/Dockerfile")),
				Context:    aws.String(filepath.Join(mockWsRoot, "my")),
				Args: map[string]string{
					"goodDog":  "bowie",
					"badGoose": "HONK",
				},
			},
		},
		"including build options": {
			inBuild: BuildArgsOrString{
				BuildArgs: DockerBuildArgs{
					Target: aws.String("foobar"),
					CacheFrom: []string{
						"foo/bar:latest",
						"foo/bar/baz:1.2.3",
					},
				},
			},
			wantedBuild: DockerBuildArgs{
				Dockerfile: aws.String(filepath.Join(mockWsRoot, "Dockerfile")),
				Context:    aws.String(mockWsRoot),
				Target:     aws.String("foobar"),
				CacheFrom: []string{
					"foo/bar:latest",
					"foo/bar/baz:1.2.3",
				},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			s := Image{
				Build: tc.inBuild,
			}
			got := s.BuildConfig(mockWsRoot)

			require.Equal(t, tc.wantedBuild, *got)
		})
	}
}

func TestSidecar_Options(t *testing.T) {
	mockImage := aws.String("mockImage")
	mockMap := map[string]string{"foo": "bar"}
	mockCredsParam := aws.String("mockCredsParam")
	testCases := map[string]struct {
		inPort string

		wanted    *template.SidecarOpts
		wantedErr error
	}{
		"invalid port": {
			inPort: "b/a/d/P/o/r/t",

			wantedErr: fmt.Errorf("cannot parse port mapping from b/a/d/P/o/r/t"),
		},
		"good port without protocol": {
			inPort: "2000",

			wanted: &template.SidecarOpts{
				Name:       aws.String("foo"),
				Port:       aws.String("2000"),
				CredsParam: mockCredsParam,
				Image:      mockImage,
				Secrets:    mockMap,
				Variables:  mockMap,
			},
		},
		"good port with protocol": {
			inPort: "2000/udp",

			wanted: &template.SidecarOpts{
				Name:       aws.String("foo"),
				Port:       aws.String("2000"),
				Protocol:   aws.String("udp"),
				CredsParam: mockCredsParam,
				Image:      mockImage,
				Secrets:    mockMap,
				Variables:  mockMap,
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			sidecar := Sidecar{
				Sidecars: map[string]*SidecarConfig{
					"foo": {
						CredsParam: mockCredsParam,
						Image:      mockImage,
						Secrets:    mockMap,
						Variables:  mockMap,
						Port:       aws.String(tc.inPort),
					},
				},
			}
			got, err := sidecar.Options()

			if tc.wantedErr != nil {
				require.EqualError(t, err, tc.wantedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, got[0], tc.wanted)
			}
		})
	}
}

func TestExec_Options(t *testing.T) {
	testCases := map[string]struct {
		inConfig ExecuteCommand

		wanted *template.ExecuteCommandOpts
	}{
		"without exec enabled": {
			inConfig: ExecuteCommand{},
			wanted:   nil,
		},
		"exec enabled": {
			inConfig: ExecuteCommand{
				Enable: aws.Bool(true),
			},
			wanted: &template.ExecuteCommandOpts{},
		},
		"exec enabled with config": {
			inConfig: ExecuteCommand{
				Config: ExecuteCommandConfig{
					Enable: aws.Bool(true),
				},
			},
			wanted: &template.ExecuteCommandOpts{},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			exec := tc.inConfig
			got := exec.Options()

			require.Equal(t, got, tc.wanted)
		})
	}
}
