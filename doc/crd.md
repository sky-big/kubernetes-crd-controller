# CustomResourceDefinition资源详细说明

```
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # CRD名称必须是: <plural>.<group>格式
  name: crd.example.dev
spec:
  # 决定了REST API路径：/apis/<group>/<version>
  group: example.dev
  # 此CRD支持的版本
  versions:
    - name: v1
      # 每个版本都可以被禁用或启用
      served: true
      # 只有一个版本可以被标记为true，表示以此版本来存储
      storage: true
  # 可选Namespaced或Cluster，表示此资源是命名空间限定的，还是全局的
  scope: Namespaced
  names:
    # 决定了REST API路径： /apis/<group>/<version>/<plural>
    plural: crd
    # 在CLI中的别名
    singular: crontab
    # 驼峰式大小写的，正式的资源类型
    kind: CRD1
    # 在CLI中可以使用的短名称
    shortNames:
    - cr
```