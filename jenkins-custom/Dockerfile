FROM jenkins/jenkins:lts

# Install prerequisites (wget, etc.)
USER root
RUN apt-get update && apt-get install -y \
    wget \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install Go
RUN wget https://go.dev/dl/go1.23.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz
RUN rm go1.23.1.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
RUN go version

# Remove the existing docker group (if it exists) and create it with GID 999
RUN groupdel docker || true
RUN groupadd -g 999 docker

# Install Docker
RUN apt-get update && apt-get install -y \
    docker.io \
    && rm -rf /var/lib/apt/lists/*

# Add the jenkins user to the docker group
RUN usermod -aG docker jenkins

# Switch back to the Jenkins user
USER jenkins
