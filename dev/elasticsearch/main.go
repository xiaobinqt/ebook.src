package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

func main() {
	// 初始化 MySQL 连接
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/wb-test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 初始化 Elasticsearch 连接
	esConfig := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	esClient, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 插入文章到 MySQL
	articleTitle := "悯农"
	articleContent := "锄禾日当午，汗滴禾下土。谁知盘中餐，粒粒皆辛苦。"
	_, err = db.Exec("INSERT INTO article (title, content) VALUES (?, ?)", articleTitle, articleContent)
	if err != nil {
		log.Fatal(err)
	}

	// 获取插入文章的 MySQL ID
	var articleID int
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&articleID)
	if err != nil {
		log.Fatal(err)
	}

	// 在 Elasticsearch 中写入文章
	esIndexName := "articles" // Elasticsearch 索引名
	docID := fmt.Sprintf("%d", articleID)
	//esDoc := map[string]interface{}{
	//	"title":   articleTitle,
	//	"content": articleContent,
	//	"mysql_id": articleID, // 将 MySQL ID 存储在 Elasticsearch 中
	//}

	// 创建 Elasticsearch 文档
	esRequest := esapi.IndexRequest{
		Index:      esIndexName,
		DocumentID: docID,
		Body: strings.NewReader(fmt.Sprintf(`{
            "title": "%s",
            "content": "%s",
            "mysql_id": %d
        }`, articleTitle, articleContent, articleID)),
		Refresh: "true",
	}

	res, err := esRequest.Do(context.Background(), esClient)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// 从 Elasticsearch 中查询文章，包括 MySQL ID
	//esQuery := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"match": map[string]interface{}{
	//			"title": "示例文章标题",
	//		},
	//	},
	//}

	esSearchRequest := esapi.SearchRequest{
		Index: []string{esIndexName},
		Body: strings.NewReader(fmt.Sprintf(`{
            "query": {
                "match": {
                    "title": "示例文章标题"
                }
            }
        }`)),
	}

	res, err = esSearchRequest.Do(context.Background(), esClient)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// 处理 Elasticsearch 查询结果
	if res.IsError() {
		log.Fatalf("Error: %s", res.Status())
	}

	// 解析查询结果
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		title := source["title"].(string)
		content := source["content"].(string)
		mysqlID := source["mysql_id"].(float64) // 从 Elasticsearch 中获取 MySQL ID
		fmt.Printf("标题: %s内容: %sMySQL ID: %d", title, content, int(mysqlID))
	}
}
