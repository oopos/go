// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Export guts for testing.

package runtime

var Fadd64 = fadd64
var Fsub64 = fsub64
var Fmul64 = fmul64
var Fdiv64 = fdiv64
var F64to32 = f64to32
var F32to64 = f32to64
var Fcmp64 = fcmp64
var Fintto64 = fintto64
var F64toint = f64toint

func entersyscall()
func exitsyscall()
func golockedOSThread() bool
func stackguard() (sp, limit uintptr)

var Entersyscall = entersyscall
var Exitsyscall = exitsyscall
var LockedOSThread = golockedOSThread
var Stackguard = stackguard

type LFNode struct {
	Next    *LFNode
	Pushcnt uintptr
}

func lfstackpush(head *uint64, node *LFNode)
func lfstackpop2(head *uint64) *LFNode

var LFStackPush = lfstackpush
var LFStackPop = lfstackpop2
