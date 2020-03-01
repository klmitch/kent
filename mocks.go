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

import "github.com/stretchr/testify/mock"

// MockReporter is a mock object implementing the Reporter interface.
// It is provided to facilitate testing with code that uses or
// otherwise manipulates objects implementing Reporter.
type MockReporter struct {
	mock.Mock
}

// Report is the core method of the Reporter interface.  It reports
// the error being reported, using whatever method the Reporter
// implementation uses, and passes on the error to the wrapped
// Reporter.
func (m *MockReporter) Report(err error) {
	m.MethodCalled("Report", err)
}

// Unwrap returns the Reporter or Reporters being wrapped by this
// Reporter, returning a possibly empty list of Reporters.
func (m *MockReporter) Unwrap() []Reporter {
	args := m.MethodCalled("Unwrap")

	if tmp := args.Get(0); tmp != nil {
		return tmp.([]Reporter)
	}

	return nil
}
