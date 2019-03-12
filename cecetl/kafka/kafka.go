package kafka

import (
	"zdy.local/utils/nlogs"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
	"strings"
	"zdy.local/cecetl/message"
	"encoding/json"
	"sync"
)

type KafkaClient struct {
	Producer sarama.SyncProducer
	Consumer sarama.Consumer
	PartitionList []int32
	//PartitionCon sarama.PartitionConsumer
	enqueued, errors int
	m *sync.Mutex
}

var (
	wg     sync.WaitGroup
	goroutine chan int
)

func init()  {
	goroutine_num,_ := beego.AppConfig.Int("golang::goroutine")
	goroutine = make(chan int,goroutine_num)
}

func (k *KafkaClient)InitPorduce(){
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	var err error
	k.Producer, err = sarama.NewSyncProducer(strings.Split(beego.AppConfig.String("kafka::host"),","), config)
	if err != nil {
		panic(err)
	}
	k.enqueued = 0
	k.errors = 0
}

func (k *KafkaClient)InitConsumer(){
	var err error
	k.Consumer, err = sarama.NewConsumer(strings.Split(beego.AppConfig.String("kafka::host"),","), nil)
	if err != nil {
		panic(err)
	}
	k.PartitionList, err = k.Consumer.Partitions("upload")
	if err != nil {
		panic(err)
	}
	k.m = new(sync.Mutex)
	/*k.PartitionCon, err = k.Consumer.ConsumePartition("upload", 5, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}*/
}

func (k *KafkaClient)Producter(topic,key,value string,){
	/*select {
	case k.Producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: sarama.StringEncoder(key), Value: sarama.ByteEncoder(value)}:
		//nlogs.ConsoleLogs.Debug("partition=%d, offset=%d, key=%s\n", partition, offset,msg.Key)
		k.enqueued++
	case err := <-k.Producer.Errors():
		log.Println("Failed to produce message", err)
		k.errors++
	}*/
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Partition = int32(-1)
	msg.Key = sarama.StringEncoder(key)
	msg.Value = sarama.ByteEncoder(value)
	//defer k.Producer.Close()
	partition, offset, err := k.Producer.SendMessage(msg)
	if err != nil {
		nlogs.FileLogs.Error("Failed to produce message: ", err)
	}
	nlogs.ConsoleLogs.Debug("partition=%d, offset=%d, key=%s\n", partition, offset,msg.Key)
	k.enqueued++
}

func (k *KafkaClient)CloseProducter(){
	nlogs.ConsoleLogs.Info("produce succeed message",k.enqueued)
	//nlogs.ConsoleLogs.Info("produce failed message",k.errors)
	if err := k.Producer.Close(); err != nil {
		nlogs.ConsoleLogs.Error("producer close error",err)
		nlogs.FileLogs.Error("producer close error",err)
	}
}

func (k *KafkaClient) Consume(obj CliImpl){
	num:=0
	//defer obj.Close()
	for partition := range k.PartitionList {
		pc, err := k.Consumer.ConsumePartition("upload", int32(partition), sarama.OffsetNewest)
		if err != nil {
			nlogs.ConsoleLogs.Error("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		defer pc.AsyncClose()
		goroutine <- 1
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				nlogs.ConsoleLogs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				var result message.Private
				err := json.Unmarshal(msg.Value, &result)
				if err != nil{
					nlogs.FileLogs.Error("get message error",string(msg.Value))
				}else {
					nlogs.ConsoleLogs.Debug(string(msg.Key))
					obj.GetStudyUID(string(msg.Key),result.SickInfo.PatientUID,&result)
					obj.LoadData(string(msg.Key),result)
					k.m.Lock()
					num++
					k.m.Unlock()
					//nlogs.ConsoleLogs.Alert("Do consuming topic upload",nlogs.Fmt2String(num))
				}
			}
			<-goroutine
			wg.Done()
		}(pc)
	}
	wg.Wait()
	nlogs.ConsoleLogs.Warn("Done consuming topic upload",nlogs.Fmt2String(num))
	defer k.Consumer.Close()
	/*num:=0
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	ConsumerLoop:
	for {
		select {
		case msg := <-k.PartitionCon.Messages():
			num++
			nlogs.ConsoleLogs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			wg.Add(1)
			goroutine <- 1
			go func(*sarama.ConsumerMessage) {
				var result message.Private
				err := json.Unmarshal(msg.Value, &result)
				if err != nil{
					nlogs.FileLogs.Error("get message error",string(msg.Value))
					<-goroutine
					wg.Done()
				}else {

					nlogs.ConsoleLogs.Debug(string(msg.Key))
					obj.GetStudyUID(string(msg.Key),result.SickInfo.PatientUID,&result)
					obj.LoadData(string(msg.Key),result)
					<-goroutine
					wg.Done()
					//nlogs.ConsoleLogs.Info("Do consuming topic upload",nlogs.Fmt2String(num))
				}
			}(msg)
		case <-signals:
			//obj.Close()
			break ConsumerLoop
		}
	}
	nlogs.ConsoleLogs.Warn("Done consuming topic upload",nlogs.Fmt2String(num))
	wg.Wait()
	k.Consumer.Close()*/
}