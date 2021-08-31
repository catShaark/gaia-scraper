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

func get_gaia4_votes() {
	plan, _ := ioutil.ReadFile("data4.json")
	var data []string
	err := json.Unmarshal(plan, &data)
	out, _ := os.OpenFile("/home/pegasus/gaia-scraper/gaia4_votes.csv", os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range data {
		txBytes, _ := base64.StdEncoding.DecodeString(s)
		vote := []byte{10, 84, 10, 82, 10, 27, 47, 99, 111, 115, 109, 111, 115, 46, 103, 111, 118, 46, 118, 49, 98, 101, 116, 97, 49, 46, 77, 115, 103, 86, 111, 116, 101, 18, 51, 8}
		len_vote := len(vote)
		is_vote := bytes.Equal(txBytes[:len_vote], vote)
		if is_vote {
			vote_option := txBytes[85]
			proposal_id := txBytes[36]
			voter_addr_bz := txBytes[39:84]
			voter_addr := string(voter_addr_bz[:])
			println(voter_addr + "," + strconv.Itoa(int(proposal_id)) + "," + strconv.Itoa(int(vote_option)))
			// _, err = out.WriteString(voter_addr + "," + strconv.Itoa(int(proposal_id)) + "," + strconv.Itoa(int(vote_option)) + "\n")
			// if err != nil {
			// 	fmt.Println(err)
			// }
		}
		if err != nil {
			print(err)
		}

	}
	out.Close()
}

func main() {
	get_gaia4_votes()
}
