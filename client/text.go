package client

//Text is an interface that returns text.
//IsLatest compares argument with text file version.
//MainMenu returns main menu of the application.
//SelectHeader returns select header of map list.
//GetResult returns result page after game.
//Complete returns complete message.
//GetHelp returns manual text.
//GetCredit returns credit.
//RequestMapName returns request for map name input.
//RequestWidth returns request for map height input.
//RequestHeight returns request for map height.
//SizeError returns error message when input number is wrong.
//FileNotExist returns error message when file doesn't exist.
//BlankBetweenMapNameAndTimer returns blank between map name and timer.
type Text interface {
	IsLatest(string) bool
	MainMenu() []string
	SelectHeader() []string
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
