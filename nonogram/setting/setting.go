package setting

import (
	"github.com/simp7/nonograminGo/nonogram"
	"github.com/simp7/nonograminGo/nonogram/file/loader"
	"github.com/simp7/nonograminGo/nonogram/text"
	"sync"
)

var instance *nonogram.Setting
var once sync.Once

func Get() *nonogram.Setting {

	once.Do(func() {
		loader.Setting().Load(&instance)
	})
	instance.Text = text.New(instance.Language)

	return instance

}
