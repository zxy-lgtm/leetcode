# TCP(Transmission Control Protocol)/IP

## Before TCP

**网络控制协议**（Network Control Protocol，缩写`NCP`）。

![img](https://www.tutorialspoint.com/assets/questions/media/19089/network_control.jpg)

##  What is TCP?

```tex
传输控制协议 (TCP) 是 Internet 协议套件的主要协议之一。它起源于最初的网络实现，它补充了 Internet 协议 (IP)。因此，整个套件通常称为 TCP/IP。 TCP 在运行于通过 IP 网络通信的主机上的应用程序之间提供可靠、有序且经过错误检查的八位字节流传输。万维网、电子邮件、远程管理和文件传输等主要互联网应用程序依赖于 TCP，它是 TCP/IP 套件传输层的一部分。
```

## How does TCP work？

![img](https://upload.wikimedia.org/wikipedia/commons/thumb/3/3b/UDP_encapsulation.svg/350px-UDP_encapsulation.svg.png)

![img](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/IP_stack_connections.svg/350px-IP_stack_connections.svg.png)

`osi`七层模型

![image-20220326102801779](C:\Users\zhangxinyu\AppData\Roaming\Typora\typora-user-images\image-20220326102801779.png)

三次握手

![建立 TCP 连接（三次握手）](https://www.ionos.com/digitalguide/fileadmin/DigitalGuide/Schaubilder/EN-tcp.png)

TCP连接终止

![TCP 连接终止（TCP 拆除）](https://www.ionos.com/digitalguide/fileadmin/DigitalGuide/Schaubilder/EN-tcp-verbindungsabbau.png)

`tcp`头部

![TCP 标头：结构](https://www.ionos.com/digitalguide/fileadmin/DigitalGuide/Schaubilder/EN-tcp-header.jpg)

总结

- TCP 是**面向连接的**，在**三次握手**之后，可以在两个端点之间进行双向通信。
- TCP是**可靠**的，因为该协议确保所有数据都被完全传输并且可以由接收器以正确的顺序组装。
- **TCP 允许在最大为1,500 字节**（包括标头）的单个段中发送数据。
- TCP 位于 `OSI` 模型的**传输层**（第 4 层）。
- TCP 通常与**Internet 协议**(`IP`) 结合使用，通常称为 TCP/`IP` 协议栈。
- **TCP 标头**的默认大小为 20 字节。最多可以添加 40 个字节的附加选项。