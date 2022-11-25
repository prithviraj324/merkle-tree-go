package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type transaction struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
//block is a bunch of transaction structs, its size is predetermined
type block struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}