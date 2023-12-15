// Package config allows using a yaml file to define [Trigger]s, [Chase]s and [Event]s.
package config

import (
	"fmt"

	models_config "github.com/H3rby7/dmx-web-go/internal/model/config"
	"github.com/H3rby7/dmx-web-go/internal/model/validation"
)

// Validates that the field Name contains a reasonable value
func validateName(chase models_config.Chase) (ok bool, error validation.ValidationError) {
	ok = true
	if chase.Name == "" {
		ok = false
		error.Problem = "the field 'Name' was not set properly"
		error.Source = chase
		return
	}
	return
}

// Validates that the Name is unique
func validateNameUniqueness(triggers []models_config.Chase) (ok bool, errors []validation.ValidationError) {
	ok = true
	usedNames := make(map[string]bool)
	for _, chase := range triggers {
		if usedNames[chase.Name] {
			ok = false
			errors = append(errors, validation.ValidationError{
				Problem: fmt.Sprintf("A duplicate name was found: %s.", chase.Name),
				Source:  chase,
			})
		}
		usedNames[chase.Name] = true
	}
	return
}

// ValidateFile validates the actions file with a set of validators;
//
// * RETURNS true, when no errors found
//
// * RETURNS false, if validation fails and returns the validation errors
func ValidateFile(file models_config.ConfigFile) (ok bool, errors []validation.ValidationError) {
	ok = true
	for _, chase := range file.Chases {
		if _ok, err := validateName(chase); !_ok {
			ok = false
			errors = append(errors, err)
		}
	}
	if _ok, err := validateNameUniqueness(file.Chases); !_ok {
		ok = false
		errors = append(errors, err...)
	}
	return
}
