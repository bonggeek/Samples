#!/usr/bin/env python3
import json
import urllib.request
import sys

server = "169.254.169.254"
port = "80"
mdUrl = "http://" + server + ":" + port + "/metadata/instance/"
format = "?format=text&api-version=latest_internal"

def restCall(mdUrl):
    header={'Metadata': 'True'}
    mdUrl = mdUrl+format
    request = urllib.request.Request(url=mdUrl, headers=header)
    response = urllib.request.urlopen(request)
    data = response.read()
    dataStr = data.decode("utf-8")
    return dataStr
    
def fetch(mdUrl):
    #print("Called with %s" % mdUrl)
    data = restCall(mdUrl)
    #print ("Got %s" % data)
    for e in data.splitlines():
        newUrl = mdUrl + e
        if e.endswith('/'):
            fetch(newUrl)
        else:
            r = restCall(newUrl)
            print(newUrl, r)

fetch(mdUrl)
