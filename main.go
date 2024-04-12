package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/LuisDaga07/gambituser/awsgo"
	"github.com/LuisDaga07/gambituser/bd"
	"github.com/LuisDaga07/gambituser/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros. debe enviar SecretManager")
		err := errors.New("Error en los parametros debe enviar SecrertName")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("sub = " + datos.UserUUID)

		}
	}
	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el Secret: " + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)
	return event, err

}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro

}
