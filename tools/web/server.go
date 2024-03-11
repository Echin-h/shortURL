package web

//
//import (
//	"log"
//	"net"
//	"sync"
//	"time"
//)
//
//// 1.监听控制通道，接收客户端的连接请求
//// 2.监听访问端口，接收来自用户的 http 请求
//// 3.第二步接收到请求之后需要存放一下这个连接并同时发消息给客户端，告诉客户端有用户访问了，赶紧建立隧道进行通信
//// 4.监听隧道通道，接收来自客户端的连接请求，将客户端的连接与用户的连接建立起来（也是用工具方法）
//
//const (
//	controlAddr = "10.106.251.162:8009"
//	tunnelAddr  = "10.106.251.162:8008"
//	visitAddr   = "10.106.251.162:8007"
//)
//
//var (
//	clientConn         *net.TCPConn
//	connectionPool     map[string]*ConnMatch
//	connectionPoolLock *sync.RWMutex
//)
//
//type ConnMatch struct {
//	addTime time.Time
//	accept  *net.TCPConn
//}
//
//func serverInit() {
//
//}
//
//// 创建一个控制通道，用于传递控制信息
//func CreateControlListener() (*net.TCPListener, error) {
//	listener, err := CreateTCPListener(controlAddr)
//	if err != nil {
//		log.Println("创建控制通道失败", err)
//		return nil, err
//	}
//	log.Println("已监听")
//	for {
//		tcpConn, err := listener.AcceptTCP()
//		if err != nil {
//			log.Println("接收连接失败", err)
//			continue
//		}
//
//		log.Println("接收到新连接", tcpConn.RemoteAddr().String())
//		if clientConn != nil {
//			_ = tcpConn.Close()
//		} else {
//			clientConn = tcpConn
//			go KeepAlive()
//		}
//	}
//}
//
//// 和客户端保持一个心跳连接
//func keepAlive() {
//	go func() {
//		for {
//			if clientConn == nil {
//				return
//			}
//			_, err := clientConn.Write(([]byte)(KeepAlive + "\n"))
//			if err != nil {
//				log.Println("[已断开客户端连接]", clientConn.RemoteAddr())
//				clientConn = nil
//				return
//			}
//			time.Sleep(time.Second * 3)
//		}
//	}()
//}
//
//func acceptUserRequest() {
//	tcpListener, err := CreateTCPListener(visitAddr)
//	if err != nil {
//		log.Println("监听用户请求失败", err)
//		return
//	}
//	defer tcpListener.Close()
//	for {
//
//	}
//}
