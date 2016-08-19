FROM zyfdedh/unqlite

ENV HOME /root
ENV DB /root/unqlite.db

COPY bin/* /usr/local/bin/
RUN chmod +x /usr/local/bin/startup-update

CMD ["/usr/local/bin/startup-update"]
