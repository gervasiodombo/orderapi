package config

import (
	"fmt"
	"os"
)

type Profile string

const (
	ProfileTest    Profile = "test"
	ProfileProd    Profile = "prod"
	ProfileStaging Profile = "staging"
)

func LoadProfile() error {
	profile := os.Getenv("APP_ENV")
	if profile == "" {
		profile = string(ProfileTest)
	}

	envFile := fmt.Sprintf("configs/.env.%s", profile)

	if err := readEnvFile(envFile); err != nil {
		return fmt.Errorf("failed to load profile '%s': %w", profile, err)
	}

	fmt.Printf("Profile loaded: %s\n", envFile)
	return nil
}
