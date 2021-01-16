package orenoftp

import (
	"encoding/binary"
	"fmt"
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

func (c *FtpConn) Lprt(args []string) {
	if len(args) != 1 {
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		return
	}

	parts := strings.Split(args[0], ",")

	addressFamily, _ := strconv.Atoi(parts[0])
	if addressFamily != 4 {
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		// return
	}

	addressLength, _ := strconv.Atoi(parts[1])
	if addressLength != 4 {
		fmt.Fprint(c.Conn, "500 BAD REQUEST.\n")
		// return
	}

	host := strings.Join(parts[2:2+addressLength], ".")

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

	log.Println(host)
	log.Println(port)
	// port := p1 << 8
	log.Println(len(args))
	log.Println(args)
	c.DataPort = "127.0.0.1:" + strconv.Itoa(port)
	fmt.Fprint(c.Conn, "200 OK.\n")
}

func (c *FtpConn) Type() {
	c.respond("215 UNIX system type.")
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
