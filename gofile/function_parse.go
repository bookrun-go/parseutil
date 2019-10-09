package gofile

import (
	"github.com/bookrun-go/parseutil/tree"
)

type functionContext struct {
	StartSymbolNum int
	EndSymbolNum   int
}

type functionParse struct {
}

func (f *functionParse) parse(char int32, tc *tree.TreeCursor) (bool, error) {
	// 是否需要把factory放到context,通过context执行下一步。
	currentContext := tc.CurrentTree.Data.(*ParseContext)
	switch currentContext.Data.(type) {
	case Function:
	case PreParse:
	default:
		return false, nil
	}

	if currentContext.Process == ParseProcessInit && char == 32 && currentContext.CurrentStr == "func" {
		tempCtx := &ParseContext{}
		tempCtx.Data = Function{}
		tempCtx.Process = ParseProcessPrepare
		tempCtx.CurrentStr = "" // 重置当前字符串

		tc.AddToNewLeaf(tempCtx)
		return true, nil
	}

	// 不属于当前范围
	if currentContext.Process == ParseProcessInit {
		return false, nil
	}

	// 开始解析
	if currentContext.Process == ParseProcessPrepare {
		f.pre(char, tc)
		return true, nil
	}

	// 空白字符
	if char == 9 || char == 10 || char == 32 {
		return true, nil
	}

	if char == 125 {
		// 结束
		_ = tc.ToParent()
		return true, nil
	}

	tc.AddToNewLeaf(&ParseContext{
		Data: &PreParse{},
	})

	//newCtx := &ContextTree{
	//	CurrentContext: &ParseContext{
	//		Data: PreParse{},
	//	},
	//	PreContext: &ContextTree{},
	//}
	//*newCtx.PreContext = *context
	//if len(newCtx.PreContext.ChildrenContext) == 0 {
	//	newCtx.PreContext.ChildrenContext = make([]*ContextTree, 0)
	//}
	//newCtx.PreContext.ChildrenContext = append(newCtx.PreContext.ChildrenContext, newCtx)
	//
	//*context = *newCtx

	// 新建一个初始上下文，并且切换过去。
	return true, nil
}

// 先只解析出函数名
func (f *functionParse) pre(char int32, tc *tree.TreeCursor) {
	currentContext := tc.CurrentTree.Data.(*ParseContext)
	if char == 123 {
		currentContext.Process = ParseProcessing
		currentContext.CurrentStr = ""
		return
	}

	//非空白字符和换行符
	if char != 32 && char != 41 && char != 40 {
		currentContext.CurrentStr = currentContext.CurrentStr + string(char)
		return
	}

	if currentContext.CurrentStr == "" {
		return
	}

	tempData := currentContext.Data.(Function)
	if char == 40 && tempData.Name == "" {
		tempData.Name = currentContext.CurrentStr

		currentContext.Data = tempData
		currentContext.CurrentStr = ""
		return
	}

	if char == 41 {
		currentContext.CurrentStr = ""
		return
	}
}
