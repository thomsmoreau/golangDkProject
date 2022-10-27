FROM golang:1-bullseye as builder

WORKDIR /src

COPY ./ ./

RUN go mod download -x
# -X flag to write information into the variable (package_path.variable_name=new_value)
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.release=`git rev-parse --short=8 HEAD`'" -o /bin/server ./cmd/server

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /bin/server ./

CMD ["./server"]