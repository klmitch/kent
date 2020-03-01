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
	"io"
)

// WritingReporter is a Reporter that emits errors and warnings (with
// an appropriate "ERROR:" and "WARNING:" prefix) to a specified
// io.Writer stream.
type WritingReporter struct {
	out io.Writer // The output stream to write to
	rep Reporter  // Child reporter
}

// NewWritingReporter constructs a new writing reporter.  A writing
// reporter emits error and warning messages to a specified output
// stream (an io.Writer) with appropriate "ERROR:" and "WARNING:"
// prefixes.
func NewWritingReporter(out io.Writer, rep Reporter) *WritingReporter {
	return &WritingReporter{
		out: out,
		rep: rep,
	}
}

// Report is the core method of the Reporter interface.  It reports
// the error being reported, using whatever method the Reporter
// implementation uses, and passes on the error to the wrapped
// Reporter.
func (wr *WritingReporter) Report(err error) {
	var prefix string
	if IsWarning(err) {
		prefix = "WARNING"
	} else {
		prefix = "ERROR"
	}

	fmt.Fprintf(wr.out, "%s: %s\n", prefix, err)

	wr.rep.Report(err)
}

// Unwrap returns the Reporter or Reporters being wrapped by this
// Reporter, returning a possibly empty list of Reporters.
func (wr *WritingReporter) Unwrap() []Reporter {
	return []Reporter{wr.rep}
}
