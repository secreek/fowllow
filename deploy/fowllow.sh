#!/bin/sh
THIN=/usr/local/bin/thin
SCRIPT_NAME=fowllow
# 根据实际部署情况修改
# 缺省设置供测试使用

THIN_CONFIG=/home/fowllow/fowllow.yml
FOWLLOW_CONFIG=/home/fowllow/fowllow/src/backend/config.ru
NOHUP=/usr/bin/nohup

[ -x "$THIN" ] || exit 0
echo $1
case "$1" in
	start)
		$THIN -C $THIN_CONFIG -R $FOWLLOW_CONFIG start
		;;
	stop)
		$THIN -C $THIN_CONFIG -R $FOWLLOW_CONFIG stop
		;;
	restart)
		$THIN -C $THIN_CONFIG -R $FOWLLOW_CONFIG restart
		;;
	*)
		echo "Usage : $SCRIPT_NAME {start|stop|restart}" >&2
		exit 3
		;;
esac

