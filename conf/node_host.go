package conf

var host string

// GetNodeHost 获取节点地址
func GetNodeHost() string {
	if len(host) == 0 {
		host = Get().Node.Host
	}

	return host
}
