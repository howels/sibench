package main

import (
	"bytes"
	"fmt"
	"logger"
	"os/exec"
	"runtime"
	"syscall"
)

func Open(path string, mode int, perm uint32) (FileDescriptor, error) {
	fd, err := syscall.Open(path, mode|syscall.O_SYNC, perm)
	if err != nil {
		return -1, err
	}

	// macos kernel prevents this -- TODO: look into disabling protections

	// _, _, err = syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), syscall.F_NOCACHE, 1)
	// if err != nil {
	// 	return -1, nil
	// }

	return FileDescriptor(fd), nil
}

func Mount(source string, target string, fstype string, flags uintptr, data string) error {
	var out bytes.Buffer

	cmd := exec.Command("mount", "-t", fstype, "-o", data, source, target)
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		logger.Errorf("Mount: %s\n", out)
		return err
	}

	return nil
}

func NewRadosConnection(target string, protocol ProtocolConfig, worker WorkerConnectionConfig) (Connection, error) {
	return nil, fmt.Errorf("rados not implemented on %q", runtime.GOOS)
}

func NewRbdConnection(target string, protocol ProtocolConfig, worker WorkerConnectionConfig) (Connection, error) {
	return nil, fmt.Errorf("rbd not implemented on %q", runtime.GOOS)
}
