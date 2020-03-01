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

func TestMockReporterReport(t *testing.T) {
	obj := &MockReporter{}
	obj.On("Report", assert.AnError)

	obj.Report(assert.AnError)

	obj.AssertExpectations(t)
}

func TestMockReporterUnwrapNil(t *testing.T) {
	obj := &MockReporter{}
	obj.On("Unwrap").Return(nil)

	result := obj.Unwrap()

	assert.Nil(t, result)
	obj.AssertExpectations(t)
}

func TestMockReporterUnwrapNonNil(t *testing.T) {
	expected := []Reporter{&MockReporter{}, &MockReporter{}}
	obj := &MockReporter{}
	obj.On("Unwrap").Return(expected)

	result := obj.Unwrap()

	assert.Equal(t, expected, result)
	obj.AssertExpectations(t)
}
