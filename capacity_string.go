// Code generated by "stringer -type=Capacity"; DO NOT EDIT.

package gq

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Default_Capacity-3]
	_ = x[Capacity_V0-0]
	_ = x[Capacity_V1-1]
	_ = x[Capacity_V3-3]
	_ = x[Capacity_V4-4]
	_ = x[Capacity_V5-5]
	_ = x[Capacity_V6-6]
}

const (
	_Capacity_name_0 = "Capacity_V0Capacity_V1"
	_Capacity_name_1 = "Default_CapacityCapacity_V4Capacity_V5Capacity_V6"
)

var (
	_Capacity_index_0 = [...]uint8{0, 11, 22}
	_Capacity_index_1 = [...]uint8{0, 16, 27, 38, 49}
)

func (i Capacity) String() string {
	switch {
	case i <= 1:
		return _Capacity_name_0[_Capacity_index_0[i]:_Capacity_index_0[i+1]]
	case 3 <= i && i <= 6:
		i -= 3
		return _Capacity_name_1[_Capacity_index_1[i]:_Capacity_index_1[i+1]]
	default:
		return "Capacity(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
