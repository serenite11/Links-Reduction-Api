FROM golang

RUN go version

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o links-reduction-api ./cmd/main.go

CMD ["./links-reduction-api"]