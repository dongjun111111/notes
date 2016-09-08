#!/bin/bash
sed -i 's@ajax.googleapis.com@ajax.lug.ustc.edu.cn@g' wp-includes/script-loader.php
sed -i 's@fonts.googleapis.com@fonts.lug.ustc.edu.cn@g' wp-includes/script-loader.php
sed -i 's@fonts.googleapis.com@fonts.lug.ustc.edu.cn@g' wp-includes/js/tinymce/plugins/compat3x/css/dialog.css 
sed -i 's@fonts.googleapis.com@fonts.lug.ustc.edu.cn@g' wp-admin/includes/class-wp-press-this.php 
