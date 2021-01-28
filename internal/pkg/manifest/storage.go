// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package manifest

import (
	"errors"

	"gopkg.in/yaml.v3"
)

var errUnmarshalEFSOpts = errors.New("unmarshal efs field into string or map")

// StorageConfig represents the options for external and native storage.
type StorageConfig struct {
	Volumes []Volume `yaml:",flow"`
}

// Volume is an abstraction which merges the MountPoint and Volumes concepts from the ECS Task Definition
type Volume struct {
	Name          string        `yaml:"name"`
	ContainerPath string        `yaml:"path"`
	ReadOnly      bool          `yaml:"read_only"`
	EFS           EFSIDOrConfig `yaml:"efs"`
}

// EFSIDOrConfig is a struct with a custom unmarshaler which can read either a string
// or a more detailed struct with advanced options.
type EFSIDOrConfig struct {
	EFSID     *string
	EFSConfig EFSVolumeConfiguration
}

// UnmarshalYAML overrides the default YAML unmarshaling logic for the EFSIDOrConfig
// struct, allowing it to perform more complex unmarshaling behavior.
// This method implements the yaml.Unmarshaler (v2) interface.
func (e *EFSIDOrConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := unmarshal(&e.EFSConfig); err != nil {
		switch err.(type) {
		case *yaml.TypeError:
			break
		default:
			return err
		}
	}

	if !e.EFSConfig.isEmpty() {
		// Unmarshaled successfully to e.EFSConfig; unset e.EFSID and return.
		e.EFSID = nil
		return nil
	}

	if err := unmarshal(&e.EFSID); err != nil {
		return errUnmarshalEFSOpts
	}
	return nil
}

// EFSVolumeConfiguration holds options which tell ECS how to reach out to the EFS filesystem.
type EFSVolumeConfiguration struct {
	FileSystemID      *string             `yaml:"filesystem_id"`      // Required.
	RootDirectory     *string             `yaml:"root_directory"`     // Default "/"
	TransitEncryption bool                `yaml:"transit_encryption"` // Default true
	AuthConfig        AuthorizationConfig `yaml:"authorization_config"`
}

func (e *EFSVolumeConfiguration) isEmpty() bool {
	if e.FileSystemID == nil && e.RootDirectory == nil && !e.TransitEncryption && e.AuthConfig.isEmpty() {
		return true
	}
	return false
}

// AuthorizationConfig holds options relating to access points and IAM authorization.
type AuthorizationConfig struct {
	IAM           bool    `yaml:"iam"`             // Default true
	AccessPointID *string `yaml:"access_point_id"` // Default ""
}

func (a *AuthorizationConfig) isEmpty() bool {
	if !a.IAM && a.AccessPointID == nil {
		return true
	}
	return false
}