package folders

/* import (
	"errors"

	"github.com/gofrs/uuid"
) */

/*
// Returns all folders for a given organisation.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// IMRPOVEMENT: Unused variables should be removed.
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	// Initialises a new slice of Folder type.
	f := []Folder{}
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	// Appends the folders gotten from r to the slice f.
	for k, v := range r {
		f = append(f, *v)
	}
	// Declares a slice of pointers to Folder structs
	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}
	// IMRPOVEMENT: Above block of code can just be one line but also add error checking.

	// Initialises a new FetchFolderResponse struct and stores the pointers of folders from the slice fp.
	// IMPROVEMENT: can just be written as ffr := &FetchFolderResponse{Folders: fp}
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
} */

/* func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	return &FetchFolderResponse{Folders: folders}, nil
}

// Filters through all folders and returns only those that belong to the organisation.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
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
	return resFolder, nil
}
*/
