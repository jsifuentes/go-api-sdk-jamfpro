package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
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
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Example profile to be updated
	profileToUpdate := jamfpro.ResponseMobileDeviceEnrollmentProfile{
		General: jamfpro.MobileDeviceEnrollmentProfileGeneral{
			Name:        "Configurator Enrollment Profile",
			Description: "string",
		},
		Location: jamfpro.MobileDeviceEnrollmentProfileLocation{
			// Initialize with empty or specific values if required
			Username:     "",
			Realname:     "",
			RealName:     "",
			EmailAddress: "",
			Position:     "",
			Phone:        "",
			PhoneNumber:  "",
			Department:   "",
			Building:     "",
			Room:         0, // or specific room number
		},
		Purchasing: jamfpro.MobileDeviceEnrollmentProfilePurchasing{
			IsPurchased:          true,
			IsLeased:             false,
			PONumber:             "",
			Vendor:               "",
			ApplecareID:          "",
			PurchasePrice:        "",
			PurchasingAccount:    "",
			PODate:               "",
			PODateEpoch:          0,
			PODateUTC:            "",
			WarrantyExpires:      "",
			WarrantyExpiresEpoch: 0,
			WarrantyExpiresUTC:   "",
			LeaseExpires:         "",
			LeaseExpiresEpoch:    0,
			LeaseExpiresUTC:      "",
			LifeExpectancy:       0,
			PurchasingContact:    "",
		},
	}

	inviteName := "YourInvitation" // Replace name with the actual invite name

	updatedProfile, err := client.UpdateMobileDeviceEnrollmentProfileByInvitation(inviteName, &profileToUpdate)
	if err != nil {
		log.Fatalf("Error updating profile by invitation: %v", err)
	}

	fmt.Println("Updated Profile by Invitation:", updatedProfile)
}
