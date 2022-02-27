package internal

import "errors"

var (
	ErrNoGetterFound   = errors.New("No getter found")
	ErrNoDeleterFound  = errors.New("No deleter found")
	ErrNoSelecterFound = errors.New("No getter found")
	ErrNoUpdaterFound  = errors.New("No updater found")
	ErrNoInserterFound = errors.New("No inserter found")
)
