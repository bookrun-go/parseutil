package gofile

import (
	"github.com/bookrun-go/parseutil/tree"
)

type defaultParse struct {
}

func (o *defaultParse) parse(char int32, tc *tree.TreeCursor) (bool, error) {
	currentContext := tc.CurrentTree.Data.(*ParseContext)
	// 不等于初始状态，表示已经又在处理了。
	if currentContext.Process != ParseProcessInit {
		return false, nil
	}
	//jj := hhh
	//fmt.Println(jj)

	//空白字符和换行符
	if char == 32 || char == 41 || char == 40 || char == 10 || char == 9 {
		currentContext.CurrentStr = ""
		_ = tc.ToParent()
		//fmt.Println(err)
		return true, nil
	}

	currentContext.CurrentStr = currentContext.CurrentStr + string(char)
	tc.CurrentTree.Data = currentContext
	return false, nil
}
