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

func TestAsBase(t *testing.T) {
	parent := root
	rep := &MockReporter{}
	rep.On("Unwrap").Return([]Reporter{parent})

	var target *rootReporter
	result := As(rep, &target)

	assert.True(t, result)
	assert.Same(t, parent, target)
	rep.AssertExpectations(t)
}

func TestAsNotFound(t *testing.T) {
	rep := &MockReporter{}
	rep.On("Unwrap").Return([]Reporter{})

	target := root
	result := As(rep, &target)

	assert.False(t, result)
	assert.Same(t, root, target)
	rep.AssertExpectations(t)
}

func TestAsNil(t *testing.T) {
	rep := &MockReporter{}

	assert.PanicsWithValue(t, "must have a target to assign to", func() { As(rep, nil) })
	rep.AssertExpectations(t)
}

func TestAsNilPointer(t *testing.T) {
	rep := &MockReporter{}

	assert.PanicsWithValue(t, "target must be a non-nil pointer", func() { As(rep, (*rootReporter)(nil)) })
	rep.AssertExpectations(t)
}

func TestAsBadTarget(t *testing.T) {
	rep := &MockReporter{}

	var target struct{}
	assert.PanicsWithValue(t, "*target must be interface or implement Reporter", func() { As(rep, &target) })
	rep.AssertExpectations(t)
}
