package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/agungfir98/mini-redis/handler"
	"github.com/agungfir98/mini-redis/proto"
)

type Server struct {
	ln         net.Listener
	ctx        context.Context
	wg         sync.WaitGroup
	clients    []Client
	mu         sync.Mutex
	shutdownCh chan int
}

type Client struct {
	conn net.Conn
	wg   *sync.WaitGroup
}

func New(ctx context.Context) (*Server, error) {
	ln, err := net.Listen("tcp", ":6380")
	if err != nil {
		return nil, err
	}

	log.Printf("TCP running on %v\n", ln.Addr())

	s := &Server{
		ln:         ln,
		clients:    []Client{},
		ctx:        ctx,
		shutdownCh: make(chan int),
	}
	go s.shutDownListener()

	return s, nil
}

func (s *Server) Run() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Println("Listener closed, stopping accept loop...")
				break
			}
			log.Println("failed to accept connection: ", err)
			continue
		}

		s.wg.Add(1)
		client := NewClient(conn, &s.wg)
		go client.handleConnection()
	}

	go func() {
		s.wg.Wait()
		close(s.shutdownCh)
	}()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	select {
	case <-s.shutdownCh:
		log.Printf("all client connection closed.")
	case <-shutdownCtx.Done():
		log.Printf("Shutwodn timeout reached. Forcing shutdown.")
	}

	fmt.Println("All connection closed. Server exited.")
}

func (s *Server) shutDownListener() {
	<-s.ctx.Done()
	log.Println("Shutdown signal received, closing listener...")
	s.ln.Close()
}

func NewClient(conn net.Conn, wg *sync.WaitGroup) *Client {
	c := &Client{wg: wg, conn: conn}
	return c
}

func (c *Client) handleConnection() {
	defer func() {
		c.wg.Done()
		c.conn.Close()
		log.Printf("client %s closed the connection\n", c.conn.RemoteAddr())
	}()

	log.Printf("new connection from %s\n", c.conn.RemoteAddr())

	for {
		resp := proto.NewResp(c.conn)
		writer := proto.NewWriter(c.conn)

		msg, err := resp.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			msg := proto.RespMessage{Typ: "error", Error: fmt.Sprintf("ERR %v\n", err)}
			writer.Write(msg)
			continue
		}

		fmt.Println("marshaled: ", strconv.Quote(string(msg.Marshal())))
		cmd := strings.ToUpper(msg.Array[0].String)
		args := msg.Array[1:]

		fmt.Println(args)

		if msg.Typ != "array" {
			msg := proto.RespMessage{Typ: "error", Error: "Type error, expected array"}
			writer.Write(msg)
			continue
		}

		handler, ok := handler.Message[cmd]
		if !ok {
			msg := proto.RespMessage{Typ: "error", Error: fmt.Sprintf("ERR unknown command '%v', with args beginning with:", cmd)}
			writer.Write(msg)
			continue
		}

		res := handler(args)
		writer.Write(res)
	}

}
