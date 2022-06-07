/*
Package quickjs Go bindings to QuickJS: a fast, small, and embeddable ES2020 JavaScript interpreter
*/
package quickjs

/*
#cgo CFLAGS: -I./deps/include/quickjs
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/deps/libs/darwin_amd64 -lquickjs
*/
import "C"
