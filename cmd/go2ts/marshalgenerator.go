package main

import (
	"github.com/shutej/go2ts/model"
)

type MarshalGenerator struct {
	*Generator
}

func (self *MarshalGenerator) withMarshal(name string, resume func()) {
	self.withPackage(name, func() {
		// If there's no name, we skip creating a type.  All top-level types are
		// named and all types below them are references or anonymous types.
		if name == "" {
			resume()
			return
		}

		// TODO(shutej): Is there a way to tighten the return type, or omit and infer it?
		self.printf("export function marshal(x: T): MarshalT { return ")
		resume()
		self.printf("(x); }\n")
	})
}

func (self *MarshalGenerator) VisitString(name string, resume func()) {
	self.withMarshal(name, func() {
		self.printf("_.identityString")
	})
}

func (self *MarshalGenerator) VisitInt(name string, resume func()) {
	self.withMarshal(name, func() {
		self.printf("_.identityNumber")
	})
}

func (self *MarshalGenerator) VisitFloat(name string, resume func()) {
	self.withMarshal(name, func() {
		self.printf("_.identityNumber")
	})
}

func (self *MarshalGenerator) VisitBool(name string, resume func()) {
	self.withMarshal(name, func() {
		self.printf("_.identityBoolean")
	})
}

func (self *MarshalGenerator) VisitPtr(name string, resume func()) {
	self.withMarshal(name, func() {
		self.printf("((f) => (x) => x !== null ? f(x) : null)(")
		resume()
		self.printf(")")
	})
}

func (self *MarshalGenerator) VisitBytes(name string, resume func()) {
	self.withMarshal(name, func() {
		self.printf("window.btoa")
	})
}

func (self *MarshalGenerator) VisitSlice(name string, resume func()) {
	self.withMarshal(name, func() {
		self.printf("_.fmapArray(")
		resume()
		self.printf(")")
	})
}

func (self *MarshalGenerator) VisitStruct(name string, fields []model.Field, resume func()) {
	self.withMarshal(name, func() {
		self.printf("(function(x) { return { ")
		resume()
		self.printf(" }; })")
	})
}

func (self *MarshalGenerator) VisitStructField(field model.Field, resume func()) {
	if field.Index != 0 {
		self.printf(", ")
	}
	self.printf("%s: ", field.Name)
	resume()
	self.printf("(x.%s)", LowerCamelcase(field.Name))
}

func (self *MarshalGenerator) VisitMap(name string, resume func()) {
	self.withMarshal(name, func() {
		self.printf("_.fmapObject(")
		resume()
		self.printf(")")
	})
}

func (self *MarshalGenerator) VisitCustom(name string, resume func()) {
	self.VisitReference(name, resume)
}

func (self *MarshalGenerator) VisitReference(name string, resume func()) {
	if self.stack.top() == name {
		self.printf("marshal")
	} else {
		imports := self.imports()
		imports[name] = struct{}{}
		self.printf("%s.marshal", Package(name))
	}
}
