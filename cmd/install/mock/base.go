package mock

import (
	"io/ioutil"
	"net/http/httptest"
	"strings"

	"github.com/douyu/juno/api/apiv1/confgo"
	pprofHandle "github.com/douyu/juno/api/apiv1/pprof"
	"github.com/douyu/juno/api/apiv1/resource"
	"github.com/douyu/juno/api/apiv1/system"
	"github.com/douyu/juno/api/apiv1/user"
	"github.com/labstack/echo/v4"
)

func MockData() {

	router := echo.New()
	urlAppPut := "/app/put"
	urlAppNodePut := "/app_node/put"
	urlConfigTplCreate := "/config/tpl/create"
	urlConfigFileCreate := "/config/file/create"
	urlConfigContentUpdate := "/config/content/update"
	urlConfigGeneratorParse := "/config/parse"
	urlUserCreate := "/user/create"
	urlSystemUpdate := "/system/setting/update"

	router.POST(urlAppPut, resource.AppPut)
	router.POST(urlAppNodePut, resource.AppNodePut)
	router.POST(urlConfigTplCreate, confgo.TplCreate)
	router.POST(urlConfigFileCreate, confgo.CreateConfigFile)
	router.POST(urlConfigContentUpdate, confgo.ItemCreate)
	router.POST(urlConfigGeneratorParse, confgo.ParseGenerator)
	router.POST(urlUserCreate, user.Create)
	router.POST(urlSystemUpdate, system.SettingUpdate)

	mockApp(urlAppPut, router)
	mockAppNode(urlAppNodePut, router)
	mockTplCreate(urlConfigTplCreate, router)
	mockConfig(urlConfigFileCreate, urlConfigContentUpdate, router)
	mockParse(urlConfigGeneratorParse, router)
	mockCreateAdmin(urlUserCreate, router)
	mockGrafanaSetting(urlSystemUpdate, router)
}

func PostForm(uri string, param string, router *echo.Echo) []byte {
	// 构造post请求
	req := httptest.NewRequest("POST", uri, strings.NewReader(param))
	req.Header.Set("Content-Type", "application/json")

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}

func Get(uri string, router *echo.Echo) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	return body
}
