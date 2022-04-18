// GoAst project main.go
package main

import (
	"fmt"
)

func main() {
	//jc, _ := parseFile(`D:\tmp\sqlite\astdemo\astres.json`)
	//jc, _ := parseFile(`astdemo.json`)
	jc, _ := parseFile(`astqoi2.json`)
	fmt.Println("================")
	//iterContent(jc.Data)
	var ac AstC
	ac.convert(jc)
	//iterContentLayer(jc.Data, 6)
}
