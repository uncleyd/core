package rpc

import (
	"context"
	"fmt"
	"github.com/edwingeng/doublejump"
	"github.com/smallnest/rpcx/client"
	"hash/fnv"
	"sort"
)

// consistentHashSelector selects based on JumpConsistentHash.
type consistentHashSelector struct {
	h       *doublejump.Hash
	servers []string
}

func newConsistentHashSelector(servers map[string]string) client.Selector {
	h := doublejump.NewHash()
	ss := make([]string, 0, len(servers))
	for k := range servers {
		ss = append(ss, k)
		h.Add(k)
	}

	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
	return &consistentHashSelector{servers: ss, h: h}
}

func (s consistentHashSelector) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	ss := s.servers
	if len(ss) == 0 {
		return ""
	}

	key := genKey(servicePath, serviceMethod, args.(Req).Data["ip"], args.(Req).Data["guid"])
	selected, _ := s.h.Get(key).(string)
	return selected
}

func (s *consistentHashSelector) UpdateServer(servers map[string]string) {
	if s.h == nil {
		h := doublejump.NewHash()
		ss := make([]string, 0, len(servers))
		for k := range servers {
			ss = append(ss, k)
			h.Add(k)
		}

		sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
		s.h = h
		s.servers = ss
	}

	ss := make([]string, 0, len(servers))
	for k := range servers {
		s.h.Add(k)
		ss = append(ss, k)
	}

	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })

	for _, k := range s.servers {
		if servers[k] == "" { // remove
			s.h.Remove(k)
		}
	}
	s.servers = ss
}

func genKey(options ...interface{}) uint64 {
	keyString := ""
	for _, opt := range options {
		keyString = keyString + "/" + toString(opt)
	}
	return HashString(keyString)
}

func toString(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}

// HashString get a hash value of a string
func HashString(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}
