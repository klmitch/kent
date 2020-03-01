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

func TestTeeReporterImplementsReporter(t *testing.T) {
	assert.Implements(t, (*Reporter)(nil), &TeeReporter{})
}

func TestNewTeeReporter(t *testing.T) {
	reps := []Reporter{&MockReporter{}, &MockReporter{}, &MockReporter{}}

	result := NewTeeReporter(reps...)

	assert.Equal(t, &TeeReporter{
		reps: reps,
	}, result)
}

func TestTeeReporterReport(t *testing.T) {
	reps := []Reporter{&MockReporter{}, &MockReporter{}, &MockReporter{}}
	for _, rep := range reps {
		rep.(*MockReporter).On("Report", assert.AnError)
	}
	obj := &TeeReporter{
		reps: reps,
	}

	obj.Report(assert.AnError)

	for _, rep := range reps {
		rep.(*MockReporter).AssertExpectations(t)
	}
}

func TestTeeReporterUnwrap(t *testing.T) {
	reps := []Reporter{&MockReporter{}, &MockReporter{}, &MockReporter{}}
	obj := &TeeReporter{
		reps: reps,
	}

	result := obj.Unwrap()

	assert.Equal(t, reps, result)
}

func TestTeeReporterAdd(t *testing.T) {
	reps := []Reporter{&MockReporter{}, &MockReporter{}, &MockReporter{}}
	obj := &TeeReporter{}

	obj.Add(reps...)

	assert.Equal(t, reps, obj.reps)
}

func TestTeeReporterRemove(t *testing.T) {
	reps := []Reporter{&MockReporter{}, &rootReporter{}, &MockReporter{}}
	obj := &TeeReporter{
		reps: reps,
	}

	obj.Remove(&rootReporter{})

	assert.Equal(t, []Reporter{&MockReporter{}, &MockReporter{}}, obj.reps)
}
