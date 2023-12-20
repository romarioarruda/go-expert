FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o go-expert

FROM scratch
COPY --from=builder /app/go-expert .
CMD [ "./go-expert" ]