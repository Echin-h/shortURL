package web

//
////连接服务端的控制通道
////等待服务端从控制通道中发来建立连接的消息
////收到建立连接的消息时，将本地服务和远端隧道建立连接（这里就要用到我们的工具方法了）
//
//import (
//	"bufio"
//	"io"
//	"log"
//	"net"
//)
////https://www.jianshu.com/p/ecda849a49bd
//var (
//	// 本地需要暴露的服务端口
//	localServerAddr = "127.0.0.1:8080"
//
//	remoteIP = "10.106.251.162"
//	// 远端的服务控制通道，用来传递控制信息，如出现新连接和心跳
//	remoteControlAddr = remoteIP + ":8009"
//	// 远端服务端口，用来建立隧道
//	remoteServerAddr = remoteIP + ":8008"
//)
//
//func clientInit() {
//	tcpConn, err := CreateTCPConn(remoteControlAddr)
//	if err != nil {
//		log.Println("连接远端控制通道失败", err)
//		return
//	}
//	log.Println("连接远端控制通道成功")
//
//	reader := bufio.NewReader(tcpConn)
//	for {
//		s, err := reader.ReadString('\n')
//		if err != nil || err == io.EOF {
//			break
//		}
//		if s == NewConnection+"\n" {
//			// 创建一个新连接
//			go connectLocalAndRemote()
//		}
//	}
//	log.Println("远端控制通道关闭", remoteControlAddr)
//}
//
//func connectLocalAndRemote() {
//	local := connectLocal()
//	remote := connectRemote()
//	if local == nil && remote == nil {
//		Join2Conn(remote, local)
//	} else {
//		if local != nil {
//			_ = local.Close()
//		}
//		if remote != nil {
//			_ = remote.Close()
//		}
//	}
//}
//
//func connectLocal() *net.TCPConn {
//	conn, err := CreateTCPConn(localServerAddr)
//	if err != nil {
//		log.Println("连接本地服务失败", err)
//		return nil
//	}
//	return conn
//}
//
//func connectRemote() *net.TCPConn {
//	conn, err := CreateTCPConn(remoteServerAddr)
//	if err != nil {
//		log.Println("连接远端服务失败", err)
//		return nil
//	}
//	return conn
//}
