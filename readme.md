
# OpenConfig Feature Profiles running with fp-opentestbedops

## Prerequisites
* Netbox is used for inventory workflow and dynamic updation of database is yet to be done. So, netbox should be updated with database.
* Netbox ip and token are hard-coded in fp.compose.yml.

## Install components

1. Clone this repository.
    git clone https://github.com/ANISH-GOTTAPU/fp-opentestbedops.git


## Steps to run the test
1. On the home directory, clone featureprofiles repo.

    ```cd $HOME
    git clone https://github.com/openconfig/featureprofiles.git
    ```

2. Initialize Environment variables, Replace `license_server_name` with actual hostname/IP address of license server.

    ```Shell
    export FEATUREPROFILES_HOME=$HOME/featureprofiles
    export LICENSE_SERVERS=license_server_name
    ```

3. Launch the deployment.
    ```Shell
    cd docker-compose
    docker compose -p keng1 --file fp.compose.yml --file fp.compose.ports.yml up -d
    ```

4. Update the testbed.go file in testbed folder as per test requirement.

5. Run the featureprofile as below
    ```Shell
    Ex: ./testbedops.sh run_fp ~/featureprofiles/interface/staticarp/otg_tests/static_arp_test/static_arp_test.go
    ```