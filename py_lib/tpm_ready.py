#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import logging,platform,os

log = logging.getLogger('./tpm_test')

# check tpm status var
os_name=""
os_arch=""
os_platform=""
tpm_status=""


# tpm does not support
# ubt18
ubt_version="hostnamectl  |grep  'Operating System'|awk -F':' '{print $2}'"
ubt_kernel_version="hostnamectl  |grep  'Kernel' |awk -F':' '{print $2}'"

# Tpm exists

def GetTpm():
    ps = platform.system()
    if (ps == "Linux"):
        #ts=os.path.isfile("/dev/tpm0")
        ts=os.path.exists("/dev/tpm0")
        return ts
#    return False
#        if (tpm0==False):
#            tpm_status = "[INFO] 未安装物理tpm 芯片"
#        else:
#            tpm_status = "[INFO] 已安装物理tpm 芯片"
#    else:
#        tpm_status = "[INFO] 非LINUX系统"

def GetPlatform():
    pp=platform.platform()
    return pp.lower()

def TestPlatform():
    print("----------Operation System--------------------------")
    #  获取Python版本
    #print
    pv=platform.python_version()
    print(pv)

    #   获取操作系统可执行程序的结构，，(’32bit’, ‘WindowsPE’)
    #print
    pa=platform.architecture()
    print(pa)

    #   计算机的网络名称，’acer-PC’
    #print
    pn=platform.node()
    print(pn)

    # 获取操作系统名称及版本号，’Windows-7-6.1.7601-SP1′
    print
    pp=platform.platform()
    print(pp)

    # 计算机处理器信息，’Intel64 Family 6 Model 42 Stepping 7, GenuineIntel’
    print
    ppr=platform.processor()
    print(ppr)

    # 获取操作系统中Python的构建日期
    print
    ppb=platform.python_build()
    print(ppb)

    #  获取系统中python解释器的信息
    print
    pc=platform.python_compiler()
    print(pc)

    if platform.python_branch() == "":
        print
        ppi=platform.python_implementation()
        print(ppi)
        print
        ppr=platform.python_revision()
        print(ppr)
    print
    pr=platform.release()
    print(pr)
    print
    ps=platform.system()
    print(ps)

    # print platform.system_alias()
    #  获取操作系统的版本
    print
    pvs=platform.version()
    print(pvs)

    #  包含上面所有的信息汇总
    print
    pu=platform.uname()
    print(pu)


def UsePlatform():
    sysstr = platform.system()
    if (sysstr == "Windows"):
        print("Call Windows tasks")
    elif (sysstr == "Linux"):
        print("Call Linux tasks")
        os.system(ubt_version)
        os.system(ubt_kernel_version)
    else:
        print("Other System tasks")

class PynisaHandler:
    def __init__(self):
        self.thread=None
        #self.callintev=self.dconf['system']['callintev']+5
        self.lastcall=0
        self.currentcall=None

    def sayHello(self):
        print("Hello ,body!')!!")
        return "Hello,world"



if __name__ == '__main__':
    # handler_ph = PynisaHandler()
    # handler_ph.sayHello()
    # TestPlatform()
    #UsePlatform()
    try:
        if (GetTpm() == False):
            tpm_status = "[INFO] 未安装物理tpm芯片或非LINUX系统"
            print(tpm_status)
            print("[INFO] 操作系统:%s" % platform.system())
        else:
            tpm_status = "[INFO] 已安装物理tpm芯片"
            print(tpm_status)
            print("[INFO] 操作系统:%s" % platform.system())

        TpmOk = ["ubuntu", "suse", "kylin", "debian", "centos", "redhat"]
        kernel_op = 3
        os_kernel = int(platform.release()[0:1])

        if (any(tok in GetPlatform() for tok in TpmOk) == True):
            os_platform = GetPlatform()
            if (os_kernel >= kernel_op):
                platform.release()
                print("[INFO] 可以安装模拟器,当前内核版本:%s" % platform.release())
                print("[INFO] 主机系统:%s" % os_platform)
            else:
                print("[INFO] 主机系统:%s,内核需要升级" % os_platform)
        else:
            os_platform = GetPlatform()
            print("[INFO] 暂时不支持您的版本")
            print("[INFO] 主机系统:%s" % os_platform)
    except ValueError:
        print("[ERROR] 请升级glibc及python3版本,或联系管理员")




