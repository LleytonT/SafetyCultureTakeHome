package folders

import (
	"encoding/base64"
	"errors"

	"strconv"

	"github.com/gofrs/uuid"
)

const MaxLimit = int(^uint(0) >> 1) // Maximum limit for pagination

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	if req.Limit <= 0 {
		req.Limit = MaxLimit
	}

	folders, err := FetchAllFoldersByOrgID(req.OrgID, req.Limit, req.Token)
	if err != nil {
		return nil, err
	}

	nextToken := ""
	if len(folders) == req.Limit {
		nextToken = generateToken(req.Token, req.Limit)
	}

	return &FetchFolderResponse{Folders: folders, Token: nextToken}, nil
}

// Filters through all folders and returns only those that belong to the organisation.
func FetchAllFoldersByOrgID(orgID uuid.UUID, limit int, token string) ([]*Folder, error) {
	// An organisation ID must be a valid UUID.
	if orgID == uuid.Nil {
		return nil, errors.New("invalid OrgID")
	}

	folders := GetSampleData()

	// Filters through all folders and returns only those that belong to the organisation.
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	startIndex := 0
	if token != "" {
		startIndex, _ = DecodeToken(token)
	}

	endIndex := startIndex + limit
	if endIndex > len(resFolder) {
		endIndex = len(resFolder)
	}

	return resFolder[startIndex:endIndex], nil
}

// Can add encyption and decryption functions here to encode and decode the token
// Thought this would be enough but if you want me to add it, I can do it.

// EncodeToken encodes an integer token into a Base64 string.
func EncodeToken(token int) string {
	tokenStr := strconv.Itoa(token)
	return base64.StdEncoding.EncodeToString([]byte(tokenStr))
}

// DecodeToken decodes a Base64 string back into an integer token.
func DecodeToken(encodedToken string) (int, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedToken)
	if err != nil {
		return 0, err
	}
	tokenStr := string(decodedBytes)
	return strconv.Atoi(tokenStr)
}

func generateToken(currentToken string, limit int) string {
	currentIndex, _ := DecodeToken(currentToken)
	nextIndex := currentIndex + limit
	return EncodeToken(nextIndex)
}

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
