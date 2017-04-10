package main

import (
	"github.com/shutej/go2ts/model"
)

type TypeGenerator struct {
	*Generator
}

func (self *TypeGenerator) withType(name string, resume func()) {
	self.withPackage(name, func() {
		// If there's no name, we skip creating a type.  All top-level types are
		// named and all types below them are references or anonymous types.
		if name == "" {
			resume()
			return
		}

		self.printf("export type T = ")
		resume()
		self.printf(";\n")
	})
}

func (self *TypeGenerator) VisitString(name string, resume func()) {
	self.withType(name, func() {
		self.printf("string")
	})
}

func (self *TypeGenerator) VisitInt(name string, resume func()) {
	self.withType(name, func() {
		self.printf("number")
	})
}

func (self *TypeGenerator) VisitFloat(name string, resume func()) {
	self.withType(name, func() {
		self.printf("number")
	})
}

func (self *TypeGenerator) VisitBool(name string, resume func()) {
	self.withType(name, func() {
		self.printf("boolean")
	})
}

func (self *TypeGenerator) VisitPtr(name string, resume func()) {
	self.withType(name, func() {
		resume()
		self.printf(" | null")
	})
}

func (self *TypeGenerator) VisitBytes(name string, resume func()) {
	self.withType(name, func() {
		self.printf("string")
	})
}

func (self *TypeGenerator) VisitSlice(name string, resume func()) {
	self.withType(name, func() {
		self.printf("Array<")
		resume()
		self.printf(">")
	})
}

func (self *TypeGenerator) VisitStruct(name string, _ []model.Field, resume func()) {
	self.withType(name, func() {
		self.printf("{ ")
		resume()
		self.printf(" }")
	})
}

func (self *TypeGenerator) VisitStructField(field model.Field, resume func()) {
	if field.Index != 0 {
		self.printf(" ")
	}
	self.printf("%s: ", LowerCamelcase(field.Name))
	resume()
	self.printf(";")
}

func (self *TypeGenerator) VisitMap(name string, resume func()) {
	self.withType(name, func() {
		self.printf("{ [k: string]: ")
		resume()
		self.printf(" }")
	})
}

func (self *TypeGenerator) VisitCustom(name string, resume func()) {
	self.VisitReference(name, resume)
}

func (self *TypeGenerator) VisitReference(name string, resume func()) {
	if self.stack.top() == name {
		self.printf("T")
	} else {
		imports := self.imports()
		imports[name] = struct{}{}
		self.printf("%s.T", Package(name))
	}
}
