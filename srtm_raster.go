package srtm

import (
	"github.com/flywave/go-raster"
)

type GeoBounds [4]float64

func newRasterConfig() *raster.RasterConfig {
	conf := raster.NewDefaultRasterConfig()
	conf.RasterFormat = raster.RT_GeoTiff
	conf.DataType = raster.DT_FLOAT32
	conf.EPSGCode = 4326
	return conf
}

func cacleBounds(latitude, longitude float64) GeoBounds {
	northSouth := 'S'
	if latitude >= 0 {
		northSouth = 'N'
	}

	eastWest := 'W'
	if longitude >= 0 {
		eastWest = 'E'
	}
	var north float64
	if northSouth == 'S' {
		north = latitude - 1
	} else {
		north = latitude + 1
	}

	var west float64
	if eastWest == 'W' {
		west = longitude + 1
	} else {
		west = longitude - 1
	}

	return [4]float64{north, latitude, longitude, west}
}

func WriteSrtmToRaster(f *SrtmFile, fileName string) error {
	bounds := cacleBounds(f.latitude, f.longitude)
	conf := newRasterConfig()
	raster, err := raster.CreateNewRaster(fileName, f.squareSize, f.squareSize, bounds[0], bounds[1], bounds[2], bounds[3], conf)
	if err != nil {
		return err
	}
	for row := 0; row < f.squareSize; row++ {
		for col := 0; col < f.squareSize; col++ {
			v := f.getElevationFromRowAndColumn(row, col)
			raster.SetValue(row, col, v)
		}
	}

	return raster.Save()
}
