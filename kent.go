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

// Package kent facilitates reporting of errors and warnings from deep
// within a call heirarchy.  The package is based on the Reporter
// interface, which allows reporting an error by calling the Report
// method.  All reporters also pass calls to Report on to their
// children, which are set up at the time the reporter is
// instantiated, allowing chains of reporters with various
// functionality.  An As function, similar to errors.As, allows a
// specific reporter to be selected from the chain, allowing access to
// additional methods that may be defined for some reporters, such as
// CountingReporter.
//
// Several Reporter implementations are provided, including
// CountingReporter, which counts errors and warnings;
// WritingReporter, which emits messages to an io.Writer with "ERROR:"
// and "WARNING:" prefixes; LoggingReporter, which is similar to
// WritingReporter except that it writes to a log.Logger; and
// TeeReporter, which allows writing to parallel reporters, with
// dynamic addition of additional reporters.  Additionally, a
// MockReporter is provided to facilitate testing of code that uses or
// manipulates Reporter instances.
//
// Both NewLoggingReporter and NewWritingReporter accept options of
// type FormatOption.  These options can be used to specify how the
// errors or warnings should be formatted for emission to the
// log.Logger or the io.Writer, respectively, overriding the prefixes
// described above.  The available format options allow setting a
// simple format string with FormatError and FormatWarning, or setting
// a function which takes the error and returns a string with
// FormatErrorFunc and FormatWarningFunc.
package kent

import (
	"container/list"
	"reflect"
)

// Reporter is the main interface of the kent package.  An object
// implementing Reporter can be used as a reporter, and can
// encapsulate other reporters.
type Reporter interface {
	// Report is the core method of the Reporter interface.  It
	// reports the error being reported, using whatever method the
	// Reporter implementation uses, and passes on the error to
	// the wrapped Reporter.
	Report(err error)

	// Unwrap returns the Reporter or Reporters being wrapped by
	// this Reporter, returning a possibly empty list of
	// Reporters.
	Unwrap() []Reporter
}

// reporterType is the type of Reporter.
var reporterType = reflect.TypeOf((*Reporter)(nil)).Elem()

// As is a helper that follows a chain of wrapped Reporters to select
// the first Reporter that is assignable to the specified target.  It
// returns a boolean true if the assignment was successful, false
// otherwise.
func As(rep Reporter, target interface{}) bool {
	// Make sure we have a target
	if target == nil {
		panic("must have a target to assign to")
	}

	// Make sure it's a pointer that can be assigned to
	value := reflect.ValueOf(target)
	valType := value.Type()
	if valType.Kind() != reflect.Ptr || value.IsNil() {
		panic("target must be a non-nil pointer")
	}

	// Make sure the thing it points to implements Reporter
	elemType := valType.Elem()
	if elemType.Kind() != reflect.Interface && !elemType.Implements(reporterType) {
		panic("*target must be interface or implement Reporter")
	}

	// Because Unwrap returns a list, let's use a work queue and a
	// breadth-first search to evaluate the possibilities
	q := list.List{}
	q.PushBack(rep)
	for q.Len() > 0 {
		item := q.Front().Value.(Reporter)
		q.Remove(q.Front())

		// Is it the one we're looking for?
		if reflect.TypeOf(item).AssignableTo(elemType) {
			value.Elem().Set(reflect.ValueOf(item))
			return true
		}

		// OK, add the children to the work queue
		children := item.Unwrap()
		if children != nil && len(children) > 0 {
			for _, child := range children {
				q.PushBack(child)
			}
		}
	}

	return false
}
