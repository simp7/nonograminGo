package cli

import (
	"strconv"
	"strings"
)

const (
	delimiterSize = 40
)

//TextData is an data of text that are not yet be processed but stored.
type TextData struct {
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

func (t *TextData) MainMenu() []string {
	list := listByNumber(t.Start, t.Create, t.Help, t.Credit, t.Exit)
	header := append(t.title(t.Title), "", t.SelectRequest, "")
	return append(header, list...)
}

func (t *TextData) SelectHeader() []string {
	return []string{"[ " + t.MapList + " ]", "[ <-" + t.Prev + " | " + t.Next + "-> ]    ", t.delimiter(), ""}
}

func (t *TextData) GetResult() []string {
	results := colonFormat(t.MapName, t.ClearTime, t.WrongCells)
	return append(t.title(t.Clear), results...)
}

func (t *TextData) Complete() string {
	return t.CompleteMsg
}

func (t *TextData) GetHelp() []string {
	return append(t.title(t.Help), t.keyInstruction()...)
}

func (t *TextData) GetCredit() []string {
	return append(t.title(t.Credit), t.DeveloperInfo, t.License, t.ThankYouMsg, t.delimiter())
}

func (t *TextData) RequestMapName() string {
	return t.ReqMapName
}

func (t *TextData) RequestWidth() string {
	return t.ReqWidth
}

func (t *TextData) RequestHeight() string {
	return t.ReqHeight
}

func (t *TextData) SizeError() string {
	return t.MapSizeError
}

func (t *TextData) FileNotExist() string {
	return t.MapFileNotExist
}

func (t *TextData) BlankBetweenMapNameAndTimer() string {
	return "               "
}

func (t *TextData) keyInstruction() []string {
	key := []string{t.ArrowKey, "Space/Z", "X", "Enter", "Esc"}
	instruction := []string{t.ExplArrowKey, t.ExplSpace, t.ExplX, t.ExplEnter, t.ExplEsc}
	return completeColonFormat(key, instruction)
}

func (t *TextData) IsLatest(s string) bool {
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

func (t *TextData) delimiter() string {
	return strings.Repeat("-", delimiterSize)
}

func (t *TextData) title(text string) []string {
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
