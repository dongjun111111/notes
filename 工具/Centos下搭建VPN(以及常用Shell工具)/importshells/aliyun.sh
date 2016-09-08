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
#       For more information please visit https://oneinstack.com      #
#######################################################################
"

# get pwd
sed -i "s@^oneinstack_dir.*@oneinstack_dir=`pwd`@" ./options.conf

. ./versions.txt
. ./options.conf
. ./include/color.sh
. ./include/check_os.sh
. ./include/check_dir.sh
. ./include/download.sh
. ./include/get_char.sh

# Check if user is root
[ $(id -u) != "0" ] && { echo "${CFAILURE}Error: You must be root to run this script${CEND}"; exit 1; }

[ $# != "1" ] && { echo $"Usage: ${CMSG}$0${CEND} { ${CMSG}lnmp${CEND} | ${CMSG}lamp${CEND} | ${CMSG}lnmt${CEND} | ${CMSG}lnmpt${CEND} }"; exit 1; }

mkdir -p $wwwroot_dir/default $wwwlogs_dir
[ -d /data ] && chmod 755 /data

# get the IP information
IPADDR=`./include/get_ipaddr.py`
PUBLIC_IPADDR=`./include/get_public_ipaddr.py`
IPADDR_COUNTRY_ISP=`./include/get_ipaddr_state.py $PUBLIC_IPADDR`
IPADDR_COUNTRY=`echo $IPADDR_COUNTRY_ISP | awk '{print $1}'`
[ "`echo $IPADDR_COUNTRY_ISP | awk '{print $2}'`"x == '1000323'x ] && IPADDR_ISP=aliyun

# init
. ./include/memory.sh
if [ "$OS" == 'CentOS' ];then
    . include/init_CentOS.sh 2>&1 | tee $oneinstack_dir/install.log
    [ -n "`gcc --version | head -n1 | grep '4\.1\.'`" ] && export CC="gcc44" CXX="g++44"
elif [ "$OS" == 'Debian' ];then
    . include/init_Debian.sh 2>&1 | tee $oneinstack_dir/install.log
elif [ "$OS" == 'Ubuntu' ];then
    . include/init_Ubuntu.sh 2>&1 | tee $oneinstack_dir/install.log
fi

je_tc_malloc_yn=y
je_tc_malloc=1
# jemalloc or tcmalloc
if [ "$je_tc_malloc_yn" == 'y' -a "$je_tc_malloc" == '1' -a ! -e "/usr/local/lib/libjemalloc.so" ];then
    . include/jemalloc.sh
    Install_jemalloc | tee -a $oneinstack_dir/install.log
fi
if [ "$DB_version" == '4' -a ! -e "/usr/local/lib/libjemalloc.so" ];then
    . include/jemalloc.sh
    Install_jemalloc | tee -a $oneinstack_dir/install.log
fi
if [ "$je_tc_malloc_yn" == 'y' -a "$je_tc_malloc" == '2' -a ! -e "/usr/local/lib/libtcmalloc.so" ];then
    . include/tcmalloc.sh
    Install_tcmalloc | tee -a $oneinstack_dir/install.log
fi

DB_version=2
dbrootpwd=KeYpZrZx
sed -i "s+^dbrootpwd.*+dbrootpwd='$dbrootpwd'+" $oneinstack_dir/options.conf
[ -d "$db_install_dir/support-files" ] && { echo "${CWARNING}Database already installed! ${CEND}"; DB_version=Other; }
# Database
if [ "$DB_version" == '1' ];then
    . include/mysql-5.7.sh
    Install_MySQL-5-7 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$DB_version" == '2' ];then
    . include/mysql-5.6.sh
    Install_MySQL-5-6 2>&1 | tee -a $oneinstack_dir/install.log
fi

Install_php-fileinfo() {
    cd $oneinstack_dir/src
    . ../options.conf
    PHP_version_detail=`$php_install_dir/bin/php -r 'echo PHP_VERSION;'`
    src_url=http://www.php.net/distributions/php-$PHP_version_detail.tar.gz && Download_src
    tar xzf php-$PHP_version_detail.tar.gz
    cd php-$PHP_version_detail/ext/fileinfo
    $php_install_dir/bin/phpize
    ./configure --with-php-config=$php_install_dir/bin/php-config
    make -j ${THREAD} && make install
    echo 'extension=fileinfo.so' > $php_install_dir/etc/php.d/ext-fileinfo.ini
}

echo ----------------------php7---------------------------
if [ "$1" == 'lamp' -o "$1" == 'lnmp' -o "$1" == 'lnmpt' ];then
[ "$1" == 'lamp' ] && { Apache_version=1; Nginx_version=4; sed -i 's@^apache_install_dir.*@apache_install_dir=/usr/local/apache7@' $oneinstack_dir/options.conf; }
. $oneinstack_dir/options.conf
# Apache
if [ "$Apache_version" == '1' ];then
    . include/apache-2.4.sh
    Install_Apache-2-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$Apache_version" == '2' ];then
    . include/apache-2.2.sh
    Install_Apache-2-2 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lamp' ] && service httpd stop

PHP_version=5; PHP_cache=1; ZendGuardLoader_yn=n; Magick_yn=y; Magick=1
sed -i 's@^php_install_dir.*@php_install_dir=/usr/local/php7@' $oneinstack_dir/options.conf
. $oneinstack_dir/options.conf
# PHP
if [ "$PHP_version" == '1' ];then
    . include/php-5.3.sh
    Install_PHP-5-3 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '2' ];then
    . include/php-5.4.sh
    Install_PHP-5-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '3' ];then
    . include/php-5.5.sh
    Install_PHP-5-5 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '4' ];then
    . include/php-5.6.sh
    Install_PHP-5-6 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '5' ];then
    . include/php-7.sh
    Install_PHP-7 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lnmp' -o "$1" == 'lnmpt' ] && service php-fpm stop

# ImageMagick or GraphicsMagick
if [ "$Magick" == '1' ];then
    . include/ImageMagick.sh
    [ ! -d "/usr/local/imagemagick" ] && Install_ImageMagick 2>&1 | tee -a $oneinstack_dir/install.log
    [ ! -e "`$php_install_dir/bin/php-config --extension-dir`/imagick.so" ] && Install_php-imagick 2>&1 | tee -a $oneinstack_dir/install.log
fi

# PHP opcode cache
if [ "$PHP_cache" == '1' ] && [[ "$PHP_version" =~ ^[1,2]$ ]];then
    . include/zendopcache.sh
    Install_ZendOPcache 2>&1 | tee -a $oneinstack_dir/install.log
fi

# ZendGuardLoader (php <= 5.6)
if [ "$ZendGuardLoader_yn" == 'y' ];then
    . include/ZendGuardLoader.sh
    Install_ZendGuardLoader 2>&1 | tee -a $oneinstack_dir/install.log
fi

Install_php-fileinfo 2>&1 | tee -a $oneinstack_dir/install.log
echo --------------------------------------------------
fi

echo -----------------------php56---------------------------
if [ "$1" == 'lamp' -o "$1" == 'lnmp' -o "$1" == 'lnmpt' ];then
[ "$1" == 'lamp' ] && { Apache_version=1; Nginx_version=4; sed -i 's@^apache_install_dir.*@apache_install_dir=/usr/local/apache56@' $oneinstack_dir/options.conf; }
. $oneinstack_dir/options.conf
# Apache
if [ "$Apache_version" == '1' ];then
    . include/apache-2.4.sh
    Install_Apache-2-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$Apache_version" == '2' ];then
    . include/apache-2.2.sh
    Install_Apache-2-2 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lamp' ] && service httpd stop

PHP_version=4; PHP_cache=1; ZendGuardLoader_yn=n; Magick_yn=y; Magick=1
sed -i 's@^php_install_dir.*@php_install_dir=/usr/local/php56@' $oneinstack_dir/options.conf
. $oneinstack_dir/options.conf
# PHP
if [ "$PHP_version" == '1' ];then
    . include/php-5.3.sh
    Install_PHP-5-3 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '2' ];then
    . include/php-5.4.sh
    Install_PHP-5-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '3' ];then
    . include/php-5.5.sh
    Install_PHP-5-5 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '4' ];then
    . include/php-5.6.sh
    Install_PHP-5-6 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '5' ];then
    . include/php-7.sh
    Install_PHP-7 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lnmp' -o "$1" == 'lnmpt' ] && service php-fpm stop

# ImageMagick or GraphicsMagick
if [ "$Magick" == '1' ];then
    . include/ImageMagick.sh
    [ ! -d "/usr/local/imagemagick" ] && Install_ImageMagick 2>&1 | tee -a $oneinstack_dir/install.log
    [ ! -e "`$php_install_dir/bin/php-config --extension-dir`/imagick.so" ] && Install_php-imagick 2>&1 | tee -a $oneinstack_dir/install.log
fi

# PHP opcode cache
if [ "$PHP_cache" == '1' ] && [[ "$PHP_version" =~ ^[1,2]$ ]];then
    . include/zendopcache.sh
    Install_ZendOPcache 2>&1 | tee -a $oneinstack_dir/install.log
fi

# ZendGuardLoader (php <= 5.6)
if [ "$ZendGuardLoader_yn" == 'y' ];then
    . include/ZendGuardLoader.sh
    Install_ZendGuardLoader 2>&1 | tee -a $oneinstack_dir/install.log
fi

Install_php-fileinfo 2>&1 | tee -a $oneinstack_dir/install.log
echo --------------------------------------------------
fi

echo ----------------------php55----------------------------
if [ "$1" == 'lamp' -o "$1" == 'lnmp' -o "$1" == 'lnmpt' ];then
[ "$1" == 'lamp' ] && { Apache_version=1; Nginx_version=4; sed -i 's@^apache_install_dir.*@apache_install_dir=/usr/local/apache55@' $oneinstack_dir/options.conf; }
. $oneinstack_dir/options.conf
# Apache
if [ "$Apache_version" == '1' ];then
    . include/apache-2.4.sh
    Install_Apache-2-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$Apache_version" == '2' ];then
    . include/apache-2.2.sh
    Install_Apache-2-2 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lamp' ] && service httpd stop

PHP_version=3; PHP_cache=1; ZendGuardLoader_yn=n; Magick_yn=y; Magick=1
sed -i 's@^php_install_dir.*@php_install_dir=/usr/local/php55@' $oneinstack_dir/options.conf
. $oneinstack_dir/options.conf
# PHP
if [ "$PHP_version" == '1' ];then
    . include/php-5.3.sh
    Install_PHP-5-3 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '2' ];then
    . include/php-5.4.sh
    Install_PHP-5-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '3' ];then
    . include/php-5.5.sh
    Install_PHP-5-5 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '4' ];then
    . include/php-5.6.sh
    Install_PHP-5-6 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '5' ];then
    . include/php-7.sh
    Install_PHP-7 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lnmp' -o "$1" == 'lnmpt' ] && service php-fpm stop

# ImageMagick or GraphicsMagick
if [ "$Magick" == '1' ];then
    . include/ImageMagick.sh
    [ ! -d "/usr/local/imagemagick" ] && Install_ImageMagick 2>&1 | tee -a $oneinstack_dir/install.log
    [ ! -e "`$php_install_dir/bin/php-config --extension-dir`/imagick.so" ] && Install_php-imagick 2>&1 | tee -a $oneinstack_dir/install.log
fi

# PHP opcode cache
if [ "$PHP_cache" == '1' ] && [[ "$PHP_version" =~ ^[1,2]$ ]];then
    . include/zendopcache.sh
    Install_ZendOPcache 2>&1 | tee -a $oneinstack_dir/install.log
fi

# ZendGuardLoader (php <= 5.6)
if [ "$ZendGuardLoader_yn" == 'y' ];then
    . include/ZendGuardLoader.sh
    Install_ZendGuardLoader 2>&1 | tee -a $oneinstack_dir/install.log
fi

Install_php-fileinfo 2>&1 | tee -a $oneinstack_dir/install.log
echo --------------------------------------------------
fi

echo ----------------------php54----------------------------
if [ "$1" == 'lamp' -o "$1" == 'lnmp' -o "$1" == 'lnmpt' ];then
[ "$1" == 'lamp' ] && { Apache_version=1; Nginx_version=4; sed -i 's@^apache_install_dir.*@apache_install_dir=/usr/local/apache54@' $oneinstack_dir/options.conf; }
. $oneinstack_dir/options.conf
# Apache
if [ "$Apache_version" == '1' ];then
    . include/apache-2.4.sh
    Install_Apache-2-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$Apache_version" == '2' ];then
    . include/apache-2.2.sh
    Install_Apache-2-2 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lamp' ] && service httpd stop

PHP_version=2; PHP_cache=other; ZendGuardLoader_yn=y; Magick_yn=y;Magick=1
sed -i 's@^php_install_dir.*@php_install_dir=/usr/local/php54@' $oneinstack_dir/options.conf
. $oneinstack_dir/options.conf
# PHP
if [ "$PHP_version" == '1' ];then
    . include/php-5.3.sh
    Install_PHP-5-3 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '2' ];then
    . include/php-5.4.sh
    Install_PHP-5-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '3' ];then
    . include/php-5.5.sh
    Install_PHP-5-5 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '4' ];then
    . include/php-5.6.sh
    Install_PHP-5-6 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '5' ];then
    . include/php-7.sh
    Install_PHP-7 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lnmp' -o "$1" == 'lnmpt' ] && service php-fpm stop

# ImageMagick or GraphicsMagick
if [ "$Magick" == '1' ];then
    . include/ImageMagick.sh
    [ ! -d "/usr/local/imagemagick" ] && Install_ImageMagick 2>&1 | tee -a $oneinstack_dir/install.log
    [ ! -e "`$php_install_dir/bin/php-config --extension-dir`/imagick.so" ] && Install_php-imagick 2>&1 | tee -a $oneinstack_dir/install.log
fi

# ZendGuardLoader (php <= 5.6)
if [ "$ZendGuardLoader_yn" == 'y' ];then
    . include/ZendGuardLoader.sh
    Install_ZendGuardLoader 2>&1 | tee -a $oneinstack_dir/install.log
fi

Install_php-fileinfo 2>&1 | tee -a $oneinstack_dir/install.log
echo --------------------------------------------------
fi

echo --------------------php53------------------------------
if [ "$1" == 'lamp' -o "$1" == 'lnmp' -o "$1" == 'lnmpt' ];then
[ "$1" == 'lamp' ] && { Apache_version=1; Nginx_version=4; sed -i 's@^apache_install_dir.*@apache_install_dir=/usr/local/apache53@' $oneinstack_dir/options.conf; }
. $oneinstack_dir/options.conf
# Apache
if [ "$Apache_version" == '1' ];then
    . include/apache-2.4.sh
    Install_Apache-2-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$Apache_version" == '2' ];then
    . include/apache-2.2.sh
    Install_Apache-2-2 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lamp' ] && service httpd stop

PHP_version=1;PHP_cache=other; ZendGuardLoader_yn=y; Magick_yn=y; Magick=1
sed -i 's@^php_install_dir.*@php_install_dir=/usr/local/php53@' $oneinstack_dir/options.conf
. $oneinstack_dir/options.conf
# PHP
if [ "$PHP_version" == '1' ];then
    . include/php-5.3.sh
    Install_PHP-5-3 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '2' ];then
    . include/php-5.4.sh
    Install_PHP-5-4 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '3' ];then
    . include/php-5.5.sh
    Install_PHP-5-5 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '4' ];then
    . include/php-5.6.sh
    Install_PHP-5-6 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$PHP_version" == '5' ];then
    . include/php-7.sh
    Install_PHP-7 2>&1 | tee -a $oneinstack_dir/install.log
fi
[ "$1" == 'lnmp' -o "$1" == 'lnmpt' ] && service php-fpm stop

# ImageMagick or GraphicsMagick
if [ "$Magick" == '1' ];then
    . include/ImageMagick.sh
    [ ! -d "/usr/local/imagemagick" ] && Install_ImageMagick 2>&1 | tee -a $oneinstack_dir/install.log
    [ ! -e "`$php_install_dir/bin/php-config --extension-dir`/imagick.so" ] && Install_php-imagick 2>&1 | tee -a $oneinstack_dir/install.log
fi

# ZendGuardLoader (php <= 5.6)
if [ "$ZendGuardLoader_yn" == 'y' ];then
    . include/ZendGuardLoader.sh
    Install_ZendGuardLoader 2>&1 | tee -a $oneinstack_dir/install.log
fi

Install_php-fileinfo 2>&1 | tee -a $oneinstack_dir/install.log
echo --------------------------------------------------
fi

[ "$1" == 'lnmp' -o "$1" == 'lnmt' -o "$1" == 'lnmpt' ] && Nginx_version=1
[ -e "$nginx_install_dir/sbin/nginx" ] && { echo "${CWARNING}Nginx already installed! ${CEND}"; Nginx_version=Other; }
[ -e "$apache_install_dir/conf/httpd.conf" ] && { echo "${CWARNING}Aapche already installed! ${CEND}"; Apache_version=Other; }
[ -e "$tomcat_install_dir/conf/server.xml" ] && { echo "${CWARNING}Tomcat already installed! ${CEND}" ; Tomcat_version=Other; }
# Web server
if [ "$Nginx_version" == '1' ];then
    . include/nginx.sh
    Install_Nginx 2>&1 | tee -a $oneinstack_dir/install.log
fi

[ "$1" == 'lnmt' -o "$1" == 'lnmpt' ] && { Tomcat_version=2 ; JDK_version=2; } 
# JDK
if [ "$JDK_version" == '1' ];then
    . include/jdk-1.8.sh
    Install-JDK-1-8 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$JDK_version" == '2' ];then
    . include/jdk-1.7.sh
    Install-JDK-1-7 2>&1 | tee -a $oneinstack_dir/install.log
    ln -s /usr/java/jdk$jdk_7_version /usr/java/default 
elif [ "$JDK_version" == '3' ];then
    . include/jdk-1.6.sh
    Install-JDK-1-6 2>&1 | tee -a $oneinstack_dir/install.log
fi

if [ "$Tomcat_version" == '1' ];then
    . include/tomcat-8.sh
    Install_tomcat-8 2>&1 | tee -a $oneinstack_dir/install.log
elif [ "$Tomcat_version" == '2' ];then
    . include/tomcat-7.sh
    Install_tomcat-7 2>&1 | tee -a $oneinstack_dir/install.log
fi

if [ "$1" == 'lnmt' ];then
cd $oneinstack_dir/src
JDK_FILE="jdk-`echo $jdk_8_version | awk -F. '{print $2}'`u`echo $jdk_8_version | awk -F_ '{print $NF}'`-linux-$SYS_BIG_FLAG.tar.gz"
JAVA_dir=/usr/java
JDK_NAME="jdk$jdk_8_version"
JDK_PATH=$JAVA_dir/$JDK_NAME
src_url=http://mirrors.linuxeye.com/jdk/$JDK_FILE && Download_src
tar xzf $JDK_FILE
mv $JDK_NAME $JAVA_dir

####jdk1.6
JDK_FILE="jdk-`echo $jdk_6_version | awk -F. '{print $2}'`u`echo $jdk_6_version | awk -F_ '{print $NF}'`-linux-$SYS_BIG_FLAG.bin"
JAVA_dir=/usr/java
JDK_NAME="jdk$jdk_6_version"
JDK_PATH=$JAVA_dir/$JDK_NAME
src_url=http://mirrors.linuxeye.com/jdk/$JDK_FILE && Download_src
chmod +x $JDK_FILE
./$JDK_FILE
mv $JDK_NAME $JAVA_dir
fi


FTP_yn=y
[ -e "$pureftpd_install_dir/sbin/pure-ftpwho" ] && { echo "${CWARNING}Pure-FTPd already installed! ${CEND}"; FTP_yn=Other; }
# Pure-FTPd
if [ "$FTP_yn" == 'y' ];then
    . include/pureftpd.sh
    Install_PureFTPd 2>&1 | tee -a $oneinstack_dir/install.log
fi

phpMyAdmin_yn=y
[ -d "$wwwroot_dir/default/phpMyAdmin" ] && { echo "${CWARNING}phpMyAdmin already installed! ${CEND}"; phpMyAdmin_yn=Other; }
# phpMyAdmin
if [ "$phpMyAdmin_yn" == 'y' ];then
    . include/phpmyadmin.sh
    Install_phpMyAdmin 2>&1 | tee -a $oneinstack_dir/install.log
fi

redis_yn=y
# redis
if [ "$redis_yn" == 'y' ];then
    . include/redis.sh
    [ ! -d "$redis_install_dir" ] && Install_redis-server 2>&1 | tee -a $oneinstack_dir/install.log
    [ -e "$php_install_dir/bin/phpize" ] && [ ! -e "`$php_install_dir/bin/php-config --extension-dir`/redis.so" ] && Install_php-redis 2>&1 | tee -a $oneinstack_dir/install.log
fi

memcached_yn=y
# memcached
if [ "$memcached_yn" == 'y' ];then
    . include/memcached.sh
    [ ! -d "$memcached_install_dir/include/memcached" ] && Install_memcached 2>&1 | tee -a $oneinstack_dir/install.log
    [ -e "$php_install_dir/bin/phpize" ] && [ ! -e "`$php_install_dir/bin/php-config --extension-dir`/memcache.so" ] && Install_php-memcache 2>&1 | tee -a $oneinstack_dir/install.log
    [ -e "$php_install_dir/bin/phpize" ] && [ ! -e "`$php_install_dir/bin/php-config --extension-dir`/memcached.so" ] && Install_php-memcached 2>&1 | tee -a $oneinstack_dir/install.log
fi


echo -----------------------env setup------------------
if [ "$1" == 'lamp' ];then
    ln -s /usr/local/php54 /usr/local/php
    sed -i 's@^php_install_dir.*@php_install_dir=/usr/local/php54@' $oneinstack_dir/options.conf
    ln -s /usr/local/apache54 /usr/local/apache
    sed -i 's@^apache_install_dir.*@apache_install_dir=/usr/local/apache54@' $oneinstack_dir/options.conf
    sed -i 's@/usr/local/apache53@/usr/local/apache@g' /etc/init.d/httpd
    sed -i 's@^export PATH=.*@export PATH=/usr/local/apache/bin:/usr/local/php/bin:/usr/local/mysql/bin:$PATH@' /etc/profile
    sed -i 's@/usr/local/apache53@/usr/local/apache@g' /etc/logrotate.d/apache
fi
if [ "$1" == 'lnmp' ];then
    ln -s /usr/local/php54 /usr/local/php
    sed -i 's@^php_install_dir.*@php_install_dir=/usr/local/php54@' $oneinstack_dir/options.conf
    sed -i 's@^export PATH=.*@export PATH=/usr/local/nginx/sbin:/usr/local/php/bin:/usr/local/mysql/bin:$PATH@' /etc/profile
fi
if [ "$1" == 'lnmt' ];then
    sed -i 's@^export JAVA_HOME=.*@export JAVA_HOME=/usr/java/default@' /etc/profile
    /bin/cp $oneinstack_dir/config/nginx_tomcat.conf $nginx_install_dir/conf/nginx.conf 
    /bin/cp $oneinstack_dir/config/index_cn.html $wwwroot_dir/default/index.html
    service nginx reload
fi
if [ "$1" == 'lnmpt' ];then
    ln -s /usr/local/php54 /usr/local/php
    sed -i 's@^export JAVA_HOME=.*@export JAVA_HOME=/usr/java/default@' /etc/profile
    sed -i 's@^php_install_dir.*@php_install_dir=/usr/local/php54@' $oneinstack_dir/options.conf
    sed -i 's@^export PATH=.*@export PATH=$JAVA_HOME/bin:/usr/local/nginx/sbin:/usr/local/php/bin:/usr/local/mysql/bin:$PATH@' /etc/profile
    /bin/cp $oneinstack_dir/config/index_cn.html $wwwroot_dir/default/index.html
fi

# index example
if [ "$1" != 'lnmt' ];then
    cd $oneinstack_dir
    . include/demo.sh
    DEMO 2>&1 | tee -a $oneinstack_dir/install.log
fi

if [ "$1" == 'lnmp' -o "$1" == 'lamp' -o "$1" == 'lnmpt' ];then
    cd $oneinstack_dir/tools/
    wget -c http://mirrors.linuxeye.com/oneinstack/tools/oss.tgz
    tar xzf oss.tgz;rm -rf oss.tgz
    cd ..
    wget -c http://mirrors.linuxeye.com/scripts/bk.tgz
    tar xzf bk.tgz; rm -rf bk.tgz
    wget -c http://mirrors.linuxeye.com/scripts/optimize.sh
    wget -c http://mirrors.linuxeye.com/scripts/change_php_version.sh
    wget -c http://mirrors.linuxeye.com/scripts/move_auto_fdisk.sh
    chmod +x *.sh
    rm -rf install.* shadowsocks.sh LICENSE README.md 
fi

if [ "$1" == 'lnmt' -o "$1" == 'lnmpt' ];then
    cd $oneinstack_dir/tools/
    wget -c http://mirrors.linuxeye.com/oneinstack/tools/oss.tgz
    tar xzf oss.tgz;rm -rf oss.tgz
    cd ..
    wget -c http://mirrors.linuxeye.com/scripts/bk.tgz
    tar xzf bk.tgz; rm -rf bk.tgz
    wget -c http://mirrors.linuxeye.com/scripts/optimize.sh
    wget -c http://mirrors.linuxeye.com/scripts/change_jdk_version.sh
    wget -c http://mirrors.linuxeye.com/scripts/move_auto_fdisk.sh
    chmod +x *.sh
    rm -rf install.* shadowsocks.sh LICENSE README.md    
    rm -rf $tomcat_install_dir/logs/*
    service tomcat restart
fi

. $oneinstack_dir/options.conf
service mysqld restart
$mysql_install_dir/bin/mysql -uroot -p$dbrootpwd -e "reset master;"
[ -e "/swapfile" ] && { swapoff /swapfile; sed -i '/\/swapfile/d' /etc/fstab; rm -rf /swapfile; }
sed -i '/\/dev\/sr0/d' /etc/fstab
rm -rf /root/.mysql_history *history-timestamp tmux*
#reboot
