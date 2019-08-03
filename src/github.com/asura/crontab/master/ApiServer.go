package master

import (
	"encoding/json"
	"github.com/asura/crontab/common"
	"net"
	"net/http"
	"strconv"
	"time"
)

// 任务接口
type ApiServer struct {
	httpServer *http.Server
}

// 保存接口任务
// POST job = {"name":"job1","command":"echo hello","cronExpr":"* * * * *"}
func handleJobSave(resp http.ResponseWriter, req *http.Request) {
	var (
		err      error
		postForm string
		job      common.Job
		oldJob   *common.Job
		bytes    []byte
	)

	// 解析 POST表单
	if err = req.ParseForm(); err != nil {
		goto ERR
	}
	// 取表单中的job字段
	postForm = req.PostForm.Get("job")
	// 反序列化
	if err = json.Unmarshal([]byte(postForm), &job); err != nil {
		goto ERR
	}

	// 保存到etcd
	if oldJob, err = G_JobMgr.SaveJob(&job); err != nil {
		goto ERR
	}

	// 返回正常应答
	if bytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		resp.Write(bytes)
	}
	return

ERR:
	// 返回异常应答
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}
}

// 删除任务
func handleJobDelete(resp http.ResponseWriter, req *http.Request) {
	var (
		err    error
		oldJob *common.Job
		bytes  []byte
		name   string
	)
	if err = req.ParseForm(); err != nil {
		goto ERR
	}
	name = req.PostForm.Get("name")

	if oldJob, err = G_JobMgr.DeleteJob(name); err != nil {
		goto ERR
	}

	// 正常返回
	if bytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		resp.Write(bytes)
	}

	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}
}
func handleJobKill(resp http.ResponseWriter, req *http.Request) {
	var (
		err   error
		name  string
		bytes []byte
	)

	if err = req.ParseForm(); err != nil {
		goto ERR
	}
	name = req.PostForm.Get("name")
	if err = G_JobMgr.JobKill(name); err != nil {
		goto ERR
	}
	// 正常返回
	if bytes, err = common.BuildResponse(0, "success", nil); err == nil {
		resp.Write(bytes)
	}

	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}
}

func handleJobList(resp http.ResponseWriter, req *http.Request) {
	var (
		err     error
		bytes   []byte
		jobList []*common.Job
	)
	if jobList, err = G_JobMgr.JobList(); err != nil {
		goto ERR
	}

	// 正常返回
	if bytes, err = common.BuildResponse(0, "success", jobList); err == nil {
		resp.Write(bytes)
	}

	return

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}
}

var (
	// 单例对象
	G_apiServer *ApiServer
)

// 初始化服务
func InitApiServer() (err error) {

	var (
		mux          *http.ServeMux
		listen       net.Listener
		httpServer   *http.Server
		staticDir    http.Dir     // 静态文件根目录
		staticHandle http.Handler // 静态文件的HTTP回调
	)

	// 配置路由
	mux = http.NewServeMux()

	mux.HandleFunc("/job/save", handleJobSave)
	mux.HandleFunc("/job/delete", handleJobDelete)
	mux.HandleFunc("/job/list", handleJobList)
	mux.HandleFunc("/job/kill", handleJobKill)

	// 静态文件目录
	staticDir = http.Dir(G_Config.WebRoot)
	staticHandle = http.FileServer(staticDir)
	mux.Handle("/", http.StripPrefix("/", staticHandle))

	// 启动 TCP 监听
	if listen, err = net.Listen("tcp", ":"+strconv.Itoa(G_Config.ApiPort)); err != nil {
		return
	}
	// 创建一个http服务
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(G_Config.ApiReadTimeOut) * time.Millisecond,
		WriteTimeout: time.Duration(G_Config.ApiWriteTimeOut) * time.Millisecond,
		Handler:      mux,
	}

	// 赋值单例
	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}
	// 启动了服务
	go httpServer.Serve(listen)

	return
}
