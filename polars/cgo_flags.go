package polars

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo linux LDFLAGS: -L${SRCDIR}/bin -lpolars_go -ldl -lm -lpthread -Wl,-rpath=${SRCDIR}/bin
#cgo darwin LDFLAGS: -L${SRCDIR}/bin -lpolars_go -framework CoreFoundation -framework Security -Wl,-rpath,${SRCDIR}/bin
#cgo windows LDFLAGS: -L${SRCDIR}/bin -lpolars_go -lws2_32 -luserenv -ladvapi32 -lkernel32
*/
import "C"
