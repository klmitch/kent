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
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/klmitch/patcher"
	"github.com/stretchr/testify/assert"
)

func TestLoggingReporterImplementsReporter(t *testing.T) {
	assert.Implements(t, (*Reporter)(nil), &LoggingReporter{})
}

func TestNewLoggingReporter(t *testing.T) {
	out := &log.Logger{}
	rep := &MockReporter{}

	result := NewLoggingReporter(out, rep)

	assert.Equal(t, &LoggingReporter{
		out: out,
		rep: rep,
	}, result)
}

func TestLoggingReporterEmitDefaultLogger(t *testing.T) {
	stream := &bytes.Buffer{}
	defer patcher.Log(stream).Install().Restore()
	obj := &LoggingReporter{}

	obj.emit("this is a test")

	assert.Contains(t, stream.String(), "this is a test")
}

func TestLoggingReporterEmitSpecifiedLogger(t *testing.T) {
	stream := &bytes.Buffer{}
	out := log.New(stream, "", 0)
	obj := &LoggingReporter{
		out: out,
	}

	obj.emit("this is a test")

	assert.Equal(t, "this is a test\n", stream.String())
}

func TestLoggingReporterReportError(t *testing.T) {
	rep := &MockReporter{}
	rep.On("Report", assert.AnError)
	stream := &bytes.Buffer{}
	out := log.New(stream, "", 0)
	obj := &LoggingReporter{
		out: out,
		rep: rep,
	}

	obj.Report(assert.AnError)

	assert.Equal(t, fmt.Sprintf("ERROR: %s\n", assert.AnError), stream.String())
	rep.AssertExpectations(t)
}

func TestLoggingReporterReportWarning(t *testing.T) {
	err := NewWarning("test warning")
	rep := &MockReporter{}
	rep.On("Report", err)
	stream := &bytes.Buffer{}
	out := log.New(stream, "", 0)
	obj := &LoggingReporter{
		out: out,
		rep: rep,
	}

	obj.Report(err)

	assert.Equal(t, "WARNING: test warning\n", stream.String())
	rep.AssertExpectations(t)
}

func TestLoggingReporterUnwrap(t *testing.T) {
	rep := &MockReporter{}
	obj := &LoggingReporter{
		rep: rep,
	}

	result := obj.Unwrap()

	assert.Equal(t, []Reporter{rep}, result)
}
