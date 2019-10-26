==========================

 Developing with Trias upgrader
==========================

Now that you have your nifty DevStack up and running, what can you do
with it?

Inspecting Services

|upgrade|   --Update trias server to lastest version.|

|genesis|   --Generate basic configuration.|

|new    |   --Star the new nodes for trias.|

|clean  |   --Clear the all files of the local node.|

|cdata  |   --Clear the all data of the local node.|


===================

By default most services in DevStack are running as `systemd` units
named `$servicename.service`. You can see running services
with.

.. code-block:: bash

   sudo systemctl status "trias@*"

To learn more about the basics of systemd, see :doc:`/systemd`

Get a Service
==================

If you want to get a upgrder to a running service on your trias node the easiest
way to do that is to change the code directly in Community trusted area.

.. code-block:: bash

   git clone https://github.com/trias-lab/upgrader.git
   
   or
   
   wget https://github.com/trias-lab/upgrader/raw/master/tupgrader
   

If your change impacts more than one funcation you can use args by
wildcard as well.

.. code-block:: bash

   exec command args: ./tupgrader $*

.. warning::

   All changes you are making are in checked out git trees that
   trias-upgrder thinks it has full control over. Uncommitted work, or
   work committed to the master branch, may be overwritten during
   subsequent bin runs.

Service design description
======================

![](https://github.com/trias-lab/upgrader/raw/master/doc/Trias-Upgrader.gif)


When testing a larger set of patches, or patches that will impact more
than one service within a project, it is often less confusing to use
custom git locations, and make all your changes in a dedicated git
tree.


Funcation  instance:

.. code-block:: func new()

   Create a new Trias node service to get the current version status.

   Ensure that the current version of the host is executed, it must be Ubuntu Server 18.04 LTS.
   
   Get the currently confirmed release version of the service in community trusted area.
   
   Unzip the corresponding installation path.
   
   Establish user accounts and usage rights.
   
   Get the seed nodes corresponding configuration.
   
   Strating  open service.
   ```
   systemctl status  Triasinit.service
   systemctl status  BlackBoxClientinit.service
   ```
   
   Get data synchronization from the seed nodes.
   
You can use this same approach to test patches that are up for review
in gerrit by using the ref name that gerrit assigns to each change.

.. code-block:: func upgrade()

   Ensure that the current version of the host is executed, it must be Ubuntu Server 18.04 LTS.
   
   Upgrade version that  the trias community to a clear.
   
   Output log after completion, waiting for restart service.
   
You can use this same approach to test patches that are up for review
In the case that some node accounts are abnormal, you need to execute cdata.

.. code-block:: func genesis()

   Ensure the current Genesis block version from trias community trusted.
   
   Generate current creation blocks.

   Ensure that the current version of the host is executed, it must be Ubuntu Server 18.04 LTS.
   
   Notification can turn on service, data synchronization.
   

.. code-block:: func cdata()

   Clear local ledger according to current version.

   Ensure that the current version of the host is executed, it must be Ubuntu Server 18.04 LTS.
   
Generally clear the data, start tracking data from new. You can use this same approach to test patches that are up for review.

.. code-block:: func clean()

   Uninstaller  Trias node service to loclhost.

   Ensure that the current version of the host is executed, it must be Ubuntu Server 18.04 LTS.
   
You can use this same approach to test patches that are up for review.

Testing Changes to Bin
============================

When testing changes to libraries consumed by trias services things are 
a little more complicated. By default we only test with released versions of
these libraries that are on test env.


As libraries are download by github, after you make any
local changes you will need to:

* download the Binary executable file. Define the Tupgrader10 version Upgrade at  blockheight of 10. the Tupgrader20 version Upgrade at blockheight of 20. 
```
wget https://github.com/trias-lab/upgrader/raw/master/tupgrader10
wget https://github.com/trias-lab/upgrader/raw/master/tupgrader20
```
* Ensure that the current version of the host is executed in trias localtesting env, it must be OS on Ubuntu Server 18.04 LTS. 


You can do that with on testing nodes.

* Consensus node upgrade to version 10

.. code-block:: bash
```
   ./tupgrader10 cdata
   ./tupgrader10 upgrade
   /etc/init.d/Trias start
   /etc/init.d/BlackBoxClient start
```

* Upgrade other nodes to version 10

```
   ./tupgrader10 cdata
   ./tupgrader10 upgrade
   /etc/init.d/Trias start
   /etc/init.d/BlackBoxClient start
   
```

* start up  transaction
```
http://192.168.1.141:46657/tri_bc_tx_commit?tx=%22aaaaa%22
```

* Verify Consensus node
```
http://192.168.1.141:46657/tri_block_info?height=11
```

* Verify other node
```
http://192.168.1.145:46657/tri_block_info?height=11
```

* Consensus node upgrade to version 20

.. code-block:: bash
```
   ./tupgrader20 upgrade
   /etc/init.d/Trias start
   /etc/init.d/BlackBoxClient start
```


* Upgrade other nodes to version 20

```
   ./tupgrader20 upgrade
   /etc/init.d/Trias start
   /etc/init.d/BlackBoxClient start
   
```

* Verify Consensus  node
```
http://192.168.1.141:46657/tri_block_info?height=21
```

* Verify other node
```
http://192.168.1.145:46657/tri_block_info?height=21
```
