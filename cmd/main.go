package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/bregydoc/plutus"
	"github.com/bregydoc/plutus/bridges/culqi"
	"github.com/bregydoc/plutus/delivery/smtp"
	proto "github.com/bregydoc/plutus/proto"
	"github.com/bregydoc/plutus/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const configFilename = "plutus.config.yml"

func main() {

	config, err := readConfigFile(configFilename)
	if err != nil {
		panic(err)
	}

	var bridge plutus.PaymentBridge

	switch config.Bridge.Backend {
	case "culqi":
		bridge, err = culqi.NewBridge(config.Bridge.PublicKey, config.Bridge.PrivateKey)
		if err != nil {
			panic(err)
		}
	case "stripe":
		panic("unimplemented bridge backend (WIP)")
	case "visanet":
		panic("unimplemented bridge backend (WIP)")
	default:
		panic("invalid payments bridge backend, please read the plutus documentation")
	}

	channels := make([]plutus.DeliveryChannel, 0)
	for channelName, channelConfig := range config.Delivery {
		if channelName == "smtp" {
			templateFile := "template.html"
			ch, err := smtp.NewSMTPDeliver(smtp.Config{
				Host:     channelConfig.Host,
				Username: channelConfig.Username,
				Password: channelConfig.Password,
				From:     channelConfig.From,
				Port:     587,
			}, templateFile)
			if err != nil {
				panic(err)
			}
			channels = append(channels, ch)
		}
	}

	repo, err := repository.NewBoltRepository("./plutus.db")
	if err != nil {
		panic(err)
	}

	company := &plutus.Company{
		Name: "Plutus Sales",
	}

	if config.Company.Name != "" {
		company.Name = config.Company.Name
	}

	if config.Company.OfficialWeb != "" {
		company.OfficialWeb = config.Company.OfficialWeb
	}

	if config.Company.Custom != nil {
		company.Custom = config.Company.Custom
	}

	if config.Company.SupportEmail != "" {
		company.Support.Email = config.Company.SupportEmail
	}

	if config.Company.SupportPhone != "" {
		company.Support.Phone = config.Company.SupportPhone
	}

	engine := &plutus.SalesEngine{
		Bridge:           bridge,
		DeliveryChannels: channels,
		Repository:       repo,
		Company:          company,
	}

	var servicePort int64 = 18000

	if config.Port != 0 {
		servicePort = config.Port
	}

	// Check if we have TLS a secure connection
	withTLS := true
	certificate := "/run/secrets/cert"
	key := "/run/secrets/key"

	_, err = os.Open(certificate)
	if err != nil || os.IsNotExist(err) {
		withTLS = false
	}

	if withTLS {
		_, err = os.Open(key)
		if err != nil || os.IsNotExist(err) {
			withTLS = false
		}
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", servicePort))
	if err != nil {
		log.Fatalf("[For Developers] Failed to listen: %v", err)
	}

	var grpcServer *grpc.Server

	if withTLS {
		log.Println("setting with TLS from \"" + certificate + "\" and \"" + key + "\"")
		c, err := credentials.NewServerTLSFromFile(certificate, key)
		if err != nil {
			log.Fatalf("[For Developers] Failed to setup tls: %v", err)
		}
		grpcServer = grpc.NewServer(
			grpc.Creds(c),
		)
	} else {
		log.Println("setting without security")
		grpcServer = grpc.NewServer()
	}

	proto.RegisterPlutusServer(grpcServer, engine)

	log.Printf("[For Developers] GRPC listening on :%d\n", servicePort)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
