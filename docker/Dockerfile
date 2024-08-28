FROM golang as builder
WORKDIR /builder

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build src/golang-extract-api.go 

FROM golang as application
WORKDIR /app

COPY --from=builder /builder/golang-extract-api /app/

EXPOSE 9999

CMD [ "./golang-extract-api" ]