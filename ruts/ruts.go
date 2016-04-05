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

// Package ruts provides utility functions on io.ByteReader.
package ruts

import (
	"io"
)

import (
	"guts/ucon"
)

// Read until the delim byte. The returned slice of bytes does _not_ include the
// delim byte. However, the delim byte is read out from the reader.
func ReadUntil(r io.ByteReader, delim byte) ([]byte, error) {
	var str []byte

	for true {
		c, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		if c == delim {
			break
		}
		str = append(str, c)
	}

	return str, nil
}

// Read until the delim byte. The returned slice of bytes includes the
// delim byte.
func ReadBytes(r io.ByteReader, delim byte) ([]byte, error) {
	var str []byte

	for true {
		c, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		str = append(str, c)

		if c == delim {
			break
		}
	}

	return str, nil
}

// Read a null terminated string. The null bytes is _not_ included in the
// returned value.
func ReadCString(r io.ByteReader) (string, error) {
	b, err := ReadUntil(r, ucon.ASCIINull)
	return string(b), err
}
