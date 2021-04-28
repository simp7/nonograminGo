package file

type System interface {
	Map(name string, formatter Formatter) (Storage, error)
	Setting(formatter Formatter) (Storage, error)
	Language() (Updater, error)
	LanguageOf(language string, formatter Formatter) (Loader, error)
	MapList() MapList
}
