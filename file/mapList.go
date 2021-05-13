package file

//MapList is an interface that shows list of maps in somewhere.
type MapList interface {
	Current() []string                          //Current returns array of name of map in current page.
	Next()                                      //Next moves list to next page.
	Prev()                                      //Prev moves list to previous page.
	CurrentPage() int                           //CurrentPage returns current page number of list.
	LastPage() int                              //LastPage returns last page number of list.
	GetMapName(from int) (name string, ok bool) //GetMapName returns name of map by in current page by index and returns whether it exists.
	GetCachedMapName() string                   //GetCachedMapName returns name of map that previously selected.
	Refresh() error                             //Refresh refreshes map list. It should be called when files have been added/removed during running.
}
