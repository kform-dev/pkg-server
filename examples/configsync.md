## configsync config

```yaml
cat <<EOF | kubectl apply -f -
  apiVersion: configsync.gke.io/v1beta1
  kind: RootSync
  metadata: # kpt-merge: config-management-system/nephio-mgmt-cluster-sync
    name: cert-manager
    namespace: config-management-system
  spec:
    sourceFormat: unstructured
    git:
      repo: https://github.com/henderiw/repo-target.git
      branch: mgmt/base/sw/cert-manager/pv-c7806cdc1a2927aa
      revision: mgmt/base/sw/cert-manager/v1
      auth: token
      dir: mgmt/base/sw/cert-manager/out
      secretRef:
        name: git-pat
EOF
```

```yaml
apiVersion: configsync.gke.io/v1beta1
kind: RootSync
metadata:
  name: nephio-workload-cluster-sync
  namespace: config-management-system
spec:
  sourceFormat: unstructured
  git:
    repo: https://github.com/nephio-test/test-edge-01
    branch: main
    dir: /free5gc-upf
    revision: free5gc-upf/v6
    auth: ...some auth stuff not shown here...
```