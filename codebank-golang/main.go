package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/codeedu/codebank/infra/grpc/server"
	"github.com/codeedu/codebank/infra/kafka"
	"github.com/codeedu/codebank/infra/repository"
	"github.com/codeedu/codebank/usecase"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	db := setupDb()
	defer db.Close()
	producer := setupKafkaProducer()
	processTransactionUseCase := setupTransactionUseCase(db, producer)
	fmt.Println("Rodando gRPC Serve")
	serverGrpc(processTransactionUseCase)

}

func setupTransactionUseCase(db *sql.DB, producer kafka.KafkaProducer) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	useCase.KafkaProducer = producer
	return useCase
}

func setupDb() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s  port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)

	db, err := sql.Open("postgres", sqlInfo)

	if err != nil {
		log.Fatal("error connect postgres")
	}
	return db
}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KafkaBootstrapServer"))
	return producer
}

func serverGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
	grpcServer := server.NewGRPCServer()
	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
	grpcServer.Serve()
}

// cc := domain.NewCreditCard()
// 	cc.Number = "1234"
// 	cc.Name = "joao"
// 	cc.ExpirationYear = 2024
// 	cc.ExpirationMonth = 7
// 	cc.CVV = 123
// 	cc.Limit = 1000
// 	cc.Balance = 0

// 	repo := repository.NewTransactionRepositoryDb(db)
// 	err := repo.CreateCreditCard(*cc)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
