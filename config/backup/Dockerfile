FROM postgres

RUN apt update && apt install -y cron

COPY backup.sh /usr/local/bin/backup.sh
RUN chmod +x /usr/local/bin/backup.sh

COPY crontab /etc/cron.d/backup-cron
RUN chmod 0644 /etc/cron.d/backup-cron

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]