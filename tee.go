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
	"reflect"
	"sync"
)

// TeeReporter is a Reporter that sends reports to a list of other
// Reporter instances.  This allows reporting of an error or warning
// to multiple output streams, for instance.
type TeeReporter struct {
	sync.Mutex

	reps []Reporter // Child reporters
}

// NewTeeReporter constructs a new tee reporter.  A tee reporter sends
// an error report to multiple subordinate child reporters.
func NewTeeReporter(reps ...Reporter) *TeeReporter {
	return &TeeReporter{
		reps: reps,
	}
}

// Report is the core method of the Reporter interface.  It reports
// the error being reported, using whatever method the Reporter
// implementation uses, and passes on the error to the wrapped
// Reporter.
func (tr *TeeReporter) Report(err error) {
	// Lock the mutex for thread safety
	tr.Lock()
	defer tr.Unlock()

	for _, rep := range tr.reps {
		rep.Report(err)
	}
}

// Unwrap returns the Reporter or Reporters being wrapped by this
// Reporter, returning a possibly empty list of Reporters.
func (tr *TeeReporter) Unwrap() []Reporter {
	// Lock the mutex for thread safety
	tr.Lock()
	defer tr.Unlock()

	// Return the reporter list
	return tr.reps
}

// Add adds 1 or more additional reporters to the tee.
func (tr *TeeReporter) Add(reps ...Reporter) {
	// Lock the mutex for thread safety
	tr.Lock()
	defer tr.Unlock()

	// Extend the reporters list
	tr.reps = append(tr.reps, reps...)
}

// Remove removes 1 or more reporters from the tee.
func (tr *TeeReporter) Remove(reps ...Reporter) {
	// Lock the mutex for thread safety
	tr.Lock()
	defer tr.Unlock()

	// Construct the new list of reporters
	newList := make([]Reporter, 0, len(tr.reps))
	for _, rep := range tr.reps {
		for _, cmpr := range reps {
			if !reflect.DeepEqual(rep, cmpr) {
				newList = append(newList, rep)
			}
		}
	}

	// Update the list of child reporters
	tr.reps = newList
}
