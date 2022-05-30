FROM golang:1.17

# Download and install wkhtmltopdf
RUN apt-get install -y build-essential xorg libssl-dev libxrender-dev wget gdebi
RUN wget http://downloads.sourceforge.net/project/wkhtmltopdf/0.12.2.1/wkhtmltox-0.12.2.1_linux-trusty-amd64.deb
RUN gdebi --n wkhtmltox-0.12.2.1_linux-trusty-amd64.deb
ENTRYPOINT ["wkhtmltopdf"]

# create a directory /app
RUN mkdir /app

# set or make /app our working directory
WORKDIR /app

#copy all files
COPY ./ /app

RUN go build -o user-api

CMD ./user-api
