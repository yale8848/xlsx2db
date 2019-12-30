// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"github.com/plandem/xlsx/format/styles"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types/options/column"
)

//Col is a higher level object that wraps ml.Col with functionality. Inherits functionality of Range
type Col struct {
	ml *ml.Col
	*Range
}

//Cell returns cell of col at row with rowIndex
func (c *Col) Cell(rowIndex int) *Cell {
	return c.sheet.Cell(c.bounds.FromCol, rowIndex)
}

//SetOptions sets options for column
func (c *Col) SetOptions(o *options.Info) {
	if o.Width > 0 {
		c.ml.Width = o.Width
		c.ml.CustomWidth = true
	}

	if o.Format != nil {
		c.SetStyles(o.Format)
	}

	c.ml.OutlineLevel = o.OutlineLevel
	c.ml.Hidden = o.Hidden
	c.ml.Collapsed = o.Collapsed
	c.ml.Phonetic = o.Phonetic
}

//Styles returns DirectStyleID of default format for column
func (c *Col) Styles() styles.DirectStyleID {
	return c.ml.Style
}

//SetStyles sets default style for the column. Affects cells not yet allocated in the column. In other words, this style applies to new cells.
func (c *Col) SetStyles(s interface{}) {
	c.ml.Style = c.sheet.info().resolveStyleID(s)
}

//CopyTo copies col cells into another col with cIdx index.
//N.B.: Merged cells are not supported
func (c *Col) CopyTo(cIdx int, withOptions bool) {
	if withOptions {
		//TODO: copy col options
		panic(errorNotSupported)
	}

	//copy cell data
	c.Range.CopyTo(cIdx, c.Range.bounds.FromRow)
}
