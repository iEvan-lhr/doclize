package doclize

import (
	"bytes"
	tools "github.com/iEvan-lhr/exciting-tool"
)

type t struct {
	other   []string
	content []string
	steps   map[int]struct {
		model int
		index int
	}
}

type location struct {
	index int
	start int
	end   int
}

func (t *t) replaceOnce(old, new string) int {
	var mark []location
	for i := 0; i < len(t.content); {
		if k := bytes.IndexByte([]byte(t.content[i]), old[0]); k != -1 {
			start := location{
				index: i,
				start: k,
			}
			hash, m, n := t.content[i][k], 0, i
			for hash == old[m] {
				if m == len(old)-1 {
					mark = append(mark, start)
					mark = append(mark, location{
						index: n,
						end:   k,
					})
					goto Find
				}
				if k != len(t.content[n])-1 {
					k++
					m++
				} else {
					n++
					k = 0
					m++
				}

				hash = t.content[n][k]
			}
		} else {
			i++
		}
	}
	return -1
Find:
	if mark[0].index == mark[1].index {
		if mark[0].start == 0 {
			t.content[mark[0].index] = tools.Make().Append(new, t.content[mark[1].index][mark[1].end+1:]).String()
		} else {
			t.content[mark[0].index] = tools.Make().Append(t.content[mark[0].index][:mark[0].start], new, t.content[mark[1].index][mark[1].end+1:]).String()
		}
	} else {
		if mark[0].start == 0 {
			t.content[mark[0].index] = tools.Make().Append(new).String()
		} else {
			t.content[mark[0].index] = tools.Make().Append(t.content[mark[0].index][:mark[0].start], new).String()
		}
		for l := mark[0].index + 1; l < mark[1].index; l++ {
			t.content[l] = ""
		}
		t.content[mark[1].index] = t.content[mark[1].index][mark[1].end+1:]
	}
	return 1
}

func (t *t) replace(old, new string, num int) (l int) {
	if num == -1 {
		for t.replaceOnce(old, new) != -1 {
			l++
		}
	} else {
		for i := 0; i < num; i++ {
			t.replaceOnce(old, new)
		}
		l = num
	}
	return
}

func (t *t) getContent() string {
	ans := tools.Make()
	for i := 0; i < len(t.steps); i++ {
		switch t.steps[i].model {
		case 0:
			ans.Append(t.content[t.steps[i].index])
		case 1:
			ans.Append(t.other[t.steps[i].index])
		}
	}
	return ans.String()
}

//func findTextIndexes(text, target string) []int {
//	indexes := []int{}
//	startTag := "<jio>"
//	endTag := "</jio>"
//	startIndex := 0
//	for {
//		startIndex = strings.Index(text, startTag)
//		if startIndex == -1 {
//			break
//		}
//		endIndex := strings.Index(text, endTag)
//		if endIndex == -1 {
//			break
//		}
//		startIndex += len(startTag)
//		subText := text[startIndex:endIndex]
//		index := strings.Index(subText, target)
//		if index != -1 {
//			indexes = append(indexes, startIndex+index)
//		}
//		text = text[endIndex+len(endTag):]
//	}
//	return indexes
//}
