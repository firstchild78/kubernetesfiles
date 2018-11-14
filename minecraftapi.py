from mcstatus import MinecraftServer

import requests
import json

url = "https://mcapi.us/server/status?"
ip = '13.75.199.109'
port = '25565'
fullUrl = url + "ip=" + ip + "&port=" + port

# Call the API
myResponse = requests.get(fullUrl, verify=True)
# Get the response
print(myResponse.status_code)

if (myResponse.ok):
    jData = json.loads(myResponse.content.decode('utf-8'))

    print("The response contains {0} properties".format(len(jData)))
    print("\n")
    for key in jData:
        print(key + ":" + str(jData[key]))
    
else:
    myResponse.raise_for_status


# Hold server just in case
server = MinecraftServer.lookup("13.75.199.109:25565")

# Status of the server
status = server.status()
print("The server has {0} players and replied in {1} ms".format(status.players.online, status.latency))

# Latency
latency = server.ping()
print("The server replied in {0} ms".format(latency))

# query
# query = server.query()
# print("The server has the following players online:{0}".format(", ".join(query.players.names)))