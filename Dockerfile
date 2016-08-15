FROM zyfdedh/unqlite

COPY bin/* /usr/local/bin/

CMD ["/usr/local/bin/startup-update"]
