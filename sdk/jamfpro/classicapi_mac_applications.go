// classicapi_mac_applications.go
// Jamf Pro Classic Api - VPP Mac Applications
// api reference: https://developer.jamf.com/jamf-pro/reference/macapplications
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriVPPMacApplications = "/JSSResource/macapplications"

type ResponseMacApplicationsList struct {
	MacApplications []MacApplication `xml:"mac_application"`
}

type MacApplication struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResponseMacApplications represents the detailed structure of a Mac Application response.
type ResponseMacApplications struct {
	General     MacAppDataSubsetGeneral     `xml:"general"`
	Scope       MacAppDataSubsetScope       `xml:"scope"`
	SelfService MacAppDataSubsetSelfService `xml:"self_service"`
}

type MacAppDataSubsetGeneral struct {
	ID       int            `xml:"id"`
	Name     string         `xml:"name"`
	Version  string         `xml:"version"`
	IsFree   bool           `xml:"is_free"`
	BundleID string         `xml:"bundle_id"`
	URL      string         `xml:"url"`
	Category MacAppCategory `xml:"category"`
	Site     MacAppSite     `xml:"site"`
}

type MacAppCategory struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppSite struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppDataSubsetScope struct {
	AllComputers   bool                       `xml:"all_computers"`
	AllJSSUsers    bool                       `xml:"all_jss_users"`
	Buildings      []MacAppScopeBuilding      `xml:"buildings>building"`
	Departments    []MacAppScopeDepartment    `xml:"departments>department"`
	Computers      []MacAppScopeComputer      `xml:"computers>computer"`
	ComputerGroups []MacAppScopeComputerGroup `xml:"computer_groups>computer_group"`
	JSSUsers       []MacAppScopeUser          `xml:"jss_users>user"`
	JSSUserGroups  []MacAppScopeUserGroup     `xml:"jss_user_groups>user_group"`
	Limitations    MacAppScopeLimitations     `xml:"limitations"`
	Exclusions     MacAppScopeExclusions      `xml:"exclusions"`
}

// Define structs for each scope component (Building, Department, Computer, etc.)
type MacAppScopeBuilding struct {
	Building MacAppBuilding `xml:"building"`
}

type MacAppBuilding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Struct definitions for Department, Computer, ComputerGroup, User, UserGroup

type MacAppScopeDepartment struct {
	Department MacAppDepartment `xml:"department"`
}

type MacAppDepartment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppScopeComputer struct {
	Computer MacAppComputer `xml:"computer"`
}

type MacAppComputer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	UDID string `xml:"udid"`
}

type MacAppScopeComputerGroup struct {
	ComputerGroup MacAppComputerGroup `xml:"computer_group"`
}

type MacAppComputerGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppScopeUser struct {
	User MacAppUser `xml:"user"`
}

type MacAppUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppScopeUserGroup struct {
	UserGroup MacAppUserGroup `xml:"user_group"`
}

type MacAppUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppScopeLimitations struct {
	Users           []MacAppScopeUser           `xml:"users>user"`
	UserGroups      []MacAppScopeUserGroup      `xml:"user_groups>user_group"`
	NetworkSegments []MacAppScopeNetworkSegment `xml:"network_segments>network_segment"`
}

type MacAppScopeExclusions struct {
	Buildings       []MacAppScopeBuilding       `xml:"buildings>building"`
	Departments     []MacAppScopeDepartment     `xml:"departments>department"`
	Users           []MacAppScopeUser           `xml:"users>user"`
	UserGroups      []MacAppScopeUserGroup      `xml:"user_groups>user_group"`
	NetworkSegments []MacAppScopeNetworkSegment `xml:"network_segments>network_segment"`
	Computers       []MacAppScopeComputer       `xml:"computers>computer"`
	ComputerGroups  []MacAppScopeComputerGroup  `xml:"computer_groups>computer_group"`
	JSSUsers        []MacAppScopeUser           `xml:"jss_users>user"`
	JSSUserGroups   []MacAppScopeUserGroup      `xml:"jss_user_groups>user_group"`
}

type MacAppScopeNetworkSegment struct {
	ID   int    `xml:"id"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name"`
}

type MacAppDataSubsetSelfService struct {
	InstallButtonText           string                      `xml:"install_button_text"`
	SelfServiceDescription      string                      `xml:"self_service_description"`
	ForceUsersToViewDescription bool                        `xml:"force_users_to_view_description"`
	SelfServiceIcon             MacAppSelfServiceIcon       `xml:"self_service_icon"`
	FeatureOnMainPage           bool                        `xml:"feature_on_main_page"`
	SelfServiceCategories       []MacAppSelfServiceCategory `xml:"self_service_categories>category"`
	Notification                string                      `xml:"notification"`
	NotificationSubject         string                      `xml:"notification_subject"`
	NotificationMessage         string                      `xml:"notification_message"`
	VPP                         MacAppVPP                   `xml:"vpp"`
}

type MacAppSelfServiceIcon struct {
	ID   int    `xml:"id"`
	URI  string `xml:"uri"`
	Data string `xml:"data"`
}

type MacAppSelfServiceCategory struct {
	ID        int    `xml:"id"`
	Name      string `xml:"name"`
	DisplayIn bool   `xml:"display_in"`
	FeatureIn bool   `xml:"feature_in"`
}

type MacAppVPP struct {
	AssignVPPDeviceBasedLicenses bool `xml:"assign_vpp_device_based_licenses"`
	VPPAdminAccountID            int  `xml:"vpp_admin_account_id"`
}

// GetDockItems retrieves a serialized list of vpp mac applications.
func (c *Client) GetMacApplications() (*ResponseMacApplicationsList, error) {
	endpoint := uriVPPMacApplications

	var macApps ResponseMacApplicationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApps)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Applications: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApps, nil
}

// GetMacApplicationByID retrieves a single Mac application by its ID.
func (c *Client) GetMacApplicationByID(id int) (*ResponseMacApplications, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPMacApplications, id)

	var macApp ResponseMacApplications
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Application by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApp, nil
}

// GetMacApplicationByName retrieves a single Mac application by its name.
func (c *Client) GetMacApplicationByName(name string) (*ResponseMacApplications, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriVPPMacApplications, name)

	var macApp ResponseMacApplications
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Application by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApp, nil
}

// GetMacApplicationByNameAndDataSubset retrieves a specific Mac Application by its ID and filters by a specific data subset.
// Subset values can be General, Scope, SelfService, VPPCodes and VPP.
func (c *Client) GetMacApplicationByIDAndDataSubset(id int, subset string) (*ResponseMacApplications, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriVPPMacApplications, id, subset)

	var macApp ResponseMacApplications
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Application by Name and Subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApp, nil
}

// GetMacApplicationByNameAndDataSubset retrieves a specific Mac Application by its name and filters by a specific data subset.
// Subset values can be General, Scope, SelfService, VPPCodes and VPP.
func (c *Client) GetMacApplicationByNameAndDataSubset(name, subset string) (*ResponseMacApplications, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriVPPMacApplications, name, subset)

	var macApp ResponseMacApplications
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Application by Name and Subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApp, nil
}

// CreateMacApplication creates a new Mac Application.
func (c *Client) CreateMacApplication(macApp ResponseMacApplications) (*ResponseMacApplications, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriVPPMacApplications) // '0' typically used for creation in APIs

	// Set default values for site if not included within request
	if macApp.General.Site.ID == 0 && macApp.General.Site.Name == "" {
		macApp.General.Site = MacAppSite{
			ID:   -1,
			Name: "None",
		}
	}

	// The requestBody struct should mirror the ResponseMacApplications struct, including all nested structs
	requestBody := struct {
		XMLName xml.Name `xml:"mac_application"`
		ResponseMacApplications
	}{
		ResponseMacApplications: macApp,
	}

	var response ResponseMacApplications
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create Mac Application: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateMacApplicationByID updates an existing Mac Application by its ID.
func (c *Client) UpdateMacApplicationByID(id int, macApp ResponseMacApplications) (*ResponseMacApplications, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPMacApplications, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mac_application"`
		ResponseMacApplications
	}{
		ResponseMacApplications: macApp,
	}

	var response ResponseMacApplications
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update Mac Application by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateMacApplicationByName updates an existing Mac Application by its name.
func (c *Client) UpdateMacApplicationByName(name string, macApp ResponseMacApplications) (*ResponseMacApplications, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriVPPMacApplications, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mac_application"`
		ResponseMacApplications
	}{
		ResponseMacApplications: macApp,
	}

	var response ResponseMacApplications
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update Mac Application by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteMacApplicationByID deletes a MacApplication by its ID.
func (c *Client) DeleteMacApplicationByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPMacApplications, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete VPP Mac Application Item by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMacApplicationByName deletes a MacApplication by its name.
func (c *Client) DeleteMacApplicationByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriVPPMacApplications, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete VPP Mac Application Item by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}