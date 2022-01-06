package api

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	numMaxClients uint = 10240
	writeTimeout       = time.Second * 2
	pingPeriod         = time.Second * 10
)

type MessageFormatter interface {
	FormatMsg(interface{}) string
}

type eventClient struct {
	conn      net.Conn
	formatter MessageFormatter
	wait      chan struct{}
	recv      chan string
	closeOnce sync.Once
}

func (ec *eventClient) send(msg string) {
	select {
	case <-ec.wait:
		// channel closed
		ec.closeOnce.Do(func() {
			close(ec.recv)
		})
	case ec.recv <- msg:
	default:
		<-ec.recv // Discard the oldest message when buffer is full.
		ec.send(msg)
	}
}

type eventsHandler struct {
	sync.RWMutex
	ws  map[string]*eventClient
	max uint
}

func NewEventsHandler() *eventsHandler {
	eh := &eventsHandler{
		ws:  make(map[string]*eventClient),
		max: numMaxClients,
	}

	return eh
}

func (eh *eventsHandler) Add(remoteAddr string, w http.ResponseWriter) error {
	hi, ok := w.(http.Hijacker)
	if !ok {
		return fmt.Errorf("%#v is not an http.Hijacker", w)
	}
	conn, _, err := hi.Hijack()
	if err != nil {
		return err
	}

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(time.Second * 30)
		tcpConn.SetLinger(3)
	}

	c := &eventClient{
		conn:      conn,
		formatter: EventsourceEvent{},
		wait:      make(chan struct{}),
		recv:      make(chan string, 1024),
	}

	go func(eh *eventsHandler, c *eventClient, remoteAddr string) {
		defer func() {
			c.conn.Close()
			eh.Delete(remoteAddr)
		}()

		for {
			select {
			case msg, ok := <-c.recv:
				if !ok {
					return
				}

				c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
				if _, err := c.conn.Write([]byte(msg)); err != nil {
					logrus.Errorf("write event message to client [%s] error: [%v]", remoteAddr, err)
					return
				}
			case <-time.After(pingPeriod):
				c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
				ping := fmt.Sprintf("event: ping\ndata: \"it is a ping\"\n\n")
				if _, err := c.conn.Write([]byte(ping)); err != nil {
					logrus.Errorf("Failed to write ping to %s: %v", remoteAddr, err)
					return
				}
			}
		}
	}(eh, c, remoteAddr)

	eh.Lock()
	eh.ws[remoteAddr] = c
	eh.Unlock()

	return nil
}

func (eh *eventsHandler) Wait(remoteAddr string) {
	eh.RLock()
	c, ok := eh.ws[remoteAddr]
	eh.RUnlock()
	if !ok {
		return
	}
	<-c.wait
}

func (eh *eventsHandler) Delete(remoteAddr string) {
	logrus.Debug("delete event listener ", remoteAddr)
	eh.Lock()
	defer eh.Unlock()
	c, ok := eh.ws[remoteAddr]
	if !ok {
		return
	}
	close(c.wait)
	delete(eh.ws, remoteAddr)
}

func (eh *eventsHandler) full() bool {
	return eh.Size() >= int(eh.max)
}

// broadcast message to all eventsource clients
func (eh *eventsHandler) SendMessage(msg map[string]interface{}) error {
	byts, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	msg["time"] = time.Now()

	eid := time.Now().UnixNano()
	str := fmt.Sprintf("event: message\nid: %d\ndata: %s\n\n", eid, string(byts))

	for _, c := range eh.clients() {
		c.send(str)
	}

	return nil
}

func (eh *eventsHandler) SendDockerMessage(msg string) error {
	eid := time.Now().UnixNano()
	str := fmt.Sprintf("event: docker\nid: %d\ndata: %s\n\n", eid, msg)

	for _, c := range eh.clients() {
		c.send(str)
	}

	return nil
}

func (eh *eventsHandler) clients() []*eventClient {
	eh.RLock()
	ret := make([]*eventClient, 0, len(eh.ws))
	for _, c := range eh.ws {
		ret = append(ret, c)
	}
	eh.RUnlock()
	return ret
}

func (eh *eventsHandler) Handle(e interface{}) error {
	//eid := time.Now().UnixNano()

	for _, c := range eh.clients() {
		str := c.formatter.FormatMsg("dd")
		c.send(str)
	}

	return nil
}

func (eh *eventsHandler) Size() int {
	eh.RLock()
	defer eh.RUnlock()
	return len(eh.ws)
}

type EventsourceEvent struct {
}

func (ee EventsourceEvent) FormatMsg(data interface{}) string {

	return ""
}
