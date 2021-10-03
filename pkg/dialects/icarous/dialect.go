//autogenerated:yes
//nolint:golint,misspell,govet,lll
// Package icarous contains the icarous dialect.
package icarous

import (
	"errors"
	"strconv"

	"github.com/aler9/gomavlib/pkg/dialect"
	"github.com/aler9/gomavlib/pkg/msg"
)

// Dialect contains the dialect object that can be passed to the library.
var Dialect = dial

// dialect is not exposed directly such that it is not displayed in godoc.
var dial = &dialect.Dialect{0, []msg.Message{
	// icarous.xml
	&MessageIcarousHeartbeat{},
	&MessageIcarousKinematicBands{},
}}

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

// MarshalText implements the encoding.TextMarshaler interface.
func (e ICAROUS_FMS_STATE) MarshalText() ([]byte, error) {
	switch e { //nolint:gocritic
	case ICAROUS_FMS_STATE_IDLE:
		return []byte("ICAROUS_FMS_STATE_IDLE"), nil
	case ICAROUS_FMS_STATE_TAKEOFF:
		return []byte("ICAROUS_FMS_STATE_TAKEOFF"), nil
	case ICAROUS_FMS_STATE_CLIMB:
		return []byte("ICAROUS_FMS_STATE_CLIMB"), nil
	case ICAROUS_FMS_STATE_CRUISE:
		return []byte("ICAROUS_FMS_STATE_CRUISE"), nil
	case ICAROUS_FMS_STATE_APPROACH:
		return []byte("ICAROUS_FMS_STATE_APPROACH"), nil
	case ICAROUS_FMS_STATE_LAND:
		return []byte("ICAROUS_FMS_STATE_LAND"), nil
	}
	return nil, errors.New("invalid value")
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ICAROUS_FMS_STATE) UnmarshalText(text []byte) error {
	switch string(text) { //nolint:gocritic
	case "ICAROUS_FMS_STATE_IDLE":
		*e = ICAROUS_FMS_STATE_IDLE
		return nil
	case "ICAROUS_FMS_STATE_TAKEOFF":
		*e = ICAROUS_FMS_STATE_TAKEOFF
		return nil
	case "ICAROUS_FMS_STATE_CLIMB":
		*e = ICAROUS_FMS_STATE_CLIMB
		return nil
	case "ICAROUS_FMS_STATE_CRUISE":
		*e = ICAROUS_FMS_STATE_CRUISE
		return nil
	case "ICAROUS_FMS_STATE_APPROACH":
		*e = ICAROUS_FMS_STATE_APPROACH
		return nil
	case "ICAROUS_FMS_STATE_LAND":
		*e = ICAROUS_FMS_STATE_LAND
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ICAROUS_FMS_STATE) String() string {
	byts, err := e.MarshalText()
	if err == nil {
		return string(byts)
	}
	return strconv.FormatInt(int64(e), 10)
}

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

// MarshalText implements the encoding.TextMarshaler interface.
func (e ICAROUS_TRACK_BAND_TYPES) MarshalText() ([]byte, error) {
	switch e { //nolint:gocritic
	case ICAROUS_TRACK_BAND_TYPE_NONE:
		return []byte("ICAROUS_TRACK_BAND_TYPE_NONE"), nil
	case ICAROUS_TRACK_BAND_TYPE_NEAR:
		return []byte("ICAROUS_TRACK_BAND_TYPE_NEAR"), nil
	case ICAROUS_TRACK_BAND_TYPE_RECOVERY:
		return []byte("ICAROUS_TRACK_BAND_TYPE_RECOVERY"), nil
	}
	return nil, errors.New("invalid value")
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ICAROUS_TRACK_BAND_TYPES) UnmarshalText(text []byte) error {
	switch string(text) { //nolint:gocritic
	case "ICAROUS_TRACK_BAND_TYPE_NONE":
		*e = ICAROUS_TRACK_BAND_TYPE_NONE
		return nil
	case "ICAROUS_TRACK_BAND_TYPE_NEAR":
		*e = ICAROUS_TRACK_BAND_TYPE_NEAR
		return nil
	case "ICAROUS_TRACK_BAND_TYPE_RECOVERY":
		*e = ICAROUS_TRACK_BAND_TYPE_RECOVERY
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ICAROUS_TRACK_BAND_TYPES) String() string {
	byts, err := e.MarshalText()
	if err == nil {
		return string(byts)
	}
	return strconv.FormatInt(int64(e), 10)
}

// icarous.xml

// ICAROUS heartbeat
type MessageIcarousHeartbeat struct {
	// See the FMS_STATE enum.
	Status ICAROUS_FMS_STATE `mavenum:"uint8"`
}

// GetID implements the msg.Message interface.
func (*MessageIcarousHeartbeat) GetID() uint32 {
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

// GetID implements the msg.Message interface.
func (*MessageIcarousKinematicBands) GetID() uint32 {
	return 42001
}
