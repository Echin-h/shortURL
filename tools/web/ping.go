package web

// 这边是将私网转到公网的一个工具
// 1. 通过配置文件读取私网的配置
// 2. 通过配置文件读取公网的配置
// 3. 通过配置文件读取需要转发的端口
// 4. 通过配置文件读取需要转发的域名
// 5. 通过配置文件读取需要转发的协议
// 6. 通过配置文件读取需要转发的路径
// 7. 通过配置文件读取需要转发的方法
// 8. 通过配置文件读取需要转发的参数
// 9. 通过配置文件读取需要转发的请求头
// 写出代码
// 代码呈现

//func Ping() {
//
//}
//func proxy(src net.Conn, dst net.Addr) {
//	defer src.Close()
//	dstConn, err := net.Dial("tcp", dst.String())
//	if err != nil {
//		fmt.Println("Error connecting to destination:", err)
//		return
//	}
//	defer dstConn.Close()
//
//	// Copy data from srcConn to dstConn and vice versa.
//	go func() {
//		_, err := io.Copy(dstConn, src)
//		if err != nil {
//			fmt.Println("Error copying to destination:", err)
//		}
//	}()
//
//	_, err = io.Copy(src, dstConn)
//	if err != nil {
//		fmt.Println("Error copying from source:", err)
//	}
//}

// 建立一个tcp的代理服务器
//func proxy(src net.Conn, dst net.Addr) {
//	defer src.Close()
//	dstConn, err := net.Dial("tcp", dst.String())
//	if err != nil {
//		fmt.Println("Error connecting to destination:", err)
//		return
//	}
//	defer dstConn.Close()
//
//	// Copy data from srcConn to dstConn and vice versa.
//	go func() {
//		_, err := io.Copy(dstConn, src)
//		if err != nil {
//			fmt.Println("Error copying to destination:", err)
//		}
//	}()
//
//	_, err = io.Copy(src, dstConn)
//	if err != nil {
//		fmt.Println("Error copying from source:", err)
//	}
//}
//
//func main() {
//	// Listen for incoming connections on a local port.
//	listener, err := net.Listen("tcp", ":8080") // Change the port as needed.
//	if err != nil {
//		fmt.Println("Error listening:", err)
//		return
//	}
//	defer listener.Close()
//
//	fmt.Println("Proxy server is running on port 8080...")
//
//	for {
//		srcConn, err := listener.Accept()
//		if err != nil {
//			fmt.Println("Error accepting connection:", err)
//			continue
//		}
//
//		// Get the destination address and port.
//		// This should be replaced with your actual logic to determine the target.
//		privateIP := net.ParseIP("192.168.1.10") // Replace with your private IP address.
//		destination := &net.TCPAddr{
//			IP:   privateIP,
//			Port: 8081, // Replace with your private server's port.
//		}
//
//		// Start a goroutine to handle the connection.
//		go proxy(srcConn, destination)
//	}
//}
