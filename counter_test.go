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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingReporterImplementsReporter(t *testing.T) {
	assert.Implements(t, (*Reporter)(nil), &CountingReporter{})
}

func TestNewCountingReporter(t *testing.T) {
	rep := &MockReporter{}

	result := NewCountingReporter(rep)

	assert.Equal(t, &CountingReporter{
		rep: rep,
	}, result)
}

func TestCountingReporterReportError(t *testing.T) {
	rep := &MockReporter{}
	rep.On("Report", assert.AnError)
	obj := &CountingReporter{
		rep: rep,
	}

	obj.Report(assert.AnError)

	assert.Equal(t, &CountingReporter{
		errors: int64(1),
		rep:    rep,
	}, obj)
	rep.AssertExpectations(t)
}

func TestCountingReporterReportWarning(t *testing.T) {
	err := NewWarning("a warning")
	rep := &MockReporter{}
	rep.On("Report", err)
	obj := &CountingReporter{
		rep: rep,
	}

	obj.Report(err)

	assert.Equal(t, &CountingReporter{
		warnings: int64(1),
		rep:      rep,
	}, obj)
	rep.AssertExpectations(t)
}

func TestCountingReporterUnwrap(t *testing.T) {
	rep := &MockReporter{}
	obj := &CountingReporter{
		rep: rep,
	}

	result := obj.Unwrap()

	assert.Equal(t, []Reporter{rep}, result)
}

func TestCountingReporterErrors(t *testing.T) {
	obj := &CountingReporter{
		errors: int64(42),
	}

	result := obj.Errors()

	assert.Equal(t, 42, result)
}

func TestCountingReporterWarnings(t *testing.T) {
	obj := &CountingReporter{
		warnings: int64(42),
	}

	result := obj.Warnings()

	assert.Equal(t, 42, result)
}
