package connector

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// LoadDetails returns the config values that allow connection to the database
func LoadDetails() string {

	// Details holds our db authentication details
	type Details struct {
		Host     string
		Port     int32
		User     string
		Password string
		Database string
	}

	var config Details

	if len(os.Args) == 0 {
		toml.DecodeFile("./config.toml", &config)
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)
	} else {
		psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("Host"), os.Getenv("User"), os.Getenv("Password"), os.Getenv("Database"))

	}
	return psqlInfo
}
