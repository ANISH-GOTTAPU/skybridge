package main

import (
	"fmt"

	"github.com/open-traffic-generator/opentestbed/goopentestbed"
)

func main() {

	api := goopentestbed.NewApi()
	api.NewHttpTransport().SetLocation("http://127.0.0.1:8080")

	testbed := goopentestbed.NewTestbed()

	d1 := testbed.Devices().Add().SetId("d1").SetRole(goopentestbed.DeviceRole.ATE)
	d2 := testbed.Devices().Add().SetId("d2").SetRole(goopentestbed.DeviceRole.ATE)

	d1Port1 := d1.Ports().Add().SetId("intf1").SetSpeed(goopentestbed.PortSpeed.S_400GB)
	d2Port1 := d2.Ports().Add().SetId("intf1").SetSpeed(goopentestbed.PortSpeed.S_400GB)

	link1 := testbed.Links().Add()
	link1.Src().SetDevice(d1.Id()).SetPort(d1Port1.Id())
	link1.Dst().SetDevice(d2.Id()).SetPort(d2Port1.Id())

	fmt.Println(testbed.Marshal().ToJson())

	// api.Reserve(testbed)

}
