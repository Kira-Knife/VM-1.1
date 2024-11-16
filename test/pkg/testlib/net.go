package testlib

import (
	"net"
	"time"

	. "github.com/onsi/gomega"
)

func ЖдатьОткрытияПорта(port string, tout int) (ret net.Conn) {
	addr := "127.0.0.1:" + port
	return ЖдатьОткрытияАдресПорта(addr, tout)
}
func ЖдатьОткрытияАдресПорта(addr string, tout int) (ret net.Conn) {
	ждать := time.Duration(tout) * time.Millisecond

	Eventually(func() bool {
		conn, err := net.Dial("tcp4", addr)
		if err != nil {
			return false
		}

		ret = conn

		return true
	}).
		Within(ждать).ProbeEvery(50*time.Millisecond).
		Should(BeTrue(), "ждатьПорт: "+addr)
	if ret == nil {
		panic("ЖдатьОткрытияПорта")
	}
	return
}
