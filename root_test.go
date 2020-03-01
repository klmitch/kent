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

func TestRootReporterReport(t *testing.T) {
	obj := &rootReporter{}

	obj.Report(assert.AnError)

	// Verifies that no panics occur
}

func TestRootReporterUnwrap(t *testing.T) {
	obj := &rootReporter{}

	result := obj.Unwrap()

	assert.Equal(t, []Reporter{}, result)
}

func TestRoot(t *testing.T) {
	result := Root()

	assert.Same(t, root, result)
}

func TestTODO(t *testing.T) {
	result := TODO()

	assert.Same(t, todo, result)
}
