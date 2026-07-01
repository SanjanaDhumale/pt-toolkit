# Base Image
FROM eclipse-temurin:21-jdk

# Install required packages
RUN apt-get update && apt-get install -y \
    python3 \
    wget \
    unzip \
    curl \
    git

# Set working directory
WORKDIR /opt

# Download Apache JMeter
RUN wget https://archive.apache.org/dist/jmeter/binaries/apache-jmeter-5.6.3.zip

# Extract JMeter
RUN unzip apache-jmeter-5.6.3.zip

# Environment Variables
ENV JMETER_HOME=/opt/apache-jmeter-5.6.3
ENV PATH=$JMETER_HOME/bin:$PATH

# Default Command
CMD ["jmeter", "--version"]