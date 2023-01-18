package conf

// TBD ...
const (
	defaultMaxRecvAndSendByteSize = 1024 * 1024 * 1024 * 2
)

// ServerConf config of Server
type ServerConf struct {
	MaxRecvByteSize int
	MaxSendByteSize int
	Port            int
}

// GatewayConf config of Gateway
type GatewayConf struct {
	MaxRecvByteSize int
	MaxSendByteSize int
	Port            int
	ServerMeta
}

// Conf .
type Conf struct {
	Server   *ServerConf
	Getaway  *GatewayConf
	Agent    *AgentConf
	LogLevel uint8
	LogPath  string
}

// AgentConf config of Agent
type AgentConf struct {
	HeartbeatPeriodSeconds int
	ServerMeta
}

// ServerMeta .
type ServerMeta struct {
	ServerAddr string
	Insecure   bool
	CertFile   string
	CertData   string
}
