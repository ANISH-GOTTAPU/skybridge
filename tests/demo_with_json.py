import opentestbed
import json

api = opentestbed.api(location="http://10.39.71.18:8080", transport="http")

testbed = opentestbed.Testbed()
file_path = "testbed.json"

with open(file_path, 'r') as file:
    data = json.load(file)


output = api.reserve(testbed.deserialize(data))
print(output)

