import os
import sys

import grpc
from grpc._channel import _Rendezvous

import data_pb2
import data_pb2_grpc

current_fonder_path = os.path.split(os.path.realpath(__file__))[0]
print(current_fonder_path)
protocal_path = os.path.join(current_fonder_path, "pb")
print(protocal_path)
sys.path.append(protocal_path)

_HOST = '127.0.0.1'
_PORT = '8080'


def run():
    conn = grpc.insecure_channel(_HOST + ':' + _PORT)
    client = data_pb2_grpc.FormatDataStub(channel=conn)
    response = client.DoFormat(data_pb2.actionRequest(text='hello,world!', corpus="NEWS"))
    for i in response.result:
        print(i)


if __name__ == '__main__':
    try:
        run()
    except _Rendezvous as e:
        print("connect error")
