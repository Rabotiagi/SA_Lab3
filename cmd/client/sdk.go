package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"os"
	"lab3/pkg/sdk"
	"github.com/joho/godotenv"
)

var err = godotenv.Load(".env")

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := &sdk.Client{BaseUrl: "http://localhost:"+os.Getenv("PORT")}

	fmt.Println("=== Scenario 1 ===")
	balancers, err := client.ListBalancers(ctx)
	if err != nil {
		log.Fatal("Cannot list balancers: ", err)
	}
	fmt.Println("List of the balancers:")
	fmt.Println(balancers)

	fmt.Println("=== Scenario 2 ===")
	err = client.ChangeMachineState(ctx, 54, true)
	if err != nil {
		log.Fatal("Cannot change state: ", err)
	}
	fmt.Println("Status changed successfully")

	balancers, err = client.ListBalancers(ctx)
	if err != nil {
		log.Fatal("Cannot list balancers: ", err)
	}
	fmt.Println("New balancers list: ")
	fmt.Println(balancers)
}