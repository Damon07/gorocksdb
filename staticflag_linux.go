// +build static

package gorocksdb

// #cgo LDFLAGS: -l:librocksdb.a -lstdc++ -lm -lz -bz2 -lsnappy
import "C"
