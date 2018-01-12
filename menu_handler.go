package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BASE_URL string = "https://backend-challenge-summer-2018.herokuapp.com/challenges.json?id=1&page="

func getApiResponseByPage(pageNum int, allMenus *map[int64]Menu) *MerchantApi {
	url := fmt.Sprintf("%s%v", BASE_URL, pageNum)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		merchantApi, _ := getMerchantInformation([]byte(content))
		addMenusToMap(merchantApi.Menu, allMenus)
		return merchantApi
	}
	return nil
}

func getMerchantInformation(content []byte) (*MerchantApi, error) {
	apiResp := MerchantApi{}
	err := json.Unmarshal(content, &apiResp)
	if err != nil {
		log.Fatal(err)
	}
	return &apiResp, err
}

func addMenusToMap(apiMenus []Menu, allMenus *map[int64]Menu) {
	for _, menu := range apiMenus {
		(*allMenus)[menu.ID] = menu
	}
}

func findCyclicAndNonCyclicMenus(allMenus *map[int64]Menu) ([]MenuResult, []MenuResult) {
	valid := make([]MenuResult, 0)
	invalid := make([]MenuResult, 0)
	for _, menu := range *allMenus {
		if menu.ParentID != 0 {
			continue
		}
		tempMenu := MenuResult{}
		tempMenu.RootID = menu.ID

		tempChildren := make([]int64, 0)
		tempMenu.Children = *(findMenuChildren(tempMenu.RootID, tempMenu.RootID, &tempChildren, allMenus))
		if isMenuCyclic(&tempMenu) {
			valid = append(valid, tempMenu)
		} else {
			invalid = append(invalid, tempMenu)
		}
	}
	return valid, invalid
}

func findMenuChildren(rootID int64, menuID int64, children *[]int64, allMenus *map[int64]Menu) *[]int64 {
	if len((*allMenus)[menuID].ChildIDs) == 0 {
		return children
	}
	for _, child := range (*allMenus)[menuID].ChildIDs {
		(*children) = append(*children, child)
		if child == rootID {
			break
		}
		findMenuChildren(rootID, child, children, allMenus)
	}
	return children
}

func isMenuCyclic(menu *MenuResult) bool {
	for _, childID := range (*menu).Children {
		if childID == (*menu).RootID {
			return false
		}
	}
	return true
}
