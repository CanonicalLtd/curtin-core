// -*- Mode: Go; indent-tabs-mode: t -*-
// Flashback
// Copyright 2018 Canonical Ltd.  All rights reserved.

package config

import (
	"fmt"
	"io/ioutil"

	"github.com/CanonicalLtd/flashback/audit"
	yaml "gopkg.in/yaml.v2"
)

// Config defines the configuration parameters
type Config struct {
	RestorePartitionLabel  string `yaml:"restore-partition"`
	WritablePartitionLabel string `yaml:"writable-partition"`
	LogFile                string `yaml:"logfile"`
}

const (
	restorePartitionLabel  = "restore"
	writablePartitionLabel = "writable"
)

// Store the stored configuration from the file
var Store Config

// Read parses the yaml config file
func Read(path string) error {
	Store = Config{}

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading config parameters: %v\n", err)
		return err
	}

	err = yaml.Unmarshal(dat, &Store)
	if err != nil {
		fmt.Printf("Error parsing config parameters: %v\n", err)
		return err
	}

	// Default the missing parameters
	setDefaults()

	return nil
}

func setDefaults() {
	if len(Store.LogFile) == 0 {
		audit.Printf("Default the LogFile to `%s`\n", audit.LogFile)
		Store.LogFile = audit.LogFile
	} else {
		// Overwrite the log file path from the config
		audit.LogFile = Store.LogFile
	}
	if len(Store.RestorePartitionLabel) == 0 {
		audit.Printf("Default the RestorePartitionLabel to `%s`\n", restorePartitionLabel)
		Store.RestorePartitionLabel = restorePartitionLabel
	}
	if len(Store.WritablePartitionLabel) == 0 {
		audit.Printf("Default the WritablePartitionLabel to `%s`\n", writablePartitionLabel)
		Store.WritablePartitionLabel = writablePartitionLabel
	}
}
