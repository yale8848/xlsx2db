// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xlsx

import (
	"encoding/xml"
	"fmt"
	"github.com/plandem/ooxml"
	"github.com/plandem/ooxml/drawing/vml"
	"github.com/plandem/ooxml/drawing/vml/css"
	"github.com/plandem/ooxml/index"
	sharedML "github.com/plandem/ooxml/ml"
	"github.com/plandem/xlsx/internal"
	"github.com/plandem/xlsx/internal/ml"
	"github.com/plandem/xlsx/types"
	"github.com/plandem/xlsx/types/comment"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type drawingsVML struct {
	sheet               *sheetInfo
	ml                  vml.Excel
	chunks              []int
	file                *ooxml.PackageFile
	initializedFile     bool
	initializedComments bool
	updated             bool
	nextShapeId         int
	nextShapeIdMax      int
	shapeIndex          index.Index
}

//capacity of chunk. vml file store shapes in chunks
const vmlChunkSize = 1024

//office's internal id of vml shape type for comments
const commentShapeTypeSpt = 202

var (
	regexpDrawings   = regexp.MustCompile(`xl/drawings/[[:alpha:]]+[\d]+\.vml`)
	regexpShapeID    = regexp.MustCompile(`_x0000_s([\d]+)`)
	commentShapeType = fmt.Sprintf("#_x0000_t%d", commentShapeTypeSpt)
)

func newDrawingsVML(sheet *sheetInfo) *drawingsVML {
	return &drawingsVML{
		sheet: sheet,
	}
}

//resolve chunks info of this VML drawings file
func (d *drawingsVML) resolveChunks() {
	if len(d.chunks) == 0 {
		d.attachFileIfRequired()

		var idMap *vml.IdMap

		//only non loaded existing files will return stream, otherwise chunks info can be gathered from existing ml info
		if stream, err := d.file.ReadStream(); err == nil {
			//locate chunks info
			for {
				t, _ := stream.Token()
				if t == nil {
					break
				}

				if start, ok := t.(xml.StartElement); ok && start.Name.Local == "shapelayout" {
					shapeLayout := &vml.ShapeLayout{}
					if err := stream.DecodeElement(shapeLayout, &start); err != nil {
						_ = stream.Close()
						panic(err)
					}

					idMap = shapeLayout.IdMap
					break
				}
			}

			_ = stream.Close()
		} else {
			idMap = d.ml.ShapeLayout.IdMap
		}

		//parse chunks info
		if idMap != nil && idMap.Data != "" {
			chunks := strings.Split(idMap.Data, ",")
			for _, s := range chunks {
				if n, err := strconv.Atoi(strings.TrimSpace(s)); err != nil {
					panic(fmt.Errorf("can't load chunks info from VML idmap: %s", err))
				} else {
					d.chunks = append(d.chunks, n)
				}
			}
		}
	}
}

func (d *drawingsVML) nextChunkID() int {
	nextChunk := 0

	//get maximum chunk - each VML file can have few non serial chunks (1024 shapes per chunk), e.g.: 1,3,9
	for _, s := range d.sheet.workbook.doc.sheets {
		s.drawingsVML.resolveChunks()

		for _, chunkID := range s.drawingsVML.chunks {
			nextChunk = int(math.Max(float64(nextChunk), float64(chunkID)))
		}
	}

	nextChunk++
	d.chunks = append(d.chunks, nextChunk)
	d.nextShapeId = nextChunk * vmlChunkSize
	d.nextShapeIdMax = d.nextShapeId + vmlChunkSize

	d.ml.ShapeLayout.IdMap.Data = strings.Trim(strings.Replace(fmt.Sprint(d.chunks), " ", ",", -1), "[]")

	return nextChunk
}

func (d *drawingsVML) nextShapeID() int {
	if d.nextShapeId >= d.nextShapeIdMax {
		d.nextChunkID()
	}

	d.nextShapeId++
	return d.nextShapeId
}

func (d *drawingsVML) addComment(bounds types.Bounds, info *comment.Info) error {
	d.initCommentsIfRequired()

	shape := &vml.Shape{}
	shape.ID = fmt.Sprintf("_x0000_s%d", d.nextShapeID())
	shape.Type = commentShapeType
	shape.FillColor = info.Background
	shape.InsetMode = vml.InsetModeAuto

	//TODO: add support for margins - right now no idea how to calculate it
	//"margin-left:242.25pt;margin-top:22.5pt"
	style := css.Style{
		ZIndex:   len(d.ml.Shape) + 1,
		Position: css.PositionAbsolute,
		Width:    css.NewNumber(info.Width),
		Height:   css.NewNumber(info.Height),
	}

	if info.Visible {
		style.Visible = css.VisibilityVisible
	} else {
		style.Visible = css.VisibilityHidden
	}

	shape.Style = style.String()
	shape.Fill = &vml.Fill{Color2: info.Background}
	shape.PathSettings = &vml.Path{ConnectType: vml.ConnectTypeNone}

	if len(info.Stroke) != 0 {
		shape.StrokeColor = info.Stroke
	}

	if len(info.Shadow) != 0 {
		shape.Shadow = &vml.Shadow{
			Color:    info.Shadow,
			On:       sharedML.TriStateTrue,
			Obscured: sharedML.TriStateTrue,
		}
	}

	//"1, 15, 0,  2, 3, 15, 3, 16"
	//"3, 15, 1, 10, 5, 15, 2, 64",
	//"1, 15, 0,  2, 2, 54, 5, 3"
	//"2, 15, 2, 14, 4, 23, 6, 19"
	//"1, 15, x, 5, 3, 15, x, 0"

	//TODO: add support for anchors - right now no idea how to calculate it
	//anchor := vml.ClientDataAnchor{
	//	LeftColumn:   3,
	//	LeftOffset:   15,
	//	TopRow:       1,
	//	TopOffset:    10,
	//	RightColumn:  5,
	//	RightOffset:  15,
	//	BottomRow:    2,
	//	BottomOffset: 16,
	//}

	shape.ClientData = &vml.ClientData{
		Row:    bounds.FromRow,
		Column: bounds.FromCol,
		Type:   vml.ObjectTypeNote,
		//Anchor:        anchor.String(),
		SizeWithCells: sharedML.TriStateBlankTrue(sharedML.TriStateTrue),
		MoveWithCells: sharedML.TriStateBlankTrue(sharedML.TriStateTrue),
		AutoFill:      sharedML.TriStateBlankTrue(sharedML.TriStateFalse),
	}

	if info.Visible {
		shape.ClientData.Visible = sharedML.TriStateBlankTrue(sharedML.TriStateTrue)
	}

	if err := d.shapeIndex.Add(shape, len(d.ml.Shape)); err != nil {
		return err
	}

	d.ml.Shape = append(d.ml.Shape, shape)
	d.file.MarkAsUpdated()
	d.updated = true
	return nil
}

//Remove removes comment info for bounds
func (d *drawingsVML) removeComment(bounds types.Bounds) {
	d.initCommentsIfRequired()

	shape := &vml.Shape{}

	//TODO: theoretically we can have few shape types for comments, but right now we use index as - shapeType+row+col
	shape.Type = commentShapeType
	shape.ClientData = &vml.ClientData{
		Column: bounds.FromCol,
		Row:    bounds.FromRow,
	}

	if id, ok := d.shapeIndex.Get(shape); ok {
		d.ml.Shape = append(d.ml.Shape[:id], d.ml.Shape[id+1:]...)
		d.shapeIndex.Remove(shape)
	}
}

//load all content if required or add minimal required
func (d *drawingsVML) initIfRequired() {
	if d.initializedFile {
		return
	}

	d.attachFileIfRequired()

	if !d.file.IsNew() {
		d.file.LoadIfRequired(nil)

		//resolve chunks info
		d.resolveChunks()

		//resolve next shapeID info
		nextShapeID := 0
		for _, shape := range d.ml.Shape {
			if matched := regexpShapeID.FindSubmatch([]byte(shape.ID)); len(matched) > 0 {
				if id, err := strconv.Atoi(string(matched[1])); err != nil {
					panic(fmt.Errorf("can't get ID of shape: %s", matched))
				} else {
					//TODO: theoretically we should take maximum shape_id of the lowest chunk from file if possible.
					// But it's premature optimization, so right now we just take maximum ID of shape, even if "lower" chunks can have more shapes.
					// E.g.: file had few chunks, but later few shapes were deleted from chunk1, so theoretically we could add next shapes to chunk1 again
					nextShapeID = int(math.Max(float64(nextShapeID), float64(id)))
				}
			}
		}

		d.nextShapeId = nextShapeID
		d.nextShapeIdMax = (nextShapeID/vmlChunkSize)*vmlChunkSize + vmlChunkSize
	} else {
		d.ml.ShapeLayout = &vml.ShapeLayout{
			Ext: vml.ExtTypeEdit,
			IdMap: &vml.IdMap{
				Ext: vml.ExtTypeEdit,
			},
		}

		//our resolving mechanism of next chunk relies on non nil object, that's why we should assign it manually
		d.ml.ShapeLayout.IdMap.Data = strconv.Itoa(d.nextChunkID())
	}

	d.file.MarkAsUpdated()
	d.initializedFile = true
	d.updated = true
	d.buildIndexes()
}

//attach LegacyDrawing info into sheet
func (d *drawingsVML) attachDrawingsRID() {
	if d.updated && (d.sheet.ml.LegacyDrawing == nil || d.sheet.ml.LegacyDrawing.RID == "") {
		fileName := d.sheet.relationships.GetTargetByType(internal.RelationTypeVmlDrawing)
		rid := d.sheet.relationships.GetIdByTarget(fileName)
		d.sheet.ml.LegacyDrawing = &ml.LegacyDrawing{RID: rid}
	}
}

//add shape type for comments if required
func (d *drawingsVML) initCommentsIfRequired() {
	if d.initializedComments {
		return
	}

	d.initIfRequired()

	//attach shape type if required
	for _, shapeType := range d.ml.ShapeType {
		if shapeType.Spt == commentShapeTypeSpt {
			return
		}
	}

	shapeType := &vml.ShapeType{}
	shapeType.ID = fmt.Sprintf("_x0000_t%d", commentShapeTypeSpt)
	shapeType.Spt = commentShapeTypeSpt
	shapeType.Path = "m,l,21600r21600,l21600,xe"
	shapeType.CoordSize = "21600,21600"

	shapeType.Stroke = &vml.Stroke{}
	shapeType.Stroke.JoinStyle = vml.StrokeJoinStyleMiter

	shapeType.PathSettings = &vml.Path{}
	shapeType.PathSettings.GradientShapeOK = sharedML.TriStateTrue
	shapeType.PathSettings.ConnectType = vml.ConnectTypeRect

	d.ml.ShapeType = append(d.ml.ShapeType, shapeType)
	d.initializedComments = true
}

//only attach files, no content is loading
func (d *drawingsVML) attachFileIfRequired() {
	//attach sheet relations file
	d.sheet.attachRelationshipsIfRequired()

	if d.file == nil {
		fileName := d.sheet.relationships.GetTargetByType(internal.RelationTypeVmlDrawing)
		if fileName != "" {
			//transform relative path to absolute
			fileName = strings.Replace(fileName, "../", "xl/", 1)

			if file := d.sheet.workbook.doc.pkg.File(fileName); file != nil {
				d.file = ooxml.NewPackageFile(d.sheet.workbook.doc.pkg, file, &d.ml, nil)
				return
			}

			panic(fmt.Sprintf("can't load VML file: %s", fileName))
		}

		totalFiles := len(d.sheet.workbook.doc.pkg.Files(regexpDrawings))
		fileName = fmt.Sprintf("xl/drawings/vmlDrawing%d.vml", totalFiles+1)

		//register a VML content type, if required
		d.sheet.workbook.doc.pkg.ContentTypes().RegisterType("vml", ooxml.ContentTypeVmlDrawing)

		//attach file to package
		d.file = ooxml.NewPackageFile(d.sheet.workbook.doc.pkg, fileName, &d.ml, nil)

		//add file to sheet relations
		d.sheet.relationships.AddFile(internal.RelationTypeVmlDrawing, fileName)
	}
}

//build indexes for shapes
func (d *drawingsVML) buildIndexes() {
	for id, s := range d.ml.Shape {
		_ = d.shapeIndex.Add(s, id)
	}
}
