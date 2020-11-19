package orenoftp

import (
	"bufio"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
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

	fmt.Fprint(c, "220 welcome oreno ftp server!\n")

	s := bufio.NewScanner(c)
	fmt.Println(s.Text())
	fmt.Println("hoge")
	for s.Scan() {
		input := strings.Fields(s.Text())
		log.Println(input)
		if len(input) == 0 {
			continue
		}
		command, args := input[0], input[1:]
		log.Printf("<< %s %v", command, args)

		switch command {
		case "USER":
			fmt.Fprint(c, "230 User kotaroooo0 logged in, proceed.\n")

		case "CWD": // cd
			// cwd(args)
			fmt.Fprint(c, "230 "+strings.Join(args, " "))

		case "LIST": // ls
			// list(args)
			target := filepath.Join("public", "/")
			files, err := ioutil.ReadDir(target)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Fprint(c, "150 gasgkhaskgali\n")

			for _, file := range files {
				_, err := fmt.Fprint(c, file.Name(), "\n")
				if err != nil {
					log.Print(err)
				}
			}
			fmt.Fprint(c, "230 "+strings.Join(args, " "))
		case "QUIT":
			// quit()
			fmt.Fprint(c, "221 BYE.\n")
		default:
			fmt.Fprint(c, "invalid command\n")
		}
	}

	// if s.Err() != nil {
	// 	log.Print(s.Err())
	// }
}
