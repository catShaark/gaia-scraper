package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func get_gaia3_votes() {
	plan, _ := ioutil.ReadFile("data3.json")
	var data []string
	err := json.Unmarshal(plan, &data)
	out, _ := os.Create("/home/pegasus/gaia-scraper/gaia3_votes.csv")
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range data {

		txBytes, _ := base64.StdEncoding.DecodeString(s)
		vote := []byte{150, 1, 40, 40, 22, 169, 10, 30, 161, 202, 221, 54, 8}
		if bytes.Equal(txBytes[:13], vote) {
			vote_option := txBytes[13]
			proposal_id := txBytes[37]
			voter_addr_bz := sdk.AccAddress(txBytes[16:36])
			voter_addr := voter_addr_bz.String()
			println(voter_addr + "," + strconv.Itoa(int(proposal_id)) + "," + strconv.Itoa(int(vote_option)))
			out.WriteString(voter_addr + "," + strconv.Itoa(int(proposal_id)) + "," + strconv.Itoa(int(vote_option)) + "\n")
		}
	}
	out.Close()
}

func get_gaia2_votes() {
	plan, _ := ioutil.ReadFile("data2.json")
	var data []string
	err := json.Unmarshal(plan, &data)
	out, _ := os.Create("/home/pegasus/gaia-scraper/gaia2_votes.csv")
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range data {

		txBytes, _ := base64.StdEncoding.DecodeString(s)
		vote := []byte{150, 1, 240, 98, 93, 238, 10, 30, 161, 202, 221, 54, 8}
		if bytes.Equal(txBytes[:13], vote) {
			vote_option := txBytes[13]
			proposal_id := txBytes[37]
			voter_addr_bz := sdk.AccAddress(txBytes[16:36])
			voter_addr := voter_addr_bz.String()
			println(voter_addr + "," + strconv.Itoa(int(proposal_id)) + "," + strconv.Itoa(int(vote_option)))
			out.WriteString(voter_addr + "," + strconv.Itoa(int(proposal_id)) + "," + strconv.Itoa(int(vote_option)) + "\n")
		}
	}
	out.Close()
}

func get_gaia1_votes() {
	plan, _ := ioutil.ReadFile("data1.json")
	var data []string
	err := json.Unmarshal(plan, &data)
	out, _ := os.Create("/home/pegasus/gaia-scraper/gaia1_votes.csv")
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range data {

		txBytes, _ := base64.StdEncoding.DecodeString(s)
		vote := []byte{150, 1, 240, 98, 93, 238, 10, 30, 161, 202, 221, 54, 8}
		if bytes.Equal(txBytes[:13], vote) {
			vote_option := txBytes[13]
			proposal_id := txBytes[37]
			voter_addr_bz := sdk.AccAddress(txBytes[16:36])
			voter_addr := voter_addr_bz.String()
			println(voter_addr + "," + strconv.Itoa(int(proposal_id)) + "," + strconv.Itoa(int(vote_option)))
			out.WriteString(voter_addr + "," + strconv.Itoa(int(proposal_id)) + "," + strconv.Itoa(int(vote_option)) + "\n")
		}
	}
	out.Close()
}

func main() {
	get_gaia1_votes()
	get_gaia2_votes()
	get_gaia3_votes()
}
