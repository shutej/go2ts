package main

import (
	"bytes"
	"fmt"
	"strings"

	camelcase "github.com/segmentio/go-camelcase"
	"github.com/shutej/go2ts/model"
)

func UpperCamelcase(s string) string {
	s = camelcase.Camelcase(s)
	return strings.ToUpper(s[0:1]) + s[1:len(s)]
}

func LowerCamelcase(s string) string {
	s = camelcase.Camelcase(s)
	return strings.ToLower(s[0:1]) + s[1:len(s)]
}

func Package(s string) string {
	tmp := strings.SplitN(s, ".", 2)
	return fmt.Sprintf("%s_%s", tmp[0], tmp[1])
}

func PackageFile(s string) string {
	return fmt.Sprintf("%s.ts", Package(s))
}

func PackageReference(s string) string {
	p := Package(s)
	return fmt.Sprintf("import * as %s from %q;\n", p, "./"+p)
}

func Name(s string) string {
	tmp := strings.SplitN(s, ".", 2)
	return tmp[1]
}

type Generator struct {
	buffers  map[string]*bytes.Buffer
	imports_ map[string]map[string]struct{}
	stack    StringStack
}

func (self *Generator) Visit(types model.Types) {
	types.Visit(&TypeGenerator{self})
	types.Visit(&EmptyGenerator{self})
	types.Visit(&MarshalTypeGenerator{self})
	types.Visit(&MarshalGenerator{self})
	types.Visit(&UnmarshalGenerator{self})
}

func (self *Generator) Buffers() map[string]*bytes.Buffer {
	retval := map[string]*bytes.Buffer{}
	for pkg, epilogue := range self.buffers {
		buffer := bytes.NewBuffer([]byte{})
		fmt.Fprintf(buffer, prelude)
		imports := self.imports_[pkg]
		newline := false
		for import_ := range imports {
			fmt.Fprintf(buffer, PackageReference(import_))
			newline = true
		}
		if newline {
			fmt.Fprintf(buffer, "\n")
		}
		epilogue.WriteTo(buffer)
		retval[pkg] = buffer
	}
	return retval
}

func (self *Generator) imports() map[string]struct{} {
	if self.imports_ == nil {
		self.imports_ = map[string]map[string]struct{}{}
	}
	pkg := self.stack.top()
	imports, ok := self.imports_[pkg]
	if !ok {
		imports = map[string]struct{}{}
		self.imports_[pkg] = imports
	}
	return imports
}

func (self *Generator) buffer() *bytes.Buffer {
	if self.buffers == nil {
		self.buffers = map[string]*bytes.Buffer{}
	}

	pkg := self.stack.top()
	buffer, ok := self.buffers[pkg]
	if !ok {
		buffer = bytes.NewBuffer([]byte{})
		self.buffers[pkg] = buffer
	}
	return buffer
}

func (self *Generator) printf(format string, args ...interface{}) {
	fmt.Fprintf(self.buffer(), format, args...)
}

func (self *Generator) withPackage(name string, fn func()) {
	if name == "" {
		fn()
		return
	}
	self.stack.with(name, fn)
}
