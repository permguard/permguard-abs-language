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

package objects

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestObjectManager tests the functions of ObjectManager.
func TestObjectManager(t *testing.T) {
	objectManager, _ := NewObjectManager()

	t.Run("Test CreateCommitObject and GetObjectInfo", func(t *testing.T) {
		assert := assert.New(t)
		commit := &Commit{
			tree:   "3b18e17a0e8664d3dffab99ebf6d730ddc6e8649",
			parent: "a1b2c3d4e5f678901234567890abcdef12345678",
			info: CommitInfo{
				author: "Nicola Gallo",
				authorTimestamp: time.Unix(1628704800, 0),
				committer: "Nicola Gallo",
				committerTimestamp: time.Unix(1628704800, 0),
			},
			message: "Initial commit",
		}

		// Create commit object
		commitObj, err := objectManager.CreateCommitObject(commit)
		assert.Nil(err)
		assert.NotEmpty(commitObj.oid, "OID should not be empty")
		assert.NotEmpty(commitObj.content, "Commit content should not be empty")

		// Get object info
		objectInfo, err := objectManager.GetObjectInfo(commitObj)
		assert.Nil(err)
		assert.Equal(ObjectTypeCommit, objectInfo.otype, "Expected commit type")
		assert.NotNil(objectInfo.instance, "Commit instance should not be nil")

		// Cast to commit and validate fields
		retrievedCommit := objectInfo.instance.(*Commit)
		assert.Equal(commit.tree, retrievedCommit.tree, "Tree mismatch")
		assert.Equal(commit.parent, retrievedCommit.parent, "Parents mismatch")
		assert.Equal(commit.info.author, retrievedCommit.info.author, "Author mismatch")
		assert.Equal(commit.info.authorTimestamp.Unix(), retrievedCommit.info.authorTimestamp.Unix(), "Author timestamp mismatch")
		assert.Equal(commit.info.committer, retrievedCommit.info.committer, "Committer mismatch")
		assert.Equal(commit.info.committerTimestamp.Unix(), retrievedCommit.info.committerTimestamp.Unix(), "Committer timestamp mismatch")
		assert.Equal(commit.message, retrievedCommit.message, "Message mismatch")
	})

	// Test for CreateTreeObject and GetObjectInfo
	t.Run("Test CreateTreeObject and GetObjectInfo", func(t *testing.T) {
		assert := assert.New(t)
		tree := &Tree{
			entries: []TreeEntry{
				{otype: "blob", oid: "6eb715b073c6b28e03715129e03a0d52c8e21b73", oname: "name1"},
				{otype: "blob", oid: "a7fdb22705a5e6145b6a8b1fa947825c5e97a51c", oname: "name2"},
				{otype: "tree", oid: "a7fdb33705a5e6145b6a8b1fa947825c5e97a51c", oname: "name3"},
			},
		}

		// Create tree object
		treeObj, err := objectManager.CreateTreeObject(tree)
		assert.Nil(err)
		assert.NotEmpty(treeObj.oid, "OID should not be empty")
		assert.NotEmpty(treeObj.content, "Tree content should not be empty")

		// Get object info
		objectInfo, err := objectManager.GetObjectInfo(treeObj)
		assert.Nil(err)
		assert.Equal(ObjectTypeTree, objectInfo.otype, "Expected tree type")
		assert.NotNil(objectInfo.instance, "Tree instance should not be nil")

		// Cast to tree and validate fields
		retrievedTree := objectInfo.instance.(*Tree)
		assert.Equal(len(tree.entries), len(retrievedTree.entries), "Entries length mismatch")
		for i, entry := range tree.entries {
			assert.Equal(entry.otype, retrievedTree.entries[i].otype, "Type mismatch for entry %d", i)
			assert.Equal(entry.oid, retrievedTree.entries[i].oid, "OID mismatch for entry %d", i)
			assert.Equal(entry.oname, retrievedTree.entries[i].oname, "Name mismatch for entry %d", i)
		}
	})

	// Test for CreateBlobObject and GetObjectInfo (new test for blob type)
	t.Run("Test CreateBlobObject and GetObjectInfo", func(t *testing.T) {
		assert := assert.New(t)
		blobData := []byte("This is the content of the blob object")

		// Create blob object
		blobObj, err := objectManager.CreateBlobObject(blobData)
		assert.Nil(err)
		assert.NotEmpty(blobObj.oid, "OID should not be empty")
		assert.NotEmpty(blobObj.content, "Blob content should not be empty")

		// Get object info
		objectInfo, err := objectManager.GetObjectInfo(blobObj)
		assert.Nil(err)
		assert.Equal(ObjectTypeBlob, objectInfo.otype, "Expected blob type")
		assert.NotNil(objectInfo.instance, "Blob instance should not be nil")

		// Validate the content of the blob
		retrievedBlob := objectInfo.instance.([]byte)
		assert.Equal(blobData, retrievedBlob, "Blob content mismatch")
	})

	// Test for invalid data
	t.Run("Test invalid object", func(t *testing.T) {
		assert := assert.New(t)
		invalidObj := &Object{content: []byte{}}
		_, err := objectManager.GetObjectInfo(invalidObj)
		assert.NotNil(err, "Expected error for empty object content")

		// Test for incorrect object type
		invalidObj.content = []byte("xx 12\000some content")
		_, err = objectManager.GetObjectInfo(invalidObj)
		assert.NotNil(err, "Expected error for wrong object type")
		assert.Contains(err.Error(), "objects: invalid object format: no NUL separator found", "Expected objects: unsupported object type ")
	})
}
