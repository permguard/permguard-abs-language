// Copyright 2024 Nitro Agility S.r.l.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package permcode

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	aztypes "github.com/permguard/permguard-abs-language/pkg/permcode/types"
	azcrypto "github.com/permguard/permguard-core/pkg/extensions/crypto"
)

// TestStringify tests the Stringify function.
func TestMashalingOfPolicies(t *testing.T) {
	tests := []struct {
		Path string
	}{
		{
			Path: "./testdata/mashaling",
		},
	}
	for _, test := range tests {
		testCases, _ := os.ReadDir(test.Path)
		for _, testCase := range testCases {
			testCaseFile := filepath.Join(test.Path, testCase.Name())
			testCaseData, _ := os.ReadFile(testCaseFile)
			var data map[string]any
			json.Unmarshal(testCaseData, &data)
			t.Run(data["testcase"].(string), func(t *testing.T) {
				assert := assert.New(t)
				pm := NewPermCodeManager()

				sanitize := data["sanitize"].(bool)
				validate := data["validate"].(bool)
				optimize := data["optimize"].(bool)

				inputData, _ := json.Marshal(data["input"])
				policyInInfo, _ := pm.UnmarshalClass(inputData, sanitize, validate, optimize)
				policyInData, _ := pm.MarshalClass(policyInInfo.Type, false, false, false)
				policyInDataSha := azcrypto.ComputeSHA256(policyInData)

				outpuData, err := json.Marshal(data["output"])
				policyOutInfo, _ := pm.UnmarshalClass(outpuData, false, false, false)
				policyOutData, _ := pm.MarshalClass(policyOutInfo.Type, false, false, false)
				policyOutDataSha := azcrypto.ComputeSHA256(policyOutData)

				assert.Nil(err)
				assert.Equal(policyInDataSha, policyOutDataSha)
			})
		}
	}
}

// TestMashalingOfPoliciesWithArgumentsErrors tests marshaling of policies with arguments errors.
func TestMashalingOfPoliciesWithArgumentsErrors(t *testing.T) {
	t.Run("marshaling with nil value", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		result, err := pm.MarshalClass(nil, false, false, false)
		assert.NotNil(err)
		assert.Nil(result)
	})
	t.Run("marshaling with invalid json 1", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		jsonStr := `{"id":"1", "color":"red"}`
		jsonBytes := []byte(jsonStr)
		result, err := pm.MarshalClass(jsonBytes, false, false, false)
		assert.NotNil(err)
		assert.Nil(result)
	})
	t.Run("marshaling with invalid json 2", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		jsonStr := `{"syntax":"permguard1", "type":"acpolicy"}`
		jsonBytes := []byte(jsonStr)
		result, err := pm.MarshalClass(jsonBytes, false, false, false)
		assert.NotNil(err)
		assert.Nil(result)
	})
	t.Run("unmarshaling with nil value", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		result, err := pm.UnmarshalClass(nil, false, false, false)
		assert.NotNil(err)
		assert.Nil(result)
	})
	t.Run("marshaling with invalid object type", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		obj := "sorry"
		result, err := pm.MarshalClass(obj, false, false, false)
		assert.NotNil(err)
		assert.Nil(result)
	})
	t.Run("marshaling with invalid policy", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		obj := aztypes.Policy{
			Name: "this is a wr@ng name",
		}
		result, err := pm.MarshalClass(obj, false, false, false)
		assert.NotNil(err)
		assert.Nil(result)
	})
}

// TestMashalingOfPoliciesWithErrors tests the Stringify function.
func TestMashalingOfPoliciesWithErrors(t *testing.T) {
	tests := []struct {
		Path string
	}{
		{
			Path: "./testdata/mashaling-with-errors",
		},
	}
	for _, test := range tests {
		testCases, _ := os.ReadDir(test.Path)
		for _, testCase := range testCases {
			testCaseFile := filepath.Join(test.Path, testCase.Name())
			testCaseData, _ := os.ReadFile(testCaseFile)
			var data map[string]any
			json.Unmarshal(testCaseData, &data)
			t.Run(data["testcase"].(string), func(t *testing.T) {
				assert := assert.New(t)
				pm := NewPermCodeManager()

				sanitize := data["sanitize"].(bool)
				validate := data["validate"].(bool)
				optimize := data["optimize"].(bool)

				inputData, _ := json.Marshal(data["input"])
				policyInInfo, err := pm.UnmarshalClass(inputData, sanitize, validate, optimize)

				assert.NotNil(err)
				assert.Nil(policyInInfo)
			})
		}
	}
}