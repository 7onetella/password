# build ember
FROM node:10-alpine AS node-builder
RUN npm install -g ember-cli
WORKDIR /build
COPY ./ui /build/
RUN ls -l /build  && \
    cd /build     && \
    ./build.sh


# build go
FROM golang:1.15.6-alpine3.12 as go-builder
WORKDIR /build
COPY ./api /build/
COPY --from=node-builder /build/dist ./build/ui/
RUN ls -l /build  && \
    cd /build     && \
    ./build.sh   

# ship
FROM alpine:3.12.2
COPY --from=go-builder /build/dist/password_linux_amd64/password /password
RUN chmod +x /password
CMD [ "/password" ]
