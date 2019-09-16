#!/bin/bash
#####create user for trias ubuntu verfiy
add_user(){
    grep ubuntu /etc/passwd > /dev/null
    if [ $? != 0 ];then
    /usr/sbin/groupadd ubuntu
    /usr/sbin/useradd ubuntu -g ubuntu -m 
    fi
    grep verfiy /etc/passwd > /dev/null
    if [ $? != 0 ];then
    /usr/sbin/useradd verfiy -g root -m
    fi    
}
#####install deb and pip
deb_install(){
grep "octa18 test" /etc/apt/sources.list > /dev/null
if [ $? != 0 ];then
cat >> /etc/apt/sources.list<<EOF
deb [trusted=yes] http://${octa_apt}/ octa18 test 
EOF
fi
    echo "[INFO: is installing common deb]"
    apt-get update &> /dev/null 
    apt-get install -y openssl python3-pip 8lab-zeromq4  
    echo "[INFO: installing pypi]"
    pip3 install -r requirements.txt > /dev/null
}

####set ima
setup_ima(){
    grep "ima_tcb"  /boot/grub/grub.cfg > /dev/null
    if [ $? != 0  ];then
	echo "[INFO:Setting up ima]"
        sed -i "/linux\t/s/$/& ima_tcb ima_template=\"ima\" ima_hash=\"sha1\"/g" /boot/grub/grub.cfg
    fi
    echo "[INFO:Ima has been added successfully, please remember to restart the host]"
}
#####dirt and start service
setup_trias(){
echo "[INFO:Set the trias directory structure]"	
tar xzf trias.tar.gz -C / 
chown -R ubuntu:ubuntu /trias
chown -R verfiy:root /8lab


}
#chmod chown
#####download binary
update_trias(){
trias_binary=(tendermint trias_accs triascode_app)
    for i in "${trias_binary[@]}" 
#    for i in ${trias_binary[*]} 
    do 
        wget -c -t 0 http://${octa_apt}/packs/files/$i -P /usr/local/bin/ &>/dev/null
	chmod +x /usr/local/bin/$i
    done

trias_binary2=(txmodule attestation)
    for i in ${trias_binary2[@]}
    do 
#        wget -c -t 0 -r -nd -np -nH   -p --level=3 -E -R html  http://${octa_apt}/packs/files/$i/ -P /$i	    
        wget -c -t 0  http://${octa_apt}/packs/files/$i.tar.gz &>/dev/null
	tar xzf $i.tar.gz -C / 
    done
trias_black=(blackbox blackbox_agent) 
    for i in ${trias_black[@]}
    do 
        wget -c -t 0 http://${octa_apt}/packs/files/$i -P /8lab &>/dev/null
        chmod +x /8lab/$i	
    done
chown -R verfiy:root /attestation
chown -R ubuntu:ubuntu /txmodule
}

#####start service
restart_service(){
echo "setup service"
systemctl enable BlackBoxClientinit.service 
systemctl enable Triasinit.service 
}

##run
if [ -f install-trias.conf  ];then
    . install-trias.conf
    add_user
    deb_install
    setup_ima
    setup_trias
    update_trias
    restart_service

else
    echo "lose config"
    exit 1 
fi    
read -p "The installation is complete, Do you want to reboot now?[yes/no]:" input
if [ $input == "yes" ];then
    reboot
fi
	
