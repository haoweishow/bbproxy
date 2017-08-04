package log

import (
	seelog "github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func init() {
	logConfig := `
		<seelog minlevel="debug" maxlevel="error">
			<outputs formatid="fmt1">
				<rollingfile type="size" filename="./bbproxy.log" maxsize="10240000" maxrolls="500" />
			</outputs>
			<formats>
				<format id="fmt1" format="%Date %Time [%LEV] %Msg%n"/>
			</formats>
		</seelog>`
	Logger, _ = seelog.LoggerFromConfigAsBytes([]byte(logConfig))

	seelog.ReplaceLogger(Logger)
}
