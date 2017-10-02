package main

import(
"github.com/stianeikeland/go-rpio"
"time"
"fmt"
)
//DC1: 16-19
//DC2: 20-26
//LED: 21
func Initiate(){
        dc1_a = rpio.Pin(16)
        dc1_b = rpio.Pin(19)
        dc2_a = rpio.Pin(20)
        dc2_b = rpio.Pin(26)
        led = rpio.Pin(21)
        dc1_a.Output()
        dc1_b.Output()
        dc2_a.Output()
        dc2_b.Output()
        led.Output()
}

func Lightson(){
		lit = true
		led.High()
}

func Lightsoff(){
		lit = false
		led.Low()
}

func Forward(wait int){
      fmt.Println("Going forward")
			dc1_a.High()
			dc1_b.Low()
			dc2_a.High()
			dc2_b.Low()
			time.Sleep(wait*time.Millisecond)
			reset()
}

func Right(wait int){
      fmt.Println("Going right")
			dc1_a.Low()
			dc1_b.High()
			dc2_a.High()
			dc2_b.Low()
			time.Sleep(wait*time.Millisecond)
			reset()
}

func Left(wait int){
      fmt.Println("Going left")
			dc1_a.High()
			dc1_b.Low()
			dc2_a.Low()
			dc2_b.High()
			time.Sleep(wait*time.Millisecond)
			reset()
}

func Backwards(wait int){
      fmt.Println("Going backwards")
			dc1_a.Low()
			dc1_b.High()
			dc1_a.Low()
			dc1_b.High()
			time.Sleep(wait*time.Millisecond)
			reset()
		}
}

func reset(){
        dc1_a.Low()
        dc1_b.Low()
        dc2_a.Low()
        dc2_b.Low()
}
