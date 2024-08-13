package jamfpro

import (
	"fmt"
	"net/http"
)

const uriIcons = "/api/v1/icon"

type ResourceIcon struct {
	ID   int    `json:"id"`
	Url  string `json:"url"`
	Name string `json:"name"`
}

// GetIconByID retrieves a department by ID.
func (c *Client) GetIconInfoByID(id string) (*ResourceIcon, error) {
	endpoint := fmt.Sprintf("%s/%v", uriIcons, id)
	var out ResourceIcon
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "department", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// func (c *Client) CreateIcon(filepath string, name string) (*ResourceIcon, error) {
// 	var out ResourceIcon

// 	fileBytes, err := os.ReadFile(filepath)

// 	if err != nil {
// 		return nil, fmt.Errorf(errMsgReadFile, filepath, err)
// 	}

// 	fileInfo, err := createSingleFilePayload(filepath, "test", ".png")

// 	if err != nil {
// 		return nil, fmt.Errorf("error making payload :%v", err)
// 	}

// 	resp, err := c.HTTP.DoRequest("POST", uriIcons, fileInfo, &out)

// 	if err != nil {
// 		return nil, fmt.Errorf(errMsgFailedCreate, "icon", err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return &out, nil
// }

// UploadPackage uploads a package to the Jamf Pro server. It requires the ID of an existing package
// manifest within JamfPro and the file paths.
func (c *Client) UploadIcon(filepath string) (*ResourceIcon, error) {
	endpoint := uriIcons

	files := map[string][]string{
		"file": {filepath},
	}

	formDataFields := map[string]string{
		// "description": "a cat",
	}

	contentTypes := map[string]string{
		// "image": "image/png",
	}

	formDataPartHeaders := map[string]http.Header{
		// "image": {
		// 	"Content-Disposition": []string{`form-data; name="a new image"; filename="cat.png"`},
		// },
	}

	var response ResourceIcon
	resp, err := c.HTTP.DoMultiPartRequest("POST", endpoint, files, formDataFields, contentTypes, formDataPartHeaders, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to upload Icon: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
