// Code generated by "stringer -type=Strategy"; DO NOT EDIT.

package strategy

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RoundRobin-0]
	_ = x[WeightedRoundRobin-1]
	_ = x[LeastConnection-2]
}

const _Strategy_name = "RoundRobinWeightedRoundRobinLeastConnection"

var _Strategy_index = [...]uint8{0, 10, 28, 43}

func (i Strategy) String() string {
	if i < 0 || i >= Strategy(len(_Strategy_index)-1) {
		return "Strategy(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Strategy_name[_Strategy_index[i]:_Strategy_index[i+1]]
}
