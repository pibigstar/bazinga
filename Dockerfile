FROM loads/alpine:3.8

MAINTAINER pibigstar <pibigstar@qq.com>

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV APP_NAME bazinga
ENV WORKDIR /app/${APP_NAME}

# 添加应用可执行文件
ADD ./bin/bazinga   $WORKDIR/bazinga
# 复制配置文件
COPY ./config/* $WORKDIR/config/

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./bazinga
