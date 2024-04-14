package cache

type Provider interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
	Delete(key string)
}
