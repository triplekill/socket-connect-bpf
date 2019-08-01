package dns

var dns4EventMap = make(map[uint32]HostInfo)

// AddIP4Entry adds an entry containing a hostname to the event map
func AddIP4Entry(ip4 uint32, pid uint32, host string) {
	dns4EventMap[ip4] = HostInfo{Pid: pid, Host: host}
}

// GetHostname gets a hostname for a cached IP (network byte order) / PID combination
func GetHostname(ip4 uint32, pid uint32) string {
	if hostInfo, ok := dns4EventMap[ip4]; ok {
		if hostInfo.Pid == pid {
			return hostInfo.Host
		}
		return ""
	}
	return ""
}

// HostInfo holds a PID/Hostname combination
type HostInfo struct {
	Pid  uint32
	Host string
}