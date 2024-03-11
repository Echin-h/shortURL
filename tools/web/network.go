package web

//
////1. 监听一个地址对应的 TCP 请求 CreateTCPListener
////2. 连接一个 TCP 地址 CreateTCPConn
////3. 将一个 TCP-A 连接的数据写入另一个 TCP-B 连接，将 TCP-B 连接返回的数据写入 TCP-A 的连接中 Join2Conn （别看这短短 10 几行代码，这就是核心了）
//import (
//	"io"
//	"log"
//	"net"
//)
//
//const (
//	KeepAlive     = "keepAlive"
//	NewConnection = "newConnection"
//)
//
//// CreateTCPListener 创建一个tcp监听
//func CreateTCPListener(addr string) (*net.TCPListener, error) {
//	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
//	if err != nil {
//		return nil, err
//	}
//	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
//	if err != nil {
//		return nil, err
//	}
//	return tcpListener, nil
//}
//
//// 连接一个TCP地址
//func CreateTCPConn(addr string) (*net.TCPConn, error) {
//	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
//	if err != nil {
//		return nil, err
//	}
//	tcp, err := net.DialTCP("tcp", nil, tcpAddr)
//	if err != nil {
//		return nil, err
//	}
//	return tcp, nil
//}
//
//// 将一个 TCP-A 连接的数据写入另一个 TCP-B 连接，将 TCP-B 连接返回的数据写入 TCP-A 的连接中 Join2Conn
//func Join2Conn(remote, local *net.TCPConn) {
//	go joinConn(remote, local)
//	go joinConn(local, remote)
//}
//
//func joinConn(remote, local *net.TCPConn) {
//	defer local.Close()
//	defer remote.Close()
//	_, err := io.Copy(remote, local)
//	if err != nil {
//		log.Println("copy failed: " + err.Error())
//		return
//	}
//}
