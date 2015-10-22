package main

import (
	"log"
	_ "testing"
	"time"
)

func ExampleUser() {

	pl := NewPeriodicLimiter(time.Second / 4)
	ap, _ := NewLogAmmoProvider(8)
	rl, _ := NewLoggingResultListener()
	u := &User{
		name:       "Example user",
		ammunition: ap,
		results:    rl,
		limiter:    pl,
		done:       make(chan bool),
		gun:        &LogGun{},
	}
	go u.run()
	u.ammunition.Start()
	u.results.Start()
	u.limiter.Start()
	<-u.done

	log.Println("Done")
	// Output:
}
