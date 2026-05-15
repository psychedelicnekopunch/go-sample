
package main


import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)



type Memcahced struct {
	Clinet *memcache.Client
}


func NewMwmcached() *Memcahced {
	return &Memcahced{
		Clinet: memcache.New("go-sample-memcached:11211"),
	}
}


func main() {
	// mc := memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212")
	// mc := memcache.New("go-sample-memcached:11211")
	// mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	// it, err := mc.Get("foo")

	// if err != nil {
	// 	fmt.Print(err.Error(), "\n")
	// 	return
	// }

	// fmt.Print(it, "\n")

	mc := NewMwmcached()
	// if err := mc.DeleteAll(); err != nil {
	// 	fmt.Print(err.Error(), "\n")
	// 	return
	// }

	mc.Set("test", "hello", 5)
	value1, _ := mc.Get("test")
	fmt.Print("---- value1 ----\n", value1, "\n")
	_ = mc.Delete("test")
	value2, _ := mc.Get("test")
	fmt.Print("---- value2 ----\n", value2, "\n")
	fmt.Print("----done----", "\n")
}


func (mc *Memcahced) Get(key string) (value string, err error) {
	item, err := mc.Clinet.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

func (mc *Memcahced) Set(key string, value string, expire int32) {
	// Expiration: 秒
	mc.Clinet.Set(&memcache.Item{
		Key: key,
		Value: []byte(value),
		Expiration: expire,
	})
}

func (mc *Memcahced) Delete(key string) error {
	return mc.Clinet.Delete(key)
}

func (mc *Memcahced) DeleteAll() error {
	return mc.Clinet.DeleteAll()
}
