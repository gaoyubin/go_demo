package xclient

import (
	"geerpc"
	"log"
)

// import "geerpc"

type XClient struct {
	d      *GeeRegistryDiscovery
	mode   SelectMode
	client *geerpc.Client
}

func NewXClient(mode SelectMode, registry string) (*XClient, error) {
	d, err := NewGeeRegistryDiscovery(registry)
	if err != nil {
		log.Println("newgeeregistry err", err)
		return nil, err
	}
	rpcAddr, err := d.Get(mode)
	client, err := geerpc.NewClientWithAddr(rpcAddr)
	if err != nil {
		log.Println("new client err", err)
		return nil, err
	}
	log.Println("get addr", registry, rpcAddr)
	// client, err := NewClientWithAddr()
	xclient := &XClient{
		d:      d,
		mode:   mode,
		client: client,
	}
	return xclient, nil
}
func (xc *XClient) Call(serviceMethod string, args, reply interface{}) error {

	return xc.client.Call(serviceMethod, args, reply)
}
