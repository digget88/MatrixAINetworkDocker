# Build Gman in a stock Go builder container #shang and yang
FROM golang:1.12-alpine 

# Grab needed packages
RUN apk add --no-cache make gcc musl-dev linux-headers git gnupg

# Add repo and compile
ADD . /go-matrix
RUN cd /go-matrix && make gman

# Pull Gman into a second stage deploy alpine container

RUN apk add --no-cache ca-certificates

#No longer needed EXPOSE 8341 8546 50505 50505/udp 30304/udp
# Start node script that sets a random entrust password to start node
ENTRYPOINT ["/go-matrix/nodeConfig.sh"]


#To start your node, run "docker run -d --network matrixnet --ip <ip address> -e MAN_PORT='<modified 50505 port>' -v /path/to/host/chaindata:/go-matrix/chaindata --name <docker_name> matrix"
#The /path/to/host/chaindata directory should be a directory on the host system that contains your keystore folder, with your wallet inside it.
#The directory can be named whatever you'd like but it will get mounted as the container chaindata folder. The scripts will do the rest.
