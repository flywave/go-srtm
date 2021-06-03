package srtm

type SrtmModel uint32

const (
	/*
	 * model 1: U.S.,
	 * sampled at one arc-second lat/long intervals, 3601 lines, 3601 samples
	 */
	US SrtmModel = 3601
	/*
	* model 3: world,
	* sampled at three arc-second lat/long intervals, 1201 lines, 1201 samples
	 */
	WORLD SrtmModel = 1201
)
