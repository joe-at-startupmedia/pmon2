FROM golang:1.22-alpine

ARG GIT_BRANCH_ARG=master
ARG MAKE_TARGET_ARG=test
ARG TEST_REGEX_ARG=Test
ARG CODECOV_TOKEN_ARG

RUN apk --update add build-base && \
  apk add --no-cache git curl && \
  curl -o  /usr/local/bin/codecov https://github.com/codecov/codecov-cli/releases/download/v0.7.5/codecovcli_alpine_x86_64 && \
  chmod +x /usr/local/bin/codecov && \
  ls -al /usr/local/bin/codecov && \
  mkdir -p /usr/src/pmon3 && \
  cd /usr/src/pmon3 && \
  git clone --single-branch --branch ${GIT_BRANCH_ARG} https://github.com/joe-at-startupmedia/pmon3.git . && \
  mkdir /usr/src/pmon3/data && \
  mkdir /usr/src/pmon3/logs

ENV TEST_REGEX=${TEST_REGEX_ARG}
ENV MAKE_TARGET=${MAKE_TARGET_ARG}
ENV CODECOV_TOKEN=${CODECOV_TOKEN_ARG}

ENTRYPOINT ["/bin/sh", "-c" , "cd /usr/src/pmon3 && make ${MAKE_TARGET} && /usr/local/bin/codecov upload-process -t ${CODECOV_TOKEN} -F ${MAKE_TARGET}"]