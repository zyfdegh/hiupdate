FROM zyfdedh/unqlite
MAINTAINER Zhang Yifa <zyfdegg@gmail.com>

WORKDIR /usr/local/bin

ENV HOME /root
ENV DB /root/unqlite.db

COPY bin /usr/local/bin/
RUN chmod +x /usr/local/bin/hiupdate

CMD ["/usr/local/bin/hiupdate"]
