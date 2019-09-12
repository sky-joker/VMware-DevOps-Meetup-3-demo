#!/usr/bin/env ruby
require 'rbvmomi'

def main
  vim = RbVmomi::VIM.connect(host: '127.0.0.1',
                             user: 'user',
                             password: 'pass',
                             insecure: true)

  mobs = vim.serviceContent.viewManager.CreateContainerView({container: vim.rootFolder,
                                                             type: ['VirtualMachine'],
                                                             recursive: true})

  for vm in mobs.view do
    puts vm.name
  end

end

if __FILE__ == $0
  main
end