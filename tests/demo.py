import opentestbed
import json

api = opentestbed.api(location="http://10.39.71.143:8080", transport="http")

testbed = opentestbed.Testbed()
d1, d2 = testbed.devices.add(), testbed.devices.add()

d1.id = "d1"
d1.role = "DUT"

d2.id = "d2"
d2.role = "ATE"

d1_port1 = d1.ports.add()
d1_port1.id = "intf1"
d1_port1.speed = d1_port1.S_400GB

d2_port1 = d2.ports.add()
d2_port1.id = "intf1"
d2_port1.speed = d2_port1.S_400GB

link1 = testbed.links.add()

link1.src.device = d1.id
link1.src.port = d1_port1.id
link1.dst.device = d2.id
link1.dst.port = d2_port1.id

# print(testbed.serialize(testbed.JSON))

output = api.reserve(testbed)
print(output)
