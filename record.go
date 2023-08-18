package doclize

import (
	tools "github.com/iEvan-lhr/exciting-tool"
	"sort"
)

type record struct {
	Index []int
}

func (r *record) copy(content []byte, new *tools.String) string {
	sort.Ints(r.Index)
	remove := new.Len() - 1
	ans := tools.Make().Append(content[:r.Index[0]], new, content[r.Index[0]+1:])
	content = ans.Bytes()
	for i := 1; i < len(r.Index); i++ {
		content[r.Index[i]+remove] = ' '
	}
	return string(content)
}
