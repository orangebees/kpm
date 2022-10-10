# kpm add
安装软件包及其依赖的任何软件包。 默认情况下，任何新软件包都安装为生产依赖项。
软链接对应包到依赖目录下
## 摘要：
| 命令 | 功能 | 
| :-----| :----- | 
| kpm add konfig | 将本地中心构建的包映射在当前工作目录下，如果没有则先加载包 | 
| kpm add -g konfig | 将包构建在本地中心 | 
| kpm add konfig:konfig2 | 为konfig包设置别名konfig2，将使用别名将包安装在工作空间 | 
| kpm add konfig@alpha | 从 alpha 标签下安装 | 
| kpm add -g konfig@1.0.0 | 安装指定版本 1.0.0 | 
| kpm add -git --internal https://github.com/KusionStack/konfig | 复制安装git包在工作区internal内部， | 

# 支持的包地址

## 从kpm仓库安装

kpm add package-name 默认会从 kpm registry 安装最新的 package-name。

如果在 工作区 中执行，该命令将去检查全局包存储是否安装此名字的包，如果不存在的话则在线下载构建

你还可以通过以下方式安装包：

tag: kpm add konfig@alpha
version: kpm add konfig@1.0.0

## 从 git 安装

 kpm add <'git remote url'>



来自 master 的最新提交： kpm add https://github.com/KusionStack/konfig

提交: kpm add https://github.com/KusionStack/konfig#97edff6f525f192a3f83cea194765f769ae2678
