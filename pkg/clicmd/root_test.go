// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package clicmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
		helpFlag bool
	}{
		{
			name:     "normal",
			args:     []string{"init", "-abc", "123", "--foo", "--bar", ""},
			expected: "init",
			helpFlag: false,
		},
		{
			name:     "flags before command",
			args:     []string{"-abc", "123", "--foo", "--empty", "", "--bar", "arg", "init"},
			expected: "init",
			helpFlag: false,
		},
		{
			name:     "no command",
			args:     []string{"-abc", "123", "--foo", "--empty", "", "--bar", "arg"},
			expected: "",
			helpFlag: true,
		},
		{
			name:     "help flag",
			args:     []string{"init", "arg", "--help", "-h"},
			expected: "init",
			helpFlag: true,
		},
		{
			name:     "help command",
			args:     []string{"help", "-abc", "--foo", "--bar", ""},
			expected: "help",
			helpFlag: false,
		},
	}

	for _, tc := range tests {
		cmd, hasHelpFlag, err := parseCommand(tc.args)
		require.NoError(t, err)
		assert.Equal(t, tc.expected, cmd, tc.name)
		assert.Equal(t, tc.helpFlag, hasHelpFlag, tc.name)

	}
}
