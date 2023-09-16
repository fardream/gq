package gq

//go:generate stringer -type=Endianness

// Endianness indicates the endianness of the message.
type Endianness uint8

const (
	BigEndian    Endianness = 0
	LittleEndian Endianness = 1
)

// MessageType - see the first example on [kx.com] (only examples are provided, no specification/documentation).
//
// [kx.com]: https://code.kx.com/q/kb/serialization/
type MessageType uint8

const (
	MessageType_Async    MessageType = 0
	MessageType_Sync     MessageType = 1
	MessageType_Response MessageType = 2
)

// IPCMessage is the message for IPC protocol of q, see [examples on kx.com]
//
// It looks to be in the following format
//
//  1. first byte, endianness
//  2. second byte, message type
//  3. 2 byte of zeros.
//  4. 4 byte, integer, length of the following payload.
//
// [examples on kx.com]: https://code.kx.com/q/kb/serialization/
type IPCMessage struct{}
