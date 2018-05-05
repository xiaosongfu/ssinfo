#!/bin/sh

# start shadowsockets
echo "starting shadowsockets server..."
/etc/init.d/shadowsocks start
echo "starting shadowsockets done"

# start ssinfo server
# 1 add exec permission
chmod +x /shadowsocks/ssinfo
# 2 start exec
/shadowsocks/ssinfo