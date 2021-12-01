FROM golang:1.12.5 AS builder
RUN apt update
WORKDIR /go/src/github.com/superconsensus/matrixchain
COPY . .
RUN make clean && make

# ---
FROM ubuntu:16.04
WORKDIR /home/work/xuperunion/
COPY --from=builder /go/src/github.com/superconsensus/matrixchain/output/ .
EXPOSE 37101 47101
CMD ./xchain-cli createChain && ./xchain
