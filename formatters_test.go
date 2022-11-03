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
	"github.com/stretchr/testify/require"
)

func TestFormatFromString(t *testing.T) {
	err := errors.New("test error") //nolint:goerr113

	fmtFunc := formatFromString("test: %s")
	result := fmtFunc(err)

	assert.Equal(t, "test: test error", result)
}

func TestFormatError(t *testing.T) {
	err := errors.New("test error") //nolint:goerr113
	obj := &Formatters{}

	opt := FormatError("test: %s\n\n")
	opt(obj)

	require.NotNil(t, obj.errors)
	result := obj.errors(err)
	assert.Equal(t, "test: test error", result)
}

func TestFormatErrorFunc(t *testing.T) {
	err := errors.New("test error") //nolint:goerr113
	fmtFuncCalled := false
	fmtFunc := func(fErr error) string {
		assert.Same(t, err, fErr)
		fmtFuncCalled = true
		return "formatted"
	}
	obj := &Formatters{}

	opt := FormatErrorFunc(fmtFunc)
	opt(obj)

	require.NotNil(t, obj.errors)
	result := obj.errors(err)
	assert.Equal(t, "formatted", result)
	assert.True(t, fmtFuncCalled)
}

func TestFormatWarning(t *testing.T) {
	err := NewWarning("test warning")
	obj := &Formatters{}

	opt := FormatWarning("test: %s\n\n")
	opt(obj)

	require.NotNil(t, obj.warnings)
	result := obj.warnings(err)
	assert.Equal(t, "test: test warning", result)
}

func TestFormatWarningFunc(t *testing.T) {
	err := NewWarning("test warning")
	fmtFuncCalled := false
	fmtFunc := func(fErr error) string {
		assert.Same(t, err, fErr)
		fmtFuncCalled = true
		return "formatted"
	}
	obj := &Formatters{}

	opt := FormatWarningFunc(fmtFunc)
	opt(obj)

	require.NotNil(t, obj.warnings)
	result := obj.warnings(err)
	assert.Equal(t, "formatted", result)
	assert.True(t, fmtFuncCalled)
}

func TestNewFormatters(t *testing.T) {
	var opt1Called *Formatters
	var opt2Called *Formatters
	options := []FormatOption{
		func(f *Formatters) {
			opt1Called = f
		},
		func(f *Formatters) {
			opt2Called = f
		},
	}

	result := NewFormatters(options...)

	require.NotNil(t, result.errors)
	assert.Equal(t, "ERROR: test error", result.errors(errors.New("test error"))) //nolint:goerr113
	require.NotNil(t, result.warnings)
	assert.Equal(t, "WARNING: test warning", result.warnings(NewWarning("test warning")))
	assert.Same(t, result, opt1Called)
	assert.Same(t, result, opt2Called)
}

func TestFormattersFormatWarning(t *testing.T) {
	err := NewWarning("test warning")
	obj := &Formatters{
		warnings: formatFromString("WARNING: %s"),
	}

	result := obj.Format(err)

	assert.Equal(t, "WARNING: test warning", result)
}

func TestFormattersFormatError(t *testing.T) {
	err := errors.New("test error") //nolint:goerr113
	obj := &Formatters{
		errors: formatFromString("ERROR: %s"),
	}

	result := obj.Format(err)

	assert.Equal(t, "ERROR: test error", result)
}
