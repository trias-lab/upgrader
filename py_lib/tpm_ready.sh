#!/usr/bin/env bash

# tpm does not support
# ubt18
os_version=`hostnamectl  |grep  "Operating System" |awk -F':' '{print $2}'`
kernel_version=`hostnamectl  |grep  "Kernel" |awk -F':' '{print $2}'`
# Tpm exists
tpm_exists(){
if [ -f /dev/tpm0 ];then
   echo "[INFO] 已安装物理tpm 芯片"
   exit 0
fi
}
os_version(){
if [ -f /etc/SuSE-release ];then
   cat /etc/issue |grep SUSE|grep "12 SP3"
   if [ $? -eq 0 ];then
       #hostnamectl |grep Kernel |grep "4.4"
       echo $kernel_version |grep "4.4"
       if [ $? -eq 0 ];then
           echo "[INFO] suse 12 SP3 支持安装"
       fi
   fi
elif [ -f /etc/redhat-release ] || [ -f /etc/neokylin-release ] || [ -f /etc/centos-release ];then
    uname -r |grep "3.10"
    if [ $? -eq 0 ];then
        echo "[INFO] 可以安装模拟器"
    else
        echo "[INFO] 暂时不支持您的版本哦"
        echo "当前操作系统为：$os_version 内核版本为：$kernel_version"
    fi
## ubt18
elif [ -f /etc/issue ];then
   cat /etc/issue |grep Ubuntu|grep "18.04"
   if [ $? -eq 0 ];then
        #hostnamectl |grep Kernel |grep 4.15 > /dev/null
        echo $kernel_version |grep "4.15"
       if [ $? -eq 0 ];then
           echo "[INFO] ubuntu 18.04 支持安装模拟器"
       fi
   fi
fi
}
tpm_exists
os_version
