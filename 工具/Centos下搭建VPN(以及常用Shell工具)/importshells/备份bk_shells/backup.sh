#!/bin/bash
# Author:  yeho <lj2007331 AT gmail.com>
# BLOG:  https://blog.linuxeye.com
#
# Notes: OneinStack for CentOS/RadHat 5+ Debian 6+ and Ubuntu 12+
#
# Project home page:
#       https://oneinstack.com
#       https://github.com/lj2007331/oneinstack

cd tools
. ../options.conf

DB_Local_BK() {
    for D in `echo $db_name | tr ',' ' '`
    do
        ./db_bk.sh $D
    done
}

DB_Remote_BK() {
    for D in `echo $db_name | tr ',' ' '`
    do
        ./db_bk.sh $D
        DB_GREP="DB_${D}_`date +%Y`"
        DB_FILE=`ls -lrt $backup_dir | grep ${DB_GREP} | tail -1 | awk '{print $NF}'`
        echo "file:::$backup_dir/$DB_FILE $backup_dir push" >> config_bakcup.txt
        echo "com:::[ -e "$backup_dir/$DB_FILE" ] && rm -rf $backup_dir/DB_${D}_$(date +%Y%m%d --date="$expired_days days ago")_*.tgz" >> config_bakcup.txt
    done
}

DB_OSS_BK() {
    for D in `echo $db_name | tr ',' ' '`
    do
        ./db_bk.sh $D
        DB_GREP="DB_${D}_`date +%Y`"
        DB_FILE=`ls -lrt $backup_dir | grep ${DB_GREP} | tail -1 | awk '{print $NF}'`
        ./osscmd mkdir oss://$oss_bucket/`date +%F`
        ./osscmd put $backup_dir/$DB_FILE oss://$oss_bucket/`date +%F`/
        [ $? -eq 0 ] && ./osscmd --force=true deleteallobject oss://$oss_bucket/`date +%F --date="$expired_days days ago"`
    done
}

WEB_Local_BK() {
    for W in `echo $website_name | tr ',' ' '`
    do
        ./website_bk.sh $W
    done
}

WEB_Remote_BK() {
    for W in `echo $website_name | tr ',' ' '`
    do
        if [ `du -sm "$wwwroot_dir/$WebSite" | awk '{print $1}'` -lt 1024 ];then
            ./website_bk.sh $W
            Web_GREP="Web_${W}_`date +%Y`"
            Web_FILE=`ls -lrt $backup_dir | grep ${Web_GREP} | tail -1 | awk '{print $NF}'`
            echo "file:::$backup_dir/$Web_FILE $backup_dir push" >> config_bakcup.txt
            echo "com:::[ -e "$backup_dir/$Web_FILE" ] && rm -rf $backup_dir/Web_${W}_$(date +%Y%m%d --date="$expired_days days ago")_*.tgz" >> config_bakcup.txt
        else
            echo "file:::$wwwroot_dir/$W $backup_dir push" >> config_bakcup.txt
        fi
    done
}

WEB_OSS_BK() {
    for W in `echo $website_name | tr ',' ' '`
    do
        [ ! -e "$wwwroot_dir/$WebSite" ] && { echo "[$wwwroot_dir/$WebSite] not exist"; break; }
        Web_FILE="Web_${W}_$(date +%Y%m%d_%H).tgz"
        if [ ! -e "$backup_dir/$Web_FILE" ];then
            cd $wwwroot_dir
            tar czf $Web_FILE ./$W
            cd -
            PUSH_FILE="$wwwroot_dir/$Web_FILE"
        else
            PUSH_FILE="$backup_dir/$Web_FILE"
        fi

        ./osscmd mkdir oss://$oss_bucket/`date +%F`
        ./osscmd put $PUSH_FILE oss://$oss_bucket/`date +%F`/
        [ $? -eq 0 ] && { [ -e "$wwwroot_dir/$Web_FILE" ] && rm -rf $wwwroot_dir/$Web_FILE; ./osscmd --force=true deleteallobject oss://$oss_bucket/`date +%F --date="$expired_days days ago"`; }
    done
}

if [ "$backup_destination" == 'local' ];then
    [ -n "`echo $backup_content | grep -ow db`" ] && DB_Local_BK
    [ -n "`echo $backup_content | grep -ow web`" ] && WEB_Local_BK
elif [ "$backup_destination" == 'remote' ];then
    echo "com:::[ ! -e "$backup_dir" ] && mkdir -p $backup_dir" > config_bakcup.txt
    [ -n "`echo $backup_content | grep -ow db`" ] && DB_Remote_BK
    [ -n "`echo $backup_content | grep -ow web`" ] && WEB_Remote_BK
    ./mabs.sh -c config_bakcup.txt -T -1 | tee mabs.log
elif [ "$backup_destination" == 'oss' ];then
    [ -n "`echo $backup_content | grep -ow db`" ] && DB_OSS_BK
    [ -n "`echo $backup_content | grep -ow web`" ] && WEB_OSS_BK
elif [ "$backup_destination" == 'local,remote' ];then
    echo "com:::[ ! -e "$backup_dir" ] && mkdir -p $backup_dir" > config_bakcup.txt
    [ -n "`echo $backup_content | grep -ow db`" ] && DB_Local_BK
    [ -n "`echo $backup_content | grep -ow web`" ] && WEB_Local_BK
    [ -n "`echo $backup_content | grep -ow db`" ] && DB_Remote_BK
    [ -n "`echo $backup_content | grep -ow web`" ] && WEB_Remote_BK
    ./mabs.sh -c config_bakcup.txt -T -1 | tee mabs.log	
elif [ "$backup_destination" == 'local,oss' ];then
    [ -n "`echo $backup_content | grep -ow db`" ] && DB_Local_BK
    [ -n "`echo $backup_content | grep -ow db`" ] && DB_OSS_BK
    [ -n "`echo $backup_content | grep -ow web`" ] && WEB_Local_BK
    [ -n "`echo $backup_content | grep -ow web`" ] && WEB_OSS_BK
elif [ "$backup_destination" == 'remote,oss' ];then
    echo "com:::[ ! -e "$backup_dir" ] && mkdir -p $backup_dir" > config_bakcup.txt
    [ -n "`echo $backup_content | grep -ow db`" ] && DB_Remote_BK 
    [ -n "`echo $backup_content | grep -ow db`" ] && DB_OSS_BK
    [ -n "`echo $backup_content | grep -ow web`" ] && WEB_Remote_BK 
    [ -n "`echo $backup_content | grep -ow web`" ] && WEB_OSS_BK
    ./mabs.sh -c config_bakcup.txt -T -1 | tee mabs.log	
fi
