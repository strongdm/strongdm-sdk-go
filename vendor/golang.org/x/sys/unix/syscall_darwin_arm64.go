// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64 && darwin
// +build arm64,darwin

package unix

import (
	"syscall"
)

func setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: sec, Nsec: nsec}
}

func setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: sec, Usec: int32(usec)}
}

// sysnb	gettimeofday(tp *Timeval) (sec int64, usec int32, err error)
func Gettimeofday(tv *Timeval) (err error) {
	// The tv passed to gettimeofday must be non-nil
	// but is otherwise unused. The answers come back
	// in the two registers.
	sec, usec, err := gettimeofday(tv)
	tv.Sec = sec
	tv.Usec = usec
	return err
}

func SetKevent(k *Kevent_t, fd, mode, flags int) {
	k.Ident = uint64(fd)
	k.Filter = int16(mode)
	k.Flags = uint16(flags)
}

func (iov *Iovec) SetLen(length int) {
	iov.Len = uint64(length)
}

func (msghdr *Msghdr) SetControllen(length int) {
	msghdr.Controllen = uint32(length)
}

func (cmsg *Cmsghdr) SetLen(length int) {
	cmsg.Len = uint32(length)
}

func Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno) // sic

// SYS___SYSCTL is used by syscall_bsd.go for all BSDs, but in modern versions
// of darwin/arm64 the syscall is called sysctl instead of __sysctl.
const SYS___SYSCTL = SYS_SYSCTL

//sys	Fstat(fd int, stat *Stat_t) (err error)
//sys	Fstatat(fd int, path string, stat *Stat_t, flags int) (err error)
//sys	Fstatfs(fd int, stat *Statfs_t) (err error)
//sys	getfsstat(buf unsafe.Pointer, size uintptr, flags int) (n int, err error) = SYS_GETFSSTAT
//sys	Lstat(path string, stat *Stat_t) (err error)
//sys	Stat(path string, stat *Stat_t) (err error)
//sys	Statfs(path string, stat *Statfs_t) (err error)

func Getdirentries(fd int, buf []byte, basep *uintptr) (n int, err error) {
	return 0, ENOSYS
}
