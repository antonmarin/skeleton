package osFamily

//OsFamily is enum, which represents Operating System type
type OsFamily int

//NewFromGOOS is a OsFamily static factory to create
// from values of GOOS environment variable
func NewFromGOOS(goos string) OsFamily {
	switch goos {
	case "linux":
		return Linux
	case "darwin":
		return MacOs
	case "windows":
		return Windows
	default:
		return *new(OsFamily)
	}
}

const (
	//Undefined operating system
	Undefined OsFamily = iota
	//Linux family operating system
	Linux
	//MacOs family operating system
	MacOs
	//Windows family operating system
	Windows
)

func (os OsFamily) String() string {
	names := [...]string{
		"Undefined",
		"Linux",
		"macOS",
		"Windows"}

	return names[os]
}
