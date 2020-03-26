#!/bin/bash
# fey

basepath=$(cd `dirname $0`; pwd)
#linkcss='<link type="text/css" rel="stylesheet" href="/upgrader/style.css">'
cd $basepath
gitname=`cat .git/config |grep url |awk -F'.' '{print $2}' |awk -F'/' '{print $NF}'`

if [ -z "$gitname" ]; then
    echo "pls, git clone project to the dir"
    exit 1
else
    if [ ! -d "./docs" ]; then
        mkdir docs
    fi
fi

DOCERSC=`which godoc`
if [ -z "$DOCERSC" ];then
    echo "pls,setup go parkage! version =1.11"
    exit 1
fi

if [ ! -f "${basepath}/style.css" ]; then
    echo "pls, download style.css or contact auth"
    exit 1
else
    cp ./style.css ./docs/
fi

#echo $dirlist

function Set_doc(){
    #dirilist=`find . -type d | while read dir; do echo $dir; done |grep -v \.git |grep -v .idea |grep -v doc`
    for i in  `find . -type d |grep -v \.git |grep -v .idea |grep -v doc`
    do
    mkdir -p ./docs/${i}
    godoc -all  -html ${i} >./docs/${i}/index.html
    linkcss='<link type="text/css" rel="stylesheet" href="/'${gitname}'/style.css">'
    #sed -i 's/upgrader/${gitname}/g' $linkcss
    echo $linkcss >>./docs/$i/index.html
    #if [ -d $basepath"/"${i} ]; then
    #    read_dir $basepath"/"${i}
    #    mkdir ./docs/${basepath}
    #    godoc -all  -html ./ >./docs/index.html
    #fi
    done
}

function Set_doc14(){
    #dirilist=`find . -type d | while read dir; do echo $dir; done |grep -v \.git |grep -v .idea |grep -v doc`
    godoc -url /pkg/${gitname} >./docs/index.html
    linkcss='<link type="text/css" rel="stylesheet" href="/'${gitname}'/style.css">'
    echo $linkcss >>./docs/index.html

    for i in  `find . -type d |grep -v \.git |grep -v .idea |grep -v doc |sed 's/..//' |awk 'NR!=1'`
    do
    mkdir -p docs/${i}
    godoc -url /pkg/${gitname}/${i} >./docs/${i}/index.html
    #linkcss='<link type="text/css" rel="stylesheet" href="/'${gitname}'/style.css">'
    echo $linkcss >>./docs/$i/index.html
    #if [ -d $basepath"/"${i} ]; then
    #    read_dir $basepath"/"${i}
    #    mkdir ./docs/${basepath}
    #    godoc -all  -html ./ >./docs/index.html
    #fi
    done
}

VER=`go version |awk -F "." '{print $2}'`
if [ $VER == "11" ]; then
    echo "This go version is 1.11"
    Set_doc
else
    echo "This go version is 1.14"
    Set_doc14
fi
#Set_doc
