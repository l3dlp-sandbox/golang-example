FROM ubuntu:latest as build
RUN apt-get update
RUN apt-get install -y wget build-essential
RUN wget -P /tmp "https://go.dev/dl/go1.19.4.linux-amd64.tar.gz"
RUN tar -C /usr/local -xzf "/tmp/go1.19.4.linux-amd64.tar.gz"
RUN rm "/tmp/go1.19.4.linux-amd64.tar.gz"
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

#RUN wget https://download.libsodium.org/libsodium/releases/libsodium-1.0.17.tar.gz && tar xvf ./libsodium-1.0.17.tar.gz  && ./libsodium-1.0.17/configure && make && make check && make install
#RUN apt-get install -y libsodium-dev
#RUN apt-get install -y libtool pkg-config build-essential automake uuid-dev manpages-dev
#RUN apt-get update
##RUN apt-get install libzmq5 libkrb5-dev libnorm-dev libpgm-dev libsodium-dev libunwind-dev libunwind8-dev libunwind7-dev libnss3-dev libgnutls28-dev libgnutls-dev libbsd-dev
#RUN wget https://download.opensuse.org/repositories/network:/messaging:/zeromq:/release-stable/xUbuntu_20.04/amd64/libzmq3-dev_4.3.4-0_amd64.deb &&  dpkg -i libzmq3-dev_4.3.4-0_amd64.deb
#RUN apt-get -y install libczmq-dev
WORKDIR /app
COPY . .
#RUN go get ./... && \
#    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o golang

#FROM ubuntu:latest
#COPY --from=build /app/golang /golang
#EXPOSE 9000
#CMD [ "/golang" ]

