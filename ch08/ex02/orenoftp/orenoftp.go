package orenoftp

import (
	"log"
	"net"

	"github.com/k0kubun/pp"
)

// ディレクトリを変更する cd
// ディレクトリを列挙する ls
// ファイルの内容を送り出す get
// 接続を閉じる close

type Server struct {
	Addr string
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", "localhost"+s.Addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	pp.Print(c)
}
