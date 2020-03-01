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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringWarningImplementsError(t *testing.T) {
	assert.Implements(t, (*error)(nil), &stringWarning{})
}

func TestStringWarningImplementsWarning(t *testing.T) {
	assert.Implements(t, (*Warning)(nil), &stringWarning{})
}

func TestStringWarningError(t *testing.T) {
	obj := &stringWarning{
		msg: "some message",
	}

	result := obj.Error()

	assert.Equal(t, "some message", result)
}

func TestStringWarningWarning(t *testing.T) {
	obj := &stringWarning{
		msg: "some message",
	}

	result := obj.Warning()

	assert.Equal(t, "some message", result)
}

func TestStringWarningUnwrap(t *testing.T) {
	obj := &stringWarning{
		err: assert.AnError,
	}

	result := obj.Unwrap()

	assert.Same(t, assert.AnError, result)
}

func TestNewWarning(t *testing.T) {
	obj := NewWarning("some message")

	assert.Equal(t, &stringWarning{
		msg: "some message",
	}, obj)
}

func TestWarningWrap(t *testing.T) {
	obj := WarningWrap(assert.AnError)

	assert.Equal(t, &stringWarning{
		msg: assert.AnError.Error(),
		err: assert.AnError,
	}, obj)
}

func TestWarningf(t *testing.T) {
	obj := Warningf("this is a test %w", assert.AnError)

	assert.Equal(t, &stringWarning{
		msg: fmt.Sprintf("this is a test %s", assert.AnError),
		err: assert.AnError,
	}, obj)
}

func TestIsWarningDirect(t *testing.T) {
	err := NewWarning("test warning")

	result := IsWarning(err)

	assert.True(t, result)
}

func TestIsWarningWrapped(t *testing.T) {
	err := fmt.Errorf("this is a test %w", NewWarning("test warning"))

	result := IsWarning(err)

	assert.True(t, result)
}

func TestIsWarningErrorOnly(t *testing.T) {
	err := fmt.Errorf("this is a test %w", assert.AnError)

	result := IsWarning(err)

	assert.False(t, result)
}
