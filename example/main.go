package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/oherych/yeelight"
)

//nolint:gocritic
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	devices, err := yeelight.Discovery(ctx)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		log.Fatalln(err)
	}

	for _, device := range devices {
		fmt.Println(`------`)
		fmt.Printf("Device '%s' [ID:%s Version:%s]\n", device.Name, device.ID, device.FirmwareVersion)
		fmt.Printf("Address: %s\n", device.Location)
		fmt.Printf("Power: %s\n", power(device.Power))

		// create new client for work with device
		client := yeelight.New(device.Location)

		// read all properties
		properties, err := client.GetProperties(context.Background(), []string{yeelight.PropertyPower, yeelight.PropertyColorMode, yeelight.PropertyBright})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Properties:")
		for name, value := range properties {
			fmt.Println("> ", name, ":", value)
		}

		// change power to ON
		err = client.Power(context.Background(), true, yeelight.PowerModeDefault, yeelight.EffectSudden, time.Second)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func power(on bool) string {
	if on {
		return "ON"
	}

	return "OFF"
}
