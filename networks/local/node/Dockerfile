FROM golang:stretch as build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y $PACKAGES

WORKDIR /brczero
# Add source files
COPY . .

# build brczero
RUN make build-linux

# Final image
FROM golang:1.20 as final

WORKDIR /brczerod
# Copy over binaries from the build-env
COPY --from=build-env /brczero/build/brczerod /usr/bin/brczerod
COPY --from=build-env /brczero/build/brczerocli /usr/bin/brczerocli
COPY --from=build-env /brczero/networks/local/node/wrapper.sh /usr/bin/wrapper.sh

EXPOSE 26656 26657
ENTRYPOINT ["/usr/bin/wrapper.sh"]
CMD ["start"]
STOPSIGNAL SIGTERM

