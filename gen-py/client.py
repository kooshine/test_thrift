# -*- coding: utf-8 -*-

import sys
import json
from hello import Transmit
from hello.ttypes import *
from hello.constants import *
from thrift import Thrift
from thrift.transport import TSocket
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol


transport = TSocket.TSocket('127.0.0.1', 8000)
transport = TTransport.TBufferedTransport(transport)
protocol = TBinaryProtocol.TBinaryProtocol(transport)
client = Transmit.Client(protocol)
# Connect!
transport.open()
"""
cmd = 2
token = '1111-2222-3333-4444'
data = json.dumps({"name":"zhoujielun"})
msg = client.invoke(cmd,token,data)
print(msg)

"""
data = json.dumps({"name":"zhoujielun"})
msg = client.sayMsg(data)
print msg

transport.close()
