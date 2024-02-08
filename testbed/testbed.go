package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/open-traffic-generator/opentestbed/goopentestbed"
)

func main() {

	api := goopentestbed.NewApi()
	api.NewHttpTransport().SetLocation("http://127.0.0.1:8080")

	testbed := goopentestbed.NewTestbed()

	d1 := testbed.Devices().Add().SetId("d1").SetRole(goopentestbed.DeviceRole.ATE)
	d2 := testbed.Devices().Add().SetId("d2").SetRole(goopentestbed.DeviceRole.ATE)

	d1Port1 := d1.Ports().Add().SetId("intf1").SetSpeed(goopentestbed.PortSpeed.S_100GB)
	d2Port1 := d2.Ports().Add().SetId("intf1").SetSpeed(goopentestbed.PortSpeed.S_100GB)

	link1 := testbed.Links().Add()
	link1.Src().SetDevice(d1.Id()).SetPort(d1Port1.Id())
	link1.Dst().SetDevice(d2.Id()).SetPort(d2Port1.Id())

	// fmt.Println(testbed.Marshal().ToJson())

	// api.Reserve(testbed)
	reservationResult, err := api.Reserve(testbed)
	if err != nil {
		fmt.Printf("Error during reservation: %v\n", err)
		return
	}
	unescapedJSON := strings.ReplaceAll(*reservationResult, "\\\"", "\"")
	unescapedJSON = strings.ReplaceAll(unescapedJSON, "\\n", "\n")

	// Write the unescaped JSON string to a file
	err = ioutil.WriteFile("otg.binding", []byte(unescapedJSON), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	// Read the otg.binding file
	fileContent, err := ioutil.ReadFile("otg.binding")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Trim double quotes from the content
	trimmedContent := strings.Trim(string(fileContent), `"`)

	// Write the trimmed content back to the file
	err = ioutil.WriteFile("otg.binding", []byte(trimmedContent), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	// Print the unescaped JSON string
	fmt.Println(strings.Trim(unescapedJSON, `"`))
}
