// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/ooxml"

	//init enums for marshal/unmarshal
	_ "github.com/plandem/xlsx/format/conditional"
	_ "github.com/plandem/xlsx/format/conditional/rule"
	_ "github.com/plandem/xlsx/format/styles"
	_ "github.com/plandem/xlsx/internal/ml"
	_ "github.com/plandem/xlsx/internal/ml/primitives"
	_ "github.com/plandem/xlsx/types"
	_ "github.com/plandem/xlsx/types/comment"
	_ "github.com/plandem/xlsx/types/hyperlink"
	_ "github.com/plandem/xlsx/types/options/column"
	_ "github.com/plandem/xlsx/types/options/row"
	_ "github.com/plandem/xlsx/types/options/sheet"
)

//Open opens a XLSX file with name or io.Reader
func Open(f interface{}) (*Spreadsheet, error) {
	doc, err := ooxml.Open(f, newSpreadsheet)
	if err != nil {
		return nil, err
	}

	if xlDoc, ok := doc.(*Spreadsheet); ok {
		return xlDoc, nil
	}

	return nil, ooxml.ErrorUnknownPackage(Spreadsheet{})
}

//New creates and returns a new XLSX document
func New() *Spreadsheet {
	if doc, err := newSpreadsheet(ooxml.NewPackage(nil)); err == nil {
		if xlDoc, ok := doc.(*Spreadsheet); ok {
			return xlDoc
		}
	}

	panic("Could not create a new XLSX document.")
}
