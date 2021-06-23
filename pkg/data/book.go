package data

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
	//dbConst "github.com/PutskouDzmitry/golang-training-Library/pkg/const_db"
)

//Entity in database
type Book struct {
	BookId            primitive.ObjectID `bson:"_id"`
	AuthorId          int                `bson:"author_id"`
	BookVolume        int                `bson:"book_volume"`
	NameOfBook        string             `bson:"name_of_book"`
	Number            int                `bson:"number"`
	PublisherId       int                `bson:"publisher_id"`
	YearOfPublication string             `bson:"year_of_publication"`
}

//ReadAll output all data with table books
func (B BookData) ReadAll() ([]Book, error) {
	var books []Book
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := B.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var book Book
		err = cursor.Decode(&book)
		if err != nil {
			return nil, err
		}
		logrus.Debug(book)
		books = append(books, book)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, mongo.ErrNoDocuments
	}
	return books, nil
}

//Read read data in db
func (B BookData) Read(id string) (Book, error) {
	var book Book
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	bookId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return book, err
	}
	cur, err := B.collection.Find(ctx, bson.M{"_id": bookId})
	defer cur.Close(ctx)
	cur.Next(ctx)
	if err = cur.Decode(&book); err != nil {
		return book, err
	}
	return book, nil
}

//Add add data in db
func (B BookData) Add(book Book) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := B.collection.InsertOne(ctx, book)
	if err != nil {
		logrus.Fatal(err)
	}
	if result == nil {
		return fmt.Errorf("Error with add data")
	}
	return nil
}

//Update update number of books by the id
func (B BookData) Update(id string, value int) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{"_id", idObj}}
	update := bson.D{
		{"$set", bson.D{
			{"number", value},
		}},
	}
	_, err = B.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error with update data", err)
	}
	return nil
}

func (B BookData) Delete(id string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_ , err = B.collection.DeleteOne(ctx, bson.M{"_id": idObj})
	if err != nil {
		return err
	}
	return nil
}

func (B BookData) ClearDb() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	B.collection.Drop(ctx)
}


//String output data in console
func (B Book) String() string {
	return fmt.Sprintln(B.BookId, B.AuthorId, B.PublisherId, strings.TrimSpace(B.NameOfBook), B.YearOfPublication, B.BookVolume, B.Number)
}

//BookData create a new connection
type BookData struct {
	collection *mongo.Collection
}

//NewBookData it's imitation constructor
func NewBookData(collection *mongo.Collection) *BookData {
	return &BookData{collection: collection}
}
