package balance

// Balancer 负载均衡
type Balancer interface {
	DoBalance([]*Instance) (*Instance, error)
}
