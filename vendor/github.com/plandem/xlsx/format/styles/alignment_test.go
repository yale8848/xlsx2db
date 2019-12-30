// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"encoding/xml"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAlignment(t *testing.T) {
	style := New(
		Alignment.VAlign(VAlignBottom),
		Alignment.HAlign(HAlignFill),
		Alignment.TextRotation(90),
		Alignment.WrapText,
		Alignment.Indent(10),
		Alignment.RelativeIndent(5),
		Alignment.JustifyLastLine,
		Alignment.ShrinkToFit,
		Alignment.ReadingOrder(4),
	)

	require.IsType(t, &Info{}, style)
	require.Equal(t, createStylesAndFill(func(f *Info) {
		f.styleInfo.Alignment = &ml.CellAlignment{
			Vertical:        VAlignBottom,
			Horizontal:      HAlignFill,
			TextRotation:    90,
			WrapText:        true,
			Indent:          10,
			RelativeIndent:  5,
			JustifyLastLine: true,
			ShrinkToFit:     true,
			ReadingOrder:    4,
		}
	}), style)
}

func TestAlignmentMarshal(t *testing.T) {
	//0 must be omitted
	style := New(
		Alignment.TextRotation(0),
		Alignment.Indent(0),
		Alignment.RelativeIndent(0),
		Alignment.ReadingOrder(0),
	)
	encoded, err := xml.Marshal(&style.styleInfo.Alignment)
	require.Empty(t, err)
	require.Equal(t, `<CellAlignment></CellAlignment>`, string(encoded))

	//simple version
	style = New(
		Alignment.WrapText,
	)
	encoded, _ = xml.Marshal(&style.styleInfo.Alignment)
	require.Equal(t, `<CellAlignment wrapText="true"></CellAlignment>`, string(encoded))

	//full version
	style = New(
		Alignment.VAlign(VAlignBottom),
		Alignment.HAlign(HAlignFill),
		Alignment.TextRotation(90),
		Alignment.WrapText,
		Alignment.Indent(10),
		Alignment.RelativeIndent(5),
		Alignment.JustifyLastLine,
		Alignment.ShrinkToFit,
		Alignment.ReadingOrder(4),
	)
	encoded, _ = xml.Marshal(&style.styleInfo.Alignment)
	require.Equal(t, `<CellAlignment horizontal="fill" vertical="bottom" wrapText="true" justifyLastLine="true" shrinkToFit="true" textRotation="90" indent="10" relativeIndent="5" readingOrder="4"></CellAlignment>`, string(encoded))
}
