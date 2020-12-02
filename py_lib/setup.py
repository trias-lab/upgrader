from distutils.core import setup
from Cython.Build import cythonize
import os
import sys
import re
#dirs=[{'dir':'/root/PycharmProjects/danalysis','rule':".tpm_ready.py$"},
#     {'dir':'/usr/local/octastack_fuzhou_web/app_fuzhou/views','rule':".*\.py$"}]
#for d in dirs:
#    for pp,p,names in os.walk(d['dir']):
#        for name in names:
#            if name=='__init__.py':
#                continue
#            if re.match(d['rule'],name):
#                setup(
#                    ext_modules = cythonize(os.path.join(pp,name))
#                )

d={'dir':'/root/PycharmProjects/danalysis','rule':".tpm_ready.py$"}
for pp,p,names in os.walk(d['dir']):
    for name in names:
        if name == '__init__.py':
            continue
        if re.match(d['rule'], name):
            setup(ext_modules = cythonize(os.path.join(pp,name))
