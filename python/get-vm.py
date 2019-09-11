#!/usr/bin/env python
from pyVim.connect import SmartConnect, Disconnect
from pyVmomi import vim, vmodl
import ssl
import atexit
import sys

def main():
    context = None
    if hasattr(ssl, '_create_unverified_context'):
        context = ssl._create_unverified_context()

    si = SmartConnect(host = '127.0.0.1',
                      user = 'user',
                      pwd = 'pass',
                      sslContext = context)

    atexit.register(Disconnect, si)
    content = si.content

    try:
        mobs = content.viewManager.CreateContainerView(content.rootFolder,
                                                       [vim.VirtualMachine],
                                                       True)
    except Exception as e:
        print(e)
        sys.exit(1)

    if mobs:
        for mob in mobs.view:
            print(mob.name)
    else:
        sys.exit(1)


if __name__ == '__main__':
    main()
