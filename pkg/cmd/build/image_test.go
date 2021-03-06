// Copyright 2020 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package build

import (
	"testing"

	"github.com/okteto/okteto/pkg/model"
	"github.com/okteto/okteto/pkg/okteto"
)

func Test_GetRepoNameWithoutTag(t *testing.T) {
	var tests = []struct {
		name     string
		image    string
		expected string
	}{
		{
			name:     "official-with-tag",
			image:    "ubuntu:2",
			expected: "ubuntu",
		},
		{
			name:     "official-without-tag",
			image:    "ubuntu",
			expected: "ubuntu",
		},
		{
			name:     "repo-with-tag",
			image:    "test/ubuntu:2",
			expected: "test/ubuntu",
		},
		{
			name:     "repo-without-tag",
			image:    "test/ubuntu",
			expected: "test/ubuntu",
		},
		{
			name:     "registry-with-tag",
			image:    "registry/gitlab.com/test/ubuntu:2",
			expected: "registry/gitlab.com/test/ubuntu",
		},
		{
			name:     "registry-without-tag",
			image:    "registry/gitlab.com/test/ubuntu",
			expected: "registry/gitlab.com/test/ubuntu",
		},
		{
			name:     "localhost-with-tag",
			image:    "localhost:5000/test/ubuntu:2",
			expected: "localhost:5000/test/ubuntu",
		},
		{
			name:     "registry-without-tag",
			image:    "localhost:5000/test/ubuntu",
			expected: "localhost:5000/test/ubuntu",
		},
		{
			name:     "sha256",
			image:    "pchico83/test@sha256:e78ad0d316485b7dbffa944a92b29ea4fa26d53c63054605c4fb7a8b787a673c",
			expected: "pchico83/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetRepoNameWithoutTag(tt.image)
			if tt.expected != result {
				t.Errorf("expected %s got %s in test %s", tt.expected, result, tt.name)
			}
		})
	}
}

func Test_GetDevImageTag(t *testing.T) {
	var tests = []struct {
		name                string
		dev                 *model.Dev
		imageTag            string
		imageFromDeployment string
		oktetoRegistryURL   string
		expected            string
	}{
		{
			name:                "imageTag-not-in-okteto",
			dev:                 &model.Dev{Name: "dev", Namespace: "ns"},
			imageTag:            "imageTag",
			imageFromDeployment: "",
			oktetoRegistryURL:   "",
			expected:            "imageTag",
		},
		{
			name:                "imageTag-in-okteto",
			dev:                 &model.Dev{Name: "dev", Namespace: "ns"},
			imageTag:            "imageTag",
			imageFromDeployment: "",
			oktetoRegistryURL:   okteto.CloudRegistryURL,
			expected:            "imageTag",
		},
		{
			name:                "okteto",
			dev:                 &model.Dev{Name: "dev", Namespace: "ns"},
			imageTag:            "",
			imageFromDeployment: "",
			oktetoRegistryURL:   okteto.CloudRegistryURL,
			expected:            "registry.cloud.okteto.net/ns/dev:okteto",
		},
		{
			name:                "not-in-okteto",
			dev:                 &model.Dev{Name: "dev", Namespace: "ns"},
			imageTag:            "",
			imageFromDeployment: "okteto/test:2",
			oktetoRegistryURL:   "",
			expected:            "okteto/test:okteto",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDevImageTag(tt.dev, tt.imageTag, tt.imageFromDeployment, tt.oktetoRegistryURL)
			if tt.expected != result {
				t.Errorf("expected %s got %s in test %s", tt.expected, result, tt.name)
			}
		})
	}
}
