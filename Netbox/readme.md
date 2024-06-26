## Setup Netbox Docker.
### The following steps are not required if already running Netbox Docker.
   * Clone the Netbox repository.
       ```git clone -b release https://github.com/netbox-community/netbox-docker.git```
   * Move to netbox-docker directory.
   * Create a new file that defines the port under which NetBox will be available. The file name must be docker-compose.override.yml and its content should be as follows:
       ```
       version: '3.4'
       services:
         netbox:
           ports:
           - 8000:8080
       ```
   * Need to pull all the containers from the Docker registry and may take a while, depending on your internet connection.
       ```docker compose pull```
   * Finally, start all the required Docker containers.
       ```docker compose up```
   * To create the first admin user run this command:
       ```docker compose exec netbox /opt/netbox/netbox/manage.py createsuperuser```
   * If you want to just stop NetBox to continue work later on, use the following command.
       ```
       # Stop all the containers
       docker compose stop
       
       # Start the containers again
       docker compose start
       ```
   * Configure the basic devices on Netbox using the below link.
        ```
        https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/open-traffic-generator/testbed/testbed_model/artifacts/openapi.yaml&nocors#tag/Testbed
        ```
## Created tool to update the Devices/Interfaces status.
* Use the below command to update the state of Devices/Interfaces (The tool works with static devices/interface names).
     ```
     go run reset_device_state.go
     ```
* Use the below command to dynamically update the state of Devices/Interfaces.
     ```
     python reset_device_state.py --netbox_api_url Hostame/IP:Port --netbox_api_token token
     Example: python reset_device_state.py --netbox_api_url box.keysight.com:3000 --netbox_api_token a67af4ed233b6dfbc62621edb14444d556f88a4e
     ```
## Add Devices/Interfaces
* Create Devices with the below mandatory fields.
  * Device role, Device type, Site, Status
* Create Device Roles enums: ATE, DUT, L1S.\
  ![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/f676d5da-675a-499e-8023-341bef41d034)
* Create Interfaces with the below mandatory fields.
  * Device, Name, Type
## Follow the below steps to add Custom Field (state) for Devices/Interfaces.
* Add custom fields with the mandatory field "state" for Devices/Interfaces.\
  ![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/b58157db-a863-46d3-b67c-c3a3193e83cf)
* Login into Netbox, go to customization and click Add custom fields.\
  ![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/ff1d1060-f8ee-47c0-acf2-e268eddc2a22)

* Fill in the Content types and all the mandatory fields and click on Create.\
  ![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/fb5bcdd6-9dae-4b9b-877a-7b2585daba12)

* Goto Devices edit the node and fill the custom field values.\
  ![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/b51d8af2-be07-459d-9bbd-cc1c8e2c2235)

* We want to display the custom field on the device view -> Click on "Configure Table" on the page add which column to display on the page and save it.\
  ![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/7f7aa529-fa83-4bfd-802a-fc55a4ac0051)


