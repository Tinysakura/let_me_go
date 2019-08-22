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