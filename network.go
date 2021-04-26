package gomyblockchain

import (
	"fmt"
	"net"
	"strconv"
)

func isFoundHost(target string, port int) bool {
	conn, err := net.Dial("tcp", target+":"+fmt.Sprint(port))
	if err != nil {
		// fmt.Println(err)
		return false
	}
	conn.Close()
	return true
}

// FindNeighbours 指定されたIP,PORTにノードがある場合，[]stringを返します
func FindNeighbours(myHost string, myPort int, startIPRange int, endIPRange int, startPort int, endPort int) []string {
	myipv4 := net.ParseIP(myHost).To4()
	address := myHost + fmt.Sprint(myPort)
	var neighbours []string
	for guessIP := startIPRange; guessIP < endIPRange; guessIP++ {
		for guessPort := startPort; guessPort < endPort; guessPort++ {
			guessHost := myipv4
			guessHost[3] = guessHost[3] + byte(guessIP)
			guessAddress := guessHost.String() + ":" + fmt.Sprint(guessPort)
			if isFoundHost(guessHost.String(), guessPort) && guessAddress != address && guessAddress != myipv4.String()+":"+strconv.Itoa(myPort) {
				neighbours = append(neighbours, guessAddress)
			}
		}
	}
	// PrintFindNeighbours(neighbours)
	return neighbours
}

// GetHost HostのプライベートipをIPv4で取ってきます.
func GetHost() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
