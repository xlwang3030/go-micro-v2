package Boot

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"log"
	"os"
	"time"
)
type DataConfig struct {
	Mysql *MySqlConfig
	Redis *RedisConfig
}
type GlobalConfig struct {
    Config *struct{
		Address string
		Path string
		Port uint64
	}
    Service *struct{
		Namespace string
		Name string
	}
    Data *DataConfig
}
var JConfig *GlobalConfig
var nacosClient  config_client.IConfigClient
//初始化主文件（本地)配置 和初始化nacos连接
func InitConfig()  {
	configFile:="app.yaml"
	err:=config.LoadFile(configFile)
	if err!=nil{
		log.Fatal(err)
	}
	JConfig=&GlobalConfig{Data:new(DataConfig)}
	err=config.Get("jtthink").Scan(JConfig)
	if err!=nil{
		log.Fatal(err)
	}
	serverConfigs := []constant.ServerConfig{
		{ IpAddr: JConfig.Config.Address, ContextPath: JConfig.Config.Path, Port: JConfig.Config.Port},
	}
	nacosClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	if err!=nil{
		log.Fatal("nacos init err",err )
	}

	//开始加载数据相关的配置
	JConfig.Data.Mysql=new(MySqlConfig)
    listenNacos("jt-dataconfig-mysql","jtthink",JConfig.Data.Mysql,true )


}

func listenNacos(dataid string,group string,model ConfigInterface,reload bool){
	err:= nacosClient.ListenConfig(vo.ConfigParam{
		DataId: dataid,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			time.Sleep(time.Second*1)
			shouldReload:=reload
			if  !model.IsLoad(){
				shouldReload=false  //如果model 没有被加载过,则不需要做重载
			}

			cacheFile:=fmt.Sprintf("./runtime/configcache/%s-%s.yaml",group,dataid)
			file,err:=os.OpenFile(cacheFile,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
			if err!=nil{
				logger.Error(err)
				return
			}
			defer  file.Close()
			_,err=file.WriteString(data)
			if err!=nil{
				logger.Error(err)
				return
			}
			err=config.LoadFile(cacheFile)
			if err!=nil{
				logger.Error(err)
				return
			}
			err=config.Scan(model)
			if err!=nil{
				logger.Error(err)
				return
			}
			if shouldReload{ //重载关键代码
				err:=model.Reload()
				if err!=nil{
					logger.Error(err)
					return
				} else{
					logger.Info(dataid," 重载完成")
				}
			}

		},
	})
	if err!=nil{
		logger.Error("listen config error,dataid:",dataid,err)
	}
}