#!/bin/sh

# start shadowsockets
echo "starting shadowsockets server..."
/etc/init.d/shadowsocks start
echo "starting shadowsockets done"

# start ssinfo server
/shadowsocks/ssinfo