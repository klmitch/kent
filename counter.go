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

import "sync/atomic"

// CountingReporter is a Reporter that counts the number of errors and
// warnings that are reported using it.
type CountingReporter struct {
	errors   int64    // Number of errors
	warnings int64    // Number of warnings
	rep      Reporter // Child reporter
}

// NewCountingReporter constructs a new counting Reporter.  A counting
// reporter counts the number of errors and warnings that are reported
// using it, and makes those counts available through the Errors and
// Warnings methods.
func NewCountingReporter(rep Reporter) *CountingReporter {
	return &CountingReporter{
		rep: rep,
	}
}

// Report is the core method of the Reporter interface.  It reports
// the error being reported, using whatever method the Reporter
// implementation uses, and passes on the error to the wrapped
// Reporter.
func (cr *CountingReporter) Report(err error) {
	if IsWarning(err) {
		atomic.AddInt64(&cr.warnings, 1)
	} else {
		atomic.AddInt64(&cr.errors, 1)
	}

	cr.rep.Report(err)
}

// Unwrap returns the Reporter or Reporters being wrapped by this
// Reporter, returning a possibly empty list of Reporters.
func (cr *CountingReporter) Unwrap() []Reporter {
	return []Reporter{cr.rep}
}

// Errors returns the number of errors counted so far by the counting
// reporter.
func (cr *CountingReporter) Errors() int {
	count := atomic.LoadInt64(&cr.errors)

	return int(count)
}

// Warnings returns the number of warnings counted so far by the
// counting reporter.
func (cr *CountingReporter) Warnings() int {
	count := atomic.LoadInt64(&cr.warnings)

	return int(count)
}
