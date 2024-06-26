package main

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "strings"
)

const (
        NETBOX_URL = "http://10.39.71.52:8000/api/"
        TOKEN      = "fd5b272642d5bead785fbac32ef501bfab9e9e95"
        HEADERS    = "application/json"
)

var httpClient = &http.Client{}

func createRequest(method, url string, body []byte) (*http.Request, error) {
        req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
        if err != nil {
                return nil, err
        }
        req.Header.Set("Authorization", "Token "+TOKEN)
        req.Header.Set("Content-Type", HEADERS)
        return req, nil
}

func performRequest(req *http.Request) (*http.Response, error) {
        resp, err := httpClient.Do(req)
        if err != nil {
                return nil, err
        }
        return resp, nil
}

func updateDevicesData() {
        deviceNames := map[string]string{
                "R1": "dut",
                "R2": "dut",
                "R3": "dut",
        }
        portNames := []map[string]string{
                {"R1": "HundredGigE1/0"},
                {"R1": "FourHundredGigE1/0"},
                {"R2": "HundredGigE2/0"},
                {"R3": "FourHundredGigE3/0"},
                {"TGEN1": "10.39.0.3/1/1"},
                {"TGEN1": "10.39.0.3/1/3"},
                {"TGEN2": "10.39.0.4/1/1"},
                {"TGEN3": "10.39.0.5/1/1"},
        }
        client := httpClient
        //deviceNames := map[string]string{
        //      "R1":"dut",
        //      "R2":"dut",
        //      "R3":"dut",
        //}
        //portNames := []string{"FourHundredGigE0/0/0/16", "FourHundredGigE0/0/0/18", "FourHundredGigE0/0/0/20", "FourHundredGigE0/0/0/22", "10.36.71.246/1", "10.36.71.246/2", "10.36.71.246/3", "10.36.71.246/5"}
        for deviceName, model := range deviceNames {
                // url := NETBOX_URL + "dcim/devices/?name=" + deviceName
                url := fmt.Sprintf("%sdcim/devices/?name=%s", NETBOX_URL, deviceName)
                req, err := createRequest("GET", url, nil)
                if err != nil {
                        log.Fatal("Get request failed: ", url, "\nError:", err)
                }
                response, err := performRequest(req)
                if err != nil {
                        log.Fatal("Failed to get response: ", url, "\nError:", err)
                }
                defer response.Body.Close()

                body, err := ioutil.ReadAll(response.Body)
                if err != nil {
                        log.Fatal("Failed to read response url: ", url, "\nError:", err)
                }

                var deviceDict map[string]interface{}
                if err := json.Unmarshal(body, &deviceDict); err != nil {
                        log.Fatal("Failed to unmarshal deviceDict: ", err)
                }
                results := deviceDict["results"].([]interface{})
                if len(results) > 0 {
                        deviceDict := results[0].(map[string]interface{})
                        if strings.ToLower(model) != "ate" {
                                if strings.ToLower(deviceDict["name"].(string)) == strings.ToLower(deviceName) {
                                        deviceURL := deviceDict["url"].(string)
                                        updateData := map[string]interface{}{
                                                "name":          deviceDict["name"],
                                                "device_type":   deviceDict["device_type"].(map[string]interface{})["id"],
                                                "custom_fields": map[string]interface{}{"state": "Available"},
                                        }
                                        updateDataJSON, err := json.Marshal(updateData)
                                        if err != nil {
                                                log.Fatal("Failed Json marshal with updateData: ", err)
                                        }
                                        req, err := http.NewRequest("PATCH", deviceURL, bytes.NewBuffer(updateDataJSON))
                                        if err != nil {
                                                log.Fatal("Failed to patch the data with url:", deviceURL, "\nError:", err)
                                        }
                                        req.Header.Set("Authorization", "Token "+TOKEN)
                                        req.Header.Set("Content-Type", HEADERS)
                                        response, err := client.Do(req)
                                        if err != nil {
                                                log.Fatal("Failed to get response for deviceURL: ", deviceURL, "\nError:", err)
                                        }
                                        defer response.Body.Close()
                                        if response.StatusCode != http.StatusOK {
                                                log.Printf("Error updating device details. Status code: %d\n", response.StatusCode)
                                        }
                                } else {
                                        log.Printf("Failed to find the device: %s\n", deviceName)
                                }
                        }
                }
        }
        for _, portName := range portNames {
                for deviceID, name := range portName {
                        url := fmt.Sprintf("%sdcim/interfaces/?name=%s", NETBOX_URL, name)
                        req, err := createRequest("GET", url, nil)
                        if err != nil {
                                log.Fatal("Get request failed: ", url, "\nError:", err)
                        }
                        response, err := performRequest(req)
                        if err != nil {
                                log.Fatal("Failed to get response: ", url, "\nError:", err)
                        }
                        defer response.Body.Close()

                        body, err := ioutil.ReadAll(response.Body)
                        if err != nil {
                                log.Fatal("Failed to read response url: ", url, "\nError:", err)
                        }

                        var portDict map[string]interface{}
                        if err := json.Unmarshal(body, &portDict); err != nil {
                                log.Fatal("Failed to unmarshal deviceDict: ", err)
                        }
                        results := portDict["results"].([]interface{})
                        for _, result := range results {
                                portDict := result.(map[string]interface{})
                                var devicename string
                                if deviceData, ok := portDict["device"].(map[string]interface{}); ok {
                                        devicename = deviceData["name"].(string)
                                }
                                if strings.ToLower(portDict["name"].(string)) == strings.ToLower(name) && strings.ToLower(devicename) == strings.ToLower(deviceID) {
                                        portURL := portDict["url"].(string)
                                        updateData := map[string]interface{}{
                                                "custom_fields": map[string]interface{}{"state": "Available"},
                                        }
                                        updateDataJSON, err := json.Marshal(updateData)
                                        if err != nil {
                                                log.Fatal("Failed Json marshal with updateData: ", err)
                                        }
                                        req, err := http.NewRequest("PATCH", portURL, bytes.NewBuffer(updateDataJSON))
                                        if err != nil {
                                                log.Fatal("Failed to patch the data with url:", portURL, "\nError:", err)
                                        }
                                        req.Header.Set("Authorization", "Token "+TOKEN)
                                        req.Header.Set("Content-Type", HEADERS)
                                        response, err := client.Do(req)
                                        if err != nil {
                                                log.Fatal("Failed to get response for portURL: ", portURL, "\nError:", err)
                                        }
                                        defer response.Body.Close()
                                        if response.StatusCode != http.StatusOK {
                                                log.Printf("Error updating port details. Status code: %d\n", response.StatusCode)
                                        }
                                }
                        }
                }
        }
}

func main() {
        updateDevicesData()
        fmt.Println("Devices and Interfaces state updated successfully")
}
