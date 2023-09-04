// Package runner
// @author uangi
// @date 2023/9/4 14:57
package runner

func (r *CockRunner) OneClickStart() {

	r.EnableRedisAndSnowflake()
	r.InitDatasource()
	r.EnableAuthentication()
	r.Run()

}
