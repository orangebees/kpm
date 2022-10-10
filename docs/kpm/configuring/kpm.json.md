#包校验信息描述

```go
package main
type KpmFile struct {
	PackageName     string `json:"package_name"`
	KclvmMinVersion string `json:"kclvm_min_version"`
	//直接依赖，包名不重复
	Require []Require `json:"require"`
	//间接依赖，允许同包名不同版本
	RequireAll []Require `json:"require_all"`
}
type Require struct {
	
	PackageName    string `json:"package_name"`
	PackageVersion string `json:"package_version"`
	//校验和 sha512
	PackageIntegrity string `json:"package_integrity"`
	//对应下载类型的包地址
	PackageAddress string `json:"package_address"`
}
```