// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/index"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//sharedStrings is a higher level object that wraps ml.SharedStrings with functionality
type sharedStrings struct {
	ml    ml.SharedStrings
	index index.Index
	doc   *Spreadsheet
	file  *ooxml.PackageFile
}

func newSharedStrings(f interface{}, doc *Spreadsheet) *sharedStrings {
	ss := &sharedStrings{
		doc: doc,
	}

	ss.file = ooxml.NewPackageFile(doc.pkg, f, &ss.ml, nil)

	if ss.file.IsNew() {
		ss.doc.pkg.ContentTypes().RegisterContent(ss.file.FileName(), internal.ContentTypeSharedStrings)
		ss.doc.relationships.AddFile(internal.RelationTypeSharedStrings, ss.file.FileName())
		ss.file.MarkAsUpdated()
	}

	return ss
}

func (ss *sharedStrings) afterLoad() {
	for i, s := range ss.ml.StringItem {
		_ = ss.index.Add(s, i)
	}
}

//get returns string item stored at index
func (ss *sharedStrings) get(index int) *ml.StringItem {
	ss.file.LoadIfRequired(ss.afterLoad)

	if index < len(ss.ml.StringItem) {
		return ss.ml.StringItem[index]
	}

	return nil
}

//addString adds a new string and return index for it
func (ss *sharedStrings) addString(value string) int {
	return ss.addText(&ml.StringItem{Text: primitives.Text(value)})
}

//addText adds a new StringItem and return index for it
func (ss *sharedStrings) addText(si *ml.StringItem) int {
	ss.file.LoadIfRequired(ss.afterLoad)

	//return sid if already exists
	if sid, ok := ss.index.Get(si); ok {
		return sid
	}

	sid := len(ss.ml.StringItem)
	ss.ml.StringItem = append(ss.ml.StringItem, si)
	_ = ss.index.Add(si, sid)

	ss.file.MarkAsUpdated()

	return sid
}
