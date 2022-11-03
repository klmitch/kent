// Copyright (c) 2020 Kevin L. Mitchell
//
// Licensed under the Apache License, Version 2.0 (the "License"); you
// may not use this file except in compliance with the License.  You
// may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.  See the License for the specific language governing
// permissions and limitations under the License.

package kent

import (
	"errors"
	"fmt"
)

// Warning is a utility interface for warnings, allowing warnings to
// be differentiated from errors.  An object implementing Warning
// implements the error interface with a Warning method.
type Warning interface {
	error

	// Warning is the equivalent to Error for the error interface.
	// It should return the text of the warning, and so should be
	// equivalent to Error; the presence of this method is the
	// signal that an error is a Warning.
	Warning() string
}

// stringWarning is an implementation of Warning that is equivalent to
// that used by errors.New.
type stringWarning struct {
	msg string // The warning message
	err error  // The wrapped error
}

// Error returns the error message.
func (sw *stringWarning) Error() string {
	return sw.msg
}

// Warning is the equivalent to Error for the error interface.  It
// should return the text of the warning, and so should be equivalent
// to Error; the presence of this method is the signal that an error
// is a Warning.
func (sw *stringWarning) Warning() string {
	return sw.msg
}

// Unwrap returns the wrapped error, if there is one.
func (sw *stringWarning) Unwrap() error {
	return sw.err
}

// NewWarning constructs a new simple warning.  It is the equivalent
// of errors.New for warnings.
func NewWarning(text string) error {
	return &stringWarning{
		msg: text,
	}
}

// WarningWrap wraps an error with a Warning.
func WarningWrap(err error) error {
	return &stringWarning{
		msg: err.Error(),
		err: err,
	}
}

// Warningf constructs a new warning potentially wrapping another
// error or warning.  It is the equivalent of fmt.Errorsf for
// warnings.
func Warningf(format string, args ...interface{}) error {
	// Use Errorf to do the work
	tmp := fmt.Errorf(format, args...) //nolint:goerr113

	// Now make it a warning
	return &stringWarning{
		msg: tmp.Error(),
		err: errors.Unwrap(tmp),
	}
}

// IsWarning checks an error to determine if it is a warning.  It
// returns true if the error is a warning, false otherwise.
func IsWarning(err error) bool {
	for ; err != nil; err = errors.Unwrap(err) {
		if _, ok := err.(Warning); ok {
			return true
		}
	}

	return false
}
