package faceit_cc

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Handler struct {
	Database Database
	Kafka    KafkaInstance
}

type Entity interface {
	Add(args *User, reply *Response) error
	Update(args *User, reply *Response) error
	Delete(args *User, reply *Response) error
	List(args *User, reply *Response) error
}

func RegisterUser(server *rpc.Server, handler Handler) {
	server.RegisterName("user", &handler)
}

func (ht Handler) Add(args *User, reply *Response) error {
	log.Println("AddEntity -> Entity received")
	if !args.Validate() {
		err := fmt.Errorf("invalid input data")
		*reply = ResponseUnprocessable(err.Error())
		return err
	}
	err := args.Add(ht.Database.Db)
	if err != nil {
		*reply = ResponseError(err.Error())
		return err
	}

	if args.Kafka == nil {
		args.Kafka = &ht.Kafka
	}
	dataevent := DataEvent{
		User:   *args,
		Action: "Add",
	}
	if err := args.Kafka.Publish(dataevent.CreateString()); err != nil {
		*reply = ResponseError(err.Error())
		return err
	}

	*reply = ResponseOk(args.Id)
	return nil
}

func (ht Handler) Update(args *User, reply *Response) error {
	log.Println("UpdateEntity -> Entity received")
	if !args.Validate() {
		err := fmt.Errorf("invalid input data")
		*reply = ResponseUnprocessable(err.Error())
		return err
	}
	err := args.Update(ht.Database.Db)
	if err != nil {
		*reply = ResponseError(err.Error())
		return err
	}

	if args.Kafka == nil {
		args.Kafka = &ht.Kafka
	}
	dataevent := DataEvent{
		User:   *args,
		Action: "Update",
	}
	if err := args.Kafka.Publish(dataevent.CreateString()); err != nil {
		*reply = ResponseError(err.Error())
		return err
	}
	*reply = ResponseOk(args.Id)
	return nil
}

func (ht Handler) Delete(args *User, reply *Response) error {
	log.Println("DeleteEntity -> Entity received")
	if !ValidateUuid(args.Id) {
		err := fmt.Errorf("invalid input data")
		*reply = ResponseUnprocessable(err.Error())
		return err
	}
	err := args.Delete(ht.Database.Db)
	if err != nil {
		*reply = ResponseError(err.Error())
		return err
	}

	if args.Kafka == nil {
		args.Kafka = &ht.Kafka
	}
	dataevent := DataEvent{
		User:   *args,
		Action: "Delete",
	}
	if err := args.Kafka.Publish(dataevent.CreateString()); err != nil {
		*reply = ResponseError(err.Error())
		return err
	}
	*reply = ResponseOk(args.Id)
	return nil
}

func (ht Handler) List(args *User, reply *Response) error {
	log.Println("ListEntity -> Request received")
	args.Pagination.SetPagination()
	users, err := args.List(ht.Database.Db, args.Pagination)
	if err != nil {
		*reply = ResponseError(err.Error())
		return err
	}
	*reply = ResponseList(users)
	return nil
}

type BasicInterface interface {
	HealthCheck(args *int, reply *int) error
}

type BasicImplementation struct {
	Test int
}

func RegisterStatus(server *rpc.Server, imp BasicImplementation) {
	server.RegisterName("status", &imp)
}

func (t *BasicImplementation) HealthCheck(args *int, reply *int) error {
	log.Println("healthcheck method")
	*reply = http.StatusOK
	return nil
}
