sudo apt-get install curl
sudo apt-get jq
curl -H Metadata:True "http://169.254.169.254/metadata/instance?api-version=2017-04-02&format=json" | jq .
