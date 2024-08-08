package common

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/user"
)

type DeviceInfo struct {
	UserName  string
	Name      string
	IPAddress string
	Location  string
}

func GetDeviceInfo() DeviceInfo {
	return DeviceInfo{
		UserName:  getUserName(),
		Name:      getDeviceName(),
		IPAddress: getIPAddress(),
		Location:  getLocation(),
	}
}

func getDeviceName() (res string) {
	res, _ = os.Hostname()
	return
}

func getUserName() (res string) {
	u, _ := user.Current()

	return u.Username
}

func getIPAddress() (res string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error getting IP addresses:", err)
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				res = ipNet.IP.String()
			}
		}
	}

	return
}

func getLocation() (res string) {
	resp, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		fmt.Println("Error getting location:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	res = fmt.Sprintf("%s, %s", result["region"], result["country"])

	return
}
