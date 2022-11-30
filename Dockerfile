FROM ubuntu
RUN apt update
RUN apt search golang-go 
RUN apt install -y golang-go 
FROM mongo 
EXPOSE 27017
RUN mongosh





FROM ubuntu

RUN apt update
RUN apt-get install -y gnupg
RUN apt-get install -y wget
RUN wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc |  apt-key add - 
RUN echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/6.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-6.0.list
RUN apt-get update
RUN wget http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.1f-1ubuntu2_amd64.deb
RUN dpkg -i libssl1.1_1.1.1f-1ubuntu2_amd64.deb
RUN apt-get install -y mongodb-org
RUN apt install -y systemctl
RUN systemctl enable mongod
RUN systemctl daemon-reload
RUN systemctl start mongod
RUN mongosh


FROM ubuntu
RUN apt update
RUN apt-get install -y gnupg
RUN apt-get install -y wget
RUN wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc |  apt-key add - 
RUN echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/6.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-6.0.list
RUN apt-get update
RUN wget http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.1f-1ubuntu2_amd64.deb
RUN dpkg -i libssl1.1_1.1.1f-1ubuntu2_amd64.deb
RUN apt install sudo 
RUN sudo apt-get install -y mongodb-org
RUN sudo apt install -y systemctl
RUN sudo systemctl enable mongod
RUN sudo systemctl daemon-reload
RUN sudo systemctl start mongod