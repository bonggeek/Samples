#!/usr/bin/env python3
import json
import urllib.request
import sys

server = "169.254.169.254"
port = "80"
mdUrl = "http://" + server + ":" + port + "/metadata/instance/"
mdUrlJson =  mdUrl + "?format=json&api-version=latest_internal"

def restCall(mdUrl):
    header={'Metadata': 'True'}
    request = urllib.request.Request(url=mdUrl, headers=header) 
    response = urllib.request.urlopen(request)
    data = response.read()
    dataStr = data.decode("utf-8")
    return dataStr
    
res = restCall(mdUrlJson)
print(res)
print('======================')
print(json.dumps(res, indent=4))

