package component

type Component interface {

	Name() string

	Check() Result

	Install() Result

	Update() Result

	Repair() Result

	Uninstall() Result
}