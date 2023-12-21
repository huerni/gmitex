package etcd

import (
	"context"
	"fmt"
	"github.com/duke-git/lancet/netutil"
	"go.etcd.io/etcd/client/v3"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	ttlKey     = "GMITEX_ETCD_REGISTRY_LEASE_TTL"
	defaultTTL = 60
)

type etcdRegistry struct {
	cli      *clientv3.Client
	leaseTTL int64
	meta     *registerMeta
}

type registerMeta struct {
	leaseID clientv3.LeaseID
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewEtcdRegistry(endpoints []string) (*etcdRegistry, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})

	if err != nil {
		log.Fatal(err)
	}

	return &etcdRegistry{
		cli:      cli,
		leaseTTL: GetTTL(),
	}, nil
}

//func (e *etcdRegistry) RegisterWithInfo(info *AppInstance) error {
//	leaseID, err := e.grantLease()
//
//	if err != nil {
//		return err
//	}
//	val, err := json.Marshal(info)
//	if err != nil {
//		return err
//	}
//	if err := e.register(info.ServerKey, string(val), leaseID); err != nil {
//		return err
//	}
//	meta := registerMeta{
//		leaseID: leaseID,
//	}
//	meta.ctx, meta.cancel = context.WithCancel(context.Background())
//	if err := e.keepalive(&meta); err != nil {
//		return err
//	}
//	e.meta = &meta
//	return nil
//}

func (e *etcdRegistry) RegisterWithKV(key string, val string) error {
	leaseID, err := e.grantLease()

	if err != nil {
		return err
	}
	if err := e.register(key, val, leaseID); err != nil {
		return err
	}
	meta := registerMeta{
		leaseID: leaseID,
	}
	meta.ctx, meta.cancel = context.WithCancel(context.Background())
	if err := e.keepalive(&meta); err != nil {
		return err
	}
	e.meta = &meta
	return nil
}

func (e *etcdRegistry) register(key string, val string, leaseID clientv3.LeaseID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err := e.cli.Put(ctx, key, val, clientv3.WithLease(leaseID))

	return err
}

func (e *etcdRegistry) grantLease() (clientv3.LeaseID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := e.cli.Grant(ctx, e.leaseTTL)
	if err != nil {
		return clientv3.NoLease, err
	}
	return resp.ID, nil
}

func (e *etcdRegistry) keepalive(meta *registerMeta) error {
	keepAlive, err := e.cli.KeepAlive(meta.ctx, meta.leaseID)
	if err != nil {
		return err
	}
	go func() {
		// eat keepAlive channel to keep related lease alive.
		fmt.Printf("start keepalive lease %x for etcd registry\n", meta.leaseID)
		for range keepAlive {
			select {
			case <-meta.ctx.Done():
				fmt.Printf("stop keepalive lease %x for etcd registry\n", meta.leaseID)
				return
			default:
			}
		}
	}()
	return nil
}

func GetTTL() int64 {
	var ttl int64 = defaultTTL
	if str, ok := os.LookupEnv(ttlKey); ok {
		if t, err := strconv.Atoi(str); err == nil {
			ttl = int64(t)
		}
	}
	return ttl
}

func GetInternalAddr(addr string) string {
	ip := addr[:len(addr)-5]
	if ip == "127.0.0.1" || ip == "localhost" || ip == "" {
		return netutil.GetInternalIp() + ":" + addr[len(addr)-4:]
	}

	return addr
}
