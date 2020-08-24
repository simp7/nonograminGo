package control

import (
	"encoding/csv"
	"github.com/simp7/nonograminGo/util"
	"os"
)

type CSVManager struct {
	file   *os.File
	reader *csv.Reader
	writer *csv.Writer
	csvMap map[string]string
}

func NewCSVManager(fileName string) *CSVManager {

	m := CSVManager{}
	var err error

	m.file, err = os.Open(fileName)
	util.CheckErr(err)

	m.reader = csv.NewReader(m.file)
	m.writer = csv.NewWriter(m.file)
	m.csvMap = make(map[string]string)

	m.read()
	return &m

}

func (m *CSVManager) read() {

	rows, err := m.reader.ReadAll()
	util.CheckErr(err)

	for _, row := range rows {
		m.csvMap[row[0]] = row[1]
	}

}

func (m *CSVManager) Write(key string) {

}
