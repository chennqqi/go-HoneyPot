package tcp

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/chennqqi/go-HoneyPot/config"
	"github.com/chennqqi/go-HoneyPot/report"
)

// Server is the tcp server struct
type Server struct {
	Ports []string
	rpt   report.Reporter
}

// NewServer creates a new tcp server
func NewServer(cfg *config.Config) (*Server, error) {
	var rpt report.Reporter
	var err error
	switch cfg.Report {
	case "http":
		rpt, err = report.NewHttpReport(&cfg.Http)
	case "database":
		rpt, err = report.NewSqlReporter(&cfg.DB)
	default:
		err = errors.New(fmt.Sprintf("unspport %v", cfg.Report))
	}
	if err != nil {
		return nil, err
	}
	return &Server{cfg.TCP.Ports, rpt}, nil
}

// Start starts the tcp server
func (t *Server) Run() {
	var wg sync.WaitGroup
	wg.Add(len(t.Ports))

	for _, port := range t.Ports {
		go func(port string, wg *sync.WaitGroup, rpt report.Reporter) {
			fmt.Printf("Listening on tcp port: %v\n", port)
			listen, err := net.Listen("tcp", ":"+port)
			if err != nil {
				log.Println(err)
				wg.Done()
				return
			}
			for {
				conn, err := listen.Accept()
				if err != nil {
					logrus.Fatalf("[tcp.go] listen.Accept error: %v", err)

					// handle error
				}
				go handleConnection(conn, rpt)
			}
		}(port, &wg, t.rpt)
	}
	wg.Wait()
	logrus.Println("TCP Server Stopped")
}

func handleConnection(conn net.Conn, rpt report.Reporter) {
	fmt.Println("connection")
	data := make([]byte, 4096)
	n, err := conn.Read(data)
	if err != nil {
		logrus.Errorf("[tcp.go] Read connection data error:", err)
		conn.Close()
		return
	}
	defer conn.Close()

	logrus.Errorf("[tcp.go] Received data from %v, of length %v data is %v", conn.RemoteAddr(), n, data[:n])
	remHost, remPort, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		logrus.Errorf("[tcp.go] SplitHostPort error: %v", err)
		return
	}
	locHost, locPort, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		logrus.Errorf("[tcp.go] SFailed to split remote host and port: %v", err)
		return
	}
	var srcport, dstport int64
	fmt.Sscanf(remPort, "%d", &srcport)
	fmt.Sscanf(locPort, "%d", &dstport)
	err = rpt.Pub(&report.HoneypotRecord{
		Src:     remHost,
		Dst:     locHost,
		Srcport: srcport,
		Dstport: dstport,
		Payload: string(data[:n]),
		Raw:     data[:n],
		Time:    report.Time(time.Now()),
	})
	if err != nil {
		logrus.Errorf("[tcp.go] rpt.Pub error: %v", err)
	}
}
