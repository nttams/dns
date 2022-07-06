package network_controller

import (
	"net"
	"bufio"
)

type UdpCtrl struct {
	conn net.UDPConn
	buffer []byte
}

func NewUdpCtrl() (result UdpCtrl) {
	result.buffer = make([]byte, 1024)
	return
}

// for server
func (ctrl *UdpCtrl) Listen(address string) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		panic("cannot bind")
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic("cannot listen")
	}
	ctrl.conn = *conn
}

func (ctrl *UdpCtrl) Read() ([]byte, *net.UDPAddr) {
	count, remoteAddr, _ := ctrl.conn.ReadFromUDP(ctrl.buffer)
	return ctrl.buffer[:count], remoteAddr
}

func (ctrl *UdpCtrl) Write(data []byte, remoteAddr *net.UDPAddr) {
	ctrl.conn.WriteToUDP(data, remoteAddr)
}

// for client
func (ctrl *UdpCtrl) Send(data []byte, addr string) {
	remoteHost, _ := net.ResolveUDPAddr("udp", addr)
	conn, _ := net.DialUDP("udp", &net.UDPAddr {}, remoteHost)
	ctrl.conn = *conn
	ctrl.conn.Write(data)
}

func (ctrl *UdpCtrl) Receive() []byte{
	count, _ := bufio.NewReader(&ctrl.conn).Read(ctrl.buffer)
	return ctrl.buffer[:count]
}