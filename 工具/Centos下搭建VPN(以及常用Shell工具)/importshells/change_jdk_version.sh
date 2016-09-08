#!/bin/bash
# Author:  yeho <lj2007331 AT gmail.com>
# BLOG:  https://blog.linuxeye.com
#
# Notes: OneinStack for CentOS/RadHat 5+ Debian 6+ and Ubuntu 12+
#
# Project home page:
#       https://oneinstack.com
#       https://github.com/lj2007331/oneinstack

export PATH=/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin
clear
printf "
#######################################################################
#       OneinStack for CentOS/RadHat 5+ Debian 6+ and Ubuntu 12+      #
#                       Change your PHP version                       #
#       For more information please visit https://oneinstack.com      #
#######################################################################
"

. ./include/color.sh
. ./versions.txt

echo
OLD_JDK_version=`/usr/java/default/bin/java -version 2>&1 |awk 'NR==1{ gsub(/"/,""); print $3 }'`
echo "Current JDK Version: ${CMSG}${OLD_JDK_version}${CEND}"

while :; do echo
    echo 'Please select a version of the PHP:'
    echo -e "\t${CMSG}1${CEND}. jdk-1.6"
    echo -e "\t${CMSG}2${CEND}. jdk-1.7"
    echo -e "\t${CMSG}3${CEND}. jdk-1.8"
    read -p "Please input a number:(Default 2 press Enter) " JDK_VERSION
    [ -z "$JDK_VERSION" ] && JDK_VERSION=2
    if [[ ! $JDK_VERSION =~ ^[1-3]$ ]];then
        echo "${CWARNING}input error! Please only input number 1,2,3${CEND}"
    else
        break
    fi
done
if [ "$JDK_VERSION" == '1' ];then
    [ "${OLD_JDK_version%.*}" == '1.6' ] && { echo "${CWARNING}The version you entered is the same as the current version${CEND}"; exit 1; }
    /etc/init.d/tomcat stop > /dev/null 2>&1;[ -n "ps -ef | grep java | grep -v grep" ] && killall java > /dev/null 2>&1;rm -rf /usr/java/default
    ln -s /usr/java/jdk${jdk_6_version} /usr/java/default
    sed -i "s@^export JAVA_HOME=.*@export JAVA_HOME=/usr/java/jdk${jdk_6_version}@" /etc/init.d/tomcat
    /etc/init.d/tomcat start > /dev/null 2>&1
    echo
    echo "You have ${CMSG}successfully${CEND} changed to ${CMSG}1.6${CEND}"
    echo
elif [ "$JDK_VERSION" == '2' ];then
    [ "${OLD_JDK_version%.*}" == '1.7' ] && { echo "${CWARNING}The version you entered is the same as the current version${CEND}"; exit 1; }
    /etc/init.d/tomcat stop > /dev/null 2>&1;[ -n "ps -ef | grep java | grep -v grep" ] && killall java > /dev/null 2>&1;rm -rf /usr/java/default
    ln -s /usr/java/jdk${jdk_7_version} /usr/java/default
    sed -i "s@^export JAVA_HOME=.*@export JAVA_HOME=/usr/java/jdk${jdk_7_version}@" /etc/init.d/tomcat
    /etc/init.d/tomcat start > /dev/null 2>&1
    echo
    echo "You have ${CMSG}successfully${CEND} changed to ${CMSG}1.7${CEND}"
    echo
elif [ "$JDK_VERSION" == '3' ];then
    [ "${OLD_JDK_version%.*}" == '1.8' ] && { echo "${CWARNING}The version you entered is the same as the current version${CEND}"; exit 1; }
    /etc/init.d/tomcat stop > /dev/null 2>&1;[ -n "ps -ef | grep java | grep -v grep" ] && killall java > /dev/null 2>&1;rm -rf /usr/java/default
    ln -s /usr/java/jdk${jdk_8_version} /usr/java/default
    sed -i "s@^export JAVA_HOME=.*@export JAVA_HOME=/usr/java/jdk${jdk_8_version}@" /etc/init.d/tomcat
    /etc/init.d/tomcat start > /dev/null 2>&1
    echo
    echo "You have ${CMSG}successfully${CEND} changed to ${CMSG}1.8${CEND}"
    echo
fi
