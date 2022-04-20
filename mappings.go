package main

type ToGoMappings struct {
	funcmapping map[string]string
	typemapping map[string]string
}

var gomap ToGoMappings

func init() {
	gomap.funcmapping = map[string]string{
		"sizeof": "len",
		"memset": "Mem",    // TODO
		"malloc": "Malloc", // TODO
		"free":   "",
	}
	gomap.typemapping = map[string]string{"unsigned char": "byte"}
}
