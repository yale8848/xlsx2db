// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package styles

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for VAlignType
const (
	_ primitives.VAlignType = iota
	VAlignTop
	VAlignCenter
	VAlignBottom
	VAlignJustify
	VAlignDistributed
)

func init() {
	primitives.FromVAlignType = map[primitives.VAlignType]string{
		VAlignTop:         "top",
		VAlignCenter:      "center",
		VAlignBottom:      "bottom",
		VAlignJustify:     "justify",
		VAlignDistributed: "distributed",
	}

	primitives.ToVAlignType = make(map[string]primitives.VAlignType, len(primitives.FromVAlignType))
	for k, v := range primitives.FromVAlignType {
		primitives.ToVAlignType[v] = k
	}
}
