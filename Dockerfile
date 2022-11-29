FROM ubuntu

RUN apt update
RUN apt search golang-go 
RUN apt install golang-go 

FROM mongo 
EXPOSE 27017