/***********************************************************************************************
 ***                                          GINNY                                          ***
 ***********************************************************************************************/

package configs

// 数据库配置
func GetMysqlConfig() map[string]string {
	// 初始化数据库配置map
	MySQLConfig := make(map[string]string)

	MySQLConfig["DB_HOST"] = "localhost"
	MySQLConfig["DB_PORT"] = "3306"
	MySQLConfig["DB_NAME"] = "root"
	MySQLConfig["DB_USER"] = "root"
	MySQLConfig["DB_PWD"] = ""
	MySQLConfig["DB_CHARSET"] = "utf8mb4"

	MySQLConfig["DB_MAX_OPEN_CONNS"] = "20"       // 连接池最大连接数
	MySQLConfig["DB_MAX_IDLE_CONNS"] = "10"       // 连接池最大空闲数
	MySQLConfig["DB_MAX_LIFETIME_CONNS"] = "7200" // 连接池链接最长生命周期

	return MySQLConfig
}