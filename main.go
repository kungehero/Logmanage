package main

import "Logmanage/logcollect"

type Tweet struct {
	User     string
	Message  string
	Retweets int
}

func main() {
	/* strfile := []string{"C:/logs/log/log1.txt", "C:/logs/log/log2.txt", "C:/logs/log/log3.txt"}
	var tails *tail.Tail
	ch := make(chan bool)

	for _, file := range strfile {
		fmt.Println(file)
		tails, _ = tail.TailFile(file, tail.Config{
			ReOpen:    true,
			Follow:    true,
			MustExist: false,
			Poll:      true,
		})

		go func(tf *tail.Tail, chdemo chan bool) {
			for true {
				select {
				case msg, ok := <-tf.Lines:
					if !ok {
						time.Sleep(100 * time.Millisecond)
						continue
					}
					fmt.Println(msg.Text)
				}
			}
			chdemo <- true
		}(tails, ch)
	}
	for v := range ch {
		fmt.Println(v)
	} */

	logcollect.ReadLogWithEs()
}
