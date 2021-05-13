package client

//Text is an interface that returns text.
type Text interface {
	IsLatest(string) bool                //IsLatest compares argument with text file version.
	MainMenu() []string                  //MainMenu returns main menu of the application.
	SelectHeader() []string              //SelectHeader returns select header of map list.
	GetResult() []string                 //GetResult returns result page after game.
	Complete() string                    //Complete returns complete message.
	GetHelp() []string                   //GetHelp returns manual text.
	GetCredit() []string                 //GetCredit returns credit.
	RequestMapName() string              //RequestMapName returns request for map name input.
	RequestWidth() string                //RequestWidth returns request for map height input.
	RequestHeight() string               //RequestHeight returns request for map height.
	SizeError() string                   //SizeError returns error message when input number is wrong.
	FileNotExist() string                //FileNotExist returns error message when file doesn't exist.
	BlankBetweenMapNameAndTimer() string //BlankBetweenMapNameAndTimer returns blank between map name and timer.
}
