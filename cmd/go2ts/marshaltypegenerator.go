package main

import (
	"github.com/shutej/go2ts/model"
)

type MarshalTypeGenerator struct {
	*Generator
}

func (self *MarshalTypeGenerator) withType(name string, resume func()) {
	self.withPackage(name, func() {
		// If there's no name, we skip creating a type.  All top-level types are
		// named and all types below them are references or anonymous types.
		if name == "" {
			resume()
			return
		}

		self.printf("export type MarshalT = ")
		resume()
		self.printf(";\n")
	})
}

func (self *MarshalTypeGenerator) VisitString(name string, resume func()) {
	self.withType(name, func() {
		self.printf("string")
	})
}

func (self *MarshalTypeGenerator) VisitInt(name string, resume func()) {
	self.withType(name, func() {
		self.printf("number")
	})
}

func (self *MarshalTypeGenerator) VisitFloat(name string, resume func()) {
	self.withType(name, func() {
		self.printf("number")
	})
}

func (self *MarshalTypeGenerator) VisitBool(name string, resume func()) {
	self.withType(name, func() {
		self.printf("boolean")
	})
}

func (self *MarshalTypeGenerator) VisitPtr(name string, resume func()) {
	self.withType(name, func() {
		resume()
		self.printf(" | null")
	})
}

func (self *MarshalTypeGenerator) VisitBytes(name string, resume func()) {
	self.withType(name, func() {
		self.printf("string | null")
	})
}

func (self *MarshalTypeGenerator) VisitSlice(name string, resume func()) {
	self.withType(name, func() {
		self.printf("Array<")
		resume()
		self.printf(">")
	})
}

func (self *MarshalTypeGenerator) VisitStruct(name string, _ []model.Field, resume func()) {
	self.withType(name, func() {
		self.printf("{ ")
		resume()
		self.printf(" }")
	})
}

func (self *MarshalTypeGenerator) VisitStructField(field model.Field, resume func()) {
	if field.Index != 0 {
		self.printf(" ")
	}
	self.printf("%s: ", field.Name)
	resume()
	self.printf(";")
}

func (self *MarshalTypeGenerator) VisitMap(name string, resume func()) {
	self.withType(name, func() {
		self.printf("{ [k: string]: ")
		resume()
		self.printf(" }")
	})
}

func (self *MarshalTypeGenerator) VisitCustom(name string, resume func()) {
	self.VisitReference(name, resume)
}

func (self *MarshalTypeGenerator) VisitReference(name string, resume func()) {
	// TODO(shutej): This, probably.
	if self.stack.top() == name {
		self.printf("MarshalT")
	} else {
		imports := self.imports()
		imports[name] = struct{}{}
		self.printf("%s.MarshalT", Package(name))
	}
}
