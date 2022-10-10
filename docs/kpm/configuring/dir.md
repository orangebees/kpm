## 目录结构

kpm工作目录
$KPM_ROOT=$HOME/kpm
example :
/root/kpm

缓存目录，基于kpm工作域名分割
$KPM_ROOT+$KPM_SERVER_ADDR_PATH
example :
/root/kpm/registry/kpm.kusionstack.io
/root/kpm/registry
/root/kpm/store/

kpm包存放目录
$KPM_ROOT+"/"+$KPM_SERVER_ADDR_PATH+"/pkg"
example :
//没有被链接的单包
/root/kpm/kpm.kusionstack.io/pkg
//被链接的单包
/root/kpm/kpm.kusionstack.io/kcl_module
kpm包存放方式
$KPM_ROOT+"/"+$KPM_SERVER_ADDR_PATH+"/pkg/:pkgname/:version/"
example :konfig v0.0.1 存放方式
/root/kpm/kpm.kusionstack.io/pkg/konfig/0.0.1/

//只能在cdn侧缓存，客户端不应缓存，实时请求
/root/kpm/kpm.kusionstack.io/pkg_tag/latest&&beta&&alpha&&v1-alpha

可强缓存
/root/kpm/kpm.kusionstack.io/pkg/:pkgname/v&&index/:version/src_tgz&&hash&&info&&index

/root/kpm/store/v1/00/:sha512
kpm包校验信息存放目录
$KPM_ROOT+"/"+$KPM_SERVER_ADDR_PATH+"/sum"
example :
/root/kpm/kpm.kusionstack.io/sum

kpm包校验信息存放方式
$KPM_ROOT+"/"+$KPM_SERVER_ADDR_PATH+"/sum/:pkgname/:version"
example :konfig v0.0.1 存放方式
/root/kpm/kpm.kusionstack.io/sum/konfig/0.0.1