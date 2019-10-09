package gofile

import (
	"github.com/bookrun-go/parseutil/tree"
)

type annotateParse struct {
}

func (a *annotateParse) parse(char int32, tc *tree.TreeCursor) (bool, error) {
	//context.CurrentTree.Data
	currentContext := tc.CurrentTree.Data.(*ParseContext)
	switch currentContext.Data.(type) {
	case Annotate:
	case PreParse:
	default:
		return false, nil
	}

	tempStr := currentContext.CurrentStr + string(char)
	if currentContext.Process == ParseProcessInit && ("/*" == tempStr || "//" == tempStr) {
		tempCtx := &ParseContext{}
		tempCtx.Data = Annotate{}
		tempCtx.Process = ParseProcessPrepare
		tempCtx.CurrentStr = "" // 重置当前字符串
		tempCtx.WrapSymbol = a.getWarpSymbol(tempStr)

		tc.AddToNewLeaf(tempCtx)
		return true, nil
	}

	// 不属于当前范围
	if currentContext.Process == ParseProcessInit {
		return false, nil
	}

	// //这种注释结束 \n
	if string(char) == currentContext.WrapSymbol.End {
		_ = tc.ToParent()
		return true, nil
	}

	if currentContext.CurrentStr == "" {
		currentContext.CurrentStr = string(char)
		return true, nil
	}

	ts := currentContext.CurrentStr[len(currentContext.CurrentStr)-1:]
	if ts+string(char) == currentContext.WrapSymbol.End {
		_ = tc.ToParent()
		return true, nil
	}

	currentContext.CurrentStr = currentContext.CurrentStr + string(char)
	return true, nil
}

func (a *annotateParse) getWarpSymbol(str string) *WrapSymbol {
	if str == "/*" {
		return &WrapSymbol{
			Start: "/*",
			End:   "*/",
		}
	}

	return &WrapSymbol{
		Start: "//",
		End:   "\n",
	}
}

// // /*
