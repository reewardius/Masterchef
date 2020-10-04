package cluster

import (
	"net"
	"time"
)

func MeetChef(addr string) bool {
	_, err := net.DialTimeout("tcp", addr, 10*time.Second)
	return err == nil
}
