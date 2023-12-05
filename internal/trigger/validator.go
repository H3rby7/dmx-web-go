package trigger

import (
	"fmt"

	"github.com/H3rby7/dmx-web-go/internal/shared/validation"
)

// Validates that the field Name contains a reasonable value
func (chase Chase) validateName() (ok bool, error validation.ValidationError) {
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
func validateNameUniqueness(triggers []Chase) (ok bool, errors []validation.ValidationError) {
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

/*
validates the actions file with a set of validators;

ok = true => no errors found

ok = false => errors field contains the validation errors
*/
func ValidateFile(file ActionsFile) (ok bool, errors []validation.ValidationError) {
	ok = true
	for _, chase := range file.Chases {
		if _ok, err := chase.validateName(); !_ok {
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
