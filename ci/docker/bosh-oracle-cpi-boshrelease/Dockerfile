FROM ubuntu:15.04
MAINTAINER CF Dev <cloud_dev_cf_ww_grp@oracle.com>

# Update base image
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
RUN locale-gen en_US.UTF-8; dpkg-reconfigure locales
RUN apt-get update; apt-get -y upgrade; apt-get clean

# Install dependencies
RUN apt-get -y install build-essential git-core curl wget jq sudo; apt-get clean
RUN apt-get install -y zip unzip; apt-get clean
RUN apt-get install -y sqlite3 libsqlite3-dev; apt-get clean
RUN apt-get install -y mysql-client libmysqlclient-dev; apt-get clean
RUN apt-get install -y libpq-dev; apt-get clean
RUN apt-get install -y python python-setuptools libyaml-dev libpython-dev; apt-get clean
RUN apt-get install -y ruby vim; apt-get clean
RUN apt-get remove python-pip
RUN easy_install pip Jinja2

# Install Golang
ENV GOLANG_VERSION 1.8.1
RUN curl -sSL https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz | tar -v -C /usr/local -xz
ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin

# Install Google Cloud SDK
RUN pip install virtualenv oci-cli
ENV PATH $PATH:/usr/local/google-cloud-sdk/bin

# Install BOSH v2 CLI
RUN curl -sSL -o /usr/local/bin/bosh https://s3.amazonaws.com/bosh-cli-artifacts/bosh-cli-2.0.28-linux-amd64
RUN chmod +x /usr/local/bin/bosh
