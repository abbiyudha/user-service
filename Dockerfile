FROM golang:1.17

# create a directory /app
RUN mkdir /app

# set or make /app our working directory
WORKDIR /app

#copy all files
COPY ./ /app

RUN go build -o user-api

CMD ./user-api
