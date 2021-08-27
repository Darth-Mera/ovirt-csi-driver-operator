// Code generated for package generated by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/controller.yaml
// assets/controller_sa.yaml
// assets/csidriver.yaml
// assets/node.yaml
// assets/node_sa.yaml
// assets/rbac/attacher_binding.yaml
// assets/rbac/attacher_role.yaml
// assets/rbac/controller_privileged_binding.yaml
// assets/rbac/node_privileged_binding.yaml
// assets/rbac/privileged_role.yaml
// assets/rbac/provisioner_binding.yaml
// assets/rbac/provisioner_role.yaml
// assets/rbac/resizer_binding.yaml
// assets/rbac/resizer_role.yaml
// assets/rbac/snapshotter_binding.yaml
// assets/rbac/snapshotter_role.yaml
// assets/storageclass.yaml
package generated

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _controllerYaml = []byte(`kind: Deployment
apiVersion: apps/v1
metadata:
  name: ovirt-csi-driver-controller
  namespace: openshift-cluster-csi-drivers
  annotations:
    config.openshift.io/inject-proxy: csi-driver
spec:
  selector:
    matchLabels:
      app: ovirt-csi-driver-controller
  serviceName: ovirt-csi-driver-controller
  replicas: 1
  template:
    metadata:
      labels:
        app: ovirt-csi-driver-controller
    spec:
      hostNetwork: true
      serviceAccount: ovirt-csi-driver-controller-sa
      priorityClassName: system-cluster-critical
      nodeSelector:
        node-role.kubernetes.io/master: ""
      tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: node-role.kubernetes.io/master
          operator: Exists
          effect: "NoSchedule"
      initContainers:
        - name: prepare-ovirt-config
          env:
            - name: OVIRT_URL
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_url
            - name: OVIRT_USERNAME
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_username
            - name: OVIRT_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_password
            - name: OVIRT_CAFILE
              value: /tmp/config/ovirt-engine-ca.pem
            - name: OVIRT_INSECURE
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_insecure
            - name: OVIRT_CA_BUNDLE
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_ca_bundle
          image: ${DRIVER_IMAGE}
          command:
            - /bin/sh
            - -c
            - |
              #!/bin/sh
              cat << EOF > /tmp/config/ovirt-config.yaml
              ovirt_url: $OVIRT_URL
              ovirt_username: $OVIRT_USERNAME
              ovirt_password: $OVIRT_PASSWORD
              # set a valid path only if ca bundle has content
              ovirt_cafile: ${OVIRT_CA_BUNDLE:+$OVIRT_CAFILE}
              ovirt_insecure: $OVIRT_INSECURE
              EOF
              if [[ -n "$OVIRT_CA_BUNDLE" ]]; then echo "$OVIRT_CA_BUNDLE" > $OVIRT_CAFILE ; fi
          volumeMounts:
            - name: config
              mountPath: /tmp/config
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
      containers:
        # Warning: the operator expects the first container to be the CSI driver
        - name: csi-driver
          image: ${DRIVER_IMAGE}
          args:
            - --endpoint=$(CSI_ENDPOINT)
            - --logtostderr
            - --v=${LOG_LEVEL}
          env:
            - name: CSI_ENDPOINT
              value: unix:///var/lib/csi/sockets/pluginproxy/csi.sock
            - name: OVIRT_CONFIG
              value: /tmp/config/ovirt-config.yaml
          ports:
            - name: healthz
              # Due to hostNetwork, this port is open on a node!
              containerPort: 10301
              protocol: TCP
          volumeMounts:
            - name: socket-dir
              mountPath: /var/lib/csi/sockets/pluginproxy/
            - name: config
              mountPath: /tmp/config
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
        - name: csi-provisioner
          image: ${PROVISIONER_IMAGE}
          args:
            - --csi-address=$(ADDRESS)
            - --default-fstype=ext4
            - --v=${LOG_LEVEL}
          env:
            - name: ADDRESS
              value: /var/lib/csi/sockets/pluginproxy/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /var/lib/csi/sockets/pluginproxy/
        - name: csi-attacher
          image: ${ATTACHER_IMAGE}
          args:
            - --csi-address=$(ADDRESS)
            - --v=${LOG_LEVEL}
          env:
            - name: ADDRESS
              value: /var/lib/csi/sockets/pluginproxy/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /var/lib/csi/sockets/pluginproxy/
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
        - name: csi-liveness-probe
          image: ${LIVENESS_PROBE_IMAGE}
          args:
            - --csi-address=/csi/csi.sock
            - --probe-timeout=3s
            - --health-port=10301
          volumeMounts:
            - name: socket-dir
              mountPath: /csi
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
      volumes:
        - name: socket-dir
          emptyDir: {}
        - name: config
          emptyDir: {}
`)

func controllerYamlBytes() ([]byte, error) {
	return _controllerYaml, nil
}

func controllerYaml() (*asset, error) {
	bytes, err := controllerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controller_saYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: ovirt-csi-driver-controller-sa
  namespace: openshift-cluster-csi-drivers
`)

func controller_saYamlBytes() ([]byte, error) {
	return _controller_saYaml, nil
}

func controller_saYaml() (*asset, error) {
	bytes, err := controller_saYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller_sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _csidriverYaml = []byte(`apiVersion: storage.k8s.io/v1
kind: CSIDriver
metadata:
  name: csi.ovirt.org
  annotations:
      # This CSIDriver is managed by an OCP CSI operator
      csi.openshift.io/managed: "true"
spec:
  attachRequired: true
  podInfoOnMount: false
`)

func csidriverYamlBytes() ([]byte, error) {
	return _csidriverYaml, nil
}

func csidriverYaml() (*asset, error) {
	bytes, err := csidriverYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "csidriver.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeYaml = []byte(`kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: ovirt-csi-driver-node
  namespace: openshift-cluster-csi-drivers
  annotations:
    config.openshift.io/inject-proxy: csi-driver
spec:
  selector:
    matchLabels:
      app: ovirt-csi-driver-node
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 10%
  template:
    metadata:
      labels:
        app: ovirt-csi-driver-node
    spec:
      hostNetwork: true
      serviceAccount: ovirt-csi-driver-node-sa
      priorityClassName: system-node-critical
      tolerations:
        - operator: Exists
      initContainers:
        - name: prepare-ovirt-config
          env:
            - name: OVIRT_URL
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_url
            - name: OVIRT_USERNAME
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_username
            - name: OVIRT_PASSWORD
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_password
            - name: OVIRT_CAFILE
              value: /tmp/config/ovirt-engine-ca.pem
            - name: OVIRT_INSECURE
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_insecure
            - name: OVIRT_CA_BUNDLE
              valueFrom:
                secretKeyRef:
                  name: ovirt-credentials
                  key: ovirt_ca_bundle
          image: ${DRIVER_IMAGE}
          command:
            - /bin/sh
            - -c
            - |
              #!/bin/sh
              cat << EOF > /tmp/config/ovirt-config.yaml
              ovirt_url: $OVIRT_URL
              ovirt_username: $OVIRT_USERNAME
              ovirt_password: $OVIRT_PASSWORD
              # set a valid path only if ca bundle has content
              ovirt_cafile: ${OVIRT_CA_BUNDLE:+$OVIRT_CAFILE}
              ovirt_insecure: $OVIRT_INSECURE
              EOF
              if [[ -n "$OVIRT_CA_BUNDLE" ]]; then echo "$OVIRT_CA_BUNDLE" > $OVIRT_CAFILE ; fi
          volumeMounts:
            - name: config
              mountPath: /tmp/config
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
      containers:
        - name: csi-driver
          securityContext:
            privileged: true
          image: ${DRIVER_IMAGE}
          args:
            - --endpoint=$(CSI_ENDPOINT)
            - --logtostderr
            - --v=${LOG_LEVEL}
            - --node-name=$(KUBE_NODE_NAME)
          env:
            - name: CSI_ENDPOINT
              value: unix:/csi/csi.sock
            - name: OVIRT_CONFIG
              value: /tmp/config/ovirt-config.yaml
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
          volumeMounts:
            - name: kubelet-dir
              mountPath: /var/lib/kubelet
              mountPropagation: "Bidirectional"
            - name: plugin-dir
              mountPath: /csi
            - name: device-dir
              mountPath: /dev
            - name: udev
              mountPath: /run/udev
            - name: config
              mountPath: /tmp/config

          ports:
            - name: healthz
              # Due to hostNetwork, this port is open on a node!
              containerPort: 10300
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /healthz
              port: healthz
            initialDelaySeconds: 10
            timeoutSeconds: 3
            periodSeconds: 30
            failureThreshold: 5
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
        - name: csi-node-driver-registrar
          securityContext:
            privileged: true
          image: ${NODE_DRIVER_REGISTRAR_IMAGE}
          args:
            - --csi-address=$(ADDRESS)
            - --kubelet-registration-path=$(DRIVER_REG_SOCK_PATH)
            - --v=${LOG_LEVEL}
          lifecycle:
            preStop:
              exec:
                command: ["/bin/sh", "-c", "rm -rf /registration/csi.ovirt.org-reg.sock /csi/csi.sock"]
          env:
            - name: ADDRESS
              value: /csi/csi.sock
            - name: DRIVER_REG_SOCK_PATH
              value: /var/lib/kubelet/plugins/csi.ovirt.org/csi.sock
          volumeMounts:
            - name: plugin-dir
              mountPath: /csi
            - name: registration-dir
              mountPath: /registration
          resources:
            requests:
              memory: 20Mi
              cpu: 5m
        - name: csi-liveness-probe
          image: ${LIVENESS_PROBE_IMAGE}
          args:
            - --csi-address=/csi/csi.sock
            - --probe-timeout=3s
            - --health-port=10300
          volumeMounts:
            - name: plugin-dir
              mountPath: /csi
          resources:
            requests:
              memory: 20Mi
              cpu: 5m
      volumes:
        - name: kubelet-dir
          hostPath:
            path: /var/lib/kubelet
            type: Directory
        - name: plugin-dir
          hostPath:
            path: /var/lib/kubelet/plugins/csi.ovirt.org/
            type: DirectoryOrCreate
        - name: registration-dir
          hostPath:
            path: /var/lib/kubelet/plugins_registry/
            type: Directory
        - name: device-dir
          hostPath:
            path: /dev
            type: Directory
        - name: udev
          hostPath:
            path: /run/udev
        - name: config
          emptyDir: {}
`)

func nodeYamlBytes() ([]byte, error) {
	return _nodeYaml, nil
}

func nodeYaml() (*asset, error) {
	bytes, err := nodeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _node_saYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: ovirt-csi-driver-node-sa
  namespace: openshift-cluster-csi-drivers
`)

func node_saYamlBytes() ([]byte, error) {
	return _node_saYaml, nil
}

func node_saYaml() (*asset, error) {
	bytes, err := node_saYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node_sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacAttacher_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-csi-attacher-binding
subjects:
  - kind: ServiceAccount
    name: ovirt-csi-driver-controller-sa
    namespace: openshift-cluster-csi-drivers
roleRef:
  kind: ClusterRole
  name: ovirt-external-attacher-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacAttacher_bindingYamlBytes() ([]byte, error) {
	return _rbacAttacher_bindingYaml, nil
}

func rbacAttacher_bindingYaml() (*asset, error) {
	bytes, err := rbacAttacher_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/attacher_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacAttacher_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-external-attacher-role
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["csinodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments/status"]
    verbs: ["patch"]
`)

func rbacAttacher_roleYamlBytes() ([]byte, error) {
	return _rbacAttacher_roleYaml, nil
}

func rbacAttacher_roleYaml() (*asset, error) {
	bytes, err := rbacAttacher_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/attacher_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacController_privileged_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-controller-privileged-binding
subjects:
  - kind: ServiceAccount
    name: ovirt-csi-driver-controller-sa
    namespace: openshift-cluster-csi-drivers
roleRef:
  kind: ClusterRole
  name: ovirt-privileged-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacController_privileged_bindingYamlBytes() ([]byte, error) {
	return _rbacController_privileged_bindingYaml, nil
}

func rbacController_privileged_bindingYaml() (*asset, error) {
	bytes, err := rbacController_privileged_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/controller_privileged_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacNode_privileged_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-node-privileged-binding
subjects:
  - kind: ServiceAccount
    name: ovirt-csi-driver-node-sa
    namespace: openshift-cluster-csi-drivers
roleRef:
  kind: ClusterRole
  name: ovirt-privileged-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacNode_privileged_bindingYamlBytes() ([]byte, error) {
	return _rbacNode_privileged_bindingYaml, nil
}

func rbacNode_privileged_bindingYaml() (*asset, error) {
	bytes, err := rbacNode_privileged_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/node_privileged_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacPrivileged_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-privileged-role
rules:
  - apiGroups: ["security.openshift.io"]
    resourceNames: ["privileged"]
    resources: ["securitycontextconstraints"]
    verbs: ["use"]
  # This permission is needed by the driver NodeGetInfo
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get"]
`)

func rbacPrivileged_roleYamlBytes() ([]byte, error) {
	return _rbacPrivileged_roleYaml, nil
}

func rbacPrivileged_roleYaml() (*asset, error) {
	bytes, err := rbacPrivileged_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/privileged_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacProvisioner_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-csi-provisioner-binding
subjects:
  - kind: ServiceAccount
    name: ovirt-csi-driver-controller-sa
    namespace: openshift-cluster-csi-drivers
roleRef:
  kind: ClusterRole
  name: ovirt-external-provisioner-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacProvisioner_bindingYamlBytes() ([]byte, error) {
	return _rbacProvisioner_bindingYaml, nil
}

func rbacProvisioner_bindingYaml() (*asset, error) {
	bytes, err := rbacProvisioner_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/provisioner_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacProvisioner_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-external-provisioner-role
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "create", "delete"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["csinodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
`)

func rbacProvisioner_roleYamlBytes() ([]byte, error) {
	return _rbacProvisioner_roleYaml, nil
}

func rbacProvisioner_roleYaml() (*asset, error) {
	bytes, err := rbacProvisioner_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/provisioner_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacResizer_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-csi-resizer-binding
subjects:
  - kind: ServiceAccount
    name: ovirt-csi-driver-controller-sa
    namespace: openshift-cluster-csi-drivers
roleRef:
  kind: ClusterRole
  name: ovirt-external-resizer-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacResizer_bindingYamlBytes() ([]byte, error) {
	return _rbacResizer_bindingYaml, nil
}

func rbacResizer_bindingYaml() (*asset, error) {
	bytes, err := rbacResizer_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/resizer_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacResizer_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-external-resizer-role
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update", "patch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims/status"]
    verbs: ["update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
`)

func rbacResizer_roleYamlBytes() ([]byte, error) {
	return _rbacResizer_roleYaml, nil
}

func rbacResizer_roleYaml() (*asset, error) {
	bytes, err := rbacResizer_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/resizer_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacSnapshotter_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-csi-snapshotter-binding
subjects:
  - kind: ServiceAccount
    name: ovirt-csi-driver-controller-sa
    namespace: openshift-cluster-csi-drivers
roleRef:
  kind: ClusterRole
  name: ovirt-external-snapshotter-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacSnapshotter_bindingYamlBytes() ([]byte, error) {
	return _rbacSnapshotter_bindingYaml, nil
}

func rbacSnapshotter_bindingYaml() (*asset, error) {
	bytes, err := rbacSnapshotter_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/snapshotter_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacSnapshotter_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: ovirt-external-snapshotter-role
rules:
- apiGroups: [""]
  resources: ["persistentvolumes"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["persistentvolumeclaims"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["storage.k8s.io"]
  resources: ["storageclasses"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["events"]
  verbs: ["list", "watch", "create", "update", "patch"]
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "list"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshotclasses"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshotcontents"]
  verbs: ["create", "get", "list", "watch", "update", "delete"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshotcontents/status"]
  verbs: ["update"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshots"]
  verbs: ["get", "list", "watch", "update"]
- apiGroups: ["apiextensions.k8s.io"]
  resources: ["customresourcedefinitions"]
  verbs: ["create", "list", "watch", "delete"]
- apiGroups: ["coordination.k8s.io"]
  resources: ["leases"]
  verbs: ["get", "watch", "list", "delete", "update", "create"]
`)

func rbacSnapshotter_roleYamlBytes() ([]byte, error) {
	return _rbacSnapshotter_roleYaml, nil
}

func rbacSnapshotter_roleYaml() (*asset, error) {
	bytes, err := rbacSnapshotter_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/snapshotter_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _storageclassYaml = []byte(`apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: ovirt-csi-sc
parameters:
  # the name of the oVirt storage domain. "nfs" is just an example.
  storageDomainName: "nfs"
  thinProvisioning: "true"
provisioner: csi.ovirt.org
reclaimPolicy: "Delete"
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
`)

func storageclassYamlBytes() ([]byte, error) {
	return _storageclassYaml, nil
}

func storageclassYaml() (*asset, error) {
	bytes, err := storageclassYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "storageclass.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"controller.yaml":                         controllerYaml,
	"controller_sa.yaml":                      controller_saYaml,
	"csidriver.yaml":                          csidriverYaml,
	"node.yaml":                               nodeYaml,
	"node_sa.yaml":                            node_saYaml,
	"rbac/attacher_binding.yaml":              rbacAttacher_bindingYaml,
	"rbac/attacher_role.yaml":                 rbacAttacher_roleYaml,
	"rbac/controller_privileged_binding.yaml": rbacController_privileged_bindingYaml,
	"rbac/node_privileged_binding.yaml":       rbacNode_privileged_bindingYaml,
	"rbac/privileged_role.yaml":               rbacPrivileged_roleYaml,
	"rbac/provisioner_binding.yaml":           rbacProvisioner_bindingYaml,
	"rbac/provisioner_role.yaml":              rbacProvisioner_roleYaml,
	"rbac/resizer_binding.yaml":               rbacResizer_bindingYaml,
	"rbac/resizer_role.yaml":                  rbacResizer_roleYaml,
	"rbac/snapshotter_binding.yaml":           rbacSnapshotter_bindingYaml,
	"rbac/snapshotter_role.yaml":              rbacSnapshotter_roleYaml,
	"storageclass.yaml":                       storageclassYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"controller.yaml":    {controllerYaml, map[string]*bintree{}},
	"controller_sa.yaml": {controller_saYaml, map[string]*bintree{}},
	"csidriver.yaml":     {csidriverYaml, map[string]*bintree{}},
	"node.yaml":          {nodeYaml, map[string]*bintree{}},
	"node_sa.yaml":       {node_saYaml, map[string]*bintree{}},
	"rbac": {nil, map[string]*bintree{
		"attacher_binding.yaml":              {rbacAttacher_bindingYaml, map[string]*bintree{}},
		"attacher_role.yaml":                 {rbacAttacher_roleYaml, map[string]*bintree{}},
		"controller_privileged_binding.yaml": {rbacController_privileged_bindingYaml, map[string]*bintree{}},
		"node_privileged_binding.yaml":       {rbacNode_privileged_bindingYaml, map[string]*bintree{}},
		"privileged_role.yaml":               {rbacPrivileged_roleYaml, map[string]*bintree{}},
		"provisioner_binding.yaml":           {rbacProvisioner_bindingYaml, map[string]*bintree{}},
		"provisioner_role.yaml":              {rbacProvisioner_roleYaml, map[string]*bintree{}},
		"resizer_binding.yaml":               {rbacResizer_bindingYaml, map[string]*bintree{}},
		"resizer_role.yaml":                  {rbacResizer_roleYaml, map[string]*bintree{}},
		"snapshotter_binding.yaml":           {rbacSnapshotter_bindingYaml, map[string]*bintree{}},
		"snapshotter_role.yaml":              {rbacSnapshotter_roleYaml, map[string]*bintree{}},
	}},
	"storageclass.yaml": {storageclassYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
