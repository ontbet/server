package models

import (
	"fmt"
	db "ontbet/databases"
)

type GuessResult struct {
	Id         int    `json:"id" form:"id"`
	GameID     int    `json:"gameid" form:"gameid"`
	Address    string `json:"address" form:"address"`
	UserNumber int    `json:"usernumber" form:"usernumber"`
	SysNumber  int    `json:"sysnumber" form:"sysnumber"`
	TokenType  int    `json:"tokentype" form:"tokentype"`
	Amount     int64  `json:"amount" form:"amount"`
}

func (p *GuessResult) GeTop20Result() (guessResults []GuessResult, err error) {
	guessResults = make([]GuessResult, 0)
	rows, err := db.SqlDB.Query("SELECT id, address, amount,gameid,tokentype,usernumber,sysnumber FROM guess order by gameid desc  limit 0,20")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var guessResult GuessResult
		rows.Scan(&guessResult.Id, &guessResult.Address, &guessResult.Amount, &guessResult.GameID, &guessResult.TokenType, &guessResult.UserNumber, &guessResult.SysNumber)
		guessResults = append(guessResults, guessResult)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (p *GuessResult) GetResultForAddress(address string) (guessResults []GuessResult, err error) {
	fmt.Println(address)
	guessResults = make([]GuessResult, 0)
	rows, err := db.SqlDB.Query("SELECT id, address, amount,gameid,tokentype,usernumber,sysnumber FROM guess  where address = ? order by gameid desc  limit 0,20", address)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var guessResult GuessResult
		rows.Scan(&guessResult.Id, &guessResult.Address, &guessResult.Amount, &guessResult.GameID, &guessResult.TokenType, &guessResult.UserNumber, &guessResult.SysNumber)
		guessResults = append(guessResults, guessResult)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
