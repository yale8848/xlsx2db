// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
)

//VisibilityType is a type to encode XSD ST_Visibility and ST_SheetState
type VisibilityType byte

//nolint
var (
	ToVisibilityType   map[string]VisibilityType
	FromVisibilityType map[VisibilityType]string
)

func (e VisibilityType) String() string {
	return FromVisibilityType[e]
}

//MarshalXMLAttr marshal VisibilityType
func (e *VisibilityType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{Name: name}

	if v, ok := FromVisibilityType[*e]; ok {
		attr.Value = v
	} else {
		attr = xml.Attr{}
	}

	return attr, nil
}

//UnmarshalXMLAttr unmarshal VisibilityType
func (e *VisibilityType) UnmarshalXMLAttr(attr xml.Attr) error {
	if v, ok := ToVisibilityType[attr.Value]; ok {
		*e = v
	}

	return nil
}
