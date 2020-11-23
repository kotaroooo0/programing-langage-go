package orenoftp

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"

	"golang.org/x/exp/errors/fmt"
)

type FtpConn struct {
	Conn     net.Conn
	RootDir  string
	WorkDir  string
	DataPort string
}

func NewFtpConn(c net.Conn, rootDir, workDir string) *FtpConn {
	return &FtpConn{
		Conn:    c,
		RootDir: rootDir,
		WorkDir: workDir,
	}
}

func (c *FtpConn) Welcome() {
	fmt.Fprint(c.Conn, "220 welcome oreno ftp server!\n")
}

func (c *FtpConn) User() {
	fmt.Fprint(c.Conn, "230 User kotaroooo0 logged in, proceed.\n")
}

func (c *FtpConn) Cwd(args []string) {
	if len(args) != 1 {
		log.Fatal("hoge")
		fmt.Fprint(c.Conn, "550 BAD.\n")
		return
	}

	workDir := filepath.Join(c.WorkDir, args[0])
	absPath := filepath.Join(c.RootDir, workDir)
	log.Println(workDir)
	log.Println(absPath)
	_, err := os.Stat(absPath)
	if err != nil {
		log.Print(err)
		fmt.Fprint(c.Conn, "550 BAD.\n")
		return
	}
	c.WorkDir = workDir
	fmt.Fprint(c.Conn, "200 OK.\n")
}

func (c *FtpConn) Quit() {
	fmt.Fprint(c.Conn, "221 BYE.\n")
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
		log.Print(err)
		return
	}
	fmt.Fprint(c.Conn, "150 File status okay; about to open data connection.\n")

	for _, file := range files {
		_, err := fmt.Fprint(c.Conn, file.Name(), "\n")
		if err != nil {
			log.Print(err)
		}
	}
	_, err = fmt.Fprintf(c.Conn, "\n")
	if err != nil {
		log.Print(err)
	}

	fmt.Fprint(c.Conn, "226 Closing data connection. Requested file action successful.\n")
}

func (c *FtpConn) Pwd() {
	path := filepath.Join(c.RootDir, c.WorkDir)
	fmt.Fprint(c.Conn, "257 "+path+" is your current location\n")
}

func (c *FtpConn) Port(args []string) {
	if len(args) != 1 {
		log.Fatal("hoge")
		fmt.Fprint(c.Conn, "550 BAD.\n")
		return
	}

	var h1, h2, h3, h4, p1, p2 int
	_, err := fmt.Sscanf(args[0], "%d,%d,%d,%d,%d,%d",
		&h1, &h2, &h3, &h4, &p1, &p2)
	if err != nil {
		log.Fatal("hoge")
		fmt.Fprint(c.Conn, "550 BAD.\n")
		return
	}

	port := p1<<8 + p2
	c.DataPort = fmt.Sprintf("%d.%d.%d.%d:%d", h1, h2, h3, h4, port)

	fmt.Fprint(c.Conn, "200 OK.\n")
}

func (c *FtpConn) Type(args []string) {

}

func (c *FtpConn) Retr(args []string) {
	if len(args) != 1 {
		log.Fatal("hoge")
		fmt.Fprint(c.Conn, "550 BAD.\n")
		return
	}

	path := filepath.Join(c.RootDir, c.WorkDir, args[0])
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("hoge")

	}
	fmt.Fprint(c.Conn, "150 File status okay; about to open data connection.\n")

	dataConn, err := net.Dial("tcp", c.DataPort)
	if err != nil {
		fmt.Fprint(c.Conn, "425 BAD.\n")
	}
	defer dataConn.Close()

	_, err = io.Copy(dataConn, file)
	if err != nil {
		fmt.Fprint(c.Conn, "426 BAD.\n")
		return
	}
	io.WriteString(dataConn, "\n")
	fmt.Fprint(c.Conn, "226 GOGO.\n")
}
