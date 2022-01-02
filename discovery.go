package yeelight

import (
	"bufio"
	"net"
	"net/http"
	"strings"
	"time"
)

type DiscoveryResultItem struct {
	ID              string
	Name            string
	Location        string
	Model           string
	FirmwareVersion string
	Support         []string
	Power           bool
}

func Discovery(timeout time.Duration) ([]DiscoveryResultItem, error) {
	const discoverMSG = "M-SEARCH * HTTP/1.1\r\n HOST:239.255.255.250:1982\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n"
	const ssdpAddr = "239.255.255.250:1982"

	ssdp, _ := net.ResolveUDPAddr("udp4", ssdpAddr)
	con, _ := net.ListenPacket("udp4", ":0")
	socket := con.(*net.UDPConn)

	if _, err := socket.WriteToUDP([]byte(discoverMSG), ssdp); err != nil {
		return nil, err
	}

	if err := socket.SetReadDeadline(time.Now().Add(timeout)); err != nil {
		return nil, err
	}

	items := make([]DiscoveryResultItem, 0)
	unique := make(map[string]bool)
	for {
		rsBuf := make([]byte, 1024)

		size, _, err := socket.ReadFromUDP(rsBuf)
		if err, ok := err.(net.Error); ok && err.Timeout() {
			break
		}
		if err != nil {
			panic(err)
		}

		item := readReadDiscoveryPayload(string(rsBuf[0:size]))
		if unique[item.ID] {
			continue
		}

		items = append(items, item)

		unique[item.ID] = true
	}

	return items, nil
}

func readReadDiscoveryPayload(in string) DiscoveryResultItem {
	const crlf = "\r\n"

	if strings.HasSuffix(in, crlf) {
		in = in + crlf
	}

	resp, err := http.ReadResponse(bufio.NewReader(strings.NewReader(in)), nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return DiscoveryResultItem{
		ID:              resp.Header.Get("ID"),
		Name:            resp.Header.Get("Name"),
		Location:        strings.TrimPrefix(resp.Header.Get("LOCATION"), "yeelight://"),
		Model:           resp.Header.Get("Model"),
		FirmwareVersion: resp.Header.Get("Fw_ver"),
		Support:         strings.Split(resp.Header.Get("Support"), " "),
		Power:           resp.Header.Get("Fw_ver") == "on",
	}
}
