package main

import (
	"fmt"
)

type AstC struct {
	tud *TranslationUnitDecl
}

func (p *AstC) convert(jc *JsonContent) {
	p.tud = convertTranslationUnitDecl(jc.Data)
	fmt.Println("-=-=-=-=-=-=-=")
	fmt.Println(p.tud.t2go())
}
