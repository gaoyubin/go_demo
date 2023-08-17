package geerpc

import (
	"log"
	"reflect"
)

type methodType struct {
	method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
}
type service struct {
	name   string
	method map[string]*methodType
	typ    reflect.Type
	val    reflect.Value
}

func (m *methodType) newArgv() reflect.Value {
	var argv reflect.Value
	if m.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(m.ArgType.Elem())
	} else {
		argv = reflect.New(m.ArgType).Elem()
	}
	return argv
}
func (m *methodType) newReplyv() reflect.Value {
	replyv := reflect.New(m.ReplyType.Elem())
	return replyv
}

func newService(rcvr interface{}) *service {
	s := new(service)
	s.typ = reflect.TypeOf(rcvr)
	s.val = reflect.ValueOf(rcvr)
	s.name = reflect.Indirect(s.val).Type().Name()
	s.registerMethods()
	return s
}

func (s *service) registerMethods() {
	s.method = make(map[string]*methodType)
	for i := 0; i < s.typ.NumMethod(); i++ {
		method := s.typ.Method(i)
		mType := method.Type
		argType, replyType := mType.In(1), mType.In(2)
		s.method[method.Name] = &methodType{
			method:    method,
			ArgType:   argType,
			ReplyType: replyType,
		}
	}
	log.Println("show method", s.method)
}

func (s *service) call(m *methodType, argv, replyv reflect.Value) error {
	returnVals := m.method.Func.Call([]reflect.Value{s.val, argv, replyv})
	if errInter := returnVals[0].Interface(); errInter != nil {
		return errInter.(error)
	}
	return nil
}
