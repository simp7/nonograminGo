package client

import (
	"github.com/simp7/nonogram/setting"
	"strconv"
	"strings"
)

const (
	delimiterSize = 40
)

//Text is an data of text that are not yet be processed but stored.
type Text struct {
	setting.Text
}

func AdaptText(target setting.Text) Text {
	return Text{target}
}

func (t Text) MainMenu() []string {
	list := listByNumber(t.Start, t.Create, t.Help, t.Credit, t.Exit)
	header := append(t.title(t.Title), "", t.SelectRequest, "")
	return append(header, list...)
}

func (t Text) SelectHeader() []string {
	return []string{"[ " + t.MapList + " ]", "[ <-" + t.Prev + " | " + t.Next + "-> ]    ", t.delimiter(), ""}
}

func (t Text) GetResult() []string {
	results := colonFormat(t.MapName, t.ClearTime, t.WrongCells)
	return append(t.title(t.Clear), results...)
}

func (t Text) Complete() string {
	return t.CompleteMsg
}

func (t Text) GetHelp() []string {
	return append(t.title(t.Help), t.keyInstruction()...)
}

func (t Text) GetCredit() []string {
	return append(t.title(t.Credit), t.DeveloperInfo, t.License, t.ThankYouMsg, t.delimiter())
}

func (t Text) RequestMapName() string {
	return t.ReqMapName
}

func (t Text) RequestWidth() string {
	return t.ReqWidth
}

func (t Text) RequestHeight() string {
	return t.ReqHeight
}

func (t Text) SizeError() string {
	return t.MapSizeError
}

func (t Text) FileNotExist() string {
	return t.MapFileNotExist
}

func (t Text) BlankBetweenMapNameAndTimer() string {
	return "               "
}

func (t Text) keyInstruction() []string {
	key := []string{t.ArrowKey, "Space/Z", "X", "Enter", "Esc"}
	instruction := []string{t.ExplArrowKey, t.ExplSpace, t.ExplX, t.ExplEnter, t.ExplEsc}
	return completeColonFormat(key, instruction)
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

func (t Text) delimiter() string {
	return strings.Repeat("-", delimiterSize)
}

func (t Text) title(text string) []string {
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
