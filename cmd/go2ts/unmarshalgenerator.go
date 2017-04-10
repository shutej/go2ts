package main

import (
	"github.com/shutej/go2ts/model"
)

type UnmarshalGenerator struct {
	*Generator
}

func (self *UnmarshalGenerator) withUnmarshal(name string, resume func()) {
	self.withPackage(name, func() {
		// If there's no name, we skip creating a type.  All top-level types are
		// named and all types below them are references or anonymous types.
		if name == "" {
			resume()
			return
		}

		self.printf("export function unmarshal(x: MarshalT): T { return ")
		resume()
		self.printf("(x); }\n")
	})
}

func (self *UnmarshalGenerator) VisitString(name string, resume func()) {
	self.withUnmarshal(name, func() {
		self.printf("_.identityString")
	})
}

func (self *UnmarshalGenerator) VisitInt(name string, resume func()) {
	self.withUnmarshal(name, func() {
		self.printf("_.identityNumber")
	})
}

func (self *UnmarshalGenerator) VisitFloat(name string, resume func()) {
	self.withUnmarshal(name, func() {
		self.printf("_.identityNumber")
	})
}

func (self *UnmarshalGenerator) VisitBool(name string, resume func()) {
	self.withUnmarshal(name, func() {
		self.printf("_.identityBoolean")
	})
}

func (self *UnmarshalGenerator) VisitPtr(name string, resume func()) {
	self.withUnmarshal(name, func() {
		self.printf("((f) => (x) => x !== null ? f(x) : null)(")
		resume()
		self.printf(")")
	})
}

func (self *UnmarshalGenerator) VisitBytes(name string, resume func()) {
	self.withUnmarshal(name, func() {
		self.printf("_.bytesUnmarshal")
	})
}

func (self *UnmarshalGenerator) VisitSlice(name string, resume func()) {
	self.withUnmarshal(name, func() {
		self.printf("_.fmapArray(")
		resume()
		self.printf(")")
	})
}

func (self *UnmarshalGenerator) VisitStruct(name string, fields []model.Field, resume func()) {
	self.withUnmarshal(name, func() {
		self.printf("(function(x) { return { ")
		resume()
		self.printf(" }; })")
	})
}

func (self *UnmarshalGenerator) VisitStructField(field model.Field, resume func()) {
	if field.Index != 0 {
		self.printf(", ")
	}
	self.printf("%s: ", LowerCamelcase(field.Name))
	if field.OmitEmpty {
		self.printf("(x.hasOwnProperty(%q) ? ", field.Name)
	}
	resume()
	self.printf("(x.%s)", field.Name)
	if field.OmitEmpty {
		self.printf(" : ")
		field.Type.Visit(&EmptyGenerator{self.Generator})
		self.printf(")")
	}
}

func (self *UnmarshalGenerator) VisitMap(name string, resume func()) {
	self.withUnmarshal(name, func() {
		self.printf("_.fmapObject(")
		resume()
		self.printf(")")
	})
}

func (self *UnmarshalGenerator) VisitCustom(name string, resume func()) {
	self.VisitReference(name, resume)
}

func (self *UnmarshalGenerator) VisitReference(name string, resume func()) {
	if self.stack.top() == name {
		self.printf("unmarshal")
	} else {
		imports := self.imports()
		imports[name] = struct{}{}
		self.printf("%s.unmarshal", Package(name))
	}
}
