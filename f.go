package f

func Collect(l func(interface{}) interface{}) (a []interface{}) { return a }
func Each(f func(interface{}), l ...interface{}) {
	if f != nil {
		for _, i := range l {
			f(i)
		}
	}
}
func When(f func() bool, fns ...func()) {
	if f != nil && f() {
		if fns != nil {
			for _, f := range fns {
				if f != nil {
					f()
				}
			}
		}
	}
}
func And(fns ...func() bool) {
	if fns != nil {
		for _, f := range fns {
			if f != nil {
				if !f() {
					break
				}
			}
		}
	}
}
func Try(a func() error, b func(error), c func()) {
	if a != nil {
		if e := a(); e != nil {
			if b != nil {
				b(e)
			}
		}
		if c != nil {
			c()
		}
	}
}
func Filter(f func(interface{}), l ...interface{})           {}
func Cons(l ...interface{}) interface{}                      { return nil }
func Reverse(l ...interface{}) interface{}                   { return nil }
func Sort(f func(interface{}), l ...interface{}) interface{} { return nil }
