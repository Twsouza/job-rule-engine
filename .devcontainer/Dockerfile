
FROM golang:1.21

RUN mkdir -p /src/job-rule-engine
WORKDIR /src/job-rule-engine

ADD go.mod go.sum ./

RUN go install github.com/cosmtrek/air@latest && \
  go mod download

ENV ADDR=0.0.0.0

EXPOSE 3000
