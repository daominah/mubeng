package proxymanager

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"ktbs.dev/mubeng/pkg/mubeng"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ProxyManager defines the proxy list and current proxy position
type ProxyManager struct {
	Proxies      []string
	CurrentIndex int
}

// New initialize ProxyManager
func New(filename string) (*ProxyManager, error) {
	keys := make(map[string]bool)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	manager := &ProxyManager{}
	manager.CurrentIndex = -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxy := scanner.Text()
		if _, value := keys[proxy]; !value {
			if _, err = mubeng.Transport(proxy); err == nil {
				keys[proxy] = true
				manager.Proxies = append(manager.Proxies, proxy)
			}
		}
	}

	if len(manager.Proxies) < 1 {
		return manager, fmt.Errorf("open %s: has no valid proxy URLs", filename)
	}

	return manager, scanner.Err()
}

// GetProxy returns the existing proxy from the list with the selected offset
func (p *ProxyManager) GetProxy(offset int) (string, error) {
	var err error

	length := len(p.Proxies) - 1
	if offset > length || offset < 0 {
		err = fmt.Errorf("Out of range proxy offset '%d' with length '%d'", offset, length)
		offset = 0
	}

	return p.Proxies[offset], err
}

// NextProxy will navigate the next proxy to use
func (p *ProxyManager) NextProxy() string {
	p.CurrentIndex++
	if p.CurrentIndex > len(p.Proxies)-1 {
		p.CurrentIndex = 0
	}

	proxy, _ := p.GetProxy(p.CurrentIndex)
	return proxy
}

// RandomProxy will choose a proxy randomly from the list
func (p *ProxyManager) RandomProxy() string {
	proxy, _ := p.GetProxy(rand.Intn(len(p.Proxies) - 1))
	return proxy
}
