package doclize

import (
	tools "github.com/iEvan-lhr/exciting-tool"
	"testing"
)

func loadFile(file string) *Docx {
	r, err := ReadDocxFile(file)
	if err != nil {
		panic(err)
	}

	return r.Editable()
}

func TestCheck(t *testing.T) {
	d := loadFile("./model3_Car.docx")

	d.ReplaceC("(+NMB%)", "12345", -1)

	tools.Error(d.WriteToFile("./model3_Car1.docx"))
}
