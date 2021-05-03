package file

//System is an main interface of file that directly used in this project.
//Map returns Storage of map file.
//Setting returns Storage of setting file.
//Language returns Updater of language directory.
//LanguageOf returns Loader of language file.
//Maps returns MapList of map directory.
type System interface {
	Map(name string, formatter Formatter) (Storage, error)
	Setting(formatter Formatter) (Storage, error)
	Language() (Updater, error)
	LanguageOf(language string, formatter Formatter) (Loader, error)
	Maps() MapList
}
