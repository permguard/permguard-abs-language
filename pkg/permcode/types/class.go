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

package types

const (
	// ClassTypeSchema is the object type for domain schemas.
	ClassTypeSchema = "schema"
	// ClassTypeObject is the object type for domain objects.
	ClassTypeIDSchema = uint32(1)

	// ClassTypeACPolicy is the object type for an access control policy.
	ClassTypeACPolicy = "acpolicy"
	// ClassTypeIDACPolicy is the object type for an access control policy.
	ClassTypeIDACPolicy = uint32(1)
)

// Class is the base class.
type Class struct {
	SyntaxVersion string `json:"syntax"`
	Type          string `json:"type"`
}

// ClassInfo is the class info.
type ClassInfo struct {
	ID      string
	Type     string
	Instance any
}
