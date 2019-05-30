// Package must provides utility functions to handle unrecoverable errors in
// Go.
package must

// Context provides all the necessary information to handle errors. Currently
// only `handler` is stored.
type Context struct {
	handler func(error)
}

var defaultContext Context

func init() {
	defaultContext = Context{}
}

func Must(err error) { defaultContext.Must(err) }

func M(e error) { defaultContext.M(e) }

func MustVal(val interface{}, err error) interface{} { return defaultContext.MustVal(val, err) }

func MV(v interface{}, e error) interface{} { return defaultContext.MV(v, e) }

func MustValVoid(val interface{}, err error) {
	defaultContext.MustVal(val, err)
}

func MVV(v interface{}, e error) { defaultContext.MustValVoid(v, e) }

func SetHandler(hnd func (error)) { defaultContext.SetHandler(hnd) }

func New(handler func(error)) Context {
	return Context{
		handler: handler,
	}
}

// Handles if err != nil.
func (mc Context) Must(err error) {
	if err != nil {
		mc.handle(err)
	}
}

// Shorthand for Must.
func (mc Context) M(e error) { mc.Must(e) }

// Handles if err != nil, returns val otherwise.
func (mc Context) MustVal(val interface{}, err error) interface{} {
	if err != nil {
		mc.handle(err)
		return nil
	}

	return val
}

// Shorthand for MustVal.
func (mc Context) MV(v interface{}, e error) interface{} { return mc.MustVal(v, e) }

// Handles if err != nil and ignores val.
func (mc Context) MustValVoid(val interface{}, err error) {
	mc.MustVal(val, err)
}

// Shorthand for MustValVoid.
func (mc Context) MVV(v interface{}, e error) { mc.MustValVoid(v, e) }

// Changes the handler function.
func (mc Context) SetHandler(hnd func (error)) {
	mc.handler = hnd
}

func (mc Context) handle(err error) {
	if mc.handler == nil {
		panic(err)
	} else {
		mc.handler(err)
	}
}
