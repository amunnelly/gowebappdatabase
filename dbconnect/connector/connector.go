package connector

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
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
	var psqlInfo string

	if len(os.Getenv("PORT")) > 0 {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("Host"), string(os.Getenv("PORT")), os.Getenv("User"), os.Getenv("Password"), os.Getenv("Database"))
		fmt.Println(psqlInfo)
	} else {
		toml.DecodeFile("./config.toml", &config)
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Database)
		fmt.Println(psqlInfo)
	}
	return psqlInfo
}
