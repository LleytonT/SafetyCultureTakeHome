package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("valid OrgId with matching folders", func(t *testing.T) {
		orgID := uuid.FromStringOrNil(folders.DefaultOrgID)

		req := &folders.FetchFolderRequest{OrgID: orgID}

		resp, err := folders.GetAllFolders(req)
		assert.NoError(t, err)
		assert.Len(t, resp.Folders, 666)
	})

	t.Run("valid OrgId with no matching folders", func(t *testing.T) {
		orgID := uuid.Must(uuid.NewV4())

		req := &folders.FetchFolderRequest{OrgID: orgID}

		resp, err := folders.GetAllFolders(req)
		assert.NoError(t, err)
		assert.Len(t, resp.Folders, 0)
	})

	t.Run("invalid OrgId", func(t *testing.T) {
		req := &folders.FetchFolderRequest{OrgID: uuid.Nil}

		resp, err := folders.GetAllFolders(req)
		assert.Nil(t, resp)
		assert.Error(t, err)
	})

	t.Run("pagination with limit", func(t *testing.T) {
		orgID := uuid.FromStringOrNil(folders.DefaultOrgID)

		req1 := &folders.FetchFolderRequest{OrgID: orgID, Limit: 1}
		req2 := &folders.FetchFolderRequest{OrgID: orgID, Limit: 10}

		resp1, err1 := folders.GetAllFolders(req1)
		assert.NoError(t, err1)
		assert.Len(t, resp1.Folders, 1)
		assert.NotEmpty(t, resp1.Token)

		resp2, err2 := folders.GetAllFolders(req2)
		assert.NoError(t, err2)
		assert.Len(t, resp2.Folders, 10)
		assert.NotEmpty(t, resp2.Token)

		// Check if the first 10 folders are the same
		for i := 1; i < 10; i++ {
			req1.Token = resp1.Token
			resp1, err1 = folders.GetAllFolders(req1)
			assert.NoError(t, err1)
			assert.Equal(t, resp2.Folders[i:i+1], resp1.Folders)
		}
	})

}
