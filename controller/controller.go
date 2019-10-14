package controller

import (
	"log"
	"context"
	"net/http"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/dymzfp/tes-netmonk/model"
	f "github.com/dymzfp/tes-netmonk/function"
	cdb "github.com/dymzfp/tes-netmonk/db"
)

var newSnmp []model.Snmp

const (
	errorConnectDD       = "cannot connect to the database: "
	errorDecodingJSONReq = "decoding JSON Error"
	emptyData            = "data is empty"
)

func GetSnmp(w http.ResponseWriter, r *http.Request) {
	db, err := cdb.Connect()
    if err != nil {
        log.Println(errorConnectDD, err.Error())
		return
	}

	resp := model.NewResponseFormat()

	// snmp
	var snmp []model.Snmp
	
	collection := db.Collection("snmp")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {

		var result model.Snmp
		
		cursor.Decode(&result)
		snmp = append(snmp, result)
	}
	if err := cursor.Err(); err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}

	// snmptrap
	var snmptrap []model.Snmp

	collection = db.Collection("snmptrap")

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err = collection.Find(ctx, bson.M{})

	if err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {

		var result model.Snmp
		
		cursor.Decode(&result)
		snmptrap = append(snmptrap, result)
	}
	if err = cursor.Err(); err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}

	newSnmp := f.Join(snmp, snmptrap)

	if len(newSnmp) <= 0 {
		resp.AddError(emptyData, emptyData)
		sendResponse(http.StatusNotFound, resp, w, r)
		return
	}
	resp.SetData(newSnmp)
	sendResponse(http.StatusOK, resp, w, r)
	return
}


func Available(w http.ResponseWriter, r *http.Request) {
	db, err := cdb.Connect()
    if err != nil {
        log.Println(errorConnectDD, err.Error())
		return
	}

	resp := model.NewResponseFormat()

	// snmp
	var snmp []model.Snmp
	
	collection := db.Collection("snmp")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {

		var result model.Snmp
		
		cursor.Decode(&result)
		snmp = append(snmp, result)
	}
	if err := cursor.Err(); err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}

	// snmptrap
	var snmptrap []model.Snmp

	collection = db.Collection("snmptrap")

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err = collection.Find(ctx, bson.M{})

	if err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {

		var result model.Snmp
		
		cursor.Decode(&result)
		snmptrap = append(snmptrap, result)
	}
	if err = cursor.Err(); err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}

	newSnmp := f.Join(snmp, snmptrap)

	if len(newSnmp) <= 0 {
		resp.AddError(emptyData, emptyData)
		sendResponse(http.StatusNotFound, resp, w, r)
		return
	}

	available := f.Avege(newSnmp)

	resp.SetData(available)
	sendResponse(http.StatusOK, resp, w, r)
	return 
}

func Mttr(w http.ResponseWriter, r *http.Request) {
	db, err := cdb.Connect()
    if err != nil {
        log.Println(errorConnectDD, err.Error())
		return
	}

	resp := model.NewResponseFormat()

	// snmp
	var snmp []model.Snmp
	
	collection := db.Collection("snmp")

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {

		var result model.Snmp
		
		cursor.Decode(&result)
		snmp = append(snmp, result)
	}
	if err := cursor.Err(); err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}

	// snmptrap
	var snmptrap []model.Snmp

	collection = db.Collection("snmptrap")

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err = collection.Find(ctx, bson.M{})

	if err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {

		var result model.Snmp
		
		cursor.Decode(&result)
		snmptrap = append(snmptrap, result)
	}
	if err = cursor.Err(); err != nil {
		resp.AddError(err.Error(), err.Error())
		sendResponse(http.StatusInternalServerError, resp, w, r)
		return
	}

	newSnmp := f.Join(snmp, snmptrap)

	if len(newSnmp) <= 0 {
		resp.AddError(emptyData, emptyData)
		sendResponse(http.StatusNotFound, resp, w, r)
		return
	}

	mttr := f.DownTime(newSnmp)

	resp.SetData(mttr)
	sendResponse(http.StatusOK, resp, w, r)
	return 
}

// response
func sendResponse(statusCode int, resp *model.ResponseFormat, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	
	encodedResponse, err := resp.EncodeToJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		log.Printf("Source: %v| Destination: %v| ResponseCode: %v| ResponseLen: %v", r.RemoteAddr, r.RequestURI, statusCode, "error while encoding response")
		return fmt.Errorf("unable to encode JSON: %v", err)
	}
	w.Write(encodedResponse)
	if user := r.Header.Get("user"); user != "" {
		log.Printf("| User: %v | Source: %v | Destination: %v | Mehod: %v | ResponseCode: %v | ResponseLen: %v", user, r.RemoteAddr, r.RequestURI, r.Method, statusCode, len(encodedResponse))
	} else {
		log.Printf("| Source: %v | Destination: %v | Mehod: %v | ResponseCode: %v | ResponseLen: %v", r.RemoteAddr, r.RequestURI, r.Method, statusCode, len(encodedResponse))
	}
	return nil
}