from concurrent import futures
import grpc
import sys
import paho.mqtt.client as mqtt
import json

from proto.gen.service_pb2_grpc import *
from proto.gen.service_pb2 import *
from common.constants import *


class ServiceImpl(SatelliteServiceServicer):
    def __init__(self, username, password, ip):
        print("Connecting to MQTT Broker on {0}...".format(ip))
        self.client_ = mqtt.Client()
        self.client_.on_connect = on_connect
        self.client_.username_pw_set(username, password)
        self.client_.connect(ip)
        self.client_.loop_start()  # Start the MQTT network loop asynchronously.

        self.ip_ = getIP()

    # SatelliteServiceServicer impl:
    def SendStatusUpdate(self, request, context):
        self.SendMqttMessage(request.client_name,
                             request.command_name,
                             request.content)
        return StatusResponse()

    # MQTT handling
    def SendMqttMessage(self, client_name, command_name, request_content):
        topic = getWyomingTopic(client_name, command_name)

        infos = {
            "name":client_name,
            "ip": self.ip_,
            "request_type": command_name,
            "context": request_content
        }

        # print("Sending {2}.{0} for '{1}'".format(command_name, request_content, client_name))
        self.client_.publish(topic, json.dumps(infos), qos = 2)
        # print("Sent!")

def on_connect(client, userdata, flags, rc):
    if rc == 0:
        print("Connected to MQTT Broker!")
    else:
        print("Failed to connect, return code %d\n", rc)
        assert False
    
def serve(servicer : SatelliteServiceServicer):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_SatelliteServiceServicer_to_server(servicer, server)
    server.add_insecure_port("[::]:{0}".format(PORT))
    server.start()
    server.wait_for_termination()