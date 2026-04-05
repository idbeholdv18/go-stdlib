package main

import (
	"log"
	"os"
)

func main() {
	err := os.Setenv("MY_ENV", "hello, env")
	if err != nil {
		panic(err)
	}

	env, ok := os.LookupEnv("MY_ENV")
	if !ok {
		log.Fatal("MY_ENV not found")
	}

	log.Print(env)
}
