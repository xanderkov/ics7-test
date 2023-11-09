FROM docker:dind

ARG GOLANG_VERSION=1.21.0

RUN wget https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz &&  \
    tar -C /usr/local -zxvf go${GOLANG_VERSION}.linux-amd64.tar.gz &&  \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
ENV GOBIN /go/bin
