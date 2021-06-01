import json
import socket

rpc = socket.create_connection(("localhost", 8080))
req = dict(method = "Arith.Multiply", params = [{'A': 3, 'B': 3}])
msg = json.dumps(req).encode()
rpc.sendall(msg)
res = rpc.recv(10240)
result = json.loads(res).get("result")
print(result)
