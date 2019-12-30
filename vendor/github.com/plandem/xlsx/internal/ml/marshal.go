// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ml

import (
	"encoding/xml"
	"github.com/plandem/ooxml/ml"
)

//DiffStyleList is a direct mapping of XSD CT_Dxfs
type DiffStyleList struct {
	Count int          `xml:"count,attr"`
	Items []*DiffStyle `xml:"dxf,omitempty"`
}

//NamedStyleInfoList is a direct mapping of XSD CT_CellStyles
type NamedStyleInfoList struct {
	Count int               `xml:"count,attr"`
	Items []*NamedStyleInfo `xml:"cellStyle,omitempty"`
}

//DirectStyleList is a direct mapping of XSD CT_CellXfs
type DirectStyleList struct {
	Count int            `xml:"count,attr"`
	Items []*DirectStyle `xml:"xf,omitempty"`
}

//NamedStyleList is a direct mapping of XSD cellStyleXfs
type NamedStyleList struct {
	Count int           `xml:"count,attr"`
	Items []*NamedStyle `xml:"xf,omitempty"`
}

//BorderList is a direct mapping of XSD CT_Borders
type BorderList struct {
	Count int       `xml:"count,attr"`
	Items []*Border `xml:"border,omitempty"`
}

//FontList is a direct mapping of XSD CT_Fonts
type FontList struct {
	Count int     `xml:"count,attr"`
	Items []*Font `xml:"font,omitempty"`
}

//FillList is a direct mapping of XSD CT_Fills
type FillList struct {
	Count int     `xml:"count,attr"`
	Items []*Fill `xml:"fill,omitempty"`
}

//NumberFormatList is a direct mapping of XSD CT_NumFmts
type NumberFormatList struct {
	Count int             `xml:"count,attr"`
	Items []*NumberFormat `xml:"numFmt,omitempty"`
}

//HyperlinkList is a direct mapping of XSD CT_Hyperlinks
type HyperlinkList struct {
	Items []*Hyperlink `xml:"hyperlink,omitempty"`
}

//MergedCellList is a direct mapping of XSD CT_MergeCells
type MergedCellList struct {
	Count int          `xml:"count,attr"`
	Items []*MergeCell `xml:"mergeCell,omitempty"`
}

//SheetViewList is a direct mapping of XSD CT_SheetViews
type SheetViewList struct {
	Items  []*SheetView `xml:"sheetView,omitempty"`
	ExtLst *ml.Reserved `xml:"extLst,omitempty"`
}

//BookViewList is a direct mapping of XSD CT_BookViews
type BookViewList struct {
	Items []*BookView `xml:"workbookView,omitempty"`
}

//ExternalReferenceList is a direct mapping of XSD CT_ExternalReferences
type ExternalReferenceList struct {
	Items []*ExternalReference `xml:"workbookView,omitempty"`
}

//MarshalXML marshal DiffStyleList
func (r *DiffStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal NamedStyleInfoList
func (r *NamedStyleInfoList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal DirectStyleList
func (r *DirectStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal NamedStyleList
func (r *NamedStyleList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal BorderList
func (r *BorderList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal FontList
func (r *FontList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal FillList
func (r *FillList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal NumberFormatList
func (r *NumberFormatList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal HyperlinkList
func (r *HyperlinkList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(r.Items) > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal MergedCellList
func (r *MergedCellList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if r.Count = len(r.Items); r.Count > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal SheetViewList
func (r *SheetViewList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(r.Items) > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal BookViewList
func (r *BookViewList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(r.Items) > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal ExternalReferenceList
func (r *ExternalReferenceList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(r.Items) > 0 {
		return e.EncodeElement(*r, start)
	}

	return nil
}

//MarshalXML marshal AutoFilter
func (r *AutoFilter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !r.Bounds.IsEmpty() {
		return e.EncodeElement(*r, start)
	}

	return nil
}
