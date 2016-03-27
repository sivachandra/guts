///////////////////////////////////////////////////////////////////////////
// Copyright 2016 Siva Chandra
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
///////////////////////////////////////////////////////////////////////////

package xet

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	f, err := os.Open("test_data/simple.xml")
	if err != nil {
		t.Errorf("Unable to open test file for reading.\n%s", err.Error())
		return
	}
	defer f.Close()

	e, err := Read(f)
	if err != nil {
		t.Errorf("Error reading test file.\n%s", err.Error())
	}

	if e.Name.Local != "phonebook" {
		t.Errorf("Wrong tag of the root node.")
	}

	if len(e.Attr) != 1 {
		t.Errorf("Wrong number of attributes for the root node.")
	}

	if e.Attr[0].Name.Local != "user" {
		t.Errorf("Wrong attribute name for the attribute in the root node.")
	}

	if e.Attr[0].Value != "sivachandra" {
		t.Errorf("Wrong attribute value for the attribute in the root node.")
	}

	if len(e.Children) != 3 {
		t.Errorf("Wrong number of children to the root node.")
	}

	for i, c := range e.Children {
		if c.Type() != TypeElement {
			t.Errorf("Wrong type of child %d of the root node.", i)
		}

		ec := c.(*Node)
		if ec.Name.Local != "contact" {
			t.Errorf("Wrong tag of child %d of the root node.", i)
		}
	}
}
