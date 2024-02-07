FROM golang:1.22
#
WORKDIR /app
COPY . . 
COPY go.mod ./
RUN go mod download
RUN rm config.yml
#
RUN go build -o ./gocloak-dd .
RUN chmod +x ./gocloak-dd
#
EXPOSE 8080
#
CMD [ "./gocloak-dd" ]
#
