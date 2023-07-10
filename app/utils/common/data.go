package common

// 三元运算
func ThreeEyes(path bool, a any, b any) any {
	if path == true {
		return a
	} else {
		return b
	}
}
