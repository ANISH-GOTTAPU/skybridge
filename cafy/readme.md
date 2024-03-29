## Follow the steps to run the Client for Cafy
* Pip install opentestbed
* Import opentestbed package
* Run the python file, the file should have the below content.
* For User test run, use the dynamically generated testbed_output.
    ```
    import opentestbed
    api = opentestbed.api(location="http://127.0.0.1:8080", transport="http")    
    testbed = opentestbed.Testbed()
    d1, d2 = testbed.devices.add(), testbed.devices.add()
    
    d1.id = "d1"
    # d1.name = "R2"
    d1.role = "DUT"
    d2.id = "d2"
    # d2.name = "TGEN2"
    d2.role = "ATE"
    
    d1_port1 = d1.ports.add()
    d1_port1.id = "intf1"
    d1_port1.speed = d1_port1.S_100GB
    
    d2_port1 = d2.ports.add()
    d2_port1.id = "intf1"
    d2_port1.speed = d2_port1.S_100GB    
    link1 = testbed.links.add()    
    link1.src.device = d1.id
    link1.src.port = d1_port1.id
    link1.dst.device = d2.id
    link1.dst.port = d2_port1.id
       
    testbed_output = api.reserve(testbed)
    print(testbed_output)
    ```
![image](https://github.com/ANISH-GOTTAPU/skybridge/assets/40664949/c33b9827-9803-4092-8b26-707ef352a9f7)

