package mapper

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

func StructMapper(origin interface{}, dest interface{}) error {
	if err := mapstructure.Decode(origin, &dest); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
