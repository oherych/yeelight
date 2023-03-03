package yeelight

import (
	"bufio"
	"context"
	"net"
	"net/http"
	"strings"
	"sync"
)

// DiscoveryResultItem contain information about result of discovery
type DiscoveryResultItem struct {
	ID              string
	Name            string
	Location        string
	Model           string
	FirmwareVersion string
	Support         []string
	Power           bool
}

// Discovery scans the network for devices.
// This feature can`t stop itself. It MUST be stopped by context.
func Discovery(ctx context.Context) (items []DiscoveryResultItem, err error) {
	const discoverMSG = "M-SEARCH * HTTP/1.1\r\n HOST:239.255.255.250:1982\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n"
	const ssdpAddr = "239.255.255.250:1982"

	ssdp, err := net.ResolveUDPAddr("udp4", ssdpAddr)
	if err != nil {
		return nil, err
	}

	con, err := net.ListenPacket("udp4", ":0")
	if err != nil {
		return nil, err
	}

	socket := con.(*net.UDPConn)

	if _, err := socket.WriteToUDP([]byte(discoverMSG), ssdp); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		items, err = readFromSocket(socket)
	}()

	<-ctx.Done()

	_ = socket.Close()

	wg.Wait()

	if err == nil {
		err = ctx.Err()
	}

	return
}

func readFromSocket(socket *net.UDPConn) ([]DiscoveryResultItem, error) {
	items := make([]DiscoveryResultItem, 0)
	unique := make(map[string]bool)

	for {
		rsBuf := make([]byte, 1024)

		size, _, err := socket.ReadFromUDP(rsBuf)
		if err, ok := err.(*net.OpError); ok && err.Err.Error() == "use of closed network connection" {
			return items, nil
		}
		if err != nil {
			return nil, err
		}

		item, err := readReadDiscoveryPayload(string(rsBuf[0:size]))
		if err != nil {
			return nil, err
		}

		if unique[item.ID] {
			continue
		}

		items = append(items, item)

		unique[item.ID] = true
	}
}

func readReadDiscoveryPayload(in string) (DiscoveryResultItem, error) {
	const crlf = "\r\n"

	if strings.HasSuffix(in, crlf) {
		in += crlf
	}

	resp, err := http.ReadResponse(bufio.NewReader(strings.NewReader(in)), nil)
	if err != nil {
		return DiscoveryResultItem{}, err
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
	}, nil
}
