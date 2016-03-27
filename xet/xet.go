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
	"encoding/xml"
	"fmt"
	"io"
)

type NodeType uint8

const (
	TypeElement   = NodeType(1)
	TypeComment   = NodeType(2)
	TypeDirective = NodeType(3)
	TypeProcInst  = NodeType(4)
)

type TreeNode interface {
	Type() NodeType
}

type Node struct {
	Name     xml.Name
	Attr     []xml.Attr
	CharData xml.CharData
	Children []TreeNode
}

func (e *Node) Type() NodeType {
	return TypeElement
}

type Comment xml.Comment

func (c Comment) Type() NodeType {
	return TypeComment
}

type Directive xml.Directive

func (d Directive) Type() NodeType {
	return TypeDirective
}

type ProcInst xml.ProcInst

func (p ProcInst) Type() NodeType {
	return TypeProcInst
}

func Read(r io.Reader) (*Node, error) {
	d := xml.NewDecoder(r)
	if d == nil {
		return nil, fmt.Errorf("Error creating an XML decoder.")
	}

	return readInternal(d, nil)
}

func readInternal(d *xml.Decoder, p *Node) (*Node, error) {
	for {
		t, err := d.Token()
		if err != nil {
			if t == nil {
				break
			}

			return nil, err
		}

		switch t.(type) {
		case xml.StartElement:
			se := t.(xml.StartElement)

			e := new(Node)
			e.Name = se.Name
			e.Attr = se.Attr

			e, err := readInternal(d, e)
			if err != nil {
				return nil, err
			}

			if p == nil {
				p = e
			} else {
				p.Children = append(p.Children, e)
			}
		case xml.EndElement:
			return p, nil
		case xml.CharData:
			p.CharData = t.(xml.CharData)
		case xml.Comment:
			p.Children = append(p.Children, Comment(t.(xml.Comment)))
		case xml.Directive:
			p.Children = append(p.Children, Directive(t.(xml.Directive)))
		case xml.ProcInst:
			p.Children = append(p.Children, ProcInst(t.(xml.ProcInst)))
		default:
		}
	}

	return p, nil
}
