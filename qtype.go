package gq

//go:generate stringer -type=QTypeID

// QTypeID is signed 8 bit integer (signed char), which is a unique number defined in q that represents a type.
// The positive numbers are generally list, and negative numbers are atomic values of corresponding lists.
// See the [doc on types] on kx.com.
//
// The constants such as [QTypeID_KB] are taken from the C/C++ header [k.h]
//
// [doc on types]: https://code.kx.com/q/basics/datatypes/
// [k.h]: https://github.com/KxSystems/kdb/blob/master/c/c/k.h
type QTypeID int8

const (
	QTypeID_KB    QTypeID = 1    // 1 boolean   char   kG
	QTypeID_UU    QTypeID = 2    // 16 guid     U      kU
	QTypeID_KG    QTypeID = 4    // 1 byte      char   kG
	QTypeID_KH    QTypeID = 5    // 2 short     short  kH
	QTypeID_KI    QTypeID = 6    // 4 int       int    kI
	QTypeID_KJ    QTypeID = 7    // 8 long      long   kJ
	QTypeID_KE    QTypeID = 8    // 4 real      float  kE
	QTypeID_KF    QTypeID = 9    // 8 float     double kF
	QTypeID_KC    QTypeID = 10   // 1 char      char   kC
	QTypeID_KS    QTypeID = 11   // * symbol    char*  kS
	QTypeID_KP    QTypeID = 12   // 8 timestamp long   kJ (nanoseconds from 2000.01.01)
	QTypeID_KM    QTypeID = 13   // 4 month     int    kI (months from 2000.01.01)
	QTypeID_KD    QTypeID = 14   // 4 date      int    kI (days from 2000.01.01)
	QTypeID_KN    QTypeID = 16   // 8 timespan  long   kJ (nanoseconds)
	QTypeID_KU    QTypeID = 17   // 4 minute    int    kI
	QTypeID_KV    QTypeID = 18   // 4 second    int    kI
	QTypeID_KT    QTypeID = 19   // 4 time      int    kI (millisecond)
	QTypeID_KZ    QTypeID = 15   // 8 datetime  double kF (DO NOT USE)
	QTypeID_XT    QTypeID = 98   // table,  x->k is XD
	QTypeID_XD    QTypeID = 99   // kK(x)[0] is keys. kK(x)[1] is values.
	QTypeID_ERROR QTypeID = -128 // error
)

// GUID is byte array of length 16 - no added functionality here,
// please consider a package like [github.com/google/uuid] (this can be directly cast into UUID provied in that package).
//
// [github.com/google/uuid]: https://pkg.go.dev/github.com/google/uuid
type GUID = [16]byte

// IsAtomic returns if the type is atomic type.
func (qtypeid QTypeID) IsAtomic() bool {
	return qtypeid < 0 && qtypeid != QTypeID_ERROR
}
