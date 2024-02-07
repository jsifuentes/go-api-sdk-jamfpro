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

	// The ID of the static computer group you wish to update
	groupID := 49 // Replace with your actual group ID

	// Define the updated computers for the static group
	updatedComputers := []jamfpro.ComputerGroupSubsetComputer{
		{
			ID:            2,
			Name:          "MacBook Pro",
			MacAddress:    "",
			AltMacAddress: "",
			SerialNumber:  "D2FHXH22QB",
		},
		{
			ID:            6,
			Name:          "MacBook Pro",
			MacAddress:    "",
			AltMacAddress: "",
			SerialNumber:  "LT6M4DTF88",
		},
	}

	// Create the updated static computer group data
	updatedStaticGroup := &jamfpro.ResourceComputerGroup{
		Name:      "Updated Static Group Name",
		IsSmart:   false,
		Site:      jamfpro.SharedResourceSite{ID: -1, Name: "None"},
		Computers: updatedComputers,
	}

	// Call UpdateComputerGroupByID function
	updatedGroup, err := client.UpdateComputerGroupByID(groupID, updatedStaticGroup)
	if err != nil {
		log.Fatalf("Error updating Computer Group by ID: %v", err)
	}

	// Pretty print the updated group in XML
	groupXML, err := xml.MarshalIndent(updatedGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Updated Computer Group:\n", string(groupXML))
}
