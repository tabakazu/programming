FROM ubuntu:18.04
RUN apt-get update -qq
RUN apt-get install -y wget
RUN wget https://packages.chef.io/files/stable/chefdk/3.0.36/ubuntu/18.04/chefdk_3.0.36-1_amd64.deb
RUN dpkg -i chefdk_3.0.36-1_amd64.deb
