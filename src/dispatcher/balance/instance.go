package balance

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (p *Instance) GetHost() string {
	return p.host
}

func (p *Instance) GetPort() int {
	return p.port
}
