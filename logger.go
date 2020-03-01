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
	"log"
)

// LoggingReporter is a Reporter that emits errors and warnings (with
// an appropriate "ERROR:" and "WARNING:" prefix) to a specified
// log.Logger.  If no log.Logger is specified, the default logger will
// be used.
type LoggingReporter struct {
	out    *log.Logger // The logger to emit to
	rep    Reporter    // Child reporter
	format *Formatters // Formatters to use
}

// NewLoggingReporter constructs a new logging reporter.  A logging
// reporter emits error and warning messages to a specified logger, or
// the default logger if the logger argument is nil; an appropriate
// "ERROR:" or "WARNING:" prefix is included in the emitted message.
// To use different formats for errors and warnings, pass appropriate
// formatting options, such as FormatError or FormatWarning.
func NewLoggingReporter(logger *log.Logger, rep Reporter, formatOptions ...FormatOption) *LoggingReporter {
	return &LoggingReporter{
		out:    logger,
		rep:    rep,
		format: newFormatters(formatOptions...),
	}
}

// emit is a helper that ensures that the output goes to the correct
// logger.  If the out element of the LoggingReporter instance is nil,
// the default logger is used.
func (lr *LoggingReporter) emit(msg string) {
	if lr.out == nil {
		log.Print(msg)
	} else {
		lr.out.Print(msg)
	}
}

// Report is the core method of the Reporter interface.  It reports
// the error being reported, using whatever method the Reporter
// implementation uses, and passes on the error to the wrapped
// Reporter.
func (lr *LoggingReporter) Report(err error) {
	lr.emit(lr.format.Format(err))

	lr.rep.Report(err)
}

// Unwrap returns the Reporter or Reporters being wrapped by this
// Reporter, returning a possibly empty list of Reporters.
func (lr *LoggingReporter) Unwrap() []Reporter {
	return []Reporter{lr.rep}
}
