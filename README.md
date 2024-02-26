启动
关闭
世界日志
面板日志
更新游戏
创建备份
更新模组
模组
监控服务器负载
远程控制台指令

#!/bin/sh
# 设置服务器目录为脚本所在目录
SERVERDIR=`dirname "$0"`
cd "$SERVERDIR"

# 获取传递给脚本的参数
PARAMS=$@

# 初始化配置文件变量
CONFIGFILE=

# 解析参数中的配置文件路径
while test $# -gt 0
do
    if [ `echo $1 | cut -c 1-12` = "-configfile=" ]; then
        CONFIGFILE=`echo $1 | cut -c 13-`
    fi
    shift
done

# 检查是否指定了配置文件
if [ "$CONFIGFILE" = "" ]; then
    echo "No config file specified. Call this script like this:"
    echo "  ./startserver.sh -configfile=serverconfig.xml"
    exit 1
else
    # 检查指定的配置文件是否存在
    if [ -f "$CONFIGFILE" ]; then
        echo Using config file: $CONFIGFILE
    else
        echo "Specified config file $CONFIGFILE does not exist."
        exit 1
    fi
fi

# 设置动态链接库路径
export LD_LIBRARY_PATH=.
# 设置 MALLOC_CHECK_ 环境变量
# export MALLOC_CHECK_=0

# 启动 7 Days to Die 服务器
./7DaysToDieServer.x86_64 -logfile $SERVERDIR/7DaysToDieServer_Data/output_log__`date +%Y-%m-%d__%H-%M-%S`.txt -quit -batchmode -nographics -dedicated $PARAMS
