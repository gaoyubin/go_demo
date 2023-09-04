package main

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

/*
func main() {
	// config := sarama.NewConfig()
	// // sarama.APIKeySASLAuth()
	// fmt.Println(config)
	// sarama.
	// // sarama.

	// // config.Producer.

	// fmt.Println()

	//1. 生产者配置
	// config := sarama.NewConfig()
	// config.Producer.RequiredAcks = sarama.WaitForAll
	// config.Producer.Partitioner = sarama.NewRandomPartitioner
	// config.Producer.Return.Successes = true

	// //2. 连接kafka
	// client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	// if err != nil {
	// 	fmt.Println("produce err", err)
	// 	return
	// }
	// defer client.Close()

	// //3. 封装消息
	// msg := &sarama.ProducerMessage{}
	// msg.Topic = "test1"
	// msg.Value = sarama.StringEncoder("this is test log")
	// pid, offset, err := client.SendMessage(msg)
	// if err != nil {
	// 	fmt.Println("send err", err)
	// 	return
	// }
	// fmt.Println("send succ", pid, offset)

	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	// sarama.NewConsumerGroup()
	if err != nil {
		fmt.Println("start consumer", err)
		return
	}

	partitionList, err := consumer.Partitions("test1")
	if err != nil {
		fmt.Println("get partition err", err)
		return
	}

	fmt.Println("show partitionlist", partitionList)
	for partition := range partitionList {
		fmt.Println("show partition", partition)
		pc, err := consumer.ConsumePartition("test1", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("get con partition err", err)
			return
		}
		defer pc.AsyncClose()

		go func(sarama.PartitionConsumer) {
			fmt.Println("show msg")
			for msg := range pc.Messages() {
				fmt.Println("partiton:", msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}

	var ss string
	fmt.Scanln(&ss)
}
*/

func main() {
	config := sarama.NewConfig()
	group, err := sarama.NewConsumerGroup([]string{"127.0.0.1:9092"}, "g1", config)
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithCancel(context.Background())
	var k Kafka
	for {
		err = group.Consume(ctx, []string{"test2"}, &k)
		if err != nil {
			log.Println("group.Consume error: err", err)
			panic(err)
		}
	}
}

type Kafka struct {
}

func (k *Kafka) ConsumeClaim(ss sarama.ConsumerGroupSession, sc sarama.ConsumerGroupClaim) error {
	log.Println("kafka init...", sc)
	return nil
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (k *Kafka) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("kafka init...", session)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (k *Kafka) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("cleanup")
	return nil
}
