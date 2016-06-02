package main

/*
Sample user defined functions for postgres
Should work with postgresql 8+

These are doing using version 0 calling conventions for postgres
due to version 1 requiring heavy use of C macros that cgo
doesn't support very well.

To see what type conversions to use, see:
https://www.postgresql.org/docs/9.5/static/xfunc-c.html#XFUNC-C-TYPE-TABLE


Any writes stdout/stderr will go to the postgresql log file.
In these examples, we use the stdlib log package that will
log to stderr by default.

Any function that has "//export NAME" will be callable by postgres only
after you add a corresponding CREATE FUNCTION call to install.sql.
*/

//#include "postgres.h"
//#include "fmgr.h"
//#include <string.h>
import "C"
import (
	"log"
	"sync/atomic"
)

// Functions are scoped to the db session
var counter int64

//Inc atomically increments a session local counter by a given delta
//export Inc
func Inc(delta C.int64) C.int64 {
	log.Printf("Inc(%v) called", delta)
	return C.int64(atomic.AddInt64(&counter, int64(delta)))
}

//AddOne adds one to the given arg and retuns it
//export AddOne
func AddOne(i C.int) C.int {
	log.Printf("AddOne(%v) called", i)
	return i + 1
}

func main() {
}
