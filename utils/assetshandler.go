package utils

import "strings"

var fileSplitStr = "__$new$__"

func AddNewUrl(assets string, newAssetUrl string) (result string) {
	assetsArr := strings.Split(assets, fileSplitStr)
	newArr := strings.Join(append(assetsArr, newAssetUrl), fileSplitStr)
	return newArr
}

func GetAssetsArr(assets string) (result []string) {
	assetsArr := strings.Split(assets, fileSplitStr)
	return assetsArr
}
