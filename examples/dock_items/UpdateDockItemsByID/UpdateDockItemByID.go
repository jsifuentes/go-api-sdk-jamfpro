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

	// Define the ID and details of the dock item to update
	dockItemID := 1 // Replace with the actual ID
	updatedDockItem := &jamfpro.ResourceDockItem{
		// Set the fields you want to update
		Name:     "Updated Safari",
		Type:     "App",
		Path:     "file://localhost/Applications/Safari.app/",
		Contents: "Updated Contents",
	}

	// Call the UpdateDockItemByID function
	dockItem, err := client.UpdateDockItemByID(dockItemID, updatedDockItem)
	if err != nil {
		log.Fatalf("Error updating dock item by ID: %v", err)
	}

	// Pretty print the updated group in XML
	groupXML, err := xml.MarshalIndent(dockItem, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Updated Computer Group:\n", string(groupXML))
}
