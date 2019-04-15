package enum

type OsType int

const (
	Linux OsType = iota
	MacOs
	Windows
)

func (os OsType) String() string {
	names := [...]string{
		"Linux",
		"macOS",
		"Windows"}

	return names[os]
}
