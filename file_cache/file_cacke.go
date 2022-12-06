package cache

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

//使用场景:可以从数据库中拉去任务 缓存在本地进行操作
type fileCache struct {
	fileName string
	data     map[string]string
	dataLock sync.RWMutex
}

//创建对象
func NewFileCache(fileName string) *fileCache {
	return &fileCache{
		fileName: fileName,
		data:     make(map[string]string),
	}
}

// 存到文件
func (fc *fileCache) SaveToFile() {
	var fileBuffer = strings.Builder{}
	for k, v := range fc.data {
		fileBuffer.WriteString(fmt.Sprintf("%v=%v\n", k, v))
	}
	ioutil.WriteFile(fc.fileName, []byte(fileBuffer.String()), 0644)
}

//加载文件中的缓存
func (fc *fileCache) GetFromFile() map[string]string {
	fc.dataLock.RLock()
	defer fc.dataLock.RUnlock()
	//读取信息
	buffer, err := ioutil.ReadFile(fc.fileName)
	if err != nil {
		return map[string]string{}
	}

	var dataMap = make(map[string]string)
	bufferLines := strings.Split(string(buffer), "\n")
	for _, val := range bufferLines {
		var idx = strings.Index(val, "=")
		if idx < 0 {
			continue
		}
		dataMap[val[:idx]] = val[idx+1:]
	}
	return dataMap
}

// 获取值
func (fc *fileCache) GetStrFromFile(_key string) string {
	fc.dataLock.RLock()
	defer fc.dataLock.RUnlock()

	buffer, err := ioutil.ReadFile(fc.fileName)
	if err != nil {
		return ""
	}

	bufferLines := strings.Split(string(buffer), "\n")
	for _, val := range bufferLines {
		var idx = strings.Index(val, "=")
		if idx < 0 {
			continue
		}
		if val[:idx] == _key {
			return val[idx+1:]
		}
	}
	return ""
}

// 从缓存中获取信息
func (fc *fileCache) GetStrFromCache(_key string) string {
	if val, ok := fc.data[_key]; ok {
		return val
	}
	return ""
}

//批量删除文件中的key
func (fc *fileCache) DelKeys(_keys []string) {
	var fullMap = fc.GetFromFile()

	fc.dataLock.Lock()
	defer fc.dataLock.Unlock()
	for _, k := range _keys {
		delete(fullMap, k)
	}
	fc.data = fullMap
	fc.SaveToFile()
}

// 设置值
func (fc *fileCache) SetItem(_key string, _val interface{}) {
	var fullMap = fc.GetFromFile()

	fc.dataLock.Lock()
	defer fc.dataLock.Unlock()
	fullMap[_key] = fmt.Sprintf("%v", _val)
	fc.data = fullMap
	fc.SaveToFile()
}

// 清空
func (fc *fileCache) Clear() {
	fc.dataLock.Lock()
	defer fc.dataLock.Unlock()
	err := os.Remove(fc.fileName)

	if err != nil {
		// 删除失败
		fmt.Println("删除缓存文件失败", err.Error())

	} else {
		// 删除成功
		fmt.Println("删除缓存文件成功", fc.fileName)
	}

}
