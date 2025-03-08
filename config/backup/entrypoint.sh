#!/bin/bash

printenv | grep -E 'DB_' > /etc/environment
cron -f
