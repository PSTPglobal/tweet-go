package bd

import (
	"context"
	"time"
	"github.com/PSTPglobal/tweet-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro( usuarioModel models.Usuario ) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoCN.Database("tweet-go")
	col := db.Collection("usuarios")

	usuarioModel.Password, _ = EncriptarPassword( usuarioModel.Password )

	result, err := col.InsertOne(ctx, usuarioModel)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
