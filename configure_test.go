// © 2025 Platform Engineering Labs Inc.
//
// SPDX-License-Identifier: FSL-1.1-ALv2

//go:build unit

package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPlugin_Configure(t *testing.T) {
	p := &Plugin{}

	config := json.RawMessage(`{"defaultTimeoutSeconds": 60, "maxConcurrentConnections": 10}`)
	err := p.Configure(config)
	require.NoError(t, err)
	assert.Equal(t, 60, p.defaultTimeoutSeconds)
	assert.Equal(t, 10, p.maxConcurrentConnections)
}

func TestPlugin_Configure_Defaults(t *testing.T) {
	p := &Plugin{}

	// Empty config — fields stay at zero values
	config := json.RawMessage(`{}`)
	err := p.Configure(config)
	require.NoError(t, err)
	assert.Equal(t, 0, p.defaultTimeoutSeconds)
	assert.Equal(t, 0, p.maxConcurrentConnections)
}

func TestPlugin_Configure_InvalidJSON(t *testing.T) {
	p := &Plugin{}

	config := json.RawMessage(`{invalid}`)
	err := p.Configure(config)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid plugin config")
}
