package utils

/*
* Copyright © 2022 Allan Nava <>
* Created 08/02/2022
* Updated 14/04/2022
*
 */
type ErrorMess struct {
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

type ApiError struct {
	Message  string `json:"message"`
	Response string `json:"response"`
}
type ApiMessage struct {
	Message  string `json:"message"`
	Response string `json:"response"`
}
