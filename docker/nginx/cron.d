SHELL=/bin/sh
PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin
MAILTO=""

# m h dom mon dow user  command
7 *  *  *  *  root  logrotate -v /etc/logrotate.conf
#
