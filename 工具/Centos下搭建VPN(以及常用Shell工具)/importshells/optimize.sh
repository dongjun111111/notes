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
export LANG=en_US.UTF-8

# Check if user is root
[ $(id -u) != "0" ] && { echo "${CFAILURE}Error: You must be root to run this script${CEND}"; exit 1; }

. ./options.conf
. ./include/check_dir.sh
. ./include/memory.sh
IPADDR=`./include/get_ipaddr.py`

if [ -d "$mysql_install_dir/support-files" ];then
    sed -i "s@max_connections.*@max_connections = $(($Mem/2))@" /etc/my.cnf
    if [ $Mem -le 1500 ];then
        sed -i 's@^thread_cache_size.*@thread_cache_size = 8@' /etc/my.cnf
        sed -i 's@^query_cache_size.*@query_cache_size = 8M@' /etc/my.cnf
        sed -i 's@^myisam_sort_buffer_size.*@myisam_sort_buffer_size = 8M@' /etc/my.cnf
        sed -i 's@^key_buffer_size.*@key_buffer_size = 8M@' /etc/my.cnf
        sed -i 's@^innodb_buffer_pool_size.*@innodb_buffer_pool_size = 64M@' /etc/my.cnf
        sed -i 's@^tmp_table_size.*@tmp_table_size = 16M@' /etc/my.cnf
        sed -i 's@^table_open_cache.*@table_open_cache = 128@' /etc/my.cnf
    elif [ $Mem -gt 1500 -a $Mem -le 2500 ];then
        sed -i 's@^thread_cache_size.*@thread_cache_size = 16@' /etc/my.cnf
        sed -i 's@^query_cache_size.*@query_cache_size = 16M@' /etc/my.cnf
        sed -i 's@^myisam_sort_buffer_size.*@myisam_sort_buffer_size = 16M@' /etc/my.cnf
        sed -i 's@^key_buffer_size.*@key_buffer_size = 16M@' /etc/my.cnf
        sed -i 's@^innodb_buffer_pool_size.*@innodb_buffer_pool_size = 128M@' /etc/my.cnf
        sed -i 's@^tmp_table_size.*@tmp_table_size = 32M@' /etc/my.cnf
        sed -i 's@^table_open_cache.*@table_open_cache = 256@' /etc/my.cnf
    elif [ $Mem -gt 2500 -a $Mem -le 3500 ];then
        sed -i 's@^thread_cache_size.*@thread_cache_size = 32@' /etc/my.cnf
        sed -i 's@^query_cache_size.*@query_cache_size = 32M@' /etc/my.cnf
        sed -i 's@^myisam_sort_buffer_size.*@myisam_sort_buffer_size = 32M@' /etc/my.cnf
        sed -i 's@^key_buffer_size.*@key_buffer_size = 64M@' /etc/my.cnf
        sed -i 's@^innodb_buffer_pool_size.*@innodb_buffer_pool_size = 512M@' /etc/my.cnf
        sed -i 's@^tmp_table_size.*@tmp_table_size = 64M@' /etc/my.cnf
        sed -i 's@^table_open_cache.*@table_open_cache = 512@' /etc/my.cnf
    elif [ $Mem -gt 3500 ];then
        sed -i 's@^thread_cache_size.*@thread_cache_size = 64@' /etc/my.cnf
        sed -i 's@^query_cache_size.*@query_cache_size = 64M@' /etc/my.cnf
        sed -i 's@^myisam_sort_buffer_size.*@myisam_sort_buffer_size = 64M@' /etc/my.cnf
        sed -i 's@^key_buffer_size.*@key_buffer_size = 256M@' /etc/my.cnf
        sed -i 's@^innodb_buffer_pool_size.*@innodb_buffer_pool_size = 1024M@' /etc/my.cnf
        sed -i 's@^tmp_table_size.*@tmp_table_size = 128M@' /etc/my.cnf
        sed -i 's@^table_open_cache.*@table_open_cache = 1024@' /etc/my.cnf
    fi
    service mysqld restart
fi

if [ -e "/usr/local/php/etc/php.ini" ];then
    sed -i "s@^memory_limit.*@memory_limit = ${Memory_limit}M@g" /usr/local/php{,53,54,55,56,7}/etc/php.ini
    [ -e "/usr/local/php/etc/php.d/ext-opcache.ini" ] && sed -i "s@^opcache.memory_consumption.*@opcache.memory_consumption=$Memory_limit@g" /usr/local/php{,53,54,55,56,7}/etc/php.d/ext-opcache.ini
    [ -e "$apache_install_dir" ] && service httpd restart
fi

if [ -e "/usr/local/php/etc/php-fpm.conf" ];then
    if [ $Mem -le 3000 ];then
        sed -i "s@^pm.max_children.*@pm.max_children = $(($Mem/3/20))@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.start_servers.*@pm.start_servers = $(($Mem/3/30))@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.min_spare_servers.*@pm.min_spare_servers = $(($Mem/3/40))@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.max_spare_servers.*@pm.max_spare_servers = $(($Mem/3/20))@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
    elif [ $Mem -gt 3000 -a $Mem -le 4500 ];then
        sed -i "s@^pm.max_children.*@pm.max_children = 50@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.start_servers.*@pm.start_servers = 30@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.min_spare_servers.*@pm.min_spare_servers = 20@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.max_spare_servers.*@pm.max_spare_servers = 50@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
    elif [ $Mem -gt 4500 -a $Mem -le 6500 ];then
        sed -i "s@^pm.max_children.*@pm.max_children = 60@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.start_servers.*@pm.start_servers = 40@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.min_spare_servers.*@pm.min_spare_servers = 30@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.max_spare_servers.*@pm.max_spare_servers = 60@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
    elif [ $Mem -gt 6500 -a $Mem -le 8500 ];then
        sed -i "s@^pm.max_children.*@pm.max_children = 70@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.start_servers.*@pm.start_servers = 50@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.min_spare_servers.*@pm.min_spare_servers = 40@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.max_spare_servers.*@pm.max_spare_servers = 70@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
    elif [ $Mem -gt 8500 ];then
        sed -i "s@^pm.max_children.*@pm.max_children = 80@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.start_servers.*@pm.start_servers = 60@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.min_spare_servers.*@pm.min_spare_servers = 50@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
        sed -i "s@^pm.max_spare_servers.*@pm.max_spare_servers = 80@g" /usr/local/php{,53,54,55,56,7}/etc/php-fpm.conf
    fi
    service php-fpm restart
fi

if [ -e "$tomcat_install_dir/bin/setenv.sh" ];then
    [ $Mem -le 768 ] && Xms_Mem=`expr $Mem / 3` || Xms_Mem=256
    cat > $tomcat_install_dir/bin/setenv.sh << EOF
JAVA_OPTS='-Djava.security.egd=file:/dev/./urandom -server -Xms${Xms_Mem}m -Xmx`expr $Mem / 2`m'
CATALINA_OPTS="-Djava.library.path=/usr/local/apr/lib"
#  -Djava.rmi.server.hostname=$IPADDR
#  -Dcom.sun.management.jmxremote.password.file=\$CATALINA_BASE/conf/jmxremote.password
# -Dcom.sun.management.jmxremote.access.file=\$CATALINA_BASE/conf/jmxremote.access
#  -Dcom.sun.management.jmxremote.ssl=false"
EOF
    service tomcat restart
fi
