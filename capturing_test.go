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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapturingReporterImplementsReporter(t *testing.T) {
	assert.Implements(t, (*Reporter)(nil), &CapturingReporter{})
}

func TestMaxCaptured(t *testing.T) {
	obj := &CapturingReporter{}

	opt := MaxCaptured(5, true)
	opt(obj)

	assert.Equal(t, &CapturingReporter{
		max:          5,
		discardFirst: true,
	}, obj)
}

func TestNewCapturingReporterBase(t *testing.T) {
	rep := &MockReporter{}

	result := NewCapturingReporter(rep)

	assert.Equal(t, &CapturingReporter{
		list: []error{},
		rep:  rep,
	}, result)
}

func TestNewCapturingReporterOptions(t *testing.T) {
	rep := &MockReporter{}
	var opt1Called, opt2Called *CapturingReporter
	options := []CapturingReporterOption{
		func(cr *CapturingReporter) {
			opt1Called = cr
		},
		func(cr *CapturingReporter) {
			opt2Called = cr
		},
	}

	result := NewCapturingReporter(rep, options...)

	assert.Equal(t, &CapturingReporter{
		list: []error{},
		rep:  rep,
	}, result)
	assert.Same(t, result, opt1Called)
	assert.Same(t, result, opt2Called)
}

func TestCapturingReporterReportBase(t *testing.T) {
	rep := &MockReporter{}
	rep.On("Report", assert.AnError)
	obj := &CapturingReporter{
		rep: rep,
	}

	obj.Report(assert.AnError)

	assert.Equal(t, []error{assert.AnError}, obj.list)
	rep.AssertExpectations(t)
}

func TestCapturingReporterReportOverflowDiscardFirst(t *testing.T) {
	err := errors.New("test error") //nolint:goerr113
	rep := &MockReporter{}
	rep.On("Report", err)
	obj := &CapturingReporter{
		list:         []error{assert.AnError, assert.AnError, assert.AnError},
		max:          3,
		discardFirst: true,
		rep:          rep,
	}

	obj.Report(err)

	assert.Equal(t, []error{assert.AnError, assert.AnError, err}, obj.list)
	rep.AssertExpectations(t)
}

func TestCapturingReporterReportOverflow(t *testing.T) {
	err := errors.New("test error") //nolint:goerr113
	rep := &MockReporter{}
	rep.On("Report", err)
	obj := &CapturingReporter{
		list: []error{assert.AnError, assert.AnError, assert.AnError},
		max:  3,
		rep:  rep,
	}

	obj.Report(err)

	assert.Equal(t, []error{assert.AnError, assert.AnError, assert.AnError}, obj.list)
	rep.AssertExpectations(t)
}

func TestCapturingReporterUnwrap(t *testing.T) {
	rep := &MockReporter{}
	obj := &CapturingReporter{
		rep: rep,
	}

	result := obj.Unwrap()

	assert.Equal(t, []Reporter{rep}, result)
}

func TestCapturingReporterList(t *testing.T) {
	obj := &CapturingReporter{
		list: []error{assert.AnError, assert.AnError, assert.AnError},
	}

	result := obj.List()

	assert.Equal(t, []error{assert.AnError, assert.AnError, assert.AnError}, result)
}
