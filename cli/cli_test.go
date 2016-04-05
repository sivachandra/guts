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

package cli

import (
	"testing"
)

func TestParsingCmdNoArgs(t *testing.T) {
	cmdStr := "cmd "
	args, err := parseCmdStr(cmdStr)
	if err != nil {
		t.Errorf("Parsing command failed.\n%s", err.Error())
	}

	if len(args) != 1 {
		t.Errorf("Incrrect number of arguments in the parsed output.")
	}

	if args[0] != "cmd" {
		t.Errorf("Command name in the parsed output is incorrect.")
	}
}

func TestParsingCmdStr(t *testing.T) {
	cmdStr := "cmd arg1 arg2 arg3"
	args, err := parseCmdStr(cmdStr)
	if err != nil {
		t.Errorf("Parsing command failed.\n%s", err.Error())
	}

	if len(args) != 4 {
		t.Errorf("Incrrect number of arguments in the parsed output.")
	}

	if args[0] != "cmd" {
		t.Errorf("Command name in the parsed output is incorrect.")
	}

	if args[1] != "arg1" {
		t.Errorf("Arg in the parsed output is incorrect.")
	}

	if args[2] != "arg2" {
		t.Errorf("Arg in the parsed output is incorrect.")
	}

	if args[3] != "arg3" {
		t.Errorf("Arg name in the parsed output is incorrect.")
	}
}

func TestParsingCmdStrWithQuotedArgSimple(t *testing.T) {
	cmdStr := ("cmd qarg1=\"Hello, \\\"World\\\"\" arg=not-quoted qarg2 \"Hello, Again\" " +
                   "qarg3 \"Hello\\\\\" qarg4 \"Hello \\\\\\\"Quote\\\\\\\"\"")
	args, err := parseCmdStr(cmdStr)
	if err != nil {
		t.Errorf("Parsing command failed.\n%s", err.Error())
		return
	}

	if args[0] != "cmd" {
		t.Errorf("Command name in the parsed output is incorrect.")
	}

	expected := "qarg1=Hello, \"World\""
	if args[1] != expected {
		t.Errorf("Error. Expected: %s; Found: %s.", expected, args[1])
	}

	expected = "arg=not-quoted"
	if args[2] != expected {
		t.Errorf("Error. Expected: %s; Found: %s.", expected, args[2])
	}

	expected = "qarg2"
	if args[3] != expected {
		t.Errorf("Error. Expected: %s; Found: %s.", expected, args[3])
	}

	expected = "Hello, Again"
	if args[4] != expected {
		t.Errorf("Error. Expected: %s; Found: %s.", expected, args[4])
	}

	expected = "qarg3"
	if args[5] != expected {
		t.Errorf("Error. Expected: %s; Found: %s.", expected, args[5])
	}

	expected = "Hello\\"
	if args[6] != expected {
		t.Errorf("Error. Expected: %s; Found: %s.", expected, args[6])
	}

	expected = "qarg4"
	if args[7] != expected {
		t.Errorf("Error. Expected: %s; Found: %s.", expected, args[7])
	}

	expected = "Hello \\\"Quote\\\""
	if args[8] != expected {
		t.Errorf("Error. Expected: %s; Found: %s.", expected, args[8])
	}
}
