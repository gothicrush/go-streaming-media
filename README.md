r.Cookie("name")
http.SetCookie(w,&http.Cookie{})

t, _ := template.ParseFiles("a.html","b.html","c.html")
t.ExecuteTemplate(w,"b", para)

m.Store(k,v)
m.Delete(k)
m.Range( func(k interface{}, v interface{}) bool  )
m.Load(k)
m.LoadOrStore(k, v)

string->int：strconv.Atoi(str)
string->int64：strconv.ParseInt(str, 10, 64)
int->string：strconv.Itoa(i)
int64->string：strconv.FormatInt(int64, 10)

### API 模块

#### User API

| 功能     | URL             | 请求方法 | 状态码              |
| -------- | --------------- | -------- | ------------------- |
| 注册用户 | /user           | POST     | 201-400-500         |
| 登录用户 | /user/:username | POST     | 200-400             |
| 用户信息 | /user/:username | GET      | 200-400-401-403-500 |
| 注销用户 | /user/:username | DELETE   | 204-400-401-403-500 |

#### Resource API

| 功能         | URL                         | 请求方法 | 状体码              |
| ------------ | --------------------------- | -------- | ------------------- |
| 列出所有视频 | /user/:username/videos      | GET      | 200-400-500         |
| 获取一个视频 | /user/:username/videos/:vid | GET      | 200-400-500         |
| 删除一个视频 | /user/:username/videos/:vid | DELETE   | 204-400-401-403-500 |

#### Comment API

| 功能                 | URL                              | 请求方法 | 状态码              |
| -------------------- | -------------------------------- | -------- | ------------------- |
| 获取一个视频所有评论 | /videos/:vid/comments            | GET      | 200-400-500         |
| 对一个视频发表评论   | /videos/:vid/comments            | POST     | 201-400-500         |
| 删除一个评论         | /videos/:vid/comment/:comment-id | DELETE   | 204-400-401-403-500 |



### Database 模块

* go get go-sql-driver/mysql

* ```
  db, err := sql.Open("mysql", "user:password@/dbname")
  
  if err != nil {
      panic(err.Error())
  }
  
  defer db.Close()
  
  err = db.Ping()
  if err != nil {
      panic(err.Error())
  }
  ```

*

CREATE DATABASE IF NOT EXISTS stream_media
DEFAULT CHARSET utf8
COLLATE utf8_general_ci;

USE stream_media;

CREATE TABLE IF NOT EXISTS user (
    userid varchar(50) UNIQUE NOT NULL,
    username varchar(20) UNIQUE NOT NULL,
    password varchar(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS session (
    username varchar(20) UNIQUE NOT NULL,
    sessionid varchar(50) UNIQUE NOT NULL,
    ttl varchar(50) NOT NULL
);