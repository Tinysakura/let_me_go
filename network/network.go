package network

type IP struct {
	host string
	port int32
}

func (ip *IP) SetHost(host string)  {
	ip.host = host
}

func (ip *IP) SetPort(port int32)  {
	ip.port = port
}

func (ip *IP) GetPort() int32 {
	return ip.port
}

func (ip *IP) GetHost() string {
	return ip.host
}