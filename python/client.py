import json
import socket
import itertools


class RPCClient(object):
    def __init__(self, addr, codec=json):
        self._socket = socket.create_connection(addr)
        self._id_iter = itertools.count()
        self._codec = codec

    def _message(self, name, *params):
        return dict(id=next(self._id_iter), method=name, params=list(params))

    def call(self, name, *params):
        req = self._message(name, *params)
        id = req.get('id')

        """
        Go RPC 返回的JSON格式
        type serverResponse struct {
            Id     *json.RawMessage `json:"id"`
            Result interface{}      `json:"result"`
            Error  interface{}      `json:"error"`
        }
        """

        msg = self._codec.dumps(req).encode()
        self._socket.sendall(msg)

        # This will actually have to loop if resp is bigger
        resp = self._socket.recv(4096)
        resp = self._codec.loads(resp)

        if resp.get('id') != id:
            raise Exception("expected id=%s, received id=%s: %s"
                            % (id, resp.get('id'), resp.get('error')))

        if resp.get('error') is not None:
            raise Exception(resp.get('error'))

        return resp.get('result')


def close(self):
    self._socket.close()


if __name__ == '__main__':
    rpc = RPCClient(("127.0.0.1", 8080))
    args = {'A': 3, 'B': 3}
    print(rpc.call("Arith.Multiply", args))
    print(rpc.call("Arith.Divide", args))
