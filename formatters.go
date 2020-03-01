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
	"fmt"
	"strings"
)

// newFormatters is a patch point to allow the client Reporter
// instances to be tested in isolation from NewFormatters.
var newFormatters func(...FormatOption) *Formatters = NewFormatters

// FormatFunc describes a function that formats an error or warning.
// It will be passed the error and must return a string that will be
// used for the format of the error.
type FormatFunc func(err error) string

// Formatters contains formatters for formatting errors and warnings.
type Formatters struct {
	errors   FormatFunc // Format function for handling errors
	warnings FormatFunc // Format function for handling warnings
}

// FormatOption is an option for setting fields of a Formatters
// structure.
type FormatOption func(*Formatters)

// formatFromString constructs a FormatFunc from a format string.
func formatFromString(format string) FormatFunc {
	return func(err error) string {
		return fmt.Sprintf(format, err)
	}
}

// FormatError specifies the format string for formatting errors.
func FormatError(format string) FormatOption {
	return func(f *Formatters) {
		f.errors = formatFromString(strings.TrimRight(format, "\n"))
	}
}

// FormatErrorFunc specifies the formatting function for formatting
// errors.
func FormatErrorFunc(fmtFunc FormatFunc) FormatOption {
	return func(f *Formatters) {
		f.errors = fmtFunc
	}
}

// FormatWarning specifies the format string for formatting warnings.
func FormatWarning(format string) FormatOption {
	return func(f *Formatters) {
		f.warnings = formatFromString(strings.TrimRight(format, "\n"))
	}
}

// FormatWarningFunc specifies the formatting function for formatting
// warnings.
func FormatWarningFunc(fmtFunc FormatFunc) FormatOption {
	return func(f *Formatters) {
		f.warnings = fmtFunc
	}
}

// NewFormatters constructs a new Formatters instance with the
// specified options.  A reasonable default is used for both format
// strings.
func NewFormatters(options ...FormatOption) *Formatters {
	obj := &Formatters{
		errors:   formatFromString("ERROR: %s"),
		warnings: formatFromString("WARNING: %s"),
	}

	// Apply options
	for _, opt := range options {
		opt(obj)
	}

	return obj
}

// Format formats the specified error according to whether it is an
// error or a warning.  It returns the formatted result.
func (f *Formatters) Format(err error) string {
	if IsWarning(err) {
		return f.warnings(err)
	}

	return f.errors(err)
}
