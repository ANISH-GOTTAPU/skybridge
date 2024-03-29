# Lab Reservation Service End-To-End Workflow.
Lab Reservation Service helps in getting the actual testbed details  from the Netbox inventory based on the user's abstract input.

![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/8d5e9c11-1f61-46a0-aa19-4056ba6b1846)


## Prerequisites
* Netbox Docker should be running, to get the Inventory data (dynamic database updates have not yet been implemented currently, using manual feed inventory data)
* Installing Client Side SDK for APIs (Go/Python)
    * import goopentestbed for Go lang ```github.com/open-traffic-generator/opentestbed/goopentestbed```
    * import opentestbed For Python ```pip install opentestbed```
      
## [Setup Netbox Docker.](https://github.com/ANISH-GOTTAPU/skybridge/blob/update_readme/Netbox/readme.md)

## Setup Lab Reservation Service
* Run reservation service (server) using docker run.
* Pull the latest version from the ghrc.
    ```docker pull ghcr.io/open-traffic-generator/lab-reservation-service:0.0.3```
* Use the below command to run the server and ensure the Netbox is available.
    ```
    docker run -d -p 8080:8080 --name lrs ghcr.io/open-traffic-generator/lab-reservation-service:0.0.3 -netbox-host "netbox-host/IP" -netbox-port "netbox-port" -netbox-user-token "netbox-token" -framework-name "cafy" (generic/cafy/ondatra)
    ```
* Execute the client-side app (the above ondatra/cafy) to obtain the testbed reservation once the server is up and running.

## [How to run End-to-End test with Ondatra using Lab Reservation Service](https://github.com/ANISH-GOTTAPU/skybridge/blob/update_readme/ondatra/readme.md)

## [How to run test with Cafy using Lab Reservation Service](https://github.com/ANISH-GOTTAPU/skybridge/blob/update_readme/cafy/readme.md)



