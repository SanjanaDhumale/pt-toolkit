FROM eclipse-temurin:21-jdk

RUN apt-get update && apt-get install -y \
    wget \
    unzip \
    curl

WORKDIR /opt

RUN wget https://archive.apache.org/dist/jmeter/binaries/apache-jmeter-5.6.3.zip

RUN unzip apache-jmeter-5.6.3.zip

ENV JMETER_HOME=/opt/apache-jmeter-5.6.3

ENV PATH=$JMETER_HOME/bin:$PATH

CMD ["tail","-f","/dev/null"]