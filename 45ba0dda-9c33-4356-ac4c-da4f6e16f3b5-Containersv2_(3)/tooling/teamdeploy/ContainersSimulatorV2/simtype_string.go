// Code generated by "stringer -type=simtype"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Tripviewer-0]
	_ = x[Userprofile-1]
	_ = x[Trips-2]
	_ = x[TripPoints-3]
	_ = x[PointsOfInterest-4]
	_ = x[UserJava-5]
}

const _simtype_name = "TripviewerUserprofileTripsTripPointsPointsOfInterestUserJava"

var _simtype_index = [...]uint8{0, 10, 21, 26, 36, 52, 60}

func (i simtype) String() string {
	if i < 0 || i >= simtype(len(_simtype_index)-1) {
		return "simtype(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _simtype_name[_simtype_index[i]:_simtype_index[i+1]]
}
