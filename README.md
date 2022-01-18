> ⚡ __Please pay attention that in January Yeelight disabled some LAN features. This library like all similar will work only with the older version of the device firmware. In my case on the version, v2.0.6_0041 discovery doesn't work.__
>
> __Right now I don't have information on when it will be back. You can find more information here__ https://forum.yeelight.com/t/topic/22664/127



# Installation SDK

Yeelight SDK writen on Golang. Package supports 90% (for now) off all features present in the official documentation.

### Installation
```sh
go get github.com/oherych/yeelight
```

### Example

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/oherych/yeelight"
	"log"
	"time"
)


func main()  {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	devices, err := yeelight.Discovery(ctx)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		log.Fatalln(err)
	}

	for _, device := range devices {
		fmt.Println(`------`)
		fmt.Printf("Device '%s' [ID:%s Version:%s]\n", device.Name, device.ID, device.FirmwareVersion)
		fmt.Printf("Adress: %s\n", device.Location)
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
```
#### Output
```text
------
Device 'my_super' [ID:0x00000000157ef201 Version:20]
Address: 192.168.1.79:55443
Power: OFF
Properties:
>  bright : 1
>  power : on
>  color_mode : 1
```

### Contributing
Unfortunately, I have only one device and I don't have it possible to test all features. I will be grateful for the help with testing and feedback. This is the main priority for the current period.

### TODO
- [ ] Implement methods `set_scene` and `bg_set_scene`
- [ ] Implement updates listener
