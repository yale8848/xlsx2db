// Copyright (c) 2017 Andrey Gayvoronsky <plandem@gmail.com>
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package rule

import (
	"github.com/plandem/xlsx/internal/ml/primitives"
)

//List of all possible values for IconSetType
const (
	_ primitives.IconSetType = iota
	fromIconSetType3Icons
	IconSetType3Arrows
	IconSetType3ArrowsGray
	IconSetType3Flags
	IconSetType3TrafficLights1
	IconSetType3TrafficLights2
	IconSetType3Signs
	IconSetType3Symbols
	IconSetType3Symbols2
	fromIconSetType4Icons
	IconSetType4Arrows
	IconSetType4ArrowsGray
	IconSetType4RedToBlack
	IconSetType4Rating
	IconSetType4TrafficLights
	fromIconSetType5Icons
	IconSetType5Arrows
	IconSetType5ArrowsGray
	IconSetType5Rating
	IconSetType5Quarters
)

func init() {
	primitives.FromIconSetType = map[primitives.IconSetType]string{
		IconSetType3Arrows:         "3Arrows",
		IconSetType3ArrowsGray:     "3ArrowsGray",
		IconSetType3Flags:          "3Flags",
		IconSetType3TrafficLights1: "3TrafficLights1",
		IconSetType3TrafficLights2: "3TrafficLights2",
		IconSetType3Signs:          "3Signs",
		IconSetType3Symbols:        "3Symbols",
		IconSetType3Symbols2:       "3Symbols2",
		IconSetType4Arrows:         "4Arrows",
		IconSetType4ArrowsGray:     "4ArrowsGray",
		IconSetType4RedToBlack:     "4RedToBlack",
		IconSetType4Rating:         "4Rating",
		IconSetType4TrafficLights:  "4TrafficLights",
		IconSetType5Arrows:         "5Arrows",
		IconSetType5ArrowsGray:     "5ArrowsGray",
		IconSetType5Rating:         "5Rating",
		IconSetType5Quarters:       "5Quarters",
	}

	primitives.ToIconSetType = make(map[string]primitives.IconSetType, len(primitives.FromIconSetType))
	for k, v := range primitives.FromIconSetType {
		primitives.ToIconSetType[v] = k
	}
}
