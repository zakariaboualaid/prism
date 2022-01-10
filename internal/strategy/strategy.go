package strategy

//go:generate stringer -type=Strategy
type Strategy int

const (
	RoundRobin Strategy = iota
	WeightedRoundRobin
	LeastConnection
)
