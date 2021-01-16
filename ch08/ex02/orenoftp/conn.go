package orenoftp

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func (c *FtpConn) respond(s string) {
	fmt.Fprint(c.Conn, s, "\n")
}

func (c *FtpConn) Welcome() {
	c.respond("220 Service ready for new user.")
}

func (c *FtpConn) User() {
	c.respond("331 ok, set password.")
}

func (c *FtpConn) Pass() {
	c.respond("230 User logged in, proceed.")
}

func (c *FtpConn) Cwd(args []string) {
	if len(args) != 1 {
		c.respond("500 bad request.")
		return
	}
	workDir := filepath.Join(c.WorkDir, args[0])
	absPath := filepath.Join(c.RootDir, workDir)

	_, err := os.Stat(absPath)
	if err != nil {
		log.Println(err)
		c.respond("500 bad request.")
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
		c.respond("500 bad request.")
		return
	}
	fmt.Fprint(c.Conn, "150 File status okay; about to open data connection.\n")

	dataConn, err := c.dataConn()
	if err != nil {
		log.Println(err)
		c.respond("500 bad request.")
		return
	}
	defer dataConn.Close()

	for _, file := range files {
		_, err := fmt.Fprint(dataConn, file.Name(), "\n")
		if err != nil {
			log.Println(err)
			c.respond("500 bad request.")
			return
		}
	}
	_, err = fmt.Fprintf(dataConn, "\n")
	if err != nil {
		log.Println(err)
		c.respond("500 bad request.")
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
		c.respond("500 bad request.")
		return
	}

	var h1, h2, h3, h4, p1, p2 int
	_, err := fmt.Sscanf(args[0], "%d,%d,%d,%d,%d,%d", &h1, &h2, &h3, &h4, &p1, &p2)
	if err != nil {
		c.respond("500 bad request.")
		return
	}

	port := p1<<8 + p2
	c.DataPort = fmt.Sprintf("%d.%d.%d.%d:%d", h1, h2, h3, h4, port)
	c.respond("200 Command okay.")
}

func (c *FtpConn) Lprt(args []string) {
	if len(args) != 1 {
		c.respond("500 bad request.")
		return
	}

	parts := strings.Split(args[0], ",")
	addressLength, _ := strconv.Atoi(parts[1])
	portLength, _ := strconv.Atoi(parts[2+addressLength])
	portAddress := parts[3+addressLength : 3+addressLength+portLength]

	// Convert string[] to byte[]
	portBytes := make([]byte, portLength)
	for i := range portAddress {
		p, _ := strconv.Atoi(portAddress[i])
		portBytes[i] = byte(p)
	}

	// convert the bytes to an int
	port := int(binary.BigEndian.Uint16(portBytes))
	c.DataPort = fmt.Sprintf("[0:0:0:0:0:0:0:1]:%d", port)
	c.respond("200 Command okay.")
}

func (c *FtpConn) Type() {
	c.respond("215 UNIX system type.")
}

func (c *FtpConn) Retr(args []string) {
	if len(args) != 1 {
		c.respond("500 bad request.")
		return
	}

	path := filepath.Join(c.RootDir, c.WorkDir, args[0])
	file, _ := os.Open(path)
	fmt.Fprint(c.Conn, "150 File status okay; about to open data connection.\n")
	dataConn, err := c.dataConn()
	if err != nil {
		log.Println(err)
		c.respond("500 bad request.")
		return
	}
	defer dataConn.Close()

	io.Copy(dataConn, file)
	c.respond("200 Command okay.")
}

func (c *FtpConn) Stor(args []string) {
	if len(args) != 1 {
		c.respond("500 bad request.")
		return
	}
	dataConn, err := c.dataConn()
	defer dataConn.Close()
	if err != nil {
		log.Println(err)
		c.respond("500 bad request.")
		return
	}
	data, err := ioutil.ReadAll(dataConn)
	log.Println(data)
	if err != nil {
		log.Println(err)
		c.respond("500 bad request.")
		return
	}
	path := filepath.Join(c.RootDir, c.WorkDir, args[0])
	if err := ioutil.WriteFile(path, data, os.ModePerm); err != nil {
		log.Println(err)
		c.respond("500 bad request.")
		return
	}
	c.respond("221 Service closing control connection.")
}

func (c *FtpConn) Quit() {
	fmt.Fprint(c.Conn, "221 BYE.\n")
}

func (c *FtpConn) dataConn() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.DataPort)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
