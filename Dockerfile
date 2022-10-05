# build stage
FROM golang:1.17.13-alpine AS builder

WORKDIR /app

COPY . .

RUN [ "./build.sh", "webservice" ]

RUN [ "mv", "webservice", "/app/release/" ]

# CMD [ "./release/webservice" ]



# run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/release/ .

ENV HTTP_ADDRESS ":8080"

# HEALTHCHECK --interval=5s --timeout=3s \
#     CMD curl -fs http://localhost/ | exit 1

EXPOSE 8080

CMD [ "/app/webservice" ]
