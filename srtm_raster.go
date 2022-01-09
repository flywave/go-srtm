package srtm

import (
	"encoding/binary"
	"image"
	"os"

	vec2d "github.com/flywave/go3d/float64/vec2"

	"github.com/flywave/go-cog"
	"github.com/flywave/go-geo"
)

var epsg4326 geo.Proj

func init() {
	epsg4326 = geo.NewProj(4326)
}

type GeoBounds [4]float64

func cacleBounds(latitude, longitude float64) vec2d.Rect {
	northSouth := 'S'
	if latitude >= 0 {
		northSouth = 'N'
	}

	var rect vec2d.Rect

	eastWest := 'W'
	if longitude >= 0 {
		eastWest = 'E'
	}
	if northSouth == 'S' {
		rect.Min[1] = latitude - 1
		rect.Max[1] = latitude
	} else {
		rect.Max[1] = latitude + 1
		rect.Min[1] = latitude
	}

	if eastWest == 'W' {
		rect.Min[0] = longitude
		rect.Max[0] = longitude + 1
	} else {
		rect.Min[0] = longitude - 1
		rect.Max[0] = longitude
	}

	return rect
}

func WriteSrtmToRaster(f *SrtmFile, fileName string) error {
	bbox := cacleBounds(f.latitude, f.longitude)

	data := make([]float64, f.squareSize*f.squareSize)

	for row := 0; row < f.squareSize; row++ {
		for col := 0; col < f.squareSize; col++ {
			v := f.getElevationFromRowAndColumn(row, col)
			data[row*f.squareSize+col] = v
		}
	}

	rect := image.Rect(0, 0, int(f.squareSize), int(f.squareSize))
	src := cog.NewSource(data, &rect, cog.CTLZW)

	w := cog.NewTileWriter(src, binary.LittleEndian, false, bbox, epsg4326, [2]uint32{uint32(f.squareSize), uint32(f.squareSize)}, nil)

	fw, err := os.Create(fileName)

	if err != nil {
		return err
	}

	return w.WriteData(fw)
}
