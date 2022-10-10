# 文件目录

kpm工作目录
$KPM_ROOT=$HOME/kpm
example :
/root/kpm

缓存目录，基于kpm工作域名分割

$KPM_ROOT+"/"+$KPM_SERVER_ADDR_PATH

example :

/root/kpm/registry/kpm.kusionstack.io

/root/kpm/registry

//被链接的单包，可以直接软链接到工程目录下

/root/kpm/registry/kpm.kusionstack.io/kcl_modules

//包标签 只能在cdn侧缓存，客户端不应缓存，实时请求

/root/kpm/registry/kpm.kusionstack.io/pkg_tag/:pkgname/latest&&beta&&alpha&&v1-alpha

可强缓存


/root/kpm/registry/kpm.kusionstack.io/pkg/:pkgname/v&&index/:version/src_tgz&&hash&&info&&index

//通过git安装的包

/root/kpm/git/pkg/github.com/a/b@:version/src&&hash&&info&&index

/root/kpm/git/kcl_module/github.com/a/b@:version/src&&hash&&info

//kpm store目录

/root/kpm/v1/store/

/kpm

/root/kpm/store/00~ff/:sha512
