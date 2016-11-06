package f

func Collect(l func(interface{}) interface{}) (a []interface{}) { return a }
func Each(f func(interface{}), l ...interface{})                {}
func When(f func() bool, fns ...func())                         {}
func And(fns ...func() bool)                                    {}
func Try(a func() error, b, c func())                           {}
func Filter(f func(interface{}), l ...interface{})              {}
func Cons(l ...interface{}) interface{}                         { return nil }
func Reverse(l ...interface{}) interface{}                      { return nil }
func Sort(f func(interface{}), l ...interface{}) interface{}    { return nil }
