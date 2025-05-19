FROM nginx:stable-bookworm

LABEL maintainer="A. Lloyd Flanagan <lloyd.flanagan@proton.me>"

SHELL ["/bin/bash", "-eu", "-o", "pipefail", "-c"]

# hadolint ignore=DL3008, DL3009
RUN apt-get update \
    && apt-get install \
        --no-install-recommends -y \
        apt-utils \
        ca-certificates \
        curl \
        procps

RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash -

# hadolint ignore=DL3008
RUN apt-get install -y --no-install-recommends nodejs \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

COPY app ./

RUN cp nginx.host.conf /etc/nginx/conf.d/

CMD ["nginx", "-g", "daemon off;"]
EXPOSE 80
