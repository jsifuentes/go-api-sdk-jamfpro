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
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Specify the ID of the VPP assignment to retrieve
	vppAssignmentID := 1 // Replace with the actual ID

	// Call the GetVPPAssignmentByID function
	vppAssignment, err := client.GetVPPAssignmentByID(vppAssignmentID)
	if err != nil {
		log.Fatalf("Error retrieving VPP Assignment by ID: %v", err)
	}

	// Pretty print the VPP assignment details in XML
	vppAssignmentsXML, err := xml.MarshalIndent(vppAssignment, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error VPP assignment data: %v", err)
	}
	fmt.Println("VPP Assignment Details:\n", string(vppAssignmentsXML))
}
