package values

type Values map[string][]string

// Get 定义一个方法 接收器是Values类型 表示从map里面获取value
func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// Add ...
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
