// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package primitives

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//UnderlineType is a type to encode XSD CT_UnderlineProperty
type UnderlineType ml.Property

//MarshalXML marshal UnderlineType
func (t *UnderlineType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return (*ml.Property)(t).MarshalXML(e, start)
}

//UnmarshalXML unmarshal UnderlineType
func (t *UnderlineType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*ml.Property)(t).UnmarshalXML(d, start)
}
