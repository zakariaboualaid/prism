package strategy

import "fmt"

//go:generate stringer -type=Strategy
type Strategy int

const (
	RoundRobin Strategy = iota
	WeightedRoundRobin
	LeastConnection
)

func (s *Strategy) UnmarshalText(bs []byte) error {
	switch string(bs) {
	case RoundRobin.String():
		*s = RoundRobin
	case WeightedRoundRobin.String():
		*s = WeightedRoundRobin
	default:
		return fmt.Errorf("Unknown Strategy %q ", s.String())
	}
	return nil
}
