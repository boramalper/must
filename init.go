// Copyright (c) 2019 Mert Bora ALPER <bora@boramalper.org>
// See LICENSE for details.

package must

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

// ------

func (mc Context) Must(err error) {
	if err != nil {
		mc.handle(err)
	}
}

func (mc Context) M(e error) { mc.Must(e) }

func (mc Context) MustVal(val interface{}, err error) interface{} {
	if err != nil {
		mc.handle(err)
		return nil
	}

	return val
}

func (mc Context) MV(v interface{}, e error) interface{} { return mc.MustVal(v, e) }

func (mc Context) MustValVoid(val interface{}, err error) {
	mc.MustVal(val, err)
}

func (mc Context) MVV(v interface{}, e error) { mc.MustValVoid(v, e) }

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

