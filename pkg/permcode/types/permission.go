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

// Permission is the permission.
type Permission struct {
	Class
	Name	string   `json:"name"`
	Permit	[]string `json:"permit"`
	Forbid	[]string `json:"forbid"`
}

// PermissionInfo is the permission info.
type PermissionInfo struct {
	SID        string
	Permission *Permission
}
