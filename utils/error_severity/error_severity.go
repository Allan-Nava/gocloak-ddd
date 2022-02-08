package errorseverity
/*
* Copyright © 2022 Allan Nava <>
* Created 08/02/2022
* Updated 08/02/2022
*
 */
type ErrorSeverity int

const (
	Low ErrorSeverity = iota
	Medium
	High
	Critical
)

func (e ErrorSeverity) String() string {
	v := [...]string{"low", "medium", "high", "critical"}[e]
	// prevent panicking in case of status is out-of-range
	if e < Low || e > Critical {
		return "Unknown"
	}
	return v
}
