# Logmanage
日志管理平台


## Etcd添加Logpath和Topic
![Etcd](https://github.com/kungehero/Logmanage/blob/master/image/logweb.png)




## 日志收集平台流程图：
![log](https://github.com/kungehero/Logmanage/blob/master/image/log.png)

# sample

## LogAgent

```

func ReadLogWithTailf() {
	ch := make(chan bool)
	for _, LogpathStruct := range models.EtcdValue {
		models.NewTailf(LogpathStruct.LogPath)
		go LogRun(models.Tails, LogpathStruct.Topic, ch)
	}

	for v := range ch {
		fmt.Println(v)
	}
}

```

## LogTransfer
```
func ReadLogWithEs() {
	ch := make(chan bool)
	for _, LogpathStruct := range models.EtcdValue {
		Run(LogpathStruct.Topic, ch)
	}

	for v := range ch {
		fmt.Println(v)
	}
}
```
