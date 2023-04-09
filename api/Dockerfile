# Use the official super-linter image as a base
FROM github/super-linter:v4

# Set the environment variables
ENV GOPATH /home/orson/go
ENV PROJECT_DIR /home/orson/GolandProjects/nautilus-billing/api

# Set the working directory
WORKDIR $PROJECT_DIR

# Configure Git to use SSH instead of HTTPS
RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"

# Accept the SSH private key as an argument
ARG SSH_PRIVATE_KEY

# Set up the SSH key for authentication
RUN mkdir -p /root/.ssh \
    && echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa \
    && chmod 600 /root/.ssh/id_rsa \
    && ssh-keyscan -t rsa github.com >> /root/.ssh/known_hosts

# Copy the go.mod and go.sum files into the working directory
COPY go.mod go.sum ./

# Change directory


# Install the dependencies
RUN go mod download

# BUILD WITH:
# docker build -t nb-super-linter -f SuperLinterDockerfile --build-arg SSH_PRIVATE_KEY="$(cat ~/.ssh/id_rsa)" .

