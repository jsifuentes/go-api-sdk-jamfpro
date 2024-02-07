package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	profileUUID := "116AF1E6-7EB5-4335-B598-276CDE5E015B" // Replace with the actual UUID
	newProfile := &jamfpro.ResourceMobileDeviceProvisioningProfile{
		General: jamfpro.MobileDeviceProvisioningProfileSubsetGeneral{
			Name:        "in-house app profile",
			DisplayName: "in-house app profile",
			UUID:        "116AF1E6-7EB5-4335-B598-276CDE5E015B",
		},
	}

	createdProfile, err := client.CreateMobileDeviceProvisioningProfileByUUID(profileUUID, newProfile)
	if err != nil {
		log.Fatalf("Error creating mobile device provisioning profile: %s\n", err)
	}

	createdProfileXML, err := xml.MarshalIndent(createdProfile, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created profile data: %v", err)
	}
	fmt.Println("Created Mobile Device Provisioning Profile:\n", string(createdProfileXML))
}
