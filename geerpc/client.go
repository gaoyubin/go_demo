package geerpc

import (
	"encoding/json"
	"errors"
	"geerpc/codec"
	"log"
	"net"
	"sync"
)

type Call struct {
	Seq           uint64
	ServiceMethod string
	Args          interface{}
	Reply         interface{}
	Error         error
	Done          chan *Call
}
type Client struct {
	cc      codec.Codec
	pending map[uint64]*Call
	seq     uint64
	mu      sync.Mutex
	sending sync.Mutex
}

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

func (client *Client) receive() {
	for {
		var h codec.Header
		client.cc.ReadHeader(&h)
		call := client.pending[h.Seq]
		switch {
		case call == nil:
			panic("no call")
		default:
			err := client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body err")
			}
			call.Done <- call
		}
	}
}
func NewClient(conn net.Conn) (*Client, error) {
	f := codec.NewCodecFuncMap[codec.GobType]
	if f == nil {
		panic("invalid codec type")
	}
	if err := json.NewEncoder(conn).Encode(DefaultOption); err != nil {
		panic("send option fail")
	}
	client := &Client{
		cc:      f(conn),
		pending: make(map[uint64]*Call),
	}
	go client.receive()
	return client, nil
}

func (client *Client) Call(serviceMethod string, args, reply interface{}) error {
	done := make(chan *Call)
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}

	{
		client.mu.Lock()
		defer client.mu.Unlock()
		call.Seq = client.seq
		client.pending[call.Seq] = call
		client.seq++
	}

	{
		client.sending.Lock()
		var h codec.Header
		h.ServiceMethod = call.ServiceMethod
		h.Seq = call.Seq
		h.Error = ""
		if err := client.cc.Write(&h, call.Args); err != nil {
			log.Println("err", err, h)
		}
		client.sending.Unlock()
	}
	result_call := <-done
	return result_call.Error
}
