package component

type Status int

const (
	NotInstalled Status = iota
	Installed
	Outdated
	Corrupted
)