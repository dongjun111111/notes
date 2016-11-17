#!/bin/sh
PRO_NAME="程序名称" 
while true;
do
NUM=`ps aux | grep PRO_NAME | grep -v grep |wc -l`
if [ "${NUM}" -lt 1 ]
then
echo "${PRO_NAME} was killed"
cd /home/go/src/程序名称          /*程序所在目录*/
./程序名称                        /*启动程序*/
else
echo "ok"
fi
sleep 3 
done

