# Lab Reservation Service End-To-End Workflow.
    Lab Reservation Service gets the dynamic testbed details from the Netbox inventory based on the user input (Ondatra/Cafy test input).
## Prerequisites
* Docker Netbox should be running, to get the Inventory data (Dynamic updates have not yet been implemented, currently using manual feed inventory data)
* Testbed API SDK (Ondatra/Cafy)
    * import goopentestbed for Ondatra ```github.com/open-traffic-generator/opentestbed/goopentestbed```
    * import opentestbed For Cafy ```pip install opentestbed```
      
## [Setup Netbox Docker.](https://github.com/ANISH-GOTTAPU/skybridge/blob/update_readme/Netbox/readme.md)

## Lab Reservation Service Setup
* Run reservation service (server) using docker run.
* Pull the latest version from the ghrc.
    ```docker pull ghcr.io/open-traffic-generator/lab-reservation-service:0.0.2​```
* Use the below command to run the server and ensure the Netbox is available.
    ```
    docker run -d -p 8080:8080 --name laas -e VERSION=0.0.2 lab_reservation_service:0.0.2 -netbox-host "netbox-host/IP" -netbox-port "netbox-port" -netbox-user-token "netbox-token" -framework-name cafy (generic/cafy/ondatra)
    ```
* Execute the client-side app (the above ondatra/cafy) to obtain the testbed reservation once the server is up and running.

## [How to run with test Ondatra](https://github.com/ANISH-GOTTAPU/skybridge/blob/update_readme/ondatra/readme.md)

## [How to run with test Cafy](https://github.com/ANISH-GOTTAPU/skybridge/blob/update_readme/cafy/readme.md)

    ![image](https://github.com/ANISH-GOTTAPU/fp-opentestbedops/assets/40664949/bb55e7db-64ce-43db-8d56-3e78902eecd7)
