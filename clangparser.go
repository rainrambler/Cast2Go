package main

type AstC struct {
	tud *TranslationUnitDecl
}

func (p *AstC) convert(jc *JsonContent) {
	p.tud = convertTranslationUnitDecl(jc.Data)
}
