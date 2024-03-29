
# OpenConfig Feature Profiles running with Opentestbed

## Prerequisites
* Netbox is currently utilized for inventory workflow, but dynamic database updates have not yet been implemented. Therefore, Netbox requires updating with the database.
* Netbox-hostname, Netbox-port and Netbox-user-token are currently hard-coded in fp.compose.yml.

## Steps to run the test

1. Initialize Environment variables, Replace `license_server_name` with actual hostname/IP address of license server. This license server will be used by KENG

    ```Shell
    export LICENSE_SERVERS=license_server_name
    ```

2. Launch the deployment.
    ```Shell
    cd docker-compose
    docker compose -p keng1 --file fp.compose.yml --file fp.compose.ports.yml up -d
    ```

3. On the home directory, clone featureprofiles repo. Use the branch dynamic-binding to use dynamic-binding flag for demo purpose.

    ```cd $HOME
    git clone https://github.com/open-traffic-generator/featureprofiles.git
    git checkout dynamic-binding
    ```

4. Update the testbed(`~/featureprofiles/topologies/dynamic_testbed.json`) file as per the test requirement.

5. Run the featureprofile test as shown in example.
    ```Shell
    Ex: go test -v feature/experimental/isis/otg_tests/lsp_updates_test/*.go -dynamic-binding ~/featureprofiles/topologies/dynamic_testbed.json -testbed ~/featureprofiles/topologies/atedut_4.testbed | tee output.log
    ```
    ![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/291a4a1b-14cc-43e0-8c79-f4db994a1b91)


