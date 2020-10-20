package config

import (
	"log"
	"os"

	"git.fuyu.moe/Fuyu/validate"
	"github.com/BurntSushi/toml"
	"github.com/joeshaw/envdecode"
)

// Read reads the given config file to the destination
// When configFile is not found only environment variables will be read
// Environment variables will overwrite settings in the give configFile
func Read(configFile string, dst interface{}) {
	_, err := toml.DecodeFile(configFile, dst)
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			break
		default:
			panic(err)
		}
	}

	err = envdecode.Decode(dst)
	if err != nil && err != envdecode.ErrNoTargetFieldsAreSet {
		panic(err)
	}

	mustValidate(dst)
}

func mustValidate(v interface{}) {
	validationErrors := validate.Validate(v)
	if len(validationErrors) != 0 {
		for _, e := range validationErrors {
			log.Fatal(`invalid config - `, e)
		}

		os.Exit(1)
	}
}
