
# XSDL-FAST-SQL-FRAME
💕💕💕💕💕💕💕💕💕

### How to Use ?

#### 1.Regist one Database Connection-pool(You also can regist lots of Connect-pool,we added mysql,sqlserver and redis driver)
```` golang
gosql.DBRegister("mysql", "localhost", "test", "root", "root", 3306, "srv1")
gosql.DBRegister("mysql", "localhost", "test", "user", "root", 3309, "srv2")
gosql.DBRegister("mssql", "localhost", "plan", "sa", "W991224z", 1433, "srv3")
gosql.DBRegister("redis", "", "", "", "", 6379, "redis")
````
#### 2.When you want to Query or Exec something......
```` golang
con,err := gosql.GetDBCon("svr1") //choose the database which you want to query
con2,err := gosql.GetDBCon("svr2") //choose the database which you want to query
c := gosql.GetRedisCon("redis")

res1 := gosql.Query(con,"SELECT * FROM user WHERE id = ?",10086)
res2 := gosql.Exec(con2,"Insert into plan_list(a,b,c) values(?,?,?)","ABC",1,"2022-01-01")
value, err := redis.String(c.Do("GET", "gokey"))
````
#### 3.Do not forget CLOSE THE DATABASE CONNECT!
```` golang
con.Close()
con2.Close()
````

#### 4.Use Transation

```` golang
DB := gosql.GetDB("svr1")
tx, _ := gosql.OpenTransaction(DB)
gosql.TExec(tx, "INSERT INTO count(name,year,count) VALUES (?,?,?)", "AB1", 1, 3)
gosql.TExec(tx, "INSERT INTO count(name,year,count) VALUES (?,?,?)", "AB2", 1, 3)
gosql.TExec(tx, "INSERT INTO count(name,year,count) VALUES (?,?,?)", "ABR", 1, 3)
gosql.SubmitTransaction(tx)
````
# Enjoy hacking!