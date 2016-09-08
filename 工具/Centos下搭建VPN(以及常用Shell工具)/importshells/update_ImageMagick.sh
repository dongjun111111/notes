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
#                  upgrade ImageMagick for OneinStack                 #
#       For more information please visit https://oneinstack.com      #
#######################################################################
"

. ./options.conf
. ./include/color.sh
. ./include/download.sh

ImageMagick_version=6.9.5-8
imagick_version=3.4.1

if [ -e "/usr/local/imagemagick/bin/convert" ];then
    OLD_ImageMagick_version=`/usr/local/imagemagick/bin/Magick-config --version | awk '{print $1}'`
else
    echo "${CWARNING}You do not have to install Imagemagick! ${CEND}"
    exit 1
fi

Stop_ImageMagick() {
if [ -e "$php_install_dir/etc/php.d/ext-imagick.ini" ];then
    /bin/mv $php_install_dir/etc/php.d/ext-imagick.ini{,_bk}
elif [ ! -e "$php_install_dir/etc/php.d/ext-imagick.ini" -a -n "`grep imagick.so $php_install_dir/etc/php.ini`" ];then
    sed -i 's@extension.*imagick.so.*@;&@' $php_install_dir/etc/php.ini
fi
[ -e "$apache_install_dir/conf/httpd.conf" ] && service httpd restart || service php-fpm restart
/bin/mv /usr/local/imagemagick{,_`date +"%Y%m%d_%H%M%S"`}
}

Start_ImageMagick() {
if [ -e "$php_install_dir/etc/php.d/ext-imagick.ini_bk" ];then
    /bin/mv $php_install_dir/etc/php.d/ext-imagick.ini{_bk,}
elif [ ! -e "$php_install_dir/etc/php.d/ext-imagick.ini" -a -n "`grep imagick.so $php_install_dir/etc/php.ini`" ];then
    sed -i 's@;extension.*imagick.so.*@extension=imagick.so@' /usr/local/php/etc/php.ini
fi
[ -e "$apache_install_dir/conf/httpd.conf" ] && service httpd restart || service php-fpm restart
}

Check_ImageMagick() {
if [ -n "`/usr/local/imagemagick/bin/convert -version | grep "$ImageMagick_version"`" ];then
    echo "You have ${CMSG}successfully${CEND} upgrade from ${CWARNING}$OLD_ImageMagick_version${CEND} to ${CWARNING}$ImageMagick_version${CEND}"
else
    echo "${CWARNING}Imagemagick upgrade failed! ${CEND}"
fi
}

Install_ImageMagick() {
cd $oneinstack_dir/src
src_url=http://mirrors.linuxeye.com/oneinstack/src/ImageMagick-$ImageMagick_version.tar.gz && Download_src

tar xzf ImageMagick-$ImageMagick_version.tar.gz
cd ImageMagick-$ImageMagick_version
./configure --prefix=/usr/local/imagemagick --enable-shared --enable-static
make && make install
cd ..
rm -rf ImageMagick-$ImageMagick_version
cd ..
}

Install_php-imagick() {
cd $oneinstack_dir/src
if [ -e "$php_install_dir/bin/phpize" ];then
    if [ "`$php_install_dir/bin/php -r 'echo PHP_VERSION;' | awk -F. '{print $1"."$2}'`" == '5.3' ];then
        src_url=http://mirrors.linuxeye.com/oneinstack/src/imagick-3.3.0.tgz && Download_src
        tar xzf imagick-3.3.0.tgz
        cd imagick-3.3.0
    else
        src_url=http://mirrors.linuxeye.com/oneinstack/src/imagick-$imagick_version.tgz && Download_src
        tar xzf imagick-$imagick_version.tgz
        cd imagick-$imagick_version
    fi
    make clean
    export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
    $php_install_dir/bin/phpize
    ./configure --with-php-config=$php_install_dir/bin/php-config --with-imagick=/usr/local/imagemagick
    make && make install
    cd ..
    rm -rf imagick-$imagick_version
fi
cd ..
}

Stop_ImageMagick
Install_ImageMagick
Install_php-imagick
Start_ImageMagick
Check_ImageMagick
