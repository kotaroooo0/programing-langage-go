package orenoftp

import (
	"net"

	"golang.org/x/exp/errors/fmt"
)

type FtpConn struct {
	Conn    net.Conn
	rootDir string
	workDir string
}

func (c FtpConn) User() {
	fmt.Fprint(c.Conn, "230 User kotaroooo0 logged in, proceed.\n")
}
