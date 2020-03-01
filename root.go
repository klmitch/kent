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

// rootReporter provides the root of the Reporter tree.  This should
// be the only Reporter that returns an empty list in the Unwrap()
// method.
type rootReporter struct{}

// Report is the core method of the Reporter interface.  It reports
// the error being reported, using whatever method the Reporter
// implementation uses, and passes on the error to the wrapped
// Reporter.
func (rr *rootReporter) Report(err error) {
	// Does nothing
}

// Unwrap returns the Reporter or Reporters being wrapped by this
// Reporter, returning a possibly empty list of Reporters.
func (rr *rootReporter) Unwrap() []Reporter {
	return []Reporter{}
}

// Available root reporters.  This is modeled on the way go does
// contexts, so these are singletons that are then returned by Root()
// and TODO().
var (
	root = &rootReporter{}
	todo = &rootReporter{}
)

// Root returns a root reporter.  A root reporter does nothing in its
// Report() method, and has no wrapped reporters.
func Root() Reporter {
	return root
}

// TODO returns a root reporter, and is intended for marking where
// work is to be done to create an actual reporter.
func TODO() Reporter {
	return todo
}
