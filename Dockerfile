FROM golang
RUN go get -u github.com/golang/dep/cmd/dep
ADD . /go/src/github.com/hIMEI29A/goranger
WORKDIR /go/src/github.com/hIMEI29A/goranger
RUN make clean
RUN make all
CMD /bin/bash
