package main

import (
    "fmt"
    "net"
    "sync"
    "os"
    "strconv"
)



var wg sync.WaitGroup
var server_ip = os.Args[1]

func tryConnection(port int, c chan bool, s chan int){	
    
    defer wg.Done()
    sport := strconv.Itoa(port)
    _, err := net.Dial("tcp", server_ip + ":" + sport)
    if err != nil {
        c <- false
        return
    }
	c <- true
	s <- port
	return
}

func main(){
    args := os.Args[2:]

	maxport, er := strconv.Atoi(args[1])
	minport, er2 := strconv.Atoi(args[0])
	
	if er != nil || er2 != nil {
		fmt.Println("[host minport maxport]")
		os.Exit(-1)
	} 
	
	numports := maxport-minport
	areopen := make(chan bool, numports)
	openports := make(chan int, numports)

	for i:=minport; i<maxport; i++{
		wg.Add(1)
		go tryConnection(i, areopen, openports)
	}
	
	wg.Wait()
	numopenports := len(openports)
	fmt.Printf("Total ports openned: %d\n", numopenports)
	for i:=0; i<numopenports; i++{
		fmt.Println(<-openports)			
	}
}
