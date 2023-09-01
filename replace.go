package doclize

import (
	"bytes"
	tools "github.com/iEvan-lhr/exciting-tool"
)

var i = 0

func (d *Docx) ReplaceC(oldString string, newString string, num int) (l int) {
	var index int
	for num == -1 || l < num {
		if index = d.replace(oldString, newString, index); index == -1 {
			return l
		} else {
			l++
		}
	}
	return
}

func (d *Docx) replace(oldString, newString string, startIndex int) int {
	r := record{Index: findTextIndexes([]byte(d.content), []byte(oldString), startIndex)}
	if r.Index[0] == 0 {
		return -1
	}
	d.content = r.copy([]byte(d.content), tools.Make(newString))
	//tools.Error(d.WriteToFile("./word/" + tools.Make(i).String() + ".docx"))
	return r.LastIndex
}

func findTextIndexes(text, target []byte, startIndex int) (indexes []int) {
	k := 0
	indexes = make([]int, len(target))
	tempText := text[startIndex:]
	for {
		var start, end int
		if t1 := bytes.Index(tempText, []byte("<w:t")); t1 != -1 {
			start = bytes.IndexByte(tempText[t1:], '>') + t1 + 1
			end = bytes.Index(tempText, []byte("</w:t>"))
			if start == -1 || end == -1 {
				indexes[0] = 0
				return
			}
		} else {
			indexes[0] = 0
			return
		}
		if start < end {
			ans := tempText[start:end]
			startIndex += start
			for i := range ans {
				if target[k] == ans[i] {
					indexes[k] = startIndex + i
					k++
				} else {
					k = 0
				}
				if k == len(target) {
					return
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
