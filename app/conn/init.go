package conn

func InitConn() {
	//连接redis
	initRedis()

	//连接数据库
	initMysql()
}
