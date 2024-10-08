# Lab Reservation Service End-To-End Workflow.
Lab Reservation Service helps in getting the actual testbed details  from the Netbox inventory based on the user's abstract input.

![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/8d5e9c11-1f61-46a0-aa19-4056ba6b1846)


## Prerequisites
* Netbox Docker should be running, to get the Inventory data (dynamic database updates have not yet been implemented currently, using manual feed inventory data)
* Installing Client Side SDK for APIs (Go/Python)
    * import goopentestbed for Go lang ```github.com/open-traffic-generator/opentestbed/goopentestbed```
    * import opentestbed For Python ```pip install opentestbed```
      
## [Setup Netbox Docker.](https://github.com/ANISH-GOTTAPU/skybridge/blob/main/Netbox/readme.md)
## Netbox mandotary attributes/custom fields
<b>state</b> and <b>session_id</b> attributes on Netbox should be added as custom fields for both devices and Interfaces.

* <b>state</b> should be set to Available during initial configuration. 

## Setup Lab Reservation Service
* Run reservation service (server) using docker run.
* Pull the latest version from the ghrc.
    ```docker pull ghcr.io/open-traffic-generator/lab-reservation-service:0.0.6```
* Use the below command to run the server and ensure the Netbox is available.
l1switch location is optional, add it only if l1switch is used.
    ```
    docker run -d -p 8080:8080 --name testbed-reservation-service ghcr.io/open-traffic-generator/lab-reservation-service:0.0.6 -netbox-host <netbox-host/IP:netbox-port> -netbox-user-token <netbox-token> -framework-name generic --l1switch-location localhost:9000

    docker run -d -p 9000:9000 --name l1-service ghcr.io/open-traffic-generator/l1s-service:0.0.1 --l1switch-host <switch-host/IP> --l1switch-user <username> --l1switch-pass <password> --l1switch-model <switchmodel>
    ```

# [docker compose file with lab-reservation service and l1s-service](https://github.com/ANISH-GOTTAPU/skybridge/blob/main/docker-compose/b2b/fp.compose.yml)

* Execute the client-side app (the above ondatra/cafy) to obtain the testbed reservation once the server is up and running.

## Basic reservation sample
```
import opentestbed
import json

api = opentestbed.api(location="http://localhost:8080", transport="http")


testbed = opentestbed.Testbed()
t1, r1 = testbed.devices.add(), testbed.devices.add()

t1.id = "ate"
t1.role = "ATE"
t1.vendor = "KEYSIGHT"

r1.id = "dut"
r1.role = "DUT"
r1.vendor = "CISCO"

t1_port1 = t1.ports.add()
t1_port1.id = "t1intf1"
t1_port1.speed = t1_port1.S_400GB

r1_port1 = r1.ports.add()
r1_port1.id = "r1intf1"
r1_port1.speed = r1_port1.S_400GB

link1, link2 = testbed.links.add()

#link1
link1.src.device = t1.id
link1.src.port = t1_port1.id
link1.dst.device = r1.id
link1.dst.port = r1_port1.id

# print(testbed.serialize(testbed.JSON))
try:
    output = api.reserve(testbed)
    print(output.testbed)
    print(output.sessionid)
    # Uncomment below if you want to use the release API
    # session = opentestbed.Session()
    # session.id = output.sessionid
    # release = api.release(session)
except Exception as e:
    print(f"Error: {e}")

 ```

## [How to run End-to-End test with Ondatra using Lab Reservation Service](https://github.com/ANISH-GOTTAPU/skybridge/blob/main/ondatra/readme.md)

## [How to run test with Cafy using Lab Reservation Service](https://github.com/ANISH-GOTTAPU/skybridge/blob/main/cafy/readme.md)



