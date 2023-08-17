package geerpc

import (
	"encoding/json"
	"errors"
	"geerpc/codec"
	"io"
	"log"
	"reflect"
	"strings"
	"sync"
)

type Server struct {
	serviceMap sync.Map
}

type request struct {
	h            *codec.Header
	argv, replyv reflect.Value
	mtype        *methodType
	svr          *service
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: option err", err)
		return
	}
	if opt.MagicNumber != MagicNumber {
		log.Println("magic err", opt)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Println("can not find codec type", opt)
		return
	}
	cc := f(conn)
	for {

		// var h codec.Header
		// if err := cc.ReadHeader(&h); err != nil {
		// 	log.Println("read header err", err)
		// 	if err == io.EOF {
		// 		log.Println("rpc read eof", err)
		// 	}
		// }

		// req := &request{h: h}
		// req.argv = reflect.New(reflect.TypeOf(""))
		// if err := cc.ReadBody(req.argv.Interface()); err != nil {
		// 	log.Println("rpc readbody err", err)
		// }

		// log.Println("show req", req, *(req.argv.Interface().(*string)))
		// reply := reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq))
		// cc.Write(&h, reply.Interface())
		// log.Println("send succ", reply)
		req, err := server.readRequest(cc)
		if err != nil {
			panic(err)
		}
		// fmt.Println(req)
		server.handleRequest(cc, req)
		// server.
	}
}

func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	var h codec.Header
	err := cc.ReadHeader(&h)
	if err != nil {
		var h codec.Header
		if err := cc.ReadHeader(&h); err != nil {
			log.Println("read header err", err)
			if err == io.EOF {
				log.Println("rpc read eof", err)
			}
		}
	}
	req := &request{
		h: &h,
	}
	req.svr, req.mtype, err = server.findService(h.ServiceMethod)
	if err != nil {
		return req, err
	}

	req.argv = req.mtype.newArgv()
	req.replyv = req.mtype.newReplyv()

	argvi := req.argv.Interface()
	// log.Println(*req.h, req.replyv, req.argv, *req.mtype, *req.svr)
	if req.argv.Kind() != reflect.Ptr {
		argvi = req.argv.Addr().Interface()
	}
	if err = cc.ReadBody(argvi); err != nil {
		log.Println("rpc read body err", err)
		return req, err
	}

	return req, nil
}
func (server *Server) findService(serviceMethod string) (svr *service, mtype *methodType, err error) {
	dot := strings.LastIndex(serviceMethod, ".")
	if dot < 0 {
		err = errors.New("rpc ill " + serviceMethod)
		return
	}
	serviceName, methodName := serviceMethod[:dot], serviceMethod[dot+1:]
	svci, ok := server.serviceMap.Load(serviceName)
	if !ok {
		err = errors.New("rpc not find service" + serviceName)
		return
	}
	svr = svci.(*service)
	mtype = svr.method[methodName]
	if mtype == nil {
		err = errors.New("rpc can not find method" + methodName)
	}
	return
}

func (server *Server) handleRequest(cc codec.Codec, req *request) {
	err := req.svr.call(req.mtype, req.argv, req.replyv)
	if err != nil {
		req.h.Error = err.Error()
		cc.Write(req.h, req.replyv.Interface())
		return
	}
	cc.Write(req.h, req.replyv.Interface())
}

func (server *Server) Register(rcvr interface{}) error {
	s := newService(rcvr)
	DefaultServer.serviceMap.LoadOrStore(s.name, s)
	return nil
}

func Register(rcvr interface{}) error {
	return DefaultServer.Register(rcvr)
}
