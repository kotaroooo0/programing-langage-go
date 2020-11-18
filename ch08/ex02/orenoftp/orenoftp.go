package orenoftp

import (
	"bufio"
	"log"
	"net"
	"strings"

	"golang.org/x/exp/errors/fmt"
)

// ディレクトリを変更する cd
// ディレクトリを列挙する ls
// ファイルの内容を送り出す get
// 接続を閉じる close

type Server struct {
	Addr string
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	return s.Serve(listener)
}

func (s *Server) Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	log.Println("start handling connection.")

	fmt.Fprint(c, "welcome oreno ftp server!\n")

	s := bufio.NewScanner(c)
	for s.Scan() {
		input := strings.Fields(s.Text())
		log.Println(input)
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "CWD": // cd
			c.cwd(args)
		case "LIST": // ls
			c.list(args)
		default:
			fmt.Fprint(c, "invalid command\n")
		}
	}
	if s.Err() != nil {
		log.Print(s.Err())
	}
	fmt.Println("gasghakjghasklghalsgh")

}
