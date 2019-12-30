// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//Bounds is alias of original primitives.Bounds type make it public.
type Bounds = primitives.Bounds

//BoundsList is alias of original primitives.BoundsList type make it public.
type BoundsList = primitives.BoundsList

//Ref is alias of original primitives.Ref type make it public.
type Ref = primitives.Ref

//RefList is alias of original primitives.RefList type make it public.
type RefList = primitives.RefList

//CellRef is alias of original primitives.CellRef type make it public.
type CellRef = primitives.CellRef

//Text is alias of original primitives.Text type make it public.
type Text = primitives.Text

//RefFromCellRefs is alias of original primitives.RefFromCellRefs to make it public
func RefFromCellRefs(from CellRef, to CellRef) Ref {
	return primitives.RefFromCellRefs(from, to)
}

//RefFromIndexes is alias of original primitives.RefFromIndexes to make it public
func RefFromIndexes(colIndex, rowIndex int) Ref {
	return primitives.RefFromIndexes(colIndex, rowIndex)
}

//CellRefFromIndexes is alias of original primitives.CellRefFromIndexes to make it public
func CellRefFromIndexes(colIndex, rowIndex int) CellRef {
	return primitives.CellRefFromIndexes(colIndex, rowIndex)
}

//BoundsFromIndexes is alias of original primitives.BoundsFromIndexes to make it public
func BoundsFromIndexes(fromCol, fromRow, toCol, toRow int) Bounds {
	return primitives.BoundsFromIndexes(fromCol, fromRow, toCol, toRow)
}
