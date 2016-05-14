package Library
///文件作用：
///存放缓存的Model，需要的时候拿出来用，都是指针类型的
///	ModelCache.Set("p", func() interface{} {return &Portal_user{}})
///
var ModelCache=&_modelCache{
	cache:make(map[string]func() interface{}),
}
// model info collection
type _modelCache struct {

	cache     map[string]func() interface{}

}
// get model  by table name
func (mc *_modelCache) Get(table string) (md func() interface{}, ok bool) {
	md, ok = mc.cache[table]
	return
}



// set model  to collection
func (mc *_modelCache) Set(table string, md func() interface{}) func() interface{} {
	mii := mc.cache[table]
	mc.cache[table] = md
	return mii
}
