FROM zyfdedh/unqlite

ENV HOME /root
ENV DB /root/unqlite.db

COPY bin/* /usr/local/bin/

CMD ["/usr/local/bin/startup-update"]
