FROM golang:1.6.0-alpine
COPY ./instant-whales /instant-whales
CMD "/instant-whales"
