package main

import (
	"fmt"
	"kafka/util"
	"os"
	"sync"

	k "kafka/kafka"
)

func main() {
	configFile := "/Users/jirapat/Documents/bank/kafka/getting-started.properties"
	conf, err := util.ReadConfig(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed read config function: %s\n", err)
		os.Exit(1)
	}

	topic := "test-purchases"
	conf["group.id"] = "test-consumer-group"
	conf["auto.offset.reset"] = "earliest"

	// if err := k.RunProducer(conf, topic); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed run producer: %s\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("============= Done Run Producer ================")

	// if err := k.RunConsumer(conf, topic); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed run consumer: %s\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("============= Done Run Consumer ================")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("============= Processing Run Producer ================")
		if err := k.RunProducer(conf, topic); err != nil {
			fmt.Fprintf(os.Stderr, "Failed run producer: %s\n", err)
			os.Exit(1)
		}
		wg.Done()
		fmt.Println("============= Done Run Producer ================")
	}()

	wg.Add(1)
	go func() {
		// time.Sleep(1 * time.Second)
		fmt.Println("============= Processing Run Consumer ================")
		if err := k.RunConsumer(conf, topic); err != nil {
			fmt.Fprintf(os.Stderr, "Failed run consumer: %s\n", err)
			os.Exit(1)
		}
		wg.Done()
		fmt.Println("============= Done Run Consumer ================")
	}()

	wg.Wait()
}
