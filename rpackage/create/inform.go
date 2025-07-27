package create

import (
	"log"
)

// inform ia a  helper function to inform user on progress
func inform(s string) error {
	log.Print("writing " + s + " \n")
	return nil
}
