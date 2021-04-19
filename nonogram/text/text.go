package text

import (
	"github.com/simp7/nonograminGo/file/loader"
	"strconv"
	"strings"
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
	err := loader.Language(language).Load(&loaded)
	if err != nil {
		return nil, err
	}
	return loaded, nil
}

func (t *textData) MainMenu() []string {
	list := listByNumber(t.Start, t.Create, t.Help, t.Credit, t.Exit)
	header := append(t.title(15, t.Title), "", t.SelectRequest, "")
	return append(header, list...)
}

func (t *textData) GetSelectHeader() []string {
	return []string{"[ " + t.MapList + " ]", "[ <-" + t.Prev + " | " + t.Next + "-> ]    ", t.delimiter(30), ""}
}

func (t *textData) GetResult() []string {
	results := colonFormat(t.MapName, t.ClearTime, t.WrongCells)
	return append(t.title(15, t.Clear), results...)
}

func (t *textData) Complete() string {
	return t.CompleteMsg
}

func (t *textData) GetHelp() []string {
	return append(t.title(15, t.Help), t.keyInstruction()...)
}

func (t *textData) GetCredit() []string {
	return append(t.title(15, t.Credit), t.DeveloperInfo, t.License, t.ThankYouMsg, t.delimiter(len(t.Credit)+30))
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

func (t *textData) delimiter(amount int) string {
	return strings.Repeat("-", amount)
}

func (t *textData) title(blank int, text string) []string {
	result := make([]string, 3)
	length := len(text) + 2*blank
	result[0] = t.delimiter(length)
	result[1] = strings.Repeat(" ", blank) + text
	result[2] = t.delimiter(length)
	return result
}
