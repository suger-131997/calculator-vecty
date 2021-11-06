package model

type FilterType int

const (
	All FilterType = iota
	Completed
	Active
)
