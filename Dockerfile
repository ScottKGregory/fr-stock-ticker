FROM golang:1.18 as build-env
WORKDIR /go/src/fr-stock-ticker
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o /go/bin/fr-stock-ticker

FROM gcr.io/distroless/static
COPY --from=build-env /go/bin/fr-stock-ticker /
EXPOSE 3000
CMD ["/fr-stock-ticker"]