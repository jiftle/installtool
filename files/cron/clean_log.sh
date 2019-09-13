#!/bin/bash

logfile="/data/logs/cron.log"
# -------- 输出日志 
echo `date "+%Y-%m-%d %H:%M:%S | --> start del log file."` >> ${logfile}

# 删除目录下过期日志文件
find /data/logs/ -type f -name "*.log*" -mtime +7 -exec rm -rf {} \;

# 添加执行权限
# chmod +x auto_del_log.sh

echo `date "+%Y-%m-%d %H:%M:%S | --> end del log file."`>> ${logfile}
