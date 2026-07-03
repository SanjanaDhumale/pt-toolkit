package stack

type Component struct {
	Name        string
	Type        string
	Version     string

	Image       string
	Container   string

	Required    bool
}