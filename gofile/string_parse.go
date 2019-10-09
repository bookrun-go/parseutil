package gofile

import "github.com/bookrun-go/parseutil/tree"

type stringParse struct {
}

func (s *stringParse) parse(char int32, tc *tree.TreeCursor) (bool, error) {
	currentContext := tc.CurrentTree.Data.(*ParseContext)

	switch currentContext.Data.(type) {
	case Str:
	case PreParse:
	default:
		return false, nil
	}

	if currentContext.Process == ParseProcessInit && s.isStr(char) {
		tempCtx := &ParseContext{}
		tempCtx.Process = ParseProcessing
		tempCtx.Data = Str{}
		tempCtx.WrapSymbol = s.getWrapSymbol(char)
		tc.AddToNewLeaf(tempCtx)
		return true, nil
	}

	if currentContext.Process == ParseProcessInit {
		return false, nil
	}

	//jj := hhh
	//fmt.Println(jj)

	if string(char) == currentContext.WrapSymbol.End {
		if currentContext.CurrentStr == "" {
			// 结束
			_ = tc.ToParent()
			return true, nil
		}

		ts := currentContext.CurrentStr[len(currentContext.CurrentStr)-1:]
		if ts != "\\" { // 判断是否转义。
			// 结束
			_ = tc.ToParent()
			return true, nil
		}
	}

	currentContext.CurrentStr = currentContext.CurrentStr + string(char)
	return true, nil
}

func (s *stringParse) getWrapSymbol(char int32) *WrapSymbol {
	return &WrapSymbol{
		Start: string(char),
		End:   string(char),
	}
}

func (s *stringParse) isStr(char int32) bool {
	return char == 34 || char == 39 || char == 96
}

// " 34
//' 39
//` 96
