// Autogenerated with dialgen, do not edit.
//
// Generated from revision https://github.com/mavlink/mavlink/tree/09bc1d5317d3557f574997378a8b1cba3b675413
//
package icarous

import (
	"github.com/aler9/gomavlib"
)

// Dialect contains the dialect object that can be passed to the library.
var Dialect = dialect

// dialect is not exposed directly such that it is not displayed in godoc.
var dialect = gomavlib.MustDialect(0, []gomavlib.Message{
	// icarous.xml
	&MessageIcarousHeartbeat{},
	&MessageIcarousKinematicBands{},
})

//
type ICAROUS_FMS_STATE int

const (
	//
	ICAROUS_FMS_STATE_IDLE ICAROUS_FMS_STATE = 0
	//
	ICAROUS_FMS_STATE_TAKEOFF ICAROUS_FMS_STATE = 1
	//
	ICAROUS_FMS_STATE_CLIMB ICAROUS_FMS_STATE = 2
	//
	ICAROUS_FMS_STATE_CRUISE ICAROUS_FMS_STATE = 3
	//
	ICAROUS_FMS_STATE_APPROACH ICAROUS_FMS_STATE = 4
	//
	ICAROUS_FMS_STATE_LAND ICAROUS_FMS_STATE = 5
)

//
type ICAROUS_TRACK_BAND_TYPES int

const (
	//
	ICAROUS_TRACK_BAND_TYPE_NONE ICAROUS_TRACK_BAND_TYPES = 0
	//
	ICAROUS_TRACK_BAND_TYPE_NEAR ICAROUS_TRACK_BAND_TYPES = 1
	//
	ICAROUS_TRACK_BAND_TYPE_RECOVERY ICAROUS_TRACK_BAND_TYPES = 2
)

// icarous.xml

// ICAROUS heartbeat
type MessageIcarousHeartbeat struct {
	// See the FMS_STATE enum.
	Status ICAROUS_FMS_STATE `mavenum:"uint8"`
}

func (*MessageIcarousHeartbeat) GetId() uint32 {
	return 42000
}

// Kinematic multi bands (track) output from Daidalus
type MessageIcarousKinematicBands struct {
	// Number of track bands
	Numbands int8 `mavname:"numBands"`
	// See the TRACK_BAND_TYPES enum.
	Type1 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min1 float32
	// max angle (degrees)
	Max1 float32
	// See the TRACK_BAND_TYPES enum.
	Type2 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min2 float32
	// max angle (degrees)
	Max2 float32
	// See the TRACK_BAND_TYPES enum.
	Type3 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min3 float32
	// max angle (degrees)
	Max3 float32
	// See the TRACK_BAND_TYPES enum.
	Type4 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min4 float32
	// max angle (degrees)
	Max4 float32
	// See the TRACK_BAND_TYPES enum.
	Type5 ICAROUS_TRACK_BAND_TYPES `mavenum:"uint8"`
	// min angle (degrees)
	Min5 float32
	// max angle (degrees)
	Max5 float32
}

func (*MessageIcarousKinematicBands) GetId() uint32 {
	return 42001
}
