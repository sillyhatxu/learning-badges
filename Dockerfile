FROM xushikuan/alpine-build:1.0

ENV BUILDER_WORK_DIR=/go/src/github.com/sillyhatxu/learning-badges
ENV WORK_DIR=/app
WORKDIR $WORK_DIR

ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY . .