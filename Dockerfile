ARG UBUNTU_VERSION=20.04
FROM ubuntu:${UBUNTU_VERSION} as haskell-builder
ARG CABAL_VERSION=3.2.0.0
ARG GHC_VERSION=8.10.2
ARG IOHK_LIBSODIUM_GIT_REV=66f017f16633f2060db25e17c170c2afa0f2a8a1
ENV DEBIAN_FRONTEND=nonintercative
RUN mkdir -p /app/src
WORKDIR /app
RUN apt-get update -y && apt-get install -y \
        automake=1:1.16.* \
        build-essential=12.* \
        g++=4:9.3.* \
        git=1:2.25.* \
        jq \
        libffi-dev=3.* \
        libghc-postgresql-libpq-dev=0.9.4.* \
        libgmp-dev=2:6.2.* \
        libncursesw5=6.* \
        libpq-dev=12.* \
        libssl-dev=1.1.* \
        libsystemd-dev=245.* \
        libtinfo-dev=6.* \
        libtool=2.4.* \
        make=4.2.* \
        pkg-config=0.29.* \
        tmux=3.* \
        wget=1.20.* \
        zlib1g-dev=1:1.2.*
RUN wget --secure-protocol=TLSv1_2 \
        https://downloads.haskell.org/~cabal/cabal-install-${CABAL_VERSION}/cabal-install-${CABAL_VERSION}-x86_64-unknown-linux.tar.xz &&\
        tar -xf cabal-install-${CABAL_VERSION}-x86_64-unknown-linux.tar.xz &&\
        rm cabal-install-${CABAL_VERSION}-x86_64-unknown-linux.tar.xz cabal.sig &&\
        mv cabal /usr/local/bin/
RUN cabal update
WORKDIR /app/ghc
RUN wget --secure-protocol=TLSv1_2 \
        https://downloads.haskell.org/~ghc/${GHC_VERSION}/ghc-${GHC_VERSION}-x86_64-deb9-linux.tar.xz &&\
        tar -xf ghc-${GHC_VERSION}-x86_64-deb9-linux.tar.xz &&\
        rm ghc-${GHC_VERSION}-x86_64-deb9-linux.tar.xz
WORKDIR /app/ghc/ghc-${GHC_VERSION}
RUN ./configure && make install
WORKDIR /app/src
RUN git clone https://github.com/input-output-hk/libsodium.git &&\
        cd libsodium &&\
        git fetch --all --tags &&\
        git checkout ${IOHK_LIBSODIUM_GIT_REV}
WORKDIR /app/src/libsodium
RUN ./autogen.sh && ./configure && make && make install
ENV LD_LIBRARY_PATH="/usr/local/lib:$LD_LIBRARY_PATH"
ENV PKG_CONFIG_PATH="/usr/local/lib/pkgconfig:$PKG_CONFIG_PATH"
WORKDIR /app/src
COPY /cardano-node /app/src/cardano-node
WORKDIR /app/src/cardano-node
RUN cabal install cardano-cli \
        --install-method=copy \
        --installdir=/usr/local/bin \
        -f -systemd
RUN cabal install cardano-node \
        --install-method=copy \
        --installdir=/usr/local/bin \
        -f -systemd
# Cleanup for server copy of /usr/local/lib
RUN rm -rf /usr/local/lib/ghc-${GHC_VERSION} /usr/local/lib/pkgconfig
FROM golang as server-builder
USER root
RUN \
        apt-get update && \
        apt-get install -y --no-install-recommends \
        protobuf-compiler \
        golang-goprotobuf-dev
COPY --from=haskell-builder /usr/local/lib /usr/local/lib
COPY --from=haskell-builder /usr/local/bin/cardano-cli /usr/local/bin/
COPY --from=haskell-builder /usr/local/bin/cardano-node /usr/local/bin/
COPY . src
WORKDIR src
RUN ./build/proto-gen.sh
ENV LD_LIBRARY_PATH="/usr/local/lib:$LD_LIBRARY_PATH"
EXPOSE 50005
RUN cardano-cli --version
CMD ./build/run-server.sh 50005