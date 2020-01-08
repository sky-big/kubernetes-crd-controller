# CustomResourceDefinition资源详细说明

## 基本字段定义说明
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

## 声明kubernetes在kubectl获取该CRD资源显示额外的字段

```
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: crd.example.dev
spec:
  # 在资源名称后面，额外显示以下字段
  additionalPrinterColumns:
    - name: Status
      type: string
      JSONPath: .status.phase
```

## 声明Kubernetes规定自定义资源的字段验证规则

```
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: crd.example.dev
spec:
  validation:
    openAPIV3Schema:
      # 根对象包含的属性
      properties:
        # spec属性规则
        spec:
          # 必须具有以下子属性
          required:
            - targetRef
            - service
            - canaryAnalysis
          # spec的子属性列表
          properties:
            # 简单属性，指定类型
            progressDeadlineSeconds:
              type: number
            # 对象属性
            targetRef:
              type: object
              required: ['apiVersion', 'kind', 'name']
              properties:
                apiVersion:
                  type: string
                kind:
                  type: string
                name:
                  type: string
            # 复杂属性：可以是字符串，也可以是对象
            autoscalerRef:
              anyOf:
                - type: string
                - type: object
              required: ['apiVersion', 'kind', 'name']
              properties:
                apiVersion:
                  type: string
                kind:
                  type: string
                name:
                  type: string
            service:
              type: object
              required: ['port']
              properties:
                port:
                  type: number
            skipAnalysis:
              type: boolean
            canaryAnalysis:
              properties:
                interval:
                  type: string
                  # 基于正则式验证
                  pattern: "^[0-9]+(m|s)"
                threshold:
                  type: number
                maxWeight:
                  type: number
                stepWeight:
                  type: number
                metrics:
                  type: array
                  properties:
                    items:
                      type: object
                      required: ['name', 'threshold']
                      properties:
                        name:
                          type: string
                        interval:
                          type: string
                          pattern: "^[0-9]+(m|s)"
                        threshold:
                          type: number
                        query:
                          type: string
                webhooks:
                  type: array
                  properties:
                    items:
                      type: object
                      required: ['name', 'url', 'timeout']
                      properties:
                        name:
                          type: string
                        url:
                          type: string
                          format: url
                        timeout:
                          type: string
                          pattern: "^[0-9]+(m|s)"
```