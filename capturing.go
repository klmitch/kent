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

import "sync"

// CapturingReporter is a Reporter that captures all errors and
// warnings reported using it.
type CapturingReporter struct {
	sync.Mutex

	list         []error  // Reported errors
	max          int      // Maximum number of errors to capture
	discardFirst bool     // Discard policy for overflow
	rep          Reporter // Child reporter
}

// CapturingReporterOption describes an option for a
// CapturingReporter.
type CapturingReporterOption func(*CapturingReporter)

// MaxCaptured allows capping the maximum number of errors captured by
// the CapturingReporter.  The discardFirst option specifies how to
// handle overflow; if true, the first captured error will be
// discarded to make room, otherwise the newly reported error will not
// be added to the list.  Note that a count less than or equal to 0
// will cause the CapturingReporter to capture all errors.
func MaxCaptured(count int, discardFirst bool) CapturingReporterOption {
	return func(cr *CapturingReporter) {
		cr.max = count
		cr.discardFirst = discardFirst
	}
}

// NewCapturingReporter constructs a new CapturingReporter.
func NewCapturingReporter(rep Reporter, options ...CapturingReporterOption) *CapturingReporter {
	obj := &CapturingReporter{
		list: []error{},
		rep:  rep,
	}

	// Apply options
	for _, opt := range options {
		opt(obj)
	}

	return obj
}

// Report is the core method of the Reporter interface.  It reports
// the error being reported, using whatever method the Reporter
// implementation uses, and passes on the error to the wrapped
// Reporter.
func (cr *CapturingReporter) Report(err error) {
	// Lock the mutex for thread safety
	cr.Lock()
	defer cr.Unlock()

	// Apply the capturing policy
	if cr.max > 0 {
		if len(cr.list)+1 >= cr.max {
			if cr.discardFirst {
				cr.list = append(cr.list[1:], err)
			}

			// Pass on to child
			cr.rep.Report(err)
			return
		}
	}

	// Save the error
	cr.list = append(cr.list, err)

	// Pass on to child
	cr.rep.Report(err)
}

// Unwrap returns the Reporter or Reporters being wrapped by this
// Reporter, returning a possibly empty list of Reporters.
func (cr *CapturingReporter) Unwrap() []Reporter {
	return []Reporter{cr.rep}
}

// List returns the list of captured errors reported using the
// CapturingReporter.
func (cr *CapturingReporter) List() []error {
	// Lock the mutex for thread safety
	cr.Lock()
	defer cr.Unlock()

	return cr.list
}
