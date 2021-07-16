package main

/**
	A 引入 B ，在执行A的main方法前会先到B中执行B的init方法，如果B还引入了C，那么会再到C中执行它的init方法，一直到没有引入其它包的文件中，
	调用它的init方法，然后逐层调用init方法直到返回到A（A中的init方法也要被调用），然后再执行A.main()

	这里顺便介绍一下import导包方式
	import (
		//"fmt"          // go中 没有使用导入的包，编译无法通过，如果需要执行它的init方法而引入这个包，可以这么做  _ "fmt"
 		"5-init/lib"     // 如果想给引入的包起个别名，可以这么做 my "5-init/lib"
		"5-init/lib1"    // 如果想直接用方法名 调用引入包中的方法 可以这么做 . "5-init/lib" 但是不推荐，因为不同包会有重名方法
	)
	示例：
	import (
		. "5-init/lib"
		my "5-init/lib1"
		_ "fmt"
	)
*/import (
	"communication/base_syntax/5-init/lib"
	"communication/base_syntax/5-init/lib1"
)

func main() {

	//lib.MyDisplay()
	//lib1.MyDisplay()
	lib.MyDisplay()
	lib1.MyDisplay()
}
