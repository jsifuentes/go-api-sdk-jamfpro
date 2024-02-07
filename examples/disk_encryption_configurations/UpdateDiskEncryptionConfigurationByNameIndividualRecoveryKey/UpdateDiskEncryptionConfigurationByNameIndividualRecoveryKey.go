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

	configToUpdate := &jamfpro.ResourceDiskEncryptionConfiguration{
		Name:                  "Corporate Encryption Name Updated",
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}

	updatedConfig, err := client.UpdateDiskEncryptionConfigurationByName("Corporate Encryption", configToUpdate)
	if err != nil {
		log.Fatalf("Error updating Disk Encryption Configuration by Name: %v", err)
	}

	configXML, err := xml.MarshalIndent(updatedConfig, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated configuration to XML: %v", err)
	}

	fmt.Printf("Updated Disk Encryption Configuration by Name:\n%s\n", configXML)
}
