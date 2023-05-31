| Status | Package | Time (seconds) |
|--------|---------|----------------|
| :heavy_check_mark: | RuntimeAgentChaos Cilium agent Checking for file-descriptor leak | 0.214          |
| :heavy_check_mark: | RuntimeAgentChaos Cilium agent removing leftover Cilium interfaces | 7.109          |
| :heavy_check_mark: | RuntimeAgentChaos Connectivity over restarts Checking that during restart no traffic is dropped using Egress + Ingress Traffic | 15.078         |
| :heavy_check_mark: | RuntimeAgentChaos Endpoint Endpoint recovery on restart | 9.721          |
| :heavy_check_mark: | RuntimeAgentChaos KVStore Delete event on KVStore with CIDR identities | 2.026          |
| :heavy_check_mark: | RuntimeAgentChaos KVStore Validate that delete events on KVStore do not release in use identities | 0.510          |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies CNAME follow | 7.072          |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies Can update L7 DNS policy rules | 7.968          |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies DNS proxy policy works if Cilium stops | 28.811         |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies Enforces L3 policy even when no IPs are inserted | 16.734         |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies Enforces ToFQDNs policy | 21.206         |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies Implements matchPattern: * | 16.813         |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies Interaction with other ToCIDR rules | 6.716          |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies Roundrobin DNS | 3.471          |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies Validate dns-proxy monitor information | 6.654          |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies With verbose policy logs Validates DNSSEC responses | 21.471         |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies toFQDNs populates toCIDRSet (data from proxy) L3-dependent L7/HTTP with toFQDN updates proxy policy | 15.874         |
| :heavy_check_mark: | RuntimeAgentFQDNPolicies toFQDNs populates toCIDRSet (data from proxy) Policy addition after DNS lookup | 15.741         |
| :heavy_check_mark: | RuntimeAgentKVStoreTest KVStore tests Consul KVStore | 9.850          |
| :heavy_check_mark: | RuntimeAgentKVStoreTest KVStore tests Etcd KVStore | 11.602         |
| :heavy_check_mark: | RuntimeAgentPolicies Init Policy Default Drop Test With PolicyAuditMode tests egress | 6.134          |
| :heavy_check_mark: | RuntimeAgentPolicies Init Policy Default Drop Test With PolicyAuditMode tests ingress | 16.108         |
| :heavy_check_mark: | RuntimeAgentPolicies Init Policy Default Drop Test tests egress | 6.338          |
| :heavy_check_mark: | RuntimeAgentPolicies Init Policy Default Drop Test tests ingress | 21.706         |
| :heavy_check_mark: | RuntimeAgentPolicies Init Policy Test Init Egress Policy Test | 6.701          |
| :heavy_check_mark: | RuntimeAgentPolicies Init Policy Test Init Ingress Policy Test | 9.782          |
| :heavy_check_mark: | RuntimeAgentPolicies Tests Endpoint Connectivity Functions After Daemon Configuration Is Updated | 49.925         |
| :heavy_check_mark: | RuntimeAgentPolicies Tests EntityNone as a deny-all | 29.826         |
| :heavy_check_mark: | RuntimeAgentPolicies TestsEgressToHost Tests Egress To Host | 29.178         |
| :heavy_check_mark: | RuntimeAgentPolicies TestsEgressToHost Tests egress with CIDR+L4 policy | 44.152         |
| :heavy_check_mark: | RuntimeAgentPolicies TestsEgressToHost Tests egress with CIDR+L4 policy to external https service | 46.651         |
| :heavy_check_mark: | RuntimeAgentPolicies TestsEgressToHost Tests egress with CIDR+L7 policy | 44.327         |
| :heavy_check_mark: | RuntimeSSHTests Should fail when context times out | 3.045          |
