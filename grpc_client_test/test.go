package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "gitlab.silkrode.com.tw/golang/kbc2/proto/order"
	"google.golang.org/grpc"
)

const grpcAddr = "localhost:7070"

type GrpcClient struct {
	ctx  context.Context
	conn *grpc.ClientConn
}

func InitGrpcClient(grpcConn *grpc.ClientConn) *GrpcClient {
	return &GrpcClient{
		ctx:  context.Context(context.Background()),
		conn: grpcConn,
	}
}

func main() {
	grpcConn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to grpc server.", err)
	}

	gc := InitGrpcClient(grpcConn)

	var testCode int
	fmt.Print("Choose a case (Enter the number):" + "\n" +
		"(1) Create Main Order\n" +
		"(2) Create Sub Order\n" +
		"(3) Create Order History\n" +
		"(4) Update Main Order\n" +
		"(5) Update Sub Order\n" +
		"(6) Query Main Order\n" +
		"(7) Query Sub Order\n" +
		"(8) Query Order History\n" +
		"(9) Quit\n")
	fmt.Scanf("%d", &testCode)

	switch testCode {
	case 1:
		gc.CreateMainOrder()
	case 2:
		gc.CreateSubOrder()
	case 3:
		gc.CreateOrderHistory()
	case 4:
		gc.UpdateMainOrder()
	case 5:
		gc.UpdateSubOrder()
	case 6:
		gc.QueryMainOrder()
	case 7:
		gc.QuerySubOrder()
	case 8:
		gc.QueryOrderHistory()
	}
}

func (gc *GrpcClient) CreateMainOrder() {
	c := pb.NewOrderClient(gc.conn)
	res, err := c.CreateMainOrder(gc.ctx, &pb.MainOrder{
		PaymentNumber:   "123",
		WithdrawAccount: "1234567890",
		CompletedAt:     uint64(time.Now().Unix()),
		ExpiredAt:       uint64(time.Now().Unix()),
		DeletedAt:       uint64(time.Now().Unix()),
		PaymentType:     pb.PaymentType_BANKCARD,
	})
	if err != nil {
		fmt.Println("CreateMainOrder Err")
	}

	fmt.Println("Get PaymentNumber: ", res.PaymentNumber)
}

func (gc *GrpcClient) CreateSubOrder() {
	c := pb.NewOrderClient(gc.conn)
	res, err := c.CreateSubOrder(gc.ctx, &pb.SubOrder{
		SubTrackingNumber: "1211",
		ChannelID:         "1233333",
		ChannelName:       "321",
	})
	if err != nil {
		fmt.Println("CreateSubOrder Err")
	}

	fmt.Println("Get GetChannelID: ", res.ChannelID)
}

func (gc *GrpcClient) CreateOrderHistory() {
	c := pb.NewOrderClient(gc.conn)
	res, err := c.CreateOrderHistory(gc.ctx, &pb.OrderHistory{
		ChannelName: "321",
		IsSuccess:   true,
	})
	if err != nil {
		fmt.Println("CreateOrderHistory Err")
	}

	fmt.Println("Get ChannelName: ", res.ChannelName)
}

func (gc *GrpcClient) UpdateMainOrder() {
	c := pb.NewOrderClient(gc.conn)
	res, err := c.UpdateMainOrder(gc.ctx, &pb.MainOrder{
		TrackingNumber:  "btlfu33ipt3dejl1s69g",
		PaymentNumber:   "987",
		WithdrawAccount: "000",
	})
	if err != nil {
		fmt.Println("UpdateMainOrder Err")
	}

	fmt.Println("Get ChannelName: ", res.PaymentNumber)
}

func (gc *GrpcClient) UpdateSubOrder() {
	c := pb.NewOrderClient(gc.conn)
	res, err := c.UpdateSubOrder(gc.ctx, &pb.SubOrder{
		SubTrackingNumber: "123333",
		ChannelID:         "000",
		ChannelName:       "000",
	})
	if err != nil {
		fmt.Println("UpdateMainOrder Err")
	}

	fmt.Println("Get ChannelName: ", res.ChannelName)
}

func (gc *GrpcClient) QueryMainOrder() {
	c := pb.NewOrderClient(gc.conn)
	res, err := c.QueryMainOrder(gc.ctx, &pb.TrackingNumber{
		TrackingNumber: "btlfu33ipt3dejl1s69g",
	})
	if err != nil {
		fmt.Println("QueryMainOrder Err")
	}

	fmt.Println("Get main order: ", res)
	fmt.Println("Get WalletStatus: ", res.WalletStatus)
}

func (gc *GrpcClient) QuerySubOrder() {
	c := pb.NewOrderClient(gc.conn)
	res, err := c.QuerySubOrder(gc.ctx, &pb.SubTrackingNumber{
		SubTrackingNumber: "123333",
	})
	if err != nil {
		fmt.Println("QuerySubOrder Err")
	}

	fmt.Println("Get sub order: ", res)
	fmt.Println("Get ChannelID: ", res.ChannelID)
}

func (gc *GrpcClient) QueryOrderHistory() {
	c := pb.NewOrderClient(gc.conn)
	res, err := c.QueryOrderHistory(gc.ctx, &pb.TrackingNumber{
		TrackingNumber: "1233",
	})
	if err != nil {
		fmt.Println("QueryOrderHistory Err")
	}

	fmt.Println("Get order history: ", res)
	fmt.Println("Get ChannelName: ", res.ChannelName)
}
