# Sysctls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsAioMaxNr** | Pointer to **int32** | Sysctl setting fs.aio-max-nr. | [optional] 
**FsFileMax** | Pointer to **int32** | Sysctl setting fs.file-max. | [optional] 
**FsInotifyMaxUserWatches** | Pointer to **int32** | Sysctl setting fs.inotify.max_user_watches. | [optional] 
**FsNrOpen** | Pointer to **int32** | Sysctl setting fs.nr_open. | [optional] 
**KernelThreadsMax** | Pointer to **int32** | Sysctl setting kernel.threads-max. | [optional] 
**NetCoreNetdevMaxBacklog** | Pointer to **int32** | Sysctl setting net.core.netdev_max_backlog. | [optional] 
**NetCoreOptmemMax** | Pointer to **int32** | Sysctl setting net.core.optmem_max. | [optional] 
**NetCoreRmemDefault** | Pointer to **int32** | Sysctl setting net.core.rmem_default. | [optional] 
**NetCoreRmemMax** | Pointer to **int32** | Sysctl setting net.core.rmem_max. | [optional] 
**NetCoreSomaxconn** | Pointer to **int32** | Sysctl setting net.core.somaxconn. | [optional] 
**NetCoreWmemDefault** | Pointer to **int32** | Sysctl setting net.core.wmem_default. | [optional] 
**NetCoreWmemMax** | Pointer to **int32** | Sysctl setting net.core.wmem_max. | [optional] 
**NetIpv4IpLocalPortRange** | Pointer to **string** | Sysctl setting net.ipv4.ip_local_port_range. | [optional] 
**NetIpv4NeighDefaultGcThresh1** | Pointer to **int32** | Sysctl setting net.ipv4.neigh.default.gc_thresh1. | [optional] 
**NetIpv4NeighDefaultGcThresh2** | Pointer to **int32** | Sysctl setting net.ipv4.neigh.default.gc_thresh2. | [optional] 
**NetIpv4NeighDefaultGcThresh3** | Pointer to **int32** | Sysctl setting net.ipv4.neigh.default.gc_thresh3. | [optional] 
**NetIpv4TcpFinTimeout** | Pointer to **int32** | Sysctl setting net.ipv4.tcp_fin_timeout. | [optional] 
**NetIpv4TcpKeepaliveProbes** | Pointer to **int32** | Sysctl setting net.ipv4.tcp_keepalive_probes. | [optional] 
**NetIpv4TcpKeepaliveTime** | Pointer to **int32** | Sysctl setting net.ipv4.tcp_keepalive_time. | [optional] 
**NetIpv4TcpMaxSynBacklog** | Pointer to **int32** | Sysctl setting net.ipv4.tcp_max_syn_backlog. | [optional] 
**NetIpv4TcpMaxTwBuckets** | Pointer to **int32** | Sysctl setting net.ipv4.tcp_max_tw_buckets. | [optional] 
**NetIpv4TcpTwReuse** | Pointer to **bool** | Sysctl setting net.ipv4.tcp_tw_reuse. | [optional] 
**NetIpv4TcpkeepaliveIntvl** | Pointer to **int32** | Sysctl setting net.ipv4.tcp_keepalive_intvl. | [optional] 
**NetNetfilterNfConntrackBuckets** | Pointer to **int32** | Sysctl setting net.netfilter.nf_conntrack_buckets. | [optional] 
**NetNetfilterNfConntrackMax** | Pointer to **int32** | Sysctl setting net.netfilter.nf_conntrack_max. | [optional] 
**VmMaxMapCount** | Pointer to **int32** | Sysctl setting vm.max_map_count. | [optional] 
**VmSwappiness** | Pointer to **int32** | Sysctl setting vm.swappiness. | [optional] 
**VmVfsCachePressure** | Pointer to **int32** | Sysctl setting vm.vfs_cache_pressure. | [optional] 

## Methods

### NewSysctls

`func NewSysctls() *Sysctls`

NewSysctls instantiates a new Sysctls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSysctlsWithDefaults

`func NewSysctlsWithDefaults() *Sysctls`

NewSysctlsWithDefaults instantiates a new Sysctls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsAioMaxNr

`func (o *Sysctls) GetFsAioMaxNr() int32`

GetFsAioMaxNr returns the FsAioMaxNr field if non-nil, zero value otherwise.

### GetFsAioMaxNrOk

`func (o *Sysctls) GetFsAioMaxNrOk() (*int32, bool)`

GetFsAioMaxNrOk returns a tuple with the FsAioMaxNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsAioMaxNr

`func (o *Sysctls) SetFsAioMaxNr(v int32)`

SetFsAioMaxNr sets FsAioMaxNr field to given value.

### HasFsAioMaxNr

`func (o *Sysctls) HasFsAioMaxNr() bool`

HasFsAioMaxNr returns a boolean if a field has been set.

### GetFsFileMax

`func (o *Sysctls) GetFsFileMax() int32`

GetFsFileMax returns the FsFileMax field if non-nil, zero value otherwise.

### GetFsFileMaxOk

`func (o *Sysctls) GetFsFileMaxOk() (*int32, bool)`

GetFsFileMaxOk returns a tuple with the FsFileMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsFileMax

`func (o *Sysctls) SetFsFileMax(v int32)`

SetFsFileMax sets FsFileMax field to given value.

### HasFsFileMax

`func (o *Sysctls) HasFsFileMax() bool`

HasFsFileMax returns a boolean if a field has been set.

### GetFsInotifyMaxUserWatches

`func (o *Sysctls) GetFsInotifyMaxUserWatches() int32`

GetFsInotifyMaxUserWatches returns the FsInotifyMaxUserWatches field if non-nil, zero value otherwise.

### GetFsInotifyMaxUserWatchesOk

`func (o *Sysctls) GetFsInotifyMaxUserWatchesOk() (*int32, bool)`

GetFsInotifyMaxUserWatchesOk returns a tuple with the FsInotifyMaxUserWatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsInotifyMaxUserWatches

`func (o *Sysctls) SetFsInotifyMaxUserWatches(v int32)`

SetFsInotifyMaxUserWatches sets FsInotifyMaxUserWatches field to given value.

### HasFsInotifyMaxUserWatches

`func (o *Sysctls) HasFsInotifyMaxUserWatches() bool`

HasFsInotifyMaxUserWatches returns a boolean if a field has been set.

### GetFsNrOpen

`func (o *Sysctls) GetFsNrOpen() int32`

GetFsNrOpen returns the FsNrOpen field if non-nil, zero value otherwise.

### GetFsNrOpenOk

`func (o *Sysctls) GetFsNrOpenOk() (*int32, bool)`

GetFsNrOpenOk returns a tuple with the FsNrOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsNrOpen

`func (o *Sysctls) SetFsNrOpen(v int32)`

SetFsNrOpen sets FsNrOpen field to given value.

### HasFsNrOpen

`func (o *Sysctls) HasFsNrOpen() bool`

HasFsNrOpen returns a boolean if a field has been set.

### GetKernelThreadsMax

`func (o *Sysctls) GetKernelThreadsMax() int32`

GetKernelThreadsMax returns the KernelThreadsMax field if non-nil, zero value otherwise.

### GetKernelThreadsMaxOk

`func (o *Sysctls) GetKernelThreadsMaxOk() (*int32, bool)`

GetKernelThreadsMaxOk returns a tuple with the KernelThreadsMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelThreadsMax

`func (o *Sysctls) SetKernelThreadsMax(v int32)`

SetKernelThreadsMax sets KernelThreadsMax field to given value.

### HasKernelThreadsMax

`func (o *Sysctls) HasKernelThreadsMax() bool`

HasKernelThreadsMax returns a boolean if a field has been set.

### GetNetCoreNetdevMaxBacklog

`func (o *Sysctls) GetNetCoreNetdevMaxBacklog() int32`

GetNetCoreNetdevMaxBacklog returns the NetCoreNetdevMaxBacklog field if non-nil, zero value otherwise.

### GetNetCoreNetdevMaxBacklogOk

`func (o *Sysctls) GetNetCoreNetdevMaxBacklogOk() (*int32, bool)`

GetNetCoreNetdevMaxBacklogOk returns a tuple with the NetCoreNetdevMaxBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCoreNetdevMaxBacklog

`func (o *Sysctls) SetNetCoreNetdevMaxBacklog(v int32)`

SetNetCoreNetdevMaxBacklog sets NetCoreNetdevMaxBacklog field to given value.

### HasNetCoreNetdevMaxBacklog

`func (o *Sysctls) HasNetCoreNetdevMaxBacklog() bool`

HasNetCoreNetdevMaxBacklog returns a boolean if a field has been set.

### GetNetCoreOptmemMax

`func (o *Sysctls) GetNetCoreOptmemMax() int32`

GetNetCoreOptmemMax returns the NetCoreOptmemMax field if non-nil, zero value otherwise.

### GetNetCoreOptmemMaxOk

`func (o *Sysctls) GetNetCoreOptmemMaxOk() (*int32, bool)`

GetNetCoreOptmemMaxOk returns a tuple with the NetCoreOptmemMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCoreOptmemMax

`func (o *Sysctls) SetNetCoreOptmemMax(v int32)`

SetNetCoreOptmemMax sets NetCoreOptmemMax field to given value.

### HasNetCoreOptmemMax

`func (o *Sysctls) HasNetCoreOptmemMax() bool`

HasNetCoreOptmemMax returns a boolean if a field has been set.

### GetNetCoreRmemDefault

`func (o *Sysctls) GetNetCoreRmemDefault() int32`

GetNetCoreRmemDefault returns the NetCoreRmemDefault field if non-nil, zero value otherwise.

### GetNetCoreRmemDefaultOk

`func (o *Sysctls) GetNetCoreRmemDefaultOk() (*int32, bool)`

GetNetCoreRmemDefaultOk returns a tuple with the NetCoreRmemDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCoreRmemDefault

`func (o *Sysctls) SetNetCoreRmemDefault(v int32)`

SetNetCoreRmemDefault sets NetCoreRmemDefault field to given value.

### HasNetCoreRmemDefault

`func (o *Sysctls) HasNetCoreRmemDefault() bool`

HasNetCoreRmemDefault returns a boolean if a field has been set.

### GetNetCoreRmemMax

`func (o *Sysctls) GetNetCoreRmemMax() int32`

GetNetCoreRmemMax returns the NetCoreRmemMax field if non-nil, zero value otherwise.

### GetNetCoreRmemMaxOk

`func (o *Sysctls) GetNetCoreRmemMaxOk() (*int32, bool)`

GetNetCoreRmemMaxOk returns a tuple with the NetCoreRmemMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCoreRmemMax

`func (o *Sysctls) SetNetCoreRmemMax(v int32)`

SetNetCoreRmemMax sets NetCoreRmemMax field to given value.

### HasNetCoreRmemMax

`func (o *Sysctls) HasNetCoreRmemMax() bool`

HasNetCoreRmemMax returns a boolean if a field has been set.

### GetNetCoreSomaxconn

`func (o *Sysctls) GetNetCoreSomaxconn() int32`

GetNetCoreSomaxconn returns the NetCoreSomaxconn field if non-nil, zero value otherwise.

### GetNetCoreSomaxconnOk

`func (o *Sysctls) GetNetCoreSomaxconnOk() (*int32, bool)`

GetNetCoreSomaxconnOk returns a tuple with the NetCoreSomaxconn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCoreSomaxconn

`func (o *Sysctls) SetNetCoreSomaxconn(v int32)`

SetNetCoreSomaxconn sets NetCoreSomaxconn field to given value.

### HasNetCoreSomaxconn

`func (o *Sysctls) HasNetCoreSomaxconn() bool`

HasNetCoreSomaxconn returns a boolean if a field has been set.

### GetNetCoreWmemDefault

`func (o *Sysctls) GetNetCoreWmemDefault() int32`

GetNetCoreWmemDefault returns the NetCoreWmemDefault field if non-nil, zero value otherwise.

### GetNetCoreWmemDefaultOk

`func (o *Sysctls) GetNetCoreWmemDefaultOk() (*int32, bool)`

GetNetCoreWmemDefaultOk returns a tuple with the NetCoreWmemDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCoreWmemDefault

`func (o *Sysctls) SetNetCoreWmemDefault(v int32)`

SetNetCoreWmemDefault sets NetCoreWmemDefault field to given value.

### HasNetCoreWmemDefault

`func (o *Sysctls) HasNetCoreWmemDefault() bool`

HasNetCoreWmemDefault returns a boolean if a field has been set.

### GetNetCoreWmemMax

`func (o *Sysctls) GetNetCoreWmemMax() int32`

GetNetCoreWmemMax returns the NetCoreWmemMax field if non-nil, zero value otherwise.

### GetNetCoreWmemMaxOk

`func (o *Sysctls) GetNetCoreWmemMaxOk() (*int32, bool)`

GetNetCoreWmemMaxOk returns a tuple with the NetCoreWmemMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCoreWmemMax

`func (o *Sysctls) SetNetCoreWmemMax(v int32)`

SetNetCoreWmemMax sets NetCoreWmemMax field to given value.

### HasNetCoreWmemMax

`func (o *Sysctls) HasNetCoreWmemMax() bool`

HasNetCoreWmemMax returns a boolean if a field has been set.

### GetNetIpv4IpLocalPortRange

`func (o *Sysctls) GetNetIpv4IpLocalPortRange() string`

GetNetIpv4IpLocalPortRange returns the NetIpv4IpLocalPortRange field if non-nil, zero value otherwise.

### GetNetIpv4IpLocalPortRangeOk

`func (o *Sysctls) GetNetIpv4IpLocalPortRangeOk() (*string, bool)`

GetNetIpv4IpLocalPortRangeOk returns a tuple with the NetIpv4IpLocalPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4IpLocalPortRange

`func (o *Sysctls) SetNetIpv4IpLocalPortRange(v string)`

SetNetIpv4IpLocalPortRange sets NetIpv4IpLocalPortRange field to given value.

### HasNetIpv4IpLocalPortRange

`func (o *Sysctls) HasNetIpv4IpLocalPortRange() bool`

HasNetIpv4IpLocalPortRange returns a boolean if a field has been set.

### GetNetIpv4NeighDefaultGcThresh1

`func (o *Sysctls) GetNetIpv4NeighDefaultGcThresh1() int32`

GetNetIpv4NeighDefaultGcThresh1 returns the NetIpv4NeighDefaultGcThresh1 field if non-nil, zero value otherwise.

### GetNetIpv4NeighDefaultGcThresh1Ok

`func (o *Sysctls) GetNetIpv4NeighDefaultGcThresh1Ok() (*int32, bool)`

GetNetIpv4NeighDefaultGcThresh1Ok returns a tuple with the NetIpv4NeighDefaultGcThresh1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4NeighDefaultGcThresh1

`func (o *Sysctls) SetNetIpv4NeighDefaultGcThresh1(v int32)`

SetNetIpv4NeighDefaultGcThresh1 sets NetIpv4NeighDefaultGcThresh1 field to given value.

### HasNetIpv4NeighDefaultGcThresh1

`func (o *Sysctls) HasNetIpv4NeighDefaultGcThresh1() bool`

HasNetIpv4NeighDefaultGcThresh1 returns a boolean if a field has been set.

### GetNetIpv4NeighDefaultGcThresh2

`func (o *Sysctls) GetNetIpv4NeighDefaultGcThresh2() int32`

GetNetIpv4NeighDefaultGcThresh2 returns the NetIpv4NeighDefaultGcThresh2 field if non-nil, zero value otherwise.

### GetNetIpv4NeighDefaultGcThresh2Ok

`func (o *Sysctls) GetNetIpv4NeighDefaultGcThresh2Ok() (*int32, bool)`

GetNetIpv4NeighDefaultGcThresh2Ok returns a tuple with the NetIpv4NeighDefaultGcThresh2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4NeighDefaultGcThresh2

`func (o *Sysctls) SetNetIpv4NeighDefaultGcThresh2(v int32)`

SetNetIpv4NeighDefaultGcThresh2 sets NetIpv4NeighDefaultGcThresh2 field to given value.

### HasNetIpv4NeighDefaultGcThresh2

`func (o *Sysctls) HasNetIpv4NeighDefaultGcThresh2() bool`

HasNetIpv4NeighDefaultGcThresh2 returns a boolean if a field has been set.

### GetNetIpv4NeighDefaultGcThresh3

`func (o *Sysctls) GetNetIpv4NeighDefaultGcThresh3() int32`

GetNetIpv4NeighDefaultGcThresh3 returns the NetIpv4NeighDefaultGcThresh3 field if non-nil, zero value otherwise.

### GetNetIpv4NeighDefaultGcThresh3Ok

`func (o *Sysctls) GetNetIpv4NeighDefaultGcThresh3Ok() (*int32, bool)`

GetNetIpv4NeighDefaultGcThresh3Ok returns a tuple with the NetIpv4NeighDefaultGcThresh3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4NeighDefaultGcThresh3

`func (o *Sysctls) SetNetIpv4NeighDefaultGcThresh3(v int32)`

SetNetIpv4NeighDefaultGcThresh3 sets NetIpv4NeighDefaultGcThresh3 field to given value.

### HasNetIpv4NeighDefaultGcThresh3

`func (o *Sysctls) HasNetIpv4NeighDefaultGcThresh3() bool`

HasNetIpv4NeighDefaultGcThresh3 returns a boolean if a field has been set.

### GetNetIpv4TcpFinTimeout

`func (o *Sysctls) GetNetIpv4TcpFinTimeout() int32`

GetNetIpv4TcpFinTimeout returns the NetIpv4TcpFinTimeout field if non-nil, zero value otherwise.

### GetNetIpv4TcpFinTimeoutOk

`func (o *Sysctls) GetNetIpv4TcpFinTimeoutOk() (*int32, bool)`

GetNetIpv4TcpFinTimeoutOk returns a tuple with the NetIpv4TcpFinTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4TcpFinTimeout

`func (o *Sysctls) SetNetIpv4TcpFinTimeout(v int32)`

SetNetIpv4TcpFinTimeout sets NetIpv4TcpFinTimeout field to given value.

### HasNetIpv4TcpFinTimeout

`func (o *Sysctls) HasNetIpv4TcpFinTimeout() bool`

HasNetIpv4TcpFinTimeout returns a boolean if a field has been set.

### GetNetIpv4TcpKeepaliveProbes

`func (o *Sysctls) GetNetIpv4TcpKeepaliveProbes() int32`

GetNetIpv4TcpKeepaliveProbes returns the NetIpv4TcpKeepaliveProbes field if non-nil, zero value otherwise.

### GetNetIpv4TcpKeepaliveProbesOk

`func (o *Sysctls) GetNetIpv4TcpKeepaliveProbesOk() (*int32, bool)`

GetNetIpv4TcpKeepaliveProbesOk returns a tuple with the NetIpv4TcpKeepaliveProbes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4TcpKeepaliveProbes

`func (o *Sysctls) SetNetIpv4TcpKeepaliveProbes(v int32)`

SetNetIpv4TcpKeepaliveProbes sets NetIpv4TcpKeepaliveProbes field to given value.

### HasNetIpv4TcpKeepaliveProbes

`func (o *Sysctls) HasNetIpv4TcpKeepaliveProbes() bool`

HasNetIpv4TcpKeepaliveProbes returns a boolean if a field has been set.

### GetNetIpv4TcpKeepaliveTime

`func (o *Sysctls) GetNetIpv4TcpKeepaliveTime() int32`

GetNetIpv4TcpKeepaliveTime returns the NetIpv4TcpKeepaliveTime field if non-nil, zero value otherwise.

### GetNetIpv4TcpKeepaliveTimeOk

`func (o *Sysctls) GetNetIpv4TcpKeepaliveTimeOk() (*int32, bool)`

GetNetIpv4TcpKeepaliveTimeOk returns a tuple with the NetIpv4TcpKeepaliveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4TcpKeepaliveTime

`func (o *Sysctls) SetNetIpv4TcpKeepaliveTime(v int32)`

SetNetIpv4TcpKeepaliveTime sets NetIpv4TcpKeepaliveTime field to given value.

### HasNetIpv4TcpKeepaliveTime

`func (o *Sysctls) HasNetIpv4TcpKeepaliveTime() bool`

HasNetIpv4TcpKeepaliveTime returns a boolean if a field has been set.

### GetNetIpv4TcpMaxSynBacklog

`func (o *Sysctls) GetNetIpv4TcpMaxSynBacklog() int32`

GetNetIpv4TcpMaxSynBacklog returns the NetIpv4TcpMaxSynBacklog field if non-nil, zero value otherwise.

### GetNetIpv4TcpMaxSynBacklogOk

`func (o *Sysctls) GetNetIpv4TcpMaxSynBacklogOk() (*int32, bool)`

GetNetIpv4TcpMaxSynBacklogOk returns a tuple with the NetIpv4TcpMaxSynBacklog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4TcpMaxSynBacklog

`func (o *Sysctls) SetNetIpv4TcpMaxSynBacklog(v int32)`

SetNetIpv4TcpMaxSynBacklog sets NetIpv4TcpMaxSynBacklog field to given value.

### HasNetIpv4TcpMaxSynBacklog

`func (o *Sysctls) HasNetIpv4TcpMaxSynBacklog() bool`

HasNetIpv4TcpMaxSynBacklog returns a boolean if a field has been set.

### GetNetIpv4TcpMaxTwBuckets

`func (o *Sysctls) GetNetIpv4TcpMaxTwBuckets() int32`

GetNetIpv4TcpMaxTwBuckets returns the NetIpv4TcpMaxTwBuckets field if non-nil, zero value otherwise.

### GetNetIpv4TcpMaxTwBucketsOk

`func (o *Sysctls) GetNetIpv4TcpMaxTwBucketsOk() (*int32, bool)`

GetNetIpv4TcpMaxTwBucketsOk returns a tuple with the NetIpv4TcpMaxTwBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4TcpMaxTwBuckets

`func (o *Sysctls) SetNetIpv4TcpMaxTwBuckets(v int32)`

SetNetIpv4TcpMaxTwBuckets sets NetIpv4TcpMaxTwBuckets field to given value.

### HasNetIpv4TcpMaxTwBuckets

`func (o *Sysctls) HasNetIpv4TcpMaxTwBuckets() bool`

HasNetIpv4TcpMaxTwBuckets returns a boolean if a field has been set.

### GetNetIpv4TcpTwReuse

`func (o *Sysctls) GetNetIpv4TcpTwReuse() bool`

GetNetIpv4TcpTwReuse returns the NetIpv4TcpTwReuse field if non-nil, zero value otherwise.

### GetNetIpv4TcpTwReuseOk

`func (o *Sysctls) GetNetIpv4TcpTwReuseOk() (*bool, bool)`

GetNetIpv4TcpTwReuseOk returns a tuple with the NetIpv4TcpTwReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4TcpTwReuse

`func (o *Sysctls) SetNetIpv4TcpTwReuse(v bool)`

SetNetIpv4TcpTwReuse sets NetIpv4TcpTwReuse field to given value.

### HasNetIpv4TcpTwReuse

`func (o *Sysctls) HasNetIpv4TcpTwReuse() bool`

HasNetIpv4TcpTwReuse returns a boolean if a field has been set.

### GetNetIpv4TcpkeepaliveIntvl

`func (o *Sysctls) GetNetIpv4TcpkeepaliveIntvl() int32`

GetNetIpv4TcpkeepaliveIntvl returns the NetIpv4TcpkeepaliveIntvl field if non-nil, zero value otherwise.

### GetNetIpv4TcpkeepaliveIntvlOk

`func (o *Sysctls) GetNetIpv4TcpkeepaliveIntvlOk() (*int32, bool)`

GetNetIpv4TcpkeepaliveIntvlOk returns a tuple with the NetIpv4TcpkeepaliveIntvl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIpv4TcpkeepaliveIntvl

`func (o *Sysctls) SetNetIpv4TcpkeepaliveIntvl(v int32)`

SetNetIpv4TcpkeepaliveIntvl sets NetIpv4TcpkeepaliveIntvl field to given value.

### HasNetIpv4TcpkeepaliveIntvl

`func (o *Sysctls) HasNetIpv4TcpkeepaliveIntvl() bool`

HasNetIpv4TcpkeepaliveIntvl returns a boolean if a field has been set.

### GetNetNetfilterNfConntrackBuckets

`func (o *Sysctls) GetNetNetfilterNfConntrackBuckets() int32`

GetNetNetfilterNfConntrackBuckets returns the NetNetfilterNfConntrackBuckets field if non-nil, zero value otherwise.

### GetNetNetfilterNfConntrackBucketsOk

`func (o *Sysctls) GetNetNetfilterNfConntrackBucketsOk() (*int32, bool)`

GetNetNetfilterNfConntrackBucketsOk returns a tuple with the NetNetfilterNfConntrackBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetNetfilterNfConntrackBuckets

`func (o *Sysctls) SetNetNetfilterNfConntrackBuckets(v int32)`

SetNetNetfilterNfConntrackBuckets sets NetNetfilterNfConntrackBuckets field to given value.

### HasNetNetfilterNfConntrackBuckets

`func (o *Sysctls) HasNetNetfilterNfConntrackBuckets() bool`

HasNetNetfilterNfConntrackBuckets returns a boolean if a field has been set.

### GetNetNetfilterNfConntrackMax

`func (o *Sysctls) GetNetNetfilterNfConntrackMax() int32`

GetNetNetfilterNfConntrackMax returns the NetNetfilterNfConntrackMax field if non-nil, zero value otherwise.

### GetNetNetfilterNfConntrackMaxOk

`func (o *Sysctls) GetNetNetfilterNfConntrackMaxOk() (*int32, bool)`

GetNetNetfilterNfConntrackMaxOk returns a tuple with the NetNetfilterNfConntrackMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetNetfilterNfConntrackMax

`func (o *Sysctls) SetNetNetfilterNfConntrackMax(v int32)`

SetNetNetfilterNfConntrackMax sets NetNetfilterNfConntrackMax field to given value.

### HasNetNetfilterNfConntrackMax

`func (o *Sysctls) HasNetNetfilterNfConntrackMax() bool`

HasNetNetfilterNfConntrackMax returns a boolean if a field has been set.

### GetVmMaxMapCount

`func (o *Sysctls) GetVmMaxMapCount() int32`

GetVmMaxMapCount returns the VmMaxMapCount field if non-nil, zero value otherwise.

### GetVmMaxMapCountOk

`func (o *Sysctls) GetVmMaxMapCountOk() (*int32, bool)`

GetVmMaxMapCountOk returns a tuple with the VmMaxMapCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmMaxMapCount

`func (o *Sysctls) SetVmMaxMapCount(v int32)`

SetVmMaxMapCount sets VmMaxMapCount field to given value.

### HasVmMaxMapCount

`func (o *Sysctls) HasVmMaxMapCount() bool`

HasVmMaxMapCount returns a boolean if a field has been set.

### GetVmSwappiness

`func (o *Sysctls) GetVmSwappiness() int32`

GetVmSwappiness returns the VmSwappiness field if non-nil, zero value otherwise.

### GetVmSwappinessOk

`func (o *Sysctls) GetVmSwappinessOk() (*int32, bool)`

GetVmSwappinessOk returns a tuple with the VmSwappiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSwappiness

`func (o *Sysctls) SetVmSwappiness(v int32)`

SetVmSwappiness sets VmSwappiness field to given value.

### HasVmSwappiness

`func (o *Sysctls) HasVmSwappiness() bool`

HasVmSwappiness returns a boolean if a field has been set.

### GetVmVfsCachePressure

`func (o *Sysctls) GetVmVfsCachePressure() int32`

GetVmVfsCachePressure returns the VmVfsCachePressure field if non-nil, zero value otherwise.

### GetVmVfsCachePressureOk

`func (o *Sysctls) GetVmVfsCachePressureOk() (*int32, bool)`

GetVmVfsCachePressureOk returns a tuple with the VmVfsCachePressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVfsCachePressure

`func (o *Sysctls) SetVmVfsCachePressure(v int32)`

SetVmVfsCachePressure sets VmVfsCachePressure field to given value.

### HasVmVfsCachePressure

`func (o *Sysctls) HasVmVfsCachePressure() bool`

HasVmVfsCachePressure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


