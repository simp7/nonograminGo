package file

//MapList is an interface that shows list of maps in somewhere.
//Current returns array of name of map in current page.
//Next moves list to next page.
//Prev moves list to previous page.
//CurrentPage returns current page number of list.
//LastPage returns last page number of list.
//GetMapName returns name of map by in current page by index and returns whether it exists.
//GetCachedMapName returns name of map that previously selected.
//Refresh refreshes map list. It should be called when files have been added/removed during running.
type MapList interface {
	Current() []string
	Next()
	Prev()
	CurrentPage() int
	LastPage() int
	GetMapName(from int) (name string, ok bool)
	GetCachedMapName() string
	Refresh() error
}
