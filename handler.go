package main

import ( // import 一些library
	"encoding/json"
	"fmt"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// w是writer, r是request, 后者是pointer可以对Request本身操作。
	// 1. r是pointer, 主要是为了让它的修改能被记录（被其他组件可见？）但实际上不需要修改
	// 2. w是要修改的，但是不是pointer，它是interface不是concrete struct，没法pointer。
	// Parse from body of request to get a json object.
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.Body) // 制造json.NewDecoder(/*body*/)
	var p Post
	// Decode - JSON变object（和Marshal是一对）
	if err := decoder.Decode(&p); err != nil { // 传入Post的指针，好让decoder去修改p本身
		panic(err) // panic是一个简单而激进的错误处理，终止程序、抛出异常。以后可以修改
	}

	fmt.Fprintf(w, "Post received: %s\n", p.Message) // 写到Wrtier里面，用Fprintf
}


func searchHandler(w http.ResponseWriter, r *http.Request) {
	// 给前端返回的数据是JSON格式，要记得set一下header。
	fmt.Println("Received one request for search")
	w.Header().Set("Content-Type", "application/json")
	// 1. 从request获得url里面参数。（Query代表?以后的部分）
	user := r.URL.Query().Get("user")
	keywords := r.URL.Query().Get("keywords")

	var posts []Post
	var err error
	// 2. 根据用户提供的是user还是keywords来判断按照哪个来搜索。
	if user != "" {
		posts, err = searchPostsByUser(user)
	} else {
		posts, err = searchPostsByKeywords(keywords)
	}
	// 3. 把搜索结果返回前端就行
	// 3.1 如果有err，不要panic（panic会停止程序太激进），应该给前端返回一个HttpError比如403、500之类
	if err != nil {
		http.Error(w, "Failed to read post from Elasticsearch", http.StatusInternalServerError)
		fmt.Printf("Failed to read post from Elasticsearch %v.\n", err) // %v是
		return
	}
	// 3.2 Marshal - Object变JSON（和Decode是一对）
	js, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, "Failed to parse posts into JSON format", http.StatusInternalServerError)
		fmt.Printf("Failed to parse posts into JSON format %v.\n", err)
		return
	}
	w.Write(js)
}
