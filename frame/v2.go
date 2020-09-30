package frame

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/aler9/gomavlib/msg"
)

func uint24Decode(in []byte) uint32 {
	return uint32(in[2])<<16 | uint32(in[1])<<8 | uint32(in[0])
}

func uint24Encode(buf []byte, in uint32) []byte {
	buf[0] = byte(in)
	buf[1] = byte(in >> 8)
	buf[2] = byte(in >> 16)
	return buf[:3]
}

func uint48Decode(in []byte) uint64 {
	return uint64(in[5])<<40 | uint64(in[4])<<32 | uint64(in[3])<<24 |
		uint64(in[2])<<16 | uint64(in[1])<<8 | uint64(in[0])
}

func uint48Encode(buf []byte, in uint64) []byte {
	buf[0] = byte(in)
	buf[1] = byte(in >> 8)
	buf[2] = byte(in >> 16)
	buf[3] = byte(in >> 24)
	buf[4] = byte(in >> 32)
	buf[5] = byte(in >> 40)
	return buf[:6]
}

const (
	V2MagicByte  = 0xFD
	V2FlagSigned = 0x01
)

// V2Signature is a V2 frame signature.
type V2Signature [6]byte

// V2Frame is a V2 frame.
type V2Frame struct {
	IncompatibilityFlag byte
	CompatibilityFlag   byte
	SequenceId          byte
	SystemId            byte
	ComponentId         byte
	Message             msg.Message
	Checksum            uint16
	SignatureLinkId     byte
	SignatureTimestamp  uint64
	Signature           *V2Signature
}

// Clone implements the Frame interface.
func (f *V2Frame) Clone() Frame {
	return &V2Frame{
		IncompatibilityFlag: f.IncompatibilityFlag,
		CompatibilityFlag:   f.CompatibilityFlag,
		SequenceId:          f.SequenceId,
		SystemId:            f.SystemId,
		ComponentId:         f.ComponentId,
		Message:             f.Message,
		Checksum:            f.Checksum,
		SignatureLinkId:     f.SignatureLinkId,
		SignatureTimestamp:  f.SignatureTimestamp,
		Signature:           f.Signature,
	}
}

// GetSystemId implements the Frame interface.
func (f *V2Frame) GetSystemId() byte {
	return f.SystemId
}

// GetComponentId implements the Frame interface.
func (f *V2Frame) GetComponentId() byte {
	return f.ComponentId
}

// GetMessage implements the Frame interface.
func (f *V2Frame) GetMessage() msg.Message {
	return f.Message
}

// GetChecksum implements the Frame interface.
func (f *V2Frame) GetChecksum() uint16 {
	return f.Checksum
}

// IsSigned checks whether the frame contains a signature. It does not validate the signature.
func (f *V2Frame) IsSigned() bool {
	return (f.IncompatibilityFlag & V2FlagSigned) != 0
}

// Decode implements the Frame interface.
func (f *V2Frame) Decode(br *bufio.Reader) error {
	// header
	buf, err := br.Peek(9)
	if err != nil {
		return err
	}
	br.Discard(9)
	msgLen := buf[0]
	f.IncompatibilityFlag = buf[1]
	f.CompatibilityFlag = buf[2]
	f.SequenceId = buf[3]
	f.SystemId = buf[4]
	f.ComponentId = buf[5]
	msgId := uint24Decode(buf[6:])

	// discard frame if incompatibility flag is not understood, as in recommendations
	if f.IncompatibilityFlag != 0 && f.IncompatibilityFlag != V2FlagSigned {
		return fmt.Errorf("unknown incompatibility flag (%d)", f.IncompatibilityFlag)
	}

	// message
	var msgEncoded []byte
	if msgLen > 0 {
		msgEncoded = make([]byte, msgLen)
		_, err = io.ReadFull(br, msgEncoded)
		if err != nil {
			return err
		}
	}
	f.Message = &msg.MessageRaw{
		Id:      msgId,
		Content: msgEncoded,
	}

	// checksum
	buf, err = br.Peek(2)
	if err != nil {
		return err
	}
	br.Discard(2)
	f.Checksum = binary.LittleEndian.Uint16(buf)

	// signature
	if f.IsSigned() {
		buf, err := br.Peek(13)
		if err != nil {
			return err
		}
		br.Discard(13)
		f.SignatureLinkId = buf[0]
		f.SignatureTimestamp = uint48Decode(buf[1:])
		f.Signature = new(V2Signature)
		copy(f.Signature[:], buf[7:])
	}

	return nil
}

// Encode implements the Frame interface.
func (f *V2Frame) Encode(buf []byte, msgEncoded []byte) ([]byte, error) {
	msgLen := len(msgEncoded)
	bufLen := 10 + msgLen + 2
	if f.IsSigned() {
		bufLen += 13
	}
	buf = buf[:bufLen]

	// header
	buf[0] = V2MagicByte
	buf[1] = byte(msgLen)
	buf[2] = f.IncompatibilityFlag
	buf[3] = f.CompatibilityFlag
	buf[4] = f.SequenceId
	buf[5] = f.SystemId
	buf[6] = f.ComponentId
	uint24Encode(buf[7:], f.Message.GetId())

	// message
	if msgLen > 0 {
		copy(buf[10:], msgEncoded)
	}

	// checksum
	binary.LittleEndian.PutUint16(buf[10+msgLen:], f.Checksum)

	// signature
	if f.IsSigned() {
		buf[12+msgLen] = f.SignatureLinkId
		uint48Encode(buf[13+msgLen:], f.SignatureTimestamp)
		copy(buf[19+msgLen:], f.Signature[:])
	}

	return buf, nil
}