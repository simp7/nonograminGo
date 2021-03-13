package asset

import "github.com/simp7/nonograminGo/util"

type Text interface {
	MainMenu() []string
	GetSelectHeader() []string
	GetResult() []string
	Complete() string
	GetHelp() []string
	GetCredit() []string
	RequestMapName() string
	RequestWidth() string
	RequestHeight() string
	SizeError() string
	FileNotExist() string
	BlankBetweenMapNameAndTimer() string
}

type textData struct {
	Title                   string
	TitleDelimiter          string
	SelectRequest           string
	Start                   string
	Create                  string
	Help                    string
	Credit                  string
	Exit                    string
	MapList                 string
	Prev                    string
	Next                    string
	MapListDelimiter        string
	Result                  string
	ResultDelimiter         string
	Clear                   string
	MapName                 string
	ClearTime               string
	WrongCells              string
	CompleteMsg             string
	Manual                  string
	ManualDelimiter         string
	ExplArrowKey            string
	ExplSpace               string
	ExplX                   string
	ExplEnter               string
	ExplEsc                 string
	CreditDelimiter         string
	DeveloperInfo           string
	License                 string
	ThankYouMsg             string
	ReqMapName              string
	ReqWidth                string
	ReqHeight               string
	MapSizeError            string
	MapFileNotExist         string
	BlankBetweenMapAndTimer string
}

func NewText(data []byte) Text {
	t := new(textData)
	f := util.NewFileFormatter()
	f.GetRaw(data)
	util.CheckErr(f.Decode(&t))
	return t
}

func (t *textData) MainMenu() []string {
	return []string{t.TitleDelimiter, " " + t.Title, t.TitleDelimiter, "", t.SelectRequest, "", "1. " + t.Start, "2. " + t.Create, "3. " + t.Help, "4. " + t.Credit, "5. " + t.Exit}
}

func (t *textData) GetSelectHeader() []string {
	return []string{"[" + t.MapList + "]", "[<-" + t.Prev + " | " + t.Next + "->]    ", t.MapListDelimiter, ""}
}

func (t *textData) GetResult() []string {
	return []string{t.ResultDelimiter, "       " + t.Clear, t.ResultDelimiter, t.MapName + "    : ", t.ClearTime + "  : ", t.WrongCells + " : "}
}

func (t *textData) Complete() string {
	return t.CompleteMsg
}

func (t *textData) GetHelp() []string {
	return []string{"    " + t.Manual, t.ManualDelimiter, t.ExplArrowKey, t.ExplSpace, t.ExplX, t.ExplEnter, t.ExplEsc}
}

func (t *textData) GetCredit() []string {
	return []string{t.CreditDelimiter, "                " + t.Credit, t.CreditDelimiter, t.DeveloperInfo, t.License, t.ThankYouMsg, t.CreditDelimiter}
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
	return t.BlankBetweenMapAndTimer
}
