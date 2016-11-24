package osinfo

import (
	"github.com/shirou/gopsutil/mem"
	"math"
	"net"
	"os"
	"time"
)

var StartTime time.Time
var IPAddress, Hostname string

// var TotalMem, FreeMem, UsedMem float64

var TotalMem, FreeMem, UsedMem uint64

func GetLocalIP() (string, string) {
	addrs, err := net.InterfaceAddrs()
	hostname, _ := os.Hostname()
	if err != nil {
		return "", ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), hostname
			}
		}
	}
	return "", ""
}

func GetMemInfo() (uint64, uint64, uint64) {
	v, _ := mem.VirtualMemory()
	return v.Total / uint64(math.Pow(10, 9)), v.Available / uint64(math.Pow(10, 9)), uint64(v.UsedPercent)
}

func CheckPort(proto string, addr string) string {
	var status string
	conn, err := net.Dial(proto, addr)
	if err != nil {
		status = "Unreachable"
	} else {
		status = "Reachable"
		conn.Close()
	}
	return status
}
