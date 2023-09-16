package gq

//go:generate stringer -type=Capacity

// Capacity for the connection, see [kx.com]. Use default 3.
//
// [kx.com]: https://code.kx.com/q/basics/ipc
type Capacity uint8

const (
	Default_Capacity Capacity = 3                // DefaultCapacity is the default of [Capacity], 3.
	Capacity_V0      Capacity = 0                // (V2.5) no compression, no timestamp, no timespan, no UUID
	Capacity_V1      Capacity = 1                // (V2.6-2.8) compression, timestamp, timespan
	Capacity_V3      Capacity = Default_Capacity // (V3.0) compression, timestamp, timespan, UUID
	Capacity_V4      Capacity = 4                // reserved
	Capacity_V5      Capacity = 5                // support msgs >2GB; vectors must each have a count ≤ 2 billion
	Capacity_V6      Capacity = 6                // support msgs >2GB and vectors may each have a count > 2 billion
)
