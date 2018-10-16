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