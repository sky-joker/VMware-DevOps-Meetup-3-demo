#!powershell

Connect-VIServer -Server 127.0.0.1 -User user -Password pass -Force
Get-VM
