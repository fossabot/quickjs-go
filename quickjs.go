/*
Package quickjs Go bindings to QuickJS: a fast, small, and embeddable ES2020 JavaScript interpreter
*/
package quickjs

/*
#cgo CFLAGS: -D_GNU_SOURCE
#cgo CFLAGS: -DCONFIG_BIGNUM
#cgo CFLAGS: -DEMSCRIPTEN
#cgo CFLAGS: -I./deps/include
#cgo LDFLAGS: -lm
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/deps/libs/darwin_amd64 -lquickjs
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/deps/libs/linux_amd64 -lquickjs
*/
import "C"
