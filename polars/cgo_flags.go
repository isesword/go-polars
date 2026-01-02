package polars

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo linux LDFLAGS: -L${POLARS_BIN_DIR} -lpolars_go -ldl -lm -lpthread -Wl,-rpath=${POLARS_BIN_DIR}
#cgo darwin LDFLAGS: -L${POLARS_BIN_DIR} -lpolars_go -framework CoreFoundation -framework Security -framework IOKit -Wl,-rpath,${POLARS_BIN_DIR}
#cgo windows LDFLAGS: -L${POLARS_BIN_DIR} -lpolars_go -lws2_32 -luserenv -ladvapi32 -lkernel32
*/
import "C"
