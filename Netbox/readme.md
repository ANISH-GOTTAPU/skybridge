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
* Use the below command to update the status of Devices/Interfaces (The tool works with static devices/interface names).
     ```
     go run reset_device_state.go
     ```