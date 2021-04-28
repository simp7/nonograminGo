package config

import (
	"github.com/simp7/nonograminGo/file/formatter"
	"github.com/simp7/nonograminGo/file/localStorage"
	"strconv"
	"strings"
)

const (
	delimiterSize = 40
)

type textData struct {
	FileVersion     string
	Title           string
	SelectRequest   string
	Start           string
	Create          string
	Help            string
	Credit          string
	Exit            string
	MapList         string
	Prev            string
	Next            string
	Result          string
	Clear           string
	MapName         string
	ClearTime       string
	WrongCells      string
	CompleteMsg     string
	ExplArrowKey    string
	ExplSpace       string
	ExplX           string
	ExplEnter       string
	ExplEsc         string
	DeveloperInfo   string
	License         string
	ThankYouMsg     string
	ReqMapName      string
	ReqWidth        string
	ReqHeight       string
	MapSizeError    string
	MapFileNotExist string
	ArrowKey        string
}

func New(language string) (*textData, error) {

	loaded := new(textData)

	fs, err := localStorage.Get()
	if err != nil {
		return nil, err
	}

	languageLoader, err := fs.LanguageOf(language, formatter.Json())
	if err != nil {
		return nil, err
	}

	err = languageLoader.Load(&loaded)
	return loaded, err

}

func (t *textData) MainMenu() []string {
	list := listByNumber(t.Start, t.Create, t.Help, t.Credit, t.Exit)
	header := append(t.title(t.Title), "", t.SelectRequest, "")
	return append(header, list...)
}

func (t *textData) GetSelectHeader() []string {
	return []string{"[ " + t.MapList + " ]", "[ <-" + t.Prev + " | " + t.Next + "-> ]    ", t.delimiter(), ""}
}

func (t *textData) GetResult() []string {
	results := colonFormat(t.MapName, t.ClearTime, t.WrongCells)
	return append(t.title(t.Clear), results...)
}

func (t *textData) Complete() string {
	return t.CompleteMsg
}

func (t *textData) GetHelp() []string {
	return append(t.title(t.Help), t.keyInstruction()...)
}

func (t *textData) GetCredit() []string {
	return append(t.title(t.Credit), t.DeveloperInfo, t.License, t.ThankYouMsg, t.delimiter())
}

func (t *textData) RequestMapName() string {
	return t.ReqMapName
}

func (t *textData) RequestWidth() string {
	return t.ReqWidth
}

func (t *textData) RequestHeight() string {
	return t.ReqHeight
}

func (t *textData) SizeError() string {
	return t.MapSizeError
}

func (t *textData) FileNotExist() string {
	return t.MapFileNotExist
}

func (t *textData) BlankBetweenMapNameAndTimer() string {
	return "               "
}

func (t *textData) keyInstruction() []string {
	key := []string{t.ArrowKey, "Space/Z", "X", "Enter", "Esc"}
	instruction := []string{t.ExplArrowKey, t.ExplSpace, t.ExplX, t.ExplEnter, t.ExplEsc}
	return completeColonFormat(key, instruction)
}

func (t *textData) IsLatest(s string) bool {
	return s == t.FileVersion
}

func maxLength(texts []string) int {
	max := 0
	for _, v := range texts {
		if max < len(v) {
			max = len(v)
		}
	}
	return max
}

func unifyLength(text string, to int) string {
	text += strings.Repeat(" ", to-len(text))
	return text
}

func addColon(text string) string {
	return text + " : "
}

func colonFormat(texts ...string) []string {
	max := maxLength(texts)
	for i := range texts {
		texts[i] = unifyLength(texts[i], max)
		texts[i] = addColon(texts[i])
	}
	return texts
}

func completeColonFormat(left []string, right []string) []string {
	result := colonFormat(left...)
	for i := range result {
		result[i] += right[i]
	}
	return result
}

func listByNumber(texts ...string) []string {
	for i, v := range texts {
		texts[i] = strconv.Itoa(i+1) + ". " + v
	}
	return texts
}

func (t *textData) delimiter() string {
	return strings.Repeat("-", delimiterSize)
}

func (t *textData) title(text string) []string {
	blank := (delimiterSize - len(text)) / 2
	if blank < 0 {
		return []string{text}
	}
	result := make([]string, 3)
	result[0] = t.delimiter()
	result[1] = strings.Repeat(" ", blank) + text
	result[2] = t.delimiter()
	return result
}
