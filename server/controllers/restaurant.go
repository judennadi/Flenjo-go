package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/judennadi/flenjo-go/config"
)

var ctx = context.Background()
var redisClient = config.GetRedisCache()
var redisExp = 60 * 60 * time.Second

// func enableCors(w *http.ResponseWriter) {
// fmt.Println("here")
// (*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	termArr := strings.Split(r.URL.Query().Get("term"), " ")
	term := strings.Join(termArr, "+")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	rating := r.URL.Query().Get("rating")
	token := os.Getenv("YELP_API_KEY")
	header := fmt.Sprintf("Bearer %v", token)
	var resp map[string]interface{}
	cacheKey := fmt.Sprintf("restaurants?term=%v&page=%v", term, page)
	cached, err := redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		url := fmt.Sprintf("https://api.yelp.com/v3/businesses/search?location=USA&term=%v&categories=food,restaurants&limit=30&offset=%v", term, page*30)

		client := &http.Client{
			Timeout: time.Second * 30,
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("Authorization", header)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = json.NewDecoder(res.Body).Decode(&resp)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		respJson, _ := json.Marshal(resp)
		err = redisClient.Set(ctx, cacheKey, respJson, redisExp).Err()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err = json.Unmarshal([]byte(cached), &resp)
		if err != nil {
			fmt.Println(err)
		}
	}

	var data map[string]interface{}
	total := resp["total"].(float64)
	businesses := resp["businesses"].([]interface{})
	if rating != "" {
		ratingNum, _ := strconv.ParseFloat(rating, 64)

		var newBusinesses []interface{}
		for _, v := range businesses {
			val := v.(map[string]interface{})
			bRating := val["rating"].(float64)
			if bRating >= ratingNum {
				newBusinesses = append(newBusinesses, v)
			}
		}

		if len(newBusinesses) > 1000 {
			data = map[string]interface{}{"data": newBusinesses, "total": 1000}
		} else {
			data = map[string]interface{}{"data": newBusinesses, "total": len(newBusinesses)}
		}
	} else {
		if total > 1000 {
			data = map[string]interface{}{"data": businesses, "total": 1000}
		} else {
			data = map[string]interface{}{"data": businesses, "total": total}
		}
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}

func GetBars(w http.ResponseWriter, r *http.Request) {
	termArr := strings.Split(r.URL.Query().Get("term"), " ")
	term := strings.Join(termArr, "+")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	rating := r.URL.Query().Get("rating")
	token := os.Getenv("YELP_API_KEY")
	header := fmt.Sprintf("Bearer %v", token)
	var resp map[string]interface{}

	cacheKey := fmt.Sprintf("bars?term=%v&page=%v", term, page)
	cached, err := redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		url := fmt.Sprintf("https://api.yelp.com/v3/businesses/search?location=USA&term=%v&categories=beergardens,bars&limit=30&offset=%v", term, page*30)

		client := &http.Client{
			Timeout: time.Second * 30,
		}
		req, err1 := http.NewRequest("GET", url, nil)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		req.Header.Add("Authorization", header)
		res, err2 := client.Do(req)
		if err2 != nil {
			fmt.Println(err2)
			return
		}

		err4 := json.NewDecoder(res.Body).Decode(&resp)
		if err4 != nil {
			fmt.Println(err4)
			return
		}
		defer res.Body.Close()

		respJson, _ := json.Marshal(resp)
		err = redisClient.Set(ctx, cacheKey, respJson, redisExp).Err()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err = json.Unmarshal([]byte(cached), &resp)
		if err != nil {
			fmt.Println(err)
		}
	}

	var data map[string]interface{}
	total := resp["total"].(float64)
	businesses := resp["businesses"].([]interface{})
	if rating != "" {
		ratingNum, _ := strconv.ParseFloat(rating, 64)

		var newBusinesses []interface{}
		for _, v := range businesses {
			val := v.(map[string]interface{})
			bRating := val["rating"].(float64)
			if bRating >= ratingNum {
				newBusinesses = append(newBusinesses, v)
			}
		}

		if len(newBusinesses) > 1000 {
			data = map[string]interface{}{"data": newBusinesses, "total": 1000}
		} else {
			data = map[string]interface{}{"data": newBusinesses, "total": len(newBusinesses)}
		}
	} else {
		if total > 1000 {
			data = map[string]interface{}{"data": businesses, "total": 1000}
		} else {
			data = map[string]interface{}{"data": businesses, "total": total}
		}
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}

func GetHotels(w http.ResponseWriter, r *http.Request) {
	termArr := strings.Split(r.URL.Query().Get("term"), " ")
	term := strings.Join(termArr, "+")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	rating := r.URL.Query().Get("rating")
	token := os.Getenv("YELP_API_KEY")
	header := fmt.Sprintf("Bearer %v", token)
	var resp map[string]interface{}

	cacheKey := fmt.Sprintf("hotels?term=%v&page=%v", term, page)
	cached, err := redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		url := fmt.Sprintf("https://api.yelp.com/v3/businesses/search?location=USA&term=%v&categories=hotels&limit=30&offset=%v", term, page*30)

		client := &http.Client{
			Timeout: time.Second * 30,
		}
		req, err1 := http.NewRequest("GET", url, nil)
		if err1 != nil {
			fmt.Println(err1)
			return
		}
		req.Header.Add("Authorization", header)
		res, err2 := client.Do(req)
		if err2 != nil {
			fmt.Println(err2)
			return
		}

		err4 := json.NewDecoder(res.Body).Decode(&resp)
		if err4 != nil {
			fmt.Println(err4)
			return
		}
		defer res.Body.Close()

		respJson, _ := json.Marshal(resp)
		err = redisClient.Set(ctx, cacheKey, respJson, redisExp).Err()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err = json.Unmarshal([]byte(cached), &resp)
		if err != nil {
			fmt.Println(err)
		}
	}

	var data map[string]interface{}
	total := resp["total"].(float64)
	businesses := resp["businesses"].([]interface{})
	if rating != "" {
		ratingNum, _ := strconv.ParseFloat(rating, 64)

		var newBusinesses []interface{}
		for _, v := range businesses {
			val := v.(map[string]interface{})
			bRating := val["rating"].(float64)
			if bRating >= ratingNum {
				newBusinesses = append(newBusinesses, v)
			}
		}

		if len(newBusinesses) > 1000 {
			data = map[string]interface{}{"data": newBusinesses, "total": 1000}
		} else {
			data = map[string]interface{}{"data": newBusinesses, "total": len(newBusinesses)}
		}
	} else {
		if total > 1000 {
			data = map[string]interface{}{"data": businesses, "total": 1000}
		} else {
			data = map[string]interface{}{"data": businesses, "total": total}
		}
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}

func SearchAutocomplete(w http.ResponseWriter, r *http.Request) {
	token := os.Getenv("YELP_API_KEY")
	header := fmt.Sprintf("Bearer %v", token)
	url := fmt.Sprintf("https://api.yelp.com/v3/autocomplete?text=%v", r.URL.Query().Get("text"))

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	req.Header.Add("Authorization", header)
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	var resp map[string]interface{}

	err4 := json.NewDecoder(res.Body).Decode(&resp)
	if err4 != nil {
		fmt.Println(err4)
		return
	}

	defer res.Body.Close()

	categories := resp["categories"].([]interface{})
	terms := resp["terms"].([]interface{})
	var texts []map[string]interface{}

	for _, v := range categories {
		t := v.(map[string]interface{})
		texts = append(texts, t)
	}

	for _, v := range terms {
		t := v.(map[string]interface{})
		texts = append(texts, t)
	}

	data := map[string]interface{}{"data": resp, "terms": texts}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}

func GetBusiness(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	token := os.Getenv("YELP_API_KEY")
	header := fmt.Sprintf("Bearer %v", token)
	url1 := fmt.Sprintf("https://api.yelp.com/v3/businesses/%v", id)
	url2 := fmt.Sprintf("https://api.yelp.com/v3/businesses/%v/reviews", id)
	var resp1 map[string]interface{}
	var resp2 map[string]interface{}

	cacheBusKey := fmt.Sprintf("business?id=%v", id)
	cacheRevKey := fmt.Sprintf("reviews?id=%v", id)
	cachedRev, err := redisClient.Get(ctx, cacheRevKey).Result()
	if err != nil {
		client := &http.Client{
			Timeout: time.Second * 30,
		}
		req1, err := http.NewRequest("GET", url1, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req1.Header.Add("Authorization", header)
		res1, err := client.Do(req1)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = json.NewDecoder(res1.Body).Decode(&resp1)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res1.Body.Close()
		respJson1, _ := json.Marshal(resp1)
		err = redisClient.Set(ctx, cacheBusKey, respJson1, 30*60*time.Second).Err()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err = json.Unmarshal([]byte(cachedRev), &resp2)
		if err != nil {
			fmt.Println(err)
		}
	}

	cachedBus, err := redisClient.Get(ctx, cacheBusKey).Result()
	if err != nil {
		client := &http.Client{
			Timeout: time.Second * 30,
		}
		req2, err := http.NewRequest("GET", url2, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		req2.Header.Add("Authorization", header)
		res2, err := client.Do(req2)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = json.NewDecoder(res2.Body).Decode(&resp2)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res2.Body.Close()
		respJson2, _ := json.Marshal(resp2)
		err = redisClient.Set(ctx, cacheRevKey, respJson2, 30*60*time.Second).Err()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err = json.Unmarshal([]byte(cachedBus), &resp1)
		if err != nil {
			fmt.Println(err)
		}

	}
	data := map[string]interface{}{"restaurant": resp1, "reviews": resp2["reviews"]}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}
