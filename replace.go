package doclize

import (
	"bytes"
	tools "github.com/iEvan-lhr/exciting-tool"
)

func (d *Docx) ReplaceC(oldString string, newString string, num int) (l int) {
	l = d.replace(oldString, newString, num)
	return
}

func (d *Docx) replace(oldString, newString string, num int) int {
	r := record{Index: findTextIndexes([]byte(d.content), []byte(oldString))}
	d.content = r.copy([]byte(d.content), tools.Make(newString))
	return 1
}

func findTextIndexes(text []byte, target []byte) (indexes []int) {
	startIndex := 0
	k := 0
	indexes = make([]int, len(target))
	tempText := text
	for {
		start := bytes.Index(tempText, []byte("<w:t>")) + 5
		end := bytes.Index(tempText, []byte("</w:t>"))
		if start == -1 || end == -1 {
			return
		}
		if start < end {
			ans := tempText[start:end]
			startIndex += start
			for i := range ans {
				if k == len(target) {
					return
				}
				if target[k] == ans[i] {
					indexes[k] = startIndex + i
					k++
				} else {
					k = 0
				}
			}
			startIndex += len(ans) + 6
			tempText = text[startIndex:]
		} else {
			startIndex += end + 6
			tempText = text[startIndex:]
		}
	}
}
