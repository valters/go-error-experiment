package main

import (
	"errors"
	"fmt"
	"log"

	"vingolds.ch/errgo/validation"
	"vingolds.ch/errgo/validation_errors"
)

func main() {

	var err = validation.MemberNotFound(404, validation.RowNotFound())

	if err != nil {
		log.Println(fmt.Errorf("got %w", err))
	}

	if errors.Is(err, validation_errors.MemberNotFound) {
		log.Printf("detected MemberNotFound: %v", err)
	}
	if errors.Is(err, validation_errors.SqlNoSuchRow) {
		log.Printf("detected SqlNoSuchRow sentinel: %v", err)
	}
	if errors.Is(err, validation_errors.SqlConnectionFailed) {
		log.Printf("detected SqlConnectionFailed sentinel (should not happen): %v", err)
	}
	if errors.Is(err, validation_errors.MemberGenericError) {
		log.Printf("detected MemberGenericError (should not happen): %v", err)
	}

	log.Println("all good")

}
