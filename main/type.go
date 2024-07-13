package main

type FoundIndexBody struct {
	FoundIndex int `json:"foundIndex"`
}

type ErrorBody struct {
	ErrorMessage string `json:"errorMessage"`
}
