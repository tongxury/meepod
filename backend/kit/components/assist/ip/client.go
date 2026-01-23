package ip

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

func NewClient(dbFile string) (*Client, error) {
	db, err := geoip2.Open(dbFile)
	if err != nil {
		return nil, err
	}

	if db.Metadata().IPVersion == 0 {
		return nil, fmt.Errorf("db metadata not found: %v", db.Metadata())
	}

	return &Client{
		db: db,
	}, nil

}

type Client struct {
	db *geoip2.Reader
}

func (c *Client) City(ip string) (*City, error) {

	netIp := net.ParseIP(ip)
	if netIp == nil {
		return nil, fmt.Errorf("in valid ip: %s", ip)
	}

	city, err := c.db.City(netIp)
	if err != nil {
		return nil, err
	}

	if city == nil {
		return nil, fmt.Errorf("ip city not found: %s", ip)
	}

	return &City{}, nil
}
