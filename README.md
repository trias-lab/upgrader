# Trias-Upgrader

[Trias](https://trias.one/)

[trias-upgrader](https://github.com/trias-lab/upgrader.git).

[![CircleCI](https://circleci.com/gh/vincent0330/tmware.svg?style=svg&circle-token=14ea248e5a95af37689a2c9bde7587b6669b9d74)](https://circleci.com/gh/vincent0330/tmware)
[![Tag](https://img.shields.io/badge/tag-v0.0.8-orange.svg)](https://github.com/trias-lab/gondwana/releases/tag/v0.0.8)
[![Go version](https://img.shields.io/badge/go-1.12.1-blue.svg)](https://github.com/moovweb/gvm)
[![license](https://img.shields.io/badge/license-GPLv3-green.svg)](https://github.com/trias-lab/gondwana/blob/master/LICENSE)


| Branch  |                                 Tests                                 |                              Coverage                               |
| ------- | --------------------------------------------------------------------- | ------------------------------------------------------------------- |
| master  | ![master test](https://img.shields.io/badge/test-passing-success.svg) | ![master cov](https://img.shields.io/badge/cov-58%25-critical.svg)  |
| develop | ![develp test](https://img.shields.io/badge/test-passing-success.svg) | ![develop cov](https://img.shields.io/badge/cov-58%25-critical.svg) |

In order to upgrade the node, the application of hard fork.
It also start services as a node installation and data synchronization.
You need to specify the branch and version number for each version.


## Resources
### Research
* [The latest paper](https://www.contact@trias.one/attachment/Trias-whitepaper%20attachments.zip)
* [Project process](https://trias.one/updates/project)
* [Original Whitepaper](https://trias.one/whitepaper)
* [News room](https://trias.one/updates/recent)

## Security

To report a security vulnerability,  [bug report](mailto:contact@trias.one)

# HOW TO USE

* Download the corresponding binary tool

* Manually build bin and exec.

exec command args:
  ./tupgrader  upgrade   
  ./tupgrader  genesis    
  ./tupgrader  check    
  ./tupgrader  ver    
  ./tupgrader  syncdata    
  ./tupgrader  new   
  ./tupgrader  clean

It is still in use in the test network.

  
| Branch  |           describe                      | 
| ------- | ----------------------------------------| 
| ------- | --------------------------------------- |
| upgrade | update trias server to lastest version. | 
| ------- | ----------------------------------------| 
| genesis | generate basic configuration.           | 
| ------- | ----------------------------------------| 
| check   | check trias server version at local.    | 
| ------- | ----------------------------------------| 
| ver     | show the current version.               | 
| ------- | ----------------------------------------| 
| syncdata | whether data is synchronized or not.   | 
| ------- | ----------------------------------------| 
| new     | star the new nodes for trias.           | 
| ------- | ----------------------------------------| 
| clean   | clean nodes bin and data for trias.     | 
| ------- | ----------------------------------------| 

# status

After execution, the relevant logs are recorded in the corresponding directory.
```
operation2019/09/07 18:10:35 logger.go:98: exec faild: pip3 install 
operation2019/09/07 18:10:35 logger.go:99: ------------------info-------------------------
operation2019/09/07 18:10:35 logger.go:36: tar zxvf /trias.tar.gzcmd exec failed!
operation2019/09/07 18:10:35 logger.go:37: <nil>
```