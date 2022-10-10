##  add
---
添加包依赖


自动在仓库中搜索此包与版本，如果没有指定版本的话则自动指定最新版本，

如果搜索到则下载到临时下载校验目录中，完成校验后移动到本地缓存模块仓库中。

校验之后先递归此次操作加载依赖，先检查包最低版本，如果最低版本高于当前设置版本并小于工作版本，出现交互式信息，询问是否提升最高版本，如果最低版本高于工作版本，直接报错。

example :

> kpm add konfig

实际等于
> kpm add konfig@latest
>
latest版本即最新版本，如果此时最新版本为0.0.1，则实际等于
> kpm add konfig@0.0.1



kpm.json

```json
{
    "package_name": "testproject",
    "require": [
      {
        "package_name": "konfig",
        "package_version": "0.0.1",
        "package_integrity": "e526c92b5bfa78c47f39f03372aadaed139bd0845a6d59c94edc9e86f468bb2fa85cc45aac256051a483f377e3d1739855ee7e63377cc6edb96d5fd832ff6b89",
        "package_address": "https://github.com/KusionStack/konfig"
      }
     ]
    }
```

指定版本安装
> kpm add konfig@0.0.1


kpm.json

```json
{
  "package_name": "testproject",
  "require": [
    {
      "package_name": "konfig",
      "package_version": "0.0.1",
      "package_integrity": "e526c92b5bfa78c47f39f03372aadaed139bd0845a6d59c94edc9e86f468bb2fa85cc45aac256051a483f377e3d1739855ee7e63377cc6edb96d5fd832ff6b89",
      "package_address": "https://github.com/KusionStack/konfig"
    }
  ]
}
```
