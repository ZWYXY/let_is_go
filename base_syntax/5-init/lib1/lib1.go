package lib1

import "fmt"

func init() {
	fmt.Println("-----lib1.init() invoked!-----")
}

/*
	关于规则，
	1.首字母大写表示，这里的方法是public的，任意地方都能调，如果是小写开头的，表示是私有的，只有包内能调
	2.方法最好不要以包名开头比如：libTest()这样是不被提倡的
*/
func MyDisplay() {
	fmt.Println("lib1.MyDisplay() invoked!")
}
