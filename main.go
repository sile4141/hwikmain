package main

import (

	// elastic "go.elastic.co/apm"

	app "github.com/hwikpass/hwik-go"
	"github.com/hwikpass/hwik-sample/model"
	router "github.com/hwikpass/hwik-sample/router"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 환경별 properties 파일 테스트용
	os.Setenv("ENV", "local")
	os.Setenv("LOG_LEVEL", "DEBUG")
	// ./.vscode/launch.json 파일에 환경 설정으로 변경해서 사용.

	// Initialize hwik-go library
	// app.Init(args1, args2)
	// - args1 : 토큰 인증 및 사용자 정보 가져오기 여부(true/false)
	// - args2 : biz logger 사용여부(true/false)
	e, appLogger := app.Init(false, true)

	defer app.Close(appLogger)
	// e, f1, err := app.Init(logrus.DebugLevel, false, true)

	// 라우터 등록
	router.SetPolicyRouter(e)

	model.InitConfig() //환경변수 일부(몽고디비) init

	// 서버 시작
	app.Start(e)
}
