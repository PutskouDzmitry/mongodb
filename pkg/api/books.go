package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"

	//"strconv"

	"github.com/PutskouDzmitry/be-sd/pkg/data"

	"github.com/gorilla/mux"
)

type bookAPI struct {
	data *data.BookData
}

func ServeUserResource(r *mux.Router, data data.BookData) {
	api := &bookAPI{data: &data}
	r.HandleFunc("/books", api.getAllBooks).Methods("GET")
	r.HandleFunc("/book{id}", api.getOneBook).Methods("GET")
	r.HandleFunc("/books", api.createBook).Methods("POST")
	r.HandleFunc("/books{id}/{number}", api.updateBook).Methods("PUT")
	r.HandleFunc("/books{id}", api.deleteBook).Methods("DELETE")
}

func (a bookAPI) getAllBooks(writer http.ResponseWriter, request *http.Request) {
	users, err := a.data.ReadAll()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get users "))
		if err != nil {
			log.Println(err)
		}
	}
	logrus.Info(users)
	err = json.NewEncoder(writer).Encode(users)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a bookAPI) getOneBook(writer http.ResponseWriter, request *http.Request) {
	idRequest := mux.Vars(request)
	id := idRequest["id"]
	user, err := a.data.Read(id)
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get users"))
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	logrus.Info(user)
	if user.NameOfBook != "" {
		err = json.NewEncoder(writer).Encode(user)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (a bookAPI) createBook(writer http.ResponseWriter, request *http.Request) {
	book := new(data.Book)
	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if book == nil {
		log.Printf("failed empty JSON")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.Add(*book)
	logrus.Info(book.BookId)
	if err != nil {
		log.Println("user hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a bookAPI) updateBook(writer http.ResponseWriter, request *http.Request) {
	idRequest := mux.Vars(request)
	id := idRequest["id"]
	strNumber := idRequest["number"]
	number, err := strconv.Atoi(strNumber)
	if err != nil {
		log.Println("book hasn't been updated, because number doesn't equal int:", number)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.Update(id, number)
	logrus.Info(id)
	if err != nil {
		log.Println("book hasn't been updated")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

func (a bookAPI) deleteBook(writer http.ResponseWriter, request *http.Request) {
	idRequest := mux.Vars(request)
	id := idRequest["id"]
	err := a.data.Delete(id)
	logrus.Println(id)
	if err != nil {
		log.Println("book hasn't been deleted(")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//bookId, err := primitive.ObjectIDFromHex(id)
//if err != nil {
//return book, err
//}
//cur, err := B.collection.Find(ctx, bson.M{"_id":bookId})
//var bookBson  []bson.M
//if err = cur.All(ctx, &bookBson); err != nil {
//return book, err
//}
//if err != nil {
//return book, err
//}
