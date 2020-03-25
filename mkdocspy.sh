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

DOCERSC=`which pydoc3`
if [ -z "$DOCERSC" ];then
    echo "pls,setup pydoc parkage! pydoc3 >3.6 "
    exit 1
fi

#if [ ! -f "${basepath}/style.css" ]; then
#    echo "pls, download style.css or contact auth"
#    exit 1
#else
#    cp ./style.css ./docs/
#fi

#echo $dirlist

function Set_index(){
    cd ../
    if [ -d $gitname ]; then
	pydoc3 -w ${gitname}
        mv ${gitname}.html $basepath"/"docs/index.html
	sed -i 's@'${gitname}'\.@@'g $basepath"/"docs/index.html
    #    mkdir ./docs/${basepath}
    #    godoc -all  -html ./ >./docs/index.html
    fi
}

function Set_doc(){
    cd $basepath
    #dirilist=`find ./ |grep -v \.git |grep -v .idea |grep -v doc |grep -v main |grep -v  __|sed 's@\./@@'`
    for i in  `find ./ |grep -v \.git |grep -v .idea |grep -v doc |grep -v main |grep -v  __|sed 's@\./@@'|grep -v .py`
    do
	pydoc3 -w ${i}
	mv *.html ./docs/
        #linkcss='<link type="text/css" rel="stylesheet" href="/'${gitname}'/style.css">'
        #sed -i 's/upgrader/${gitname}/g' $linkcss
    #if [ -d $basepath"/"${i} ]; then
    #    read_dir $basepath"/"${i}
    #    mkdir ./docs/${basepath}
    #    godoc -all  -html ./ >./docs/index.html
    #fi
    done
}

function Set_docPy(){
    cd $basepath
    #dirilist=`find ./ |grep -v \.git |grep -v .idea |grep -v doc |grep -v main |grep -v  __|sed 's@\./@@'`
    for i in  `find ./ |grep -v \.git |grep -v .idea |grep -v doc |grep -v main |grep -v  __|grep .py |awk -F "." 'NF{NF-=1};1' |sed 's@/@\.@'\g |sed 's/..//'`
    do
        pydoc3 -w ${i}
        mv *.html ./docs/
        #linkcss='<link type="text/css" rel="stylesheet" href="/'${gitname}'/style.css">'
        #sed -i 's/upgrader/${gitname}/g' $linkcss
    #if [ -d $basepath"/"${i} ]; then
    #    read_dir $basepath"/"${i}
    #    mkdir ./docs/${basepath}
    #    godoc -all  -html ./ >./docs/index.html
    #fi
    done
}

Set_index
Set_doc
Set_docPy

