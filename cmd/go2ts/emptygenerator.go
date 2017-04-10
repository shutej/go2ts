package main

import (
	"github.com/shutej/go2ts/model"
)

type EmptyGenerator struct {
	*Generator
}

func (self *EmptyGenerator) withEmpty(name string, resume func()) {
	self.withPackage(name, func() {
		// If there's no name, we skip creating a type.  All top-level types are
		// named and all types below them are references or anonymous types.
		if name == "" {
			resume()
			return
		}

		self.printf("export function empty(): T { return ")
		resume()
		self.printf("; }\n")
	})
}

func (self *EmptyGenerator) VisitString(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("%q", "")
	})
}

func (self *EmptyGenerator) VisitInt(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("0")
	})
}

func (self *EmptyGenerator) VisitFloat(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("0.0")
	})
}

func (self *EmptyGenerator) VisitBool(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("false")
	})
}

func (self *EmptyGenerator) VisitPtr(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("null")
	})
}

func (self *EmptyGenerator) VisitBytes(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("%q", "")
	})
}

func (self *EmptyGenerator) VisitSlice(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("[]")
	})
}

func (self *EmptyGenerator) VisitStruct(name string, fields []model.Field, resume func()) {
	self.withEmpty(name, func() {
		self.printf("{ ")
		resume()
		self.printf(" }")
	})
}

func (self *EmptyGenerator) VisitStructField(field model.Field, resume func()) {
	if field.Index != 0 {
		self.printf(", ")
	}
	self.printf("%s: ", LowerCamelcase(field.Name))
	resume()
}

func (self *EmptyGenerator) VisitMap(name string, resume func()) {
	self.withEmpty(name, func() {
		self.printf("{}")
	})
}

func (self *EmptyGenerator) VisitCustom(name string, resume func()) {
	self.VisitReference(name, resume)
}

func (self *EmptyGenerator) VisitReference(name string, resume func()) {
	if self.stack.top() == name {
		// TODO(shutej): This would be an infinite loop.  Assertion strategy?
		self.printf("empty()")
	} else {
		imports := self.imports()
		imports[name] = struct{}{}
		self.printf("%s.empty()", Package(name))
	}
}
