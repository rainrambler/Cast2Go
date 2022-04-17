// GoAst project main.go
package main

import (
	"fmt"
)

func main() {
	jc, _ := parseFile(`astdemo.json`)
	//fmt.Printf("%+v\n", jc)
	//jc, _ := parseFile(`demo.json`)
	//fmt.Printf("%v\n", jc)
	fmt.Println("================")
	//iterContent(jc.Data)
	var ac AstC
	ac.convert(jc)
	//iterContentLayer(jc.Data, 6)
}
