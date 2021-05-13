package file

//System is an main interface of file that directly used in this project.
type System interface {
	Map(name string, formatter Formatter) (Storage, error)           //Map returns Storage of map file.
	Setting(formatter Formatter) (Storage, error)                    //Setting returns Storage of setting file.
	Language() (Updater, error)                                      //Language returns Updater of language directory.
	LanguageOf(language string, formatter Formatter) (Loader, error) //LanguageOf returns Loader of language file.
	Maps() MapList                                                   //Maps returns MapList of map directory.
}
