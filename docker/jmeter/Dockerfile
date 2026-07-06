FROM eclipse-temurin:21-jdk

LABEL maintainer="PT Toolkit"

ENV JMETER_VERSION=5.6.3
ENV JMETER_HOME=/opt/apache-jmeter
ENV PATH=$JMETER_HOME/bin:$PATH

RUN apt-get update && \
    apt-get install -y wget unzip ca-certificates && \
    rm -rf /var/lib/apt/lists/*

RUN wget https://archive.apache.org/dist/jmeter/binaries/apache-jmeter-${JMETER_VERSION}.zip \
    -O /tmp/jmeter.zip

RUN unzip /tmp/jmeter.zip -d /opt && \
    mv /opt/apache-jmeter-${JMETER_VERSION} $JMETER_HOME && \
    rm /tmp/jmeter.zip

RUN mkdir -p /scripts
RUN mkdir -p /results
RUN mkdir -p /reports
RUN mkdir -p /plugins

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

WORKDIR /scripts

ENTRYPOINT ["/entrypoint.sh"]