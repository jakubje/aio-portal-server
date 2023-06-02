package database

import (
	"database/sql"
	"fmt"
	"github.com/jakub/aioportal/server/internal/models"
	"github.com/jakub/aioportal/server/internal/utils"
	"io"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/joho/godotenv"
)

const dbName = "aio-portal"
const collectionCoins = "coins"
const collectionUsers = "users"

var secretKey string

func init() {

	//// Load env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connectionString := os.Getenv("MONGODB_URI")
	//secretKey = os.Getenv("SECRET_KEY")
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")

	_, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	// store := db.newStore(conn)

}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil

}

func GetUserByCredentials(body io.ReadCloser) (*models.UserRegister, error) {
	var loginUser models.LoginUser
	var user models.User
	//err := json.NewDecoder(body).Decode(&loginUser)
	//
	//filter := bson.M{"email": loginUser.Email}
	//
	//if err != nil {
	//	return nil, fmt.Errorf("Empty information or invalid email")
	//}
	//
	//err = userCollection.FindOne(context.Background(), filter).Decode(&user)
	//
	//if err != nil {
	//	return nil, fmt.Errorf("User does not exist, check your credentials information")
	//}

	if !utils.ComparePassword(user.Password, []byte(loginUser.Password)) {
		return nil, fmt.Errorf("Invalid password")
	}

	tokenAuth := jwtauth.New("HS256", []byte(secretKey), nil)

	claims := jwt.MapClaims{"_id": user.ID, "name": user.Name, "lastName": user.LastName, "email": user.Email}

	_, tokenString, _ := tokenAuth.Encode(claims)

	registerUser := models.UserRegister{
		Email:    user.Email,
		Name:     user.Name,
		LastName: user.LastName,
		Token:    tokenString,
	}

	return &registerUser, nil

}

func CreateUser(body io.ReadCloser) (*models.UserRegister, error) {
	var user models.User
	//err := json.NewDecoder(body).Decode(&user)
	//
	//if err != nil {
	//	return nil, fmt.Errorf("Error body request")
	//}
	//
	tokenAuth := jwtauth.New("HS256", []byte(secretKey), nil)
	//
	//user.Password, _ = utils.HashPasswod(user.Password)
	//fmt.Println(user)
	//res, err := userCollection.InsertOne(context.Background(), user)
	//fmt.Println(err)
	//
	//if err != nil {
	//	return nil, fmt.Errorf("Error registering the user")
	//}

	//user.ID = res.InsertedID.(primitive.ObjectID)

	claims := jwt.MapClaims{"_id": user.ID, "name": user.Name, "lastName": user.LastName, "email": user.Email}
	_, tokenString, _ := tokenAuth.Encode(claims)

	registerUser := models.UserRegister{
		Email:    user.Email,
		Name:     user.Name,
		LastName: user.LastName,
		Token:    tokenString,
	}

	return &registerUser, nil
}

func CreateCoin(body io.ReadCloser) {}

func AddCoin(body io.ReadCloser) {}
