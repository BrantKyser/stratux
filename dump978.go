// Copyright (c) 2015 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

// Package dump978 wraps libdump978, a 978MHz UAT demodulator.

package dump978

/*
#cgo linux LDFLAGS: -L. -ldump978
#cgo darwin LDFLAGS: -L. -ldump978
#cgo windows CFLAGS: -IC:/WINDOWS/system32
#cgo windows LDFLAGS: -L. -lrtlsdr -LC:/WINDOWS/system32

#include <stdlib.h>
#include <stdint.h>
#include "dump978/dump978.h"

extern void dump978Cb(char updown, uint8_t *data, int len);
static inline CallBack GetGoCb() {
	return (CallBack)dump978Cb;
}
*/
import "C"
import "unsafe"

// Current version.
var PackageVersion = "v0.1"

// InChan is a buffered input channel for raw data.
var InChan = make(chan []byte, 100)

type UserCbT func(C.char, *C.uint8_t, C.int)

// Dump978Init must be the first function called in this package.
func Dump978Init() {
	C.Dump978Init((C.CallBack)(C.GetGoCb()))
}

// ProcessData passes buf (modulated data) to dump978 for demodulation.
func ProcessData(buf []byte) {
	C.process_data((*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)))
}

func ProcessDataFromChannel() {
	for {
		inData := <-InChan
		ProcessData(inData)
	}
}