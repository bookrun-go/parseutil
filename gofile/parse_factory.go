package gofile

import (
	"github.com/bookrun-go/parseutil/tree"
	"io/ioutil"
)

type iParseFactory interface {
	Do(content string)
	DoByFile(filename string) error
	GetRootFunctions() []Function
}

type parseFactory struct {
	parses   []iParse
	tc       *tree.TreeCursor
	funcList []Function
}

func newParseGoFileFactory() iParseFactory {
	factory := &parseFactory{
		tc: tree.NewCursor(&ParseContext{
			Data: PreParse{},
		}),
	}
	_ = factory.register(&functionParse{})
	_ = factory.register(&annotateParse{})
	_ = factory.register(&stringParse{})
	_ = factory.register(&defaultParse{})
	return factory
}

// 可以考虑加锁。
func (f *parseFactory) register(p iParse) error {
	if f.parses == nil {
		f.parses = make([]iParse, 0)
	}

	f.parses = append(f.parses, p)
	return nil
}

func (f *parseFactory) Do(content string) {
	for _, char := range content {
		// 忽略无用的空格字符。
		currentContext := f.tc.CurrentTree.Data.(*ParseContext)
		if char == 32 && currentContext.CurrentStr == "" {
			continue
		}
		for _, p := range f.parses {
			end, _ := p.parse(char, f.tc)
			if end {
				break
			}
		}
	}
}

func (f *parseFactory) DoByFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	f.Do(string(b))
	return nil
}

func (f *parseFactory) GetRootFunctions() []Function {
	f.funcList = make([]Function, 0)
	f.recursive(f.tc.Head)
	return f.funcList
}

func (factory *parseFactory) recursive(leafTree *tree.MultiLeafTree) {
	if leafTree == nil {
		return
	}

	c := leafTree.Data.(*ParseContext)
	switch c.Data.(type) {
	case Function:
		{
			f := c.Data.(Function)
			factory.funcList = append(factory.funcList, f)
		}
	}

	for _, v := range leafTree.Children {
		factory.recursive(v)
	}
}

type iParse interface {
	parse(char int32, tc *tree.TreeCursor) (bool, error)
}

var NewParseFactory func() iParseFactory

func init() {
	NewParseFactory = newParseGoFileFactory
}
