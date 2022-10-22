package main

import (
	"context"
	pb "github.com/djedjethai/clientGeneration0/pkg/proto/keyvalue"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:50051",
		grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewKeyValueClient(conn)

	var action, key, value string

	if len(os.Args) > 2 {
		action, key = os.Args[1], os.Args[2]
		value = strings.Join(os.Args[3:], "")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch action {
	case "get":
		r, err := client.Get(ctx, &pb.GetRequest{Key: key})
		if err != nil {
			log.Fatalf("could not get value for key %s: %v", key, err)
		}

		log.Printf("Get %s returns: %s", key, r.Value)
	case "put":
		_, err := client.Put(ctx, &pb.PutRequest{Key: key, Value: value})
		if err != nil {
			log.Fatalf("could not set value for key %s: %v", key, err)
		}

		log.Printf("Put: %s", key)
	case "delete":
		_, err := client.Delete(ctx, &pb.DeleteRequest{Key: key})
		if err != nil {
			log.Fatalf("could not delete value for key %s: %v", key, err)
		}

		log.Printf("Deleted: %s", key)
	case "getkeys":
		r, err := client.GetKeys(ctx, &pb.GetKeysRequest{})
		if err != nil {
			log.Fatalf("could not get keys: %v", err)
		}

		log.Printf("GetKeys returns: %s", r.Keys)
	default:
		log.Fatalf("Syntax: go run [get|put|delete|getkeys] key value ...")
	}

}
