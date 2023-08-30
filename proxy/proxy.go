package proxy

import (
	"errors"
	"net/http"
	"net/url"
	"sync/atomic"
)

type ProxyFunc func(r *http.Request) (*url.URL, error)

func RoundRobinProxySwitcher(ProxyURLs ...string) (ProxyFunc, error) {

	if len(ProxyURLs) < 1 {
		return nil, errors.New("no proxy url")
	}

	urls := make([]*url.URL, len(ProxyURLs))

	for i, v := range ProxyURLs {
		u, err := url.Parse(v)
		if err != nil {
			return nil, err
		}
		urls[i] = u
	}

	return (&roundRobinSwitcher{urls: urls, index: 0}).GetProxy, nil

}

type roundRobinSwitcher struct {
	urls  []*url.URL
	index uint32
}

func (s *roundRobinSwitcher) GetProxy(pr *http.Request) (*url.URL, error) {
	index := atomic.AddUint32(&s.index, 1) - 1
	u := s.urls[index%uint32(len(s.urls))]
	return u, nil
}
