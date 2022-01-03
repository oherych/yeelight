This is Yeelight SDK writen on Golang. Package supports 90% (for now) off all features present in the official documentation. 

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
	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)

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
		err = client.Power(context.Background(), true, yeelight.PowerModeDefault, yeelight.AffectSudden, time.Second)
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

### Support methods


| Method  |  Yeelight  Method    |    State    |
|-----------------------|--------------|-------------|
| Get()                 |  get_prop    | implemented |
| SetColorTemperature() |  set_ct_abx  | implemented |
| SetRGB()              |  set_rgb     | implemented |
| TODO |  set_hsv  | implemented |
| TODO |  set_bright  | implemented |
| TODO |  set_power  | implemented |
| TODO |  toggle  | implemented |
| TODO |  set_default  | implemented |
| TODO |  start_cf  | implemented |
| TODO |  stop_cf  | implemented |
| TODO |  set_scene  | implemented |
| TODO |  cron_add  | implemented |
| TODO |  cron_get  | implemented |
| TODO |  cron_del  | implemented |
| TODO |  set_adjust  | implemented |
| TODO |  set_music  | implemented |
| TODO |  set_name  | implemented |
| TODO |  bg_set_rgb  | implemented |
| TODO |  bg_set_hsv  | implemented |
| TODO |  bg_set_ct_abx  | implemented |
| TODO |  bg_start_cf  | implemented |
| TODO |  bg_stop_cf  | implemented |
| TODO |  bg_set_scene  | implemented |
| TODO |  bg_set_default  | implemented |
| TODO |  bg_set_power  | implemented |
| TODO |  bg_set_bright  | implemented |
| TODO |  bg_set_adjust  | implemented |
| TODO |  bg_toggle  | implemented |
| TODO |  dev_toggle  | implemented |
| TODO |  adjust_bright  | implemented |
| TODO |  adjust_ct  | implemented |
| TODO |  adjust_color  | implemented |
| TODO |  bg_adjust_bright  | implemented |
| TODO |  bg_adjust_ct  | implemented |
| TODO |  bg_adjust_color  | implemented |


