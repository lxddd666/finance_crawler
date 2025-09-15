// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/cache"
	"hotgo/internal/library/queue"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"hotgo/utility/charset"
	"hotgo/utility/simple"
	"hotgo/utility/validate"
	"runtime"
	"strings"
	"time"

	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"
)

// ProxyInfo 代理信息结构体
type ProxyInfo struct {
	IpAddress string `json:"ip_address"`
	Port      int    `json:"port"`
}

// ParsedProxyList 解析后的代理列表
var ParsedProxyList []string

func Init(ctx context.Context) {
	// 设置gf运行模式
	SetGFMode(ctx)

	// 设置服务日志处理
	glog.SetDefaultHandler(LoggingServeLogHandler)

	// 默认上海时区
	if err := gtime.SetTimeZone("Asia/Shanghai"); err != nil {
		g.Log().Fatalf(ctx, "时区设置异常 err：%+v", err)
		return
	}

	fmt.Printf("欢迎使用HotGo！\r\n当前运行环境：%v, 运行根路径为：%v \r\nHotGo版本：v%v, gf版本：%v \n", runtime.GOOS, gfile.Pwd(), consts.VersionApp, gf.VERSION)

	// 初始化链路追踪
	InitTrace(ctx)

	// 设置缓存适配器
	cache.SetAdapter(ctx)

	// 初始化功能库配置
	service.SysConfig().InitConfig(ctx)

	// 加载超管数据
	service.AdminMember().LoadSuperAdmin(ctx)

	// 订阅集群同步
	SubscribeClusterSync(ctx)

	// 财经init
	_ = FinanceInit(ctx)

	// 获取当前时间未执行code
	_ = GetDayData(ctx)

	// 解析代理文件
	parseJsonFile(ctx)

	// 初始化代理池
	InitProxyPool()
}

func InitProxyPool() {

	ProxyList = NewSafeProxyList(proxyList())
}

func parseJsonFile(ctx context.Context) {
	// 代理文件路径
	proxyFilePath := "internal/global/proxy_file/proxies.json"

	// 检查文件是否存在
	if !gfile.Exists(proxyFilePath) {
		g.Log().Warningf(ctx, "代理文件不存在: %s", proxyFilePath)
		return
	}

	// 读取JSON文件内容
	jsonContent := gfile.GetContents(proxyFilePath)
	if jsonContent == "" {
		g.Log().Warning(ctx, "代理文件内容为空")
		return
	}

	// 解析JSON数据
	var proxyInfos []ProxyInfo
	if err := gjson.DecodeTo(jsonContent, &proxyInfos); err != nil {
		g.Log().Errorf(ctx, "解析代理JSON文件失败: %v", err)
		return
	}

	// 转换为ip:port格式的字符串数组
	ParsedProxyList = make([]string, 0, len(proxyInfos))
	for _, proxy := range proxyInfos {
		proxyStr := fmt.Sprintf("%s:%d", proxy.IpAddress, proxy.Port)
		ParsedProxyList = append(ParsedProxyList, proxyStr)
	}

	g.Log().Infof(ctx, "成功解析代理文件，共加载 %d 个代理", len(ParsedProxyList))
}

func proxyList() []string {
	// 优先使用从JSON文件解析的代理列表
	if len(ParsedProxyList) > 0 {
		g.Log().Infof(gctx.New(), "使用JSON文件中的代理列表，共 %d 个代理", len(ParsedProxyList))
		return ParsedProxyList
	}

	// 如果没有解析到代理，使用硬编码的备用代理列表
	g.Log().Warning(gctx.New(), "未找到有效的代理文件，使用硬编码的备用代理列表")
	// https://fineproxy.org/cn/free-proxy/
	return []string{
		"68.71.251.134:4145",
		"198.177.253.13:4145",
		"199.187.210.54:4145",
		"192.111.130.5:17002",
		"192.111.135.17:18302",
		"192.111.135.18:18301",
		"103.82.25.14:10001",
		"72.214.108.67:4145",
		"162.253.68.97:4145",
		"8.220.200.221:8080",
		"192.252.215.5:16137",
		"184.178.172.13:15311",
		"70.166.167.38:57728",
		"198.8.94.174:39078",
		"192.252.210.233:4145",
		"72.223.188.92:4145",
		"72.211.46.99:4145",

		"67.201.58.190:4145",
		"68.71.242.118:4145",
		"199.102.104.70:4145",
		"98.170.57.241:4145",
		"72.37.216.68:4145",
		"142.54.239.1:4145",
		"67.201.35.145:4145",
		"206.220.175.2:4145",
		"199.58.185.9:4145",
		"184.170.249.65:4145",
		"184.170.245.148:4145",
		"199.58.184.97:4145",
		"98.191.0.37:4145",
		"47.238.226.127:1024",
		"72.211.46.99:4145",
		"46.4.88.72:9050",
		"72.223.188.67:4145",
		"72.207.113.97:4145",
		"68.71.249.153:48606",
		"199.102.106.94:4145",
		"199.102.105.242:4145",
		"199.102.107.145:4145",
		"72.206.74.126:4145",
		"98.182.147.97:4145",
		"107.181.168.145:4145",
		"72.37.217.3:4145",
		"192.252.215.2:4145",
		"68.71.249.158:4145",
		"68.71.252.38:4145",
		"85.111.94.98:15833",
		"127.0.0.1:7890",
		"127.0.0.1:7890",
		"127.0.0.1:7890",
		"192.252.209.158:4145",
		"98.190.239.3:4145",
		"199.116.112.6:4145",
		"184.170.251.30:11288",
		"98.182.171.161:4145",
		"192.111.129.150:4145",
		"192.252.214.17:4145",
		"174.75.211.193:4145",
		"198.177.252.24:4145",
		"142.54.237.38:4145",
		"68.71.245.206:4145",
		"98.175.31.222:4145",
		"68.71.243.14:4145",
		"68.71.240.210:4145",
		"68.71.254.6:4145",
		"72.207.109.5:4145",
		"74.119.144.60:4145",
		"98.181.137.80:4145",
		"68.1.210.163:4145",
		"98.181.137.83:4145",
		"68.71.247.130:4145",
		"74.119.147.209:4145",
		"192.252.216.81:4145",
		"184.178.172.17:4145",
		"72.205.0.67:4145",
		"198.177.254.157:4145",
		"192.252.220.92:17328",
		"184.181.217.201:4145",
		"184.181.217.206:4145",
		"184.181.217.210:4145",
		"184.181.217.213:4145",
		"184.181.217.220:4145",
		"198.177.254.157:4145",
		"192.252.220.92:17328",
		"184.181.217.201:4145",
		"184.181.217.206:4145",
		"184.181.217.210:4145",
		"184.181.217.213:4145",
		"184.181.217.220:4145",
		"184.178.172.11:4145",
		"184.178.172.14:4145",
		"184.178.172.18:15280",
		"184.178.172.23:4145",
		"184.178.172.25:15291",
		"184.178.172.26:4145",
		"184.178.172.28:15294",
		"184.181.217.194:4145",
		"184.178.172.3:4145",
		"184.178.172.5:15303",
		"174.64.199.79:4145",
		"174.77.111.196:4145",
		"198.177.254.131:4145",
		"68.1.210.189:4145",
		"72.205.0.93:4145",
		"192.252.216.86:4145",
		"192.252.211.193:4145",
		"68.71.241.33:4145",
		"67.201.39.14:4145",
		"174.75.211.222:4145",
		"192.252.220.89:4145",
		"107.152.98.5:4145",
		"198.8.84.3:4145",
		"142.54.232.6:4145",
		"142.54.235.9:4145",
		"142.54.229.249:4145",
		"142.54.237.34:4145",
		"142.54.231.38:4145",
		"68.71.251.134:4145",
		"198.177.253.13:4145",
		"199.187.210.54:4145",
		"192.111.135.17:18302",
		"72.214.108.67:4145",
		"103.82.25.14:10001",
		"198.8.94.174:39078",
		"192.252.215.5:16137",
	}
}

// LoggingServeLogHandler 服务日志处理
// 需要将异常日志保存到服务日志时可以通过SetHandlers设置此方法
func LoggingServeLogHandler(ctx context.Context, in *glog.HandlerInput) {
	in.Next(ctx)

	err := g.Try(ctx, func(ctx context.Context) {
		var err error
		defer func() {
			if err != nil {
				panic(err)
			}
		}()

		// web服务日志不做记录，因为会导致重复记录
		r := g.RequestFromCtx(ctx)
		if r != nil && r.Server != nil && in.Logger.GetConfig().Path == r.Server.Logger().GetConfig().Path {
			return
		}

		conf, err := service.SysConfig().GetLoadServeLog(ctx)
		if err != nil {
			return
		}

		if conf == nil {
			return
		}

		if !conf.Switch {
			return
		}

		if in.LevelFormat == "" || !gstr.InArray(conf.LevelFormat, in.LevelFormat) {
			return
		}

		if in.Stack == "" {
			in.Stack = in.Logger.GetStack()
		}

		if len(in.Content) == 0 {
			in.Content = gstr.StrLimit(gvar.New(in.Values).String(), consts.MaxServeLogContentLen)
		}

		var data entity.SysServeLog
		data.TraceId = gctx.CtxId(ctx)
		data.LevelFormat = in.LevelFormat
		data.Content = in.Content
		data.Stack = gjson.New(charset.ParseStack(in.Stack))
		data.Line = strings.TrimRight(in.CallerPath, ":")
		data.TriggerNs = in.Time.UnixNano()
		data.Status = consts.StatusEnabled

		if gstr.Contains(in.Content, `exception recovered`) {
			data.LevelFormat = "PANI"
		}

		if data.Stack.IsNil() {
			data.Stack = gjson.New(consts.NilJsonToString)
		}

		if conf.Queue {
			err = queue.Push(consts.QueueServeLogTopic, data)
		} else {
			err = service.SysServeLog().RealWrite(ctx, data)
		}
	})

	if err != nil {
		g.Dump("LoggingServeLogHandler err:", err)
	}
}

// InitTrace 初始化链路追踪
func InitTrace(ctx context.Context) {
	if !g.Cfg().MustGet(ctx, "jaeger.switch").Bool() {
		return
	}

	tp, err := jaeger.Init(simple.AppName(ctx), g.Cfg().MustGet(ctx, "jaeger.endpoint").String())
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	simple.Event().Register(consts.EventServerClose, func(ctx context.Context, args ...interface{}) {
		_ = tp.Shutdown(ctx)
		g.Log().Debug(ctx, "jaeger closed ..")
	})
}

// SetGFMode 设置gf运行模式
func SetGFMode(ctx context.Context) {
	mode := g.Cfg().MustGet(ctx, "system.mode").String()
	if len(mode) == 0 {
		mode = gmode.NOT_SET
	}

	var modes = []string{gmode.DEVELOP, gmode.TESTING, gmode.STAGING, gmode.PRODUCT}

	// 如果是有效的运行模式，就进行设置
	if validate.InSlice(modes, mode) {
		gmode.Set(mode)
	}
}

func FinanceInit(ctx context.Context) (err error) {
	_, err = service.SysConfig().FinanceConfig(ctx)
	return
}

func GetDayData(ctx context.Context) (err error) {
	today := time.Now().Truncate(24 * time.Hour)

	// 获取当前时间
	Timestamp = today.Unix()

	Day = today.Format("2006-01-02")

	// 获取当前任务
	flag, _ := dao.FinanceCodeDaily.Ctx(ctx).Where(dao.FinanceCodeDaily.Columns().Day, Day).Exist()
	if !flag {
		// 删除过去的
		_, err = dao.FinanceCodeDaily.Ctx(ctx).Where("1=1").Delete()
		// 获取所有
		var codeList []entity.FinanceCode
		var codeDaily []entity.FinanceCodeDaily
		err = dao.FinanceCode.Ctx(ctx).Scan(&codeList)
		if err != nil {
			return
		}
		for _, code := range codeList {
			codeDaily = append(codeDaily, entity.FinanceCodeDaily{
				Code:      code.Code,
				Name:      code.Name,
				Exchange:  code.Exchange,
				Day:       Day,
				Timestamp: Timestamp,
				Status:    consts.TaskNotStarted,
			})
		}
		if len(codeDaily) > 0 {
			batchSize := 500

			for i := 0; i < len(codeDaily); i += batchSize {
				end := i + batchSize
				if end > len(codeDaily) {
					end = len(codeDaily)
				}

				batch := codeDaily[i:end]
				_, err = dao.FinanceCodeDaily.Ctx(ctx).Insert(batch)
				if err != nil {
					// 处理错误
					return
				}
			}
		}
	}
	return
}
