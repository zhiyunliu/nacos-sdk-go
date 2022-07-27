package clients

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhiyunliu/nacos-sdk-go/common/constant"
	"github.com/zhiyunliu/nacos-sdk-go/vo"
)

func TestSetConfigClient(t *testing.T) {

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			"console.nacos.io",
			80,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos")),
	}

	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("e525eafa-f7d7-4029-83d9-008937f9d468"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)

	t.Run("setConfig_error", func(t *testing.T) {
		nacosClient, err := setConfig(vo.NacosClientParam{})
		assert.Nil(t, nacosClient)
		assert.Equal(t, "server configs not found in properties", err.Error())
	})

	t.Run("setConfig_normal", func(t *testing.T) {
		// use map params setConfig
		param := getConfigParam(map[string]interface{}{
			"serverConfigs": sc,
			"clientConfig":  cc,
		})
		nacosClientFromMap, err := setConfig(param)
		assert.Nil(t, err)
		nacosClientFromStruct, err := setConfig(vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		})
		assert.Nil(t, err)
		assert.True(t, reflect.DeepEqual(nacosClientFromMap, nacosClientFromStruct))
	})

}
