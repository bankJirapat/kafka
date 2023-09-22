package main

// func main() {
// 	// configFile := "/Users/jirapat/Documents/bank/kafka/getting-started.properties"
// 	// conf, err := util.ReadConfig(configFile)
// 	// if err != nil {
// 	// 	fmt.Fprintf(os.Stderr, "Failed read config function: %s\n", err)
// 	// 	os.Exit(1)
// 	// }

// 	// fmt.Printf("> Running with config: %+v\n", conf)
// 	confProducer := util.SetProducerKafkaConfig()
// 	confConsumer := util.SetConsumerKafkaConfig()
// 	fmt.Printf("> Running with config producer: %+v\n", confProducer)
// 	fmt.Printf("> Running with config consumer: %+v\n", confConsumer)

// 	topic := "purchases"

// 	// if err := k.RunProducer(conf, topic); err != nil {
// 	// 	fmt.Fprintf(os.Stderr, "Failed run producer: %s\n", err)
// 	// 	os.Exit(1)
// 	// }

// 	// fmt.Println("============= Done Run Producer ================")

// 	// if err := k.RunConsumer(conf, topic); err != nil {
// 	// 	fmt.Fprintf(os.Stderr, "Failed run consumer: %s\n", err)
// 	// 	os.Exit(1)
// 	// }

// 	// fmt.Println("============= Done Run Consumer ================")
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		fmt.Println("============= Processing Run Producer ================")
// 		if err := k.RunProducer(confProducer, topic); err != nil {
// 			fmt.Fprintf(os.Stderr, "Failed run producer: %s\n", err)
// 			os.Exit(1)
// 		}
// 		wg.Done()
// 		fmt.Println("============= Done Run Producer ================")
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		// time.Sleep(1 * time.Second)
// 		fmt.Println("============= Processing Run Consumer ================")
// 		if err := k.RunConsumer(confConsumer, topic); err != nil {
// 			fmt.Fprintf(os.Stderr, "Failed run consumer: %s\n", err)
// 			os.Exit(1)
// 		}
// 		wg.Done()
// 		fmt.Println("============= Done Run Consumer ================")
// 	}()

// 	wg.Wait()
// }
