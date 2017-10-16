#!/bin/bash
#sudo apt-get install curl
#sudo apt-get jq
#curl -H Metadata:True "http://169.254.169.254/metadata/instance?api-version=2017-04-02&format=json" | jq .

getVmId()
{
    for i in `seq 1 60`
    do
        retVal=`curl -H Metadata:True --silent --write-out "~%{http_code}" http://169.254.169.254/metadata/instance/compute/vmId?api-version=2017-04-02\&format=text`
        IFS='~' read vmId http_status <<< $retVal
        if [ "$http_status" == "200" ]; then
            echo "Got vmId"
            break
        fi

        echo "Attempt $i: Failed with $http_status ($vmId)"
        sleep 1
    done
}
getVmId
echo $vmId

