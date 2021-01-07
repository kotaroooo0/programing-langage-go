package orenoftp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
)

type FtpConn struct {
	Conn     net.Conn
	RootDir  string
	WorkDir  string
	DataPort string
}

func NewFtpConn(c net.Conn, rootDir, workDir string) FtpConn {
	return FtpConn{
		Conn:    c,
		RootDir: rootDir,
		WorkDir: workDir,
	}
}

func (c *FtpConn) Welcome() {
	fmt.Fprint(c.Conn, "220 welcome oreno ftp server!\n")
}

func (c *FtpConn) User() {
	fmt.Fprint(c.Conn, "331 ok, set password.\n")
}

func (c *FtpConn) Pass() {
	fmt.Fprint(c.Conn, "230 User kotaroooo0 logged in, proceed.\n")
}

func (c *FtpConn) Cwd(args []string) {
	if len(args) != 1 {
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		return
	}

	workDir := filepath.Join(c.WorkDir, args[0])
	absPath := filepath.Join(c.RootDir, workDir)
	_, err := os.Stat(absPath)
	if err != nil {
		log.Println(err)
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		return
	}
	c.WorkDir = workDir
	fmt.Fprint(c.Conn, "200 OK.\n")
}

func (c *FtpConn) List(args []string) {
	var target string
	if len(args) > 0 {
		target = filepath.Join(c.RootDir, c.WorkDir, args[0])
	} else {
		target = filepath.Join(c.RootDir, c.WorkDir)
	}

	files, err := ioutil.ReadDir(target)
	if err != nil {
		log.Println(err)
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		return
	}
	fmt.Fprint(c.Conn, "150 File status okay; about to open data connection.\n")

	dataConn, err := c.dataConn()
	if err != nil {
		log.Println(err)
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		return
	}
	defer dataConn.Close()

	for _, file := range files {
		_, err := fmt.Fprint(dataConn, file.Name(), "\n")
		if err != nil {
			log.Println(err)
			fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
			return
		}
	}
	_, err = fmt.Fprintf(dataConn, "\n")
	if err != nil {
		log.Println(err)
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		return
	}

	fmt.Fprint(c.Conn, "226 Closing data connection. Requested file action successful.\n")
}

func (c *FtpConn) Pwd() {
	path := filepath.Join(c.RootDir, c.WorkDir)
	fmt.Fprint(c.Conn, "257 \""+path+"\" is your current directory\n")
}

func (c *FtpConn) Port(args []string) {
	if len(args) != 1 {
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		return
	}

	var h1, h2, h3, h4, p1, p2 int
	_, err := fmt.Sscanf(args[0], "%d,%d,%d,%d,%d,%d", &h1, &h2, &h3, &h4, &p1, &p2)
	if err != nil {
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		return
	}

	port := p1<<8 + p2
	c.DataPort = fmt.Sprintf("%d.%d.%d.%d:%d", h1, h2, h3, h4, port)
	fmt.Fprint(c.Conn, "200 OK.\n")
}

func (c *FtpConn) Type() {
	fmt.Fprint(c.Conn, "215 UNIX system type.\n")
}

func (c *FtpConn) Quit() {
	fmt.Fprint(c.Conn, "221 BYE.\n")
}

func (c *FtpConn) dataConn() (net.Conn, error) {
	// fmt.Println(c.Da)
	conn, err := net.Dial("tcp", c.DataPort)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
