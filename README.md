

代码参考   oasis项目的 develop_cvs_sensorskafka 分支



**使用sdk**

```	
import (
	oasispkgsensors "sensorsanalytics"
)

```

**注入**

```
type CollegeApplyUsecase struct {	
	sensorsAnalytics oasispkgsensors.SensorsAnalytics
}

...
// new
sa := oasispkgsensors.InitSensorsAnalytics(nil, cfgParams.GetSensorsProjectName(), false)
	return &CollegeApplyUsecase{
		sensorsAnalytics: sa,
	}
```

**生成 json 数据， 神策的 log agent**

topic ： sensors_data_channel

```

	// 使用公共pkg 的 sensors sdk  生成 json string bytes
	sendbytes, err := uc.sensorsAnalytics.TrackByIdGenSendBytes3(oasispkgsensors.Identity{Identities: map[string]string{
		"identity_uuid": uuid,
	}}, eventname, properties)
	
	if err != nil {
		uc.log.Errorw("SendSensorsdataTrack", "TrackByIdGenSendBytes3 生成数据失败", "project", projectname, "event", eventname, "distinctID", distinctID, "uuid", uuid, "properties", properties, "error", err)
		return nil
	}
	

	uc.sendKafkaRepo.SendSensorsData(ctx, sendbytes)


```

**发送kafka**

建立连接

* 不能用 withCodec("json") 参数  ，这会导致数据格式无法正确解析，注释掉吧

```

func (k *SensorsKafka) Run() broker.Broker {
	ctx := context.Background()
	bs := make([]broker.Option, 0, 0)
	bs = append(bs, broker.WithOptionContext(ctx))
	bs = append(bs, broker.WithAddress(k.conf.Broker.Addresses...))
	
	// 不能用 withCodec("json") 参数  ，这会导致数据格式无法正确解析，注释掉吧
	//bs = append(bs, broker.WithCodec("json"))
	
	
	bs = append(bs, kafka.WithBatchSize(1))
	bs = append(bs, kafka.WithAsync(true))
	
	...
}

```

发送消息

* withHeaders 参数传递自己的 app 名称,  

```


func (repo *sendToKafkaRepo) SendSensorsData(ctx context.Context, msg []byte) error {
	ctx, span := otel.Tracer("Data").Start(ctx, "sendToKafkaRepo.SendSensorsData", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	kafkatopic := repo.broker.SensorsKafkaConf.Topic

	err := repo.broker.SensorsKafka.Publish(kafkatopic, string(msg), kafka.WithHeaders(map[string]interface{}{
		"appName": "oasis-cvs",   // 替换成自己的 app 名称,   
	}))
	if err != nil {
		repo.logger.Errorw("SendSensorsData", string(msg), "topic", kafkatopic, "msg", "发送消息失败!!!", "error", err)
		return err
	}

	repo.logger.Infow("SendSensorsData", string(msg), "topic", kafkatopic, "msg", "success")

	return nil
}

```