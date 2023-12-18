// classicapi_computer_groups.go
// Jamf Pro Classic Api - Computer Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/computergroups
// Classic API requires the structs to support an XML data structure.
package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriComputerGroups = "/JSSResource/computergroups"

type ResponseComputerGroupsList struct {
	Size    int `xml:"size"`
	Results []struct {
		ID      int    `xml:"id,omitempty"`
		Name    string `xml:"name,omitempty"`
		IsSmart bool   `xml:"is_smart,omitempty"`
	} `xml:"computer_group"`
}

type ResourceComputerGroup struct {
	ID      int    `xml:"id"`
	Name    string `xml:"name"`
	IsSmart bool   `xml:"is_smart"`
	Site    struct {
		ID   int    `json:"id,omitempty" xml:"id,omitempty"`
		Name string `json:"name,omitempty" xml:"name,omitempty"`
	} `xml:"site"`
	Criteria []struct {
		Size      int `xml:"size"`
		Criterion struct {
			Name         string           `xml:"name"`
			Priority     int              `xml:"priority"`
			AndOr        DeviceGroupAndOr `xml:"and_or"`
			SearchType   string           `xml:"search_type"`
			SearchValue  string           `xml:"value"`
			OpeningParen bool             `xml:"opening_paren"`
			ClosingParen bool             `xml:"closing_paren"`
		} `xml:"criterion"`
	} `xml:"criteria>criterion"`
	Computers []struct {
		ID            int    `json:"id,omitempty" xml:"id,omitempty"`
		Name          string `json:"name,omitempty" xml:"name,omitempty"`
		SerialNumber  string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
		MacAddress    string `json:"mac_address,omitempty" xml:"mac_address,omitempty"`
		AltMacAddress string `json:"alt_mac_address,omitempty" xml:"alt_mac_address,omitempty"`
	} `xml:"computers>computer"`
}

type DeviceGroupAndOr string

const (
	And DeviceGroupAndOr = "and"
	Or  DeviceGroupAndOr = "or"
)

// GetComputerGroups gets a list of all computer groups
func (c *Client) GetComputerGroups() (*ResponseComputerGroupsList, error) {
	endpoint := uriComputerGroups

	var computerGroups ResponseComputerGroupsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &computerGroups)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Computer Groups: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &computerGroups, nil
}

// GetComputerGroupByID retrieves a computer group by its ID.
func (c *Client) GetComputerGroupByID(id int) (*ResourceComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerGroups, id)

	var group ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// GetComputerGroupByName retrieves a computer group by its name.
func (c *Client) GetComputerGroupByName(name string) (*ResourceComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerGroups, name)

	var group ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Computer Group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// CreateComputerGroup creates a new computer group.
func (c *Client) CreateComputerGroup(group *ResourceComputerGroup) (*ResourceComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriComputerGroups) // Using ID 0 for creation as per the pattern

	// Check if site is not provided and set default values
	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site.ID = -1
		group.Site.Name = "none"
	}

	// Wrap the group request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResourceComputerGroup
	}{
		ResourceComputerGroup: group,
	}

	var createdGroup ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to create Computer Group: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdGroup, nil
}

// UpdateComputerGroupByID updates an existing computer group by its ID.
func (c *Client) UpdateComputerGroupByID(id int, group *ResourceComputerGroup) (*ResourceComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerGroups, id)

	// Check if site is not provided and set default values
	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site.ID = -1
		group.Site.Name = "none"
	}

	// Wrap the group request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResourceComputerGroup
	}{
		ResourceComputerGroup: group,
	}

	var updatedGroup ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to update Computer Group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// UpdateComputerGroupByName updates a computer group by its name.
func (c *Client) UpdateComputerGroupByName(name string, group *ResourceComputerGroup) (*ResourceComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerGroups, name)

	// Check if site is not provided and set default values
	if group.Site.ID == 0 && group.Site.Name == "" {
		group.Site.ID = -1
		group.Site.Name = "none"
	}

	// Wrap the group request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"computer_group"`
		*ResourceComputerGroup
	}{
		ResourceComputerGroup: group,
	}

	var updatedGroup ResourceComputerGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to update Computer Group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedGroup, nil
}

// DeleteComputerGroupByID deletes a computer group by its ID.
func (c *Client) DeleteComputerGroupByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriComputerGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Computer Group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteComputerGroupByName deletes a computer group by its name.
func (c *Client) DeleteComputerGroupByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriComputerGroups, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Computer Group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
