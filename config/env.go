package config

import (
	// "log"
	"os"
)

// LoadEnv sets environment variables before anything else runs
func LoadEnv() {
	// if os.Getenv("SECRET_KEY") == "" {
	//     log.Println("SECRET_KEY is not set, using default key for development.")
	//     os.Setenv("SECRET_KEY", "MySuperSecretKeyThatShouldBeLongAndRandom123!")
	// }
	os.Setenv("SECRET_KEY", "MySuperSecretKeyThatShouldBeLongAndRandom123!")
}
