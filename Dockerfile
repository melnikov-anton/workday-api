FROM golang:alpine as builder
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o workday .

WORKDIR /dist
RUN cp /build/workday .
RUN cp -r /build/static/. /dist/static/ && \
    cp -r /build/templates/. /dist/templates/ && \
    cp -r /build/data/. /dist/data/


FROM alpine
COPY --from=builder /dist /dist/
EXPOSE 80
CMD ["/dist/workday", "-port", "80"]