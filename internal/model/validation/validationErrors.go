// Package validation provides generic utilities for validations
package validation

import log "github.com/sirupsen/logrus"

// Container for validation errors concerning Question
type ValidationError struct {
	// Speaking description of the error, that occured
	Problem string
	// The problem source, as a starting point for possible fixes.
	Source any
}

// Convenience function, iterates through a list of validation errors and logs them as errors
func LogValidationErrors(errs []ValidationError) {
	for _, e := range errs {
		log.WithField("source ", e.Source).Error(e.Problem)
	}
}
