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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWritingReporterImplementsReporter(t *testing.T) {
	assert.Implements(t, (*Reporter)(nil), &WritingReporter{})
}

func TestNewWritingReporter(t *testing.T) {
	out := &bytes.Buffer{}
	rep := &MockReporter{}

	result := NewWritingReporter(out, rep)

	assert.Equal(t, &WritingReporter{
		out: out,
		rep: rep,
	}, result)
}

func TestWritingReporterReportError(t *testing.T) {
	rep := &MockReporter{}
	rep.On("Report", assert.AnError)
	out := &bytes.Buffer{}
	obj := &WritingReporter{
		out: out,
		rep: rep,
	}

	obj.Report(assert.AnError)

	assert.Equal(t, fmt.Sprintf("ERROR: %s\n", assert.AnError), out.String())
	rep.AssertExpectations(t)
}

func TestWritingReporterReportWarning(t *testing.T) {
	err := NewWarning("test warning")
	rep := &MockReporter{}
	rep.On("Report", err)
	out := &bytes.Buffer{}
	obj := &WritingReporter{
		out: out,
		rep: rep,
	}

	obj.Report(err)

	assert.Equal(t, "WARNING: test warning\n", out.String())
	rep.AssertExpectations(t)
}

func TestWritingReporterUnwrap(t *testing.T) {
	rep := &MockReporter{}
	obj := &WritingReporter{
		rep: rep,
	}

	result := obj.Unwrap()

	assert.Equal(t, []Reporter{rep}, result)
}
