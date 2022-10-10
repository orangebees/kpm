## init
--- 
初始化工作目录，创建kpm.json文件。

设置最低运行版本为系统自带kcl版本

example :

> kpm init testproject

此时创建了kpm.json



```json
{
    "package_name": "testproject",
    "kclvm_min_version": "v0.50",
    "require": []
    }
```