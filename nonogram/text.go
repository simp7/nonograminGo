package nonogram

type Text interface {
	IsLatest(string) bool
	MainMenu() []string
	GetSelectHeader() []string
	GetResult() []string
	Complete() string
	GetHelp() []string
	GetCredit() []string
	RequestMapName() string
	RequestWidth() string
	RequestHeight() string
	SizeError() string
	FileNotExist() string
	BlankBetweenMapNameAndTimer() string
}
