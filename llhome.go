package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "567436"
	dbname   = "postgres"
)

type Item struct {
	FID                     string       `json:"fid"`
	FNameEn                 string       `json:"f_name_en"`
	FNameZh                 string       `json:"f_name_zh"`
	FSubID                  string       `json:"f_sub_id"`
	StationZh               string       `json:"station_zh"`
	StationEn               string       `json:"station_en"`
	MainCategoryEn          string       `json:"main_category_en"`
	SubCategoryEn           string       `json:"sub_category_en"`
	MainCategoryZh          string       `json:"main_category_zh"`
	SubCategoryZh           string       `json:"sub_category_zh"`
	GetCategoryZh           string       `json:"get_category_zh"`
	BestGetWayZh            string       `json:"best_get_way_zh"`
	GetCategoryEn           string       `json:"get_category_en"`
	BestGetWayEn            string       `json:"best_get_way_en"`
	PriceGold               string       `json:"price_gold"`
	PriceCrown              string       `json:"price_crown"`
	PriceAP                 string       `json:"price_ap"`
	PriceMasterWrit         string       `json:"price_master_writ"`
	PriceTelStone           string       `json:"price_tel_stone"`
	PriceGem                string       `json:"price_gem"`
	PriceEnd                string       `json:"price_end"`
	RecipeNameEn            string       `json:"recipe_name_en"`
	RecipeMaterialEn        string       `json:"recipe_material_en"`
	RecipeNameZh            string       `json:"recipe_name_zh"`
	RecipeMaterialZh        string       `json:"recipe_material_zh"`
	FurnitureImageComplete  string       `json:"furniture_image_complete"`
	FurnitureInfoComplete   string       `json:"furniture_info_complete"`
	FurnitureImageLocalPath string       `json:"furniture_image_local_path"`
	Tag1                    string       `json:"tag1"`
	Tag2                    string       `json:"tag2"`
	Tag3                    string       `json:"tag3"`
	Pic                     []byte       `json:"pic"`
	DoIHave                 sql.NullBool `json:"do_i_have"`
	DoIWant                 sql.NullBool `json:"do_i_want"`
}

func main() {
	http.HandleFunc("/search", corsMiddleware(searchHandler))
	http.HandleFunc("/mark", corsMiddleware(markHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 设置允许的域名
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// 设置允许的 HTTP 方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 设置允许的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")

		// 处理预检请求
		//if r.Method == "OPTIONS" {
		//	w.WriteHeader(http.StatusOK)
		//	return
		//}

		// 调用下一个处理函数
		next(w, r)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Build the query
	query := buildQuery(r)
	if query == "" {
		http.Error(w, "Invalid query", http.StatusBadRequest)
		return
	}

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Collect results
	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.FID, &item.FNameEn, &item.FNameZh, &item.FSubID, &item.StationZh, &item.StationEn, &item.MainCategoryEn, &item.SubCategoryEn, &item.MainCategoryZh, &item.SubCategoryZh, &item.GetCategoryZh, &item.BestGetWayZh, &item.GetCategoryEn, &item.BestGetWayEn, &item.PriceGold, &item.PriceCrown, &item.PriceAP, &item.PriceMasterWrit, &item.PriceTelStone, &item.PriceGem, &item.PriceEnd, &item.RecipeNameEn, &item.RecipeMaterialEn, &item.RecipeNameZh, &item.RecipeMaterialZh, &item.FurnitureImageComplete, &item.FurnitureInfoComplete, &item.FurnitureImageLocalPath, &item.Tag1, &item.Tag2, &item.Tag3, &item.Pic, &item.DoIHave, &item.DoIWant)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return results as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func markHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("?")
	// 解析请求参数
	fNameZh := r.URL.Query().Get("f_name_zh")
	mark_type := r.URL.Query().Get("mark_type")
	if fNameZh == "" {
		http.Error(w, "参数 f_name_zh 不能为空", http.StatusBadRequest)
		return
	}

	// 连接数据库
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	fmt.Println(mark_type)
	// 更新数据库
	if mark_type == "have" {
		_, err = db.Exec("UPDATE jj SET do_i_have = true WHERE f_name_zh = $1", fNameZh)
		fmt.Println("UPDATE jj SET do_i_have = true WHERE f_name_zh =" + fNameZh)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if mark_type == "want" {
		_, err = db.Exec("UPDATE jj SET do_i_want = true WHERE f_name_zh = $1", fNameZh)
		fmt.Println("UPDATE jj SET do_i_want = true WHERE f_name_zh =" + fNameZh)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("标记成功！"))
}

func buildQuery(r *http.Request) string {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("SELECT * FROM jj WHERE ")

	first := true
	for key, values := range r.URL.Query() {
		if len(values) == 0 || values[0] == "" {
			continue
		}

		if !first {
			queryBuilder.WriteString(" AND ")
		} else {
			first = false
		}

		switch key {
		case "f_name_zh":
			queryBuilder.WriteString(fmt.Sprintf("%s LIKE '%%%s%%'", key, values[0]))
		case "tag1":
			r.ParseForm()
			tag1s := r.Form["tag1"]
			fmt.Println(tag1s)

			if len(tag1s) == 1 {
				queryBuilder.WriteString(fmt.Sprintf("%s LIKE '%%%s%%'", key, values[0]))
			} else {
				queryBuilder.WriteString("(")
				// 打印数组参数
				for i, tag1 := range tag1s {
					queryBuilder.WriteString(fmt.Sprintf("%s LIKE '%%%s%%'", key, tag1))
					if i < len(tag1s)-1 {
						queryBuilder.WriteString(" or ")
					}
				}
				queryBuilder.WriteString(")")
			}
		case "tag2":
			r.ParseForm()
			tag2s := r.Form["tag2"]
			fmt.Println(tag2s)

			if len(tag2s) == 1 {
				queryBuilder.WriteString(fmt.Sprintf("%s LIKE '%%%s%%'", key, values[0]))
			} else {
				queryBuilder.WriteString("(")
				// 打印数组参数
				for i, tag2 := range tag2s {
					queryBuilder.WriteString(fmt.Sprintf("%s LIKE '%%%s%%'", key, tag2))
					if i < len(tag2s)-1 {
						queryBuilder.WriteString(" or ")
					}
				}
				queryBuilder.WriteString(")")
			}
		case "tag3":
			r.ParseForm()
			tag3s := r.Form["tag3"]
			fmt.Println(tag3s)

			if len(tag3s) == 1 {
				queryBuilder.WriteString(fmt.Sprintf("%s LIKE '%%%s%%'", key, values[0]))
			} else {
				queryBuilder.WriteString("(")
				// 打印数组参数
				for i, tag3 := range tag3s {
					queryBuilder.WriteString(fmt.Sprintf("%s LIKE '%%%s%%'", key, tag3))
					if i < len(tag3s)-1 {
						queryBuilder.WriteString(" or ")
					}
				}
				queryBuilder.WriteString(")")
			}
		case "best_get_way_zh":
			r.ParseForm()
			best_get_way_zh := r.Form["best_get_way_zh"]
			fmt.Println(best_get_way_zh)

			if len(best_get_way_zh) == 1 {
				queryBuilder.WriteString(fmt.Sprintf("%s LIKE '%%%s%%'", key, values[0]))
			} else {
				queryBuilder.WriteString("(")
				// 打印数组参数
				for i, way := range best_get_way_zh {
					queryBuilder.WriteString(fmt.Sprintf("%s LIKE '%%%s%%'", key, way))
					if i < len(best_get_way_zh)-1 {
						queryBuilder.WriteString(" or ")
					}
				}
				queryBuilder.WriteString(")")
			}
		default:
			queryBuilder.WriteString(fmt.Sprintf("%s = '%s'", key, values[0]))
		}
	}

	if first {
		// No valid parameters were provided
		return ""
	}

	fmt.Println(queryBuilder.String())
	return queryBuilder.String()
}
