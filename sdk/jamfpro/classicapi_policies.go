// classicapi_policies.go
// Jamf Pro Classic Api - Policies
// api reference: https://developer.jamf.com/jamf-pro/reference/policies
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriPolicies = "/JSSResource/policies"

// Policies List Structs
type ResponsePoliciesList struct {
	Size   int          `xml:"size"`
	Policy []PolicyItem `xml:"policy"`
}

type PolicyItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResourcePolicyCreateAndUpdate represents the response structure for creating or updating a policy
type ResourcePolicyCreateAndUpdate struct {
	XMLName xml.Name `xml:"policy"`
	ID      int      `xml:"id"`
}

// ResourcePolicy represents the response structure for a single policy
type ResourcePolicy struct {
	General              PolicySubsetGeneral              `xml:"general"`
	Scope                PolicySubsetScope                `xml:"scope,omitempty"`
	SelfService          PolicySubsetSelfService          `xml:"self_service"`
	PackageConfiguration PolicySubsetPackageConfiguration `xml:"package_configuration,omitempty"`
	Scripts              PolicySubsetScripts              `xml:"scripts,omitempty"`
	Printers             PolicySubsetPrinters             `xml:"printers"`
	DockItems            PolicySubsetDockItems            `xml:"dock_items"`
	AccountMaintenance   PolicySubsetAccountMaintenance   `xml:"account_maintenance"`
	Maintenance          PolicySubsetMaintenance          `xml:"maintenance"`
	FilesProcesses       PolicySubsetFilesProcesses       `xml:"files_processes"`
	UserInteraction      PolicySubsetUserInteraction      `xml:"user_interaction"`
	DiskEncryption       PolicySubsetDiskEncryption       `xml:"disk_encryption"`
	Reboot               PolicySubsetReboot               `xml:"reboot"`
}

// Subsets & Containers

// General

// PolicySubsetGeneral represents the general information of a policy
type PolicySubsetGeneral struct {
	ID                         int                                    `xml:"id"`
	Name                       string                                 `xml:"name"`
	Enabled                    bool                                   `xml:"enabled"`
	Trigger                    string                                 `xml:"trigger,omitempty"`
	TriggerCheckin             bool                                   `xml:"trigger_checkin"`
	TriggerEnrollmentComplete  bool                                   `xml:"trigger_enrollment_complete"`
	TriggerLogin               bool                                   `xml:"trigger_login"`
	TriggerLogout              bool                                   `xml:"trigger_logout"`
	TriggerNetworkStateChanged bool                                   `xml:"trigger_network_state_changed"`
	TriggerStartup             bool                                   `xml:"trigger_startup"`
	TriggerOther               string                                 `xml:"trigger_other,omitempty"`
	Frequency                  string                                 `xml:"frequency,omitempty"`
	RetryEvent                 string                                 `xml:"retry_event,omitempty"`
	RetryAttempts              int                                    `xml:"retry_attempts,omitempty"`
	NotifyOnEachFailedRetry    bool                                   `xml:"notify_on_each_failed_retry"`
	LocationUserOnly           bool                                   `xml:"location_user_only"`
	TargetDrive                string                                 `xml:"target_drive,omitempty"`
	Offline                    bool                                   `xml:"offline"`
	Category                   PolicyCategory                         `xml:"category,omitempty"`
	DateTimeLimitations        PolicySubsetGeneralDateTimeLimitations `xml:"date_time_limitations,omitempty"`
	NetworkLimitations         PolicySubsetGeneralNetworkLimitations  `xml:"network_limitations,omitempty"`
	OverrideDefaultSettings    PolicySubsetGeneralOverrideSettings    `xml:"override_default_settings,omitempty"`
	NetworkRequirements        string                                 `xml:"network_requirements,omitempty"`
	Site                       SharedResourceSite                     `xml:"site"`
}

type PolicySubsetGeneralDateTimeLimitations struct {
	ActivationDate      string                                              `xml:"activation_date,omitempty"`
	ActivationDateEpoch int                                                 `xml:"activation_date_epoch,omitempty"`
	ActivationDateUTC   string                                              `xml:"activation_date_utc,omitempty"`
	ExpirationDate      string                                              `xml:"expiration_date,omitempty"`
	ExpirationDateEpoch int                                                 `xml:"expiration_date_epoch,omitempty"`
	ExpirationDateUTC   string                                              `xml:"expiration_date_utc,omitempty"`
	NoExecuteOn         []PolicySubsetGeneralDateTimeLimitationsNoExecuteOn `xml:"no_execute_on>day,omitempty"`
	NoExecuteStart      string                                              `xml:"no_execute_start,omitempty"`
	NoExecuteEnd        string                                              `xml:"no_execute_end,omitempty"`
}

type PolicySubsetGeneralDateTimeLimitationsNoExecuteOn struct {
	Day string `xml:",chardata"`
}

type PolicySubsetGeneralNetworkLimitations struct {
	MinimumNetworkConnection string `xml:"minimum_network_connection,omitempty"`
	AnyIPAddress             bool   `xml:"any_ip_address"`
	NetworkSegments          string `xml:"network_segments"`
}

type PolicySubsetGeneralOverrideSettings struct {
	TargetDrive       string `xml:"target_drive,omitempty"`
	DistributionPoint string `xml:"distribution_point,omitempty"`
	ForceAfpSmb       bool   `xml:"force_afp_smb"`
	SUS               string `xml:"sus,omitempty"`
	NetbootServer     string `xml:"netboot_server,omitempty"`
}

// Scope

// PolicySubsetScope represents the scope of the policy
type PolicySubsetScope struct {
	AllComputers   bool                            `xml:"all_computers"`
	AllJSSUsers    bool                            `xml:"all_jss_users"`
	Computers      []PolicyDataSubsetComputer      `xml:"computers>computer,omitempty"`
	ComputerGroups []PolicyDataSubsetComputerGroup `xml:"computer_groups>computer_group,omitempty"`
	JSSUsers       []PolicyDataSubsetJSSUser       `xml:"jss_users>jss_user,omitempty"`
	JSSUserGroups  []PolicyDataSubsetJSSUserGroup  `xml:"jss_user_groups>jss_user_group,omitempty"`
	Buildings      []PolicyDataSubsetBuilding      `xml:"buildings>building,omitempty"`
	Departments    []PolicyDataSubsetDepartment    `xml:"departments>department,omitempty"`
	LimitToUsers   PolicyLimitToUsers              `xml:"limit_to_users,omitempty"`
	Limitations    PolicySubsetScopeLimitations    `xml:"limitations,omitempty"`
	Exclusions     PolicySubsetScopeExclusions     `xml:"exclusions,omitempty"`
}

type PolicySubsetScopeLimitations struct {
	Users           []PolicyDataSubsetUser           `xml:"users>user,omitempty"`
	UserGroups      []PolicyDataSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []PolicyDataSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	IBeacons        []PolicyDataSubsetIBeacon        `xml:"ibeacons>ibeacon,omitempty"`
}

type PolicySubsetScopeExclusions struct {
	Computers       []PolicyDataSubsetComputer       `xml:"computers>computer,omitempty"`
	ComputerGroups  []PolicyDataSubsetComputerGroup  `xml:"computer_groups>computer_group,omitempty"`
	Users           []PolicyDataSubsetUser           `xml:"users>user,omitempty"`
	UserGroups      []PolicyDataSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
	Buildings       []PolicyDataSubsetBuilding       `xml:"buildings>building,omitempty"`
	Departments     []PolicyDataSubsetDepartment     `xml:"departments>department,omitempty"`
	NetworkSegments []PolicyDataSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	JSSUsers        []PolicyDataSubsetJSSUser        `xml:"jss_users>jss_user,omitempty"`
	JSSUserGroups   []PolicyDataSubsetJSSUserGroup   `xml:"jss_user_groups>jss_user_group,omitempty"`
	IBeacons        []PolicyDataSubsetIBeacon        `xml:"ibeacons>ibeacon,omitempty"`
}

// PolicySubsetSelfService represents the self service settings of a policy
type PolicySubsetSelfService struct {
	UseForSelfService           bool                              `xml:"use_for_self_service"`
	SelfServiceDisplayName      string                            `xml:"self_service_display_name"`
	InstallButtonText           string                            `xml:"install_button_text"`
	ReinstallButtonText         string                            `xml:"re_install_button_text"`
	SelfServiceDescription      string                            `xml:"self_service_description"`
	ForceUsersToViewDescription bool                              `xml:"force_users_to_view_description"`
	SelfServiceIcon             SharedResourceSelfServiceIcon     `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                              `xml:"feature_on_main_page"`
	SelfServiceCategories       []PolicySubsetSelfServiceCategory `xml:"self_service_categories"`
	Notification                bool                              `xml:"notification,omitempty"`
	NotificationType            string                            `xml:"notification_type"`
	NotificationSubject         string                            `xml:"notification_subject,omitempty"`
	NotificationMessage         string                            `xml:"notification_message,omitempty"`
}

type PolicySubsetSelfServiceCategory struct {
	Category PolicyCategory `xml:"category"`
}

// Package Configuration

// PolicySubsetPackageConfiguration represents the package configuration settings of a policy
type PolicySubsetPackageConfiguration struct {
	Packages          []PolicySubsetPackageConfigurationPackage `xml:"packages>package"`
	DistributionPoint string                                    `xml:"distribution_point"`
}

type PolicySubsetPackageConfigurationPackage struct {
	ID                int    `xml:"id,omitempty"`
	Name              string `xml:"name,omitempty"`
	Action            string `xml:"action,omitempty"`
	FillUserTemplate  bool   `xml:"fut,omitempty"`
	FillExistingUsers bool   `xml:"feu,omitempty"`
	UpdateAutorun     bool   `xml:"update_autorun,omitempty"`
}

// Scripts

// PolicySubsetScripts represents the scripts settings of a policy
type PolicySubsetScripts struct {
	Size   int                  `xml:"size"`
	Script []PolicySubsetScript `xml:"script"`
}

type PolicySubsetScript struct {
	ID          string `xml:"id,omitempty"`
	Name        string `xml:"name,omitempty"`
	Priority    string `xml:"priority,omitempty"`
	Parameter4  string `xml:"parameter4,omitempty"`
	Parameter5  string `xml:"parameter5,omitempty"`
	Parameter6  string `xml:"parameter6,omitempty"`
	Parameter7  string `xml:"parameter7,omitempty"`
	Parameter8  string `xml:"parameter8,omitempty"`
	Parameter9  string `xml:"parameter9,omitempty"`
	Parameter10 string `xml:"parameter10,omitempty"`
	Parameter11 string `xml:"parameter11,omitempty"`
}

// Printers

// PolicySubsetPrinters represents the printers settings of a policy
type PolicySubsetPrinters struct {
	Size                 int                   `xml:"size"`
	LeaveExistingDefault bool                  `xml:"leave_existing_default"`
	Printer              []PolicySubsetPrinter `xml:"printer"`
}

type PolicySubsetPrinter struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Action      string `xml:"action"`
	MakeDefault bool   `xml:"make_default"`
}

// Dock Items

// PolicySubsetDockItems represents the dock items settings of a policy
type PolicySubsetDockItems struct {
	Size     int                    `xml:"size"`
	DockItem []PolicySubsetDockItem `xml:"dock_item"`
}

type PolicySubsetDockItem struct {
	ID     int    `xml:"id"`
	Name   string `xml:"name"`
	Action string `xml:"action"`
}

// Account Maintenance

// PolicySubsetAccountMaintenance represents the account maintenance settings of a policy
type PolicySubsetAccountMaintenance struct {
	Accounts                []PolicySubsetAccountMaintenanceAccount               `xml:"accounts>account"`
	DirectoryBindings       []PolicySubsetAccountMaintenanceDirectoryBindings     `xml:"directory_bindings>binding"`
	ManagementAccount       PolicySubsetAccountMaintenanceManagementAccount       `xml:"management_account"`
	OpenFirmwareEfiPassword PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword `xml:"open_firmware_efi_password"`
}

type PolicySubsetAccountMaintenanceAccount struct {
	Action                 string `xml:"action"`
	Username               string `xml:"username"`
	Realname               string `xml:"realname"`
	Password               string `xml:"password"`
	ArchiveHomeDirectory   bool   `xml:"archive_home_directory"`
	ArchiveHomeDirectoryTo string `xml:"archive_home_directory_to"`
	Home                   string `xml:"home"`
	Hint                   string `xml:"hint"`
	Picture                string `xml:"picture"`
	Admin                  bool   `xml:"admin"`
	FilevaultEnabled       bool   `xml:"filevault_enabled"`
	PasswordSha256         string `xml:"password_sha256"`
}

type PolicySubsetAccountMaintenanceDirectoryBindings struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type PolicySubsetAccountMaintenanceManagementAccount struct {
	Action                string `xml:"action"`
	ManagedPassword       string `xml:"managed_password"`
	ManagedPasswordLength int    `xml:"managed_password_length"`
}

type PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword struct {
	OfMode           string `xml:"of_mode"`
	OfPassword       string `xml:"of_password"`
	OfPasswordSHA256 string `xml:"of_password_sha256"`
}

// Maintenance

// PolicySubsetMaintenance represents the maintenance settings of a policy
type PolicySubsetMaintenance struct {
	Recon                    bool `xml:"recon"`
	ResetName                bool `xml:"reset_name"`
	InstallAllCachedPackages bool `xml:"install_all_cached_packages"`
	Heal                     bool `xml:"heal"`
	Prebindings              bool `xml:"prebindings"`
	Permissions              bool `xml:"permissions"`
	Byhost                   bool `xml:"byhost"`
	SystemCache              bool `xml:"system_cache"`
	UserCache                bool `xml:"user_cache"`
	Verify                   bool `xml:"verify"`
}

// Files Processes

// PolicySubsetFilesProcesses represents the files and processes settings of a policy
type PolicySubsetFilesProcesses struct {
	SearchByPath         string `xml:"search_by_path"`
	DeleteFile           bool   `xml:"delete_file"`
	LocateFile           string `xml:"locate_file"`
	UpdateLocateDatabase bool   `xml:"update_locate_database"`
	SpotlightSearch      string `xml:"spotlight_search"`
	SearchForProcess     string `xml:"search_for_process"`
	KillProcess          bool   `xml:"kill_process"`
	RunCommand           string `xml:"run_command"`
}

// User Interaction

// PolicySubsetUserInteraction represents the user interaction settings of a policy
type PolicySubsetUserInteraction struct {
	MessageStart          string `xml:"message_start"`
	AllowUserToDefer      bool   `xml:"allow_user_to_defer"`
	AllowDeferralUntilUtc string `xml:"allow_deferral_until_utc"`
	AllowDeferralMinutes  int    `xml:"allow_deferral_minutes"`
	MessageFinish         string `xml:"message_finish"`
}

// Disk Encryption

// PolicySubsetDiskEncryption represents the disk encryption settings of a policy
type PolicySubsetDiskEncryption struct {
	Action                                 string `xml:"action"`
	DiskEncryptionConfigurationID          int    `xml:"disk_encryption_configuration_id"`
	AuthRestart                            bool   `xml:"auth_restart"`
	RemediateKeyType                       string `xml:"remediate_key_type,omitempty"`
	RemediateDiskEncryptionConfigurationID int    `xml:"remediate_disk_encryption_configuration_id,omitempty"`
}

// Reboot

// PolicySubsetReboot represents the reboot settings of a policy
type PolicySubsetReboot struct {
	Message                     string `xml:"message"`
	StartupDisk                 string `xml:"startup_disk"`
	SpecifyStartup              string `xml:"specify_startup"`
	NoUserLoggedIn              string `xml:"no_user_logged_in"`
	UserLoggedIn                string `xml:"user_logged_in"`
	MinutesUntilReboot          int    `xml:"minutes_until_reboot"`
	StartRebootTimerImmediately bool   `xml:"start_reboot_timer_immediately"`
	FileVault2Reboot            bool   `xml:"file_vault_2_reboot"`
}

// Shared

type PolicyCategory struct {
	ID        int    `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	DisplayIn bool   `xml:"display_in"`
	FeatureIn bool   `xml:"feature_in"`
}

type PolicyDataSubsetComputer struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UDID string `xml:"udid,omitempty"`
}

type PolicyDataSubsetComputerGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetJSSUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetJSSUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
type PolicyDataSubsetBuilding struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetDepartment struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyLimitToUsers struct {
	UserGroups []string `xml:"user_groups>user_group,omitempty"`
}

type PolicyDataSubsetUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type PolicyDataSubsetNetworkSegment struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UID  string `xml:"uid,omitempty"`
}

type PolicyDataSubsetIBeacon struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// CRUD

// GetPolicies retrieves a list of all policies.
func (c *Client) GetPolicies() (*ResponsePoliciesList, error) {
	endpoint := uriPolicies

	var policiesList ResponsePoliciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policiesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all policies: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policiesList, nil
}

// GetPolicyByID retrieves the details of a policy by its ID.
func (c *Client) GetPolicyByID(id int) (*ResourcePolicy, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)

	var policyDetails ResourcePolicy
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policyDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch policy by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policyDetails, nil
}

// GetPolicyByName retrieves a policy by its name.
func (c *Client) GetPolicyByName(name string) (*ResourcePolicy, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)

	var policyDetails ResourcePolicy
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policyDetails)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch policy by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policyDetails, nil
}

// GetPolicyByCategory retrieves policies by their category.
func (c *Client) GetPolicyByCategory(category string) (*ResponsePoliciesList, error) {
	endpoint := fmt.Sprintf("%s/category/%s", uriPolicies, category)

	var policiesList ResponsePoliciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policiesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch policies by category: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policiesList, nil
}

// GetPoliciesByType retrieves policies by the type of entity that created them.
// The createdBy param can be either the value 'casper' which refers to Casper Remote. Or the value 'jss', which refers to policies created in the GUI or via the API.
func (c *Client) GetPoliciesByType(createdBy string) (*ResponsePoliciesList, error) {
	endpoint := fmt.Sprintf("%s/createdBy/%s", uriPolicies, createdBy)

	var policiesList ResponsePoliciesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &policiesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch policies by type: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &policiesList, nil
}

// CreatePolicy creates a new policy.
func (c *Client) CreatePolicy(policy *ResourcePolicy) (*ResourcePolicyCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, policy.General.ID)

	// Wrap the policy with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"policy"`
		*ResourcePolicy
	}{
		ResourcePolicy: policy,
	}

	var ResourcePolicy ResourcePolicyCreateAndUpdate
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &ResourcePolicy)
	if err != nil {
		return nil, fmt.Errorf("failed to create policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Return the ID of the newly created policy
	return &ResourcePolicy, nil
}

// UpdatePolicyByID updates an existing policy by its ID.
func (c *Client) UpdatePolicyByID(id int, policy *ResourcePolicy) (*ResourcePolicyCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)

	// Wrap the policy with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"policy"`
		*ResourcePolicy
	}{
		ResourcePolicy: policy,
	}

	var response ResourcePolicyCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdatePolicyByName updates an existing policy by its name.
func (c *Client) UpdatePolicyByName(name string, policy *ResourcePolicy) (*ResourcePolicyCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)

	// Wrap the policy with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"policy"`
		*ResourcePolicy
	}{
		ResourcePolicy: policy,
	}

	var response ResourcePolicyCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeletePolicyByID deletes a policy by its ID.
func (c *Client) DeletePolicyByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriPolicies, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeletePolicyByName deletes a policy by its name.
func (c *Client) DeletePolicyByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriPolicies, name)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete policy: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
