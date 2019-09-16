#!powershell

Connect-VIServer -Server 127.0.0.1 -User user -Password pass -Force
$vms = Get-VM

foreach($vm in $vms) {
    Write-Host $vm.Name
}
