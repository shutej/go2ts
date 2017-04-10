go2flow
=======

Generates Typescript types corresponding to Go's JSON types.

All code for Go type `package.Type` is written in `package_Type.ts`.

The type itself is in a type alias, `package_Type.T`.  We use type aliases
because Go type may include `int`, `float`, `string`, or `bool`, which cannot
be represented by Javascript classes or compound types.

We make an `empty()` function, which returns a zero value for type `T`.

We make a `marshal(t)` function, which returns an object that can then be used
with `JSON.stringify()`.  Making `marshal(t)` independent of `toJSON` allows
custom behavior to be attached to primitive types.

We make an `unmarshal(data)` function which, when passed an object from
`JSON.parse()`, returns an instance of `T` or raises an `UnmarshalException`.
