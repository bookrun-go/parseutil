package gofile

// go文件映射结构体
type (
	FileRoot struct {
		SimpleFilename string      // 文件名称，不包含路径
		PackageName    string      // 包名称
		Structs        []*Struct   // 结构体列表
		Functions      []*Function // 函数
		Vars           []*Var      // 变量列表
		Constants      []*Constant //常数列表
		Imports        []*Import   // 引用包列表
	}

	Struct struct {
		Name       string      // 结构体名称
		Permission int         // 权限类型
		Functions  []*Function // 函数列表
		Properties []*Property //属性列表
	}

	Function struct {
		Name       string // 函数名称
		Permission int    // 权限类型
		// params is future plan
		// back params is future plan
	}

	Var struct {
		Name       string // 变量
		Permission int    // 权限类型
		//type is future plan
	}

	Constant struct {
		Name       string // 常数名称
		Permission int    // 权限类型
	}

	Import struct {
		FullPath string // 应用路径
		Alias    string // 别名
	}

	Property struct {
		//this is future plan
	}

	Annotate struct {
		Content string // 注释内容
	}

	Str struct {
		Content string // 字符串内容
	}

	PreParse struct {
		Content string // 解析预处理，未知当前类型的情况使用。
	}
)

const (
	private = iota // 私有权限
	public         // 公开权限
)

//go文件映射结构体结束

type WrapSymbol struct {
	Start string // 开始符号
	End   string // 结束符号
}

type ParseContext struct {
	Data       interface{} // 当前数据
	CurrentStr string      // 当前字符串。
	Process    int         // 进度
	WrapSymbol *WrapSymbol //特定方法解析优先于包裹符号。
}

// 解析进度
const (
	ParseProcessInit    = iota // 初始阶段，等待开始，还不知当前具体类型
	ParseProcessPrepare        // 准备阶段
	ParseProcessing            // 进行中
	// 应该没有end,end是状态不会有进度
)
