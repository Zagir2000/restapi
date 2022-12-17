package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

/*
Логрус поддерживает хуки
Сейчас хотим чтобы наше логирование лилось в два места: в файл и в output. Файл может понадобится для эластика,
а output, чтобы можно было видеть в докере
*/
type writerHook struct {
	writer    []io.Writer
	LogLevels []logrus.Level
}

// метод Fire будеьт вызываться каждый раз когда мы будем что то писать куда то для каждого уровня
func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.writer {
		w.Write([]byte(line))
	}
	return err
}

//Levels будет возращать левелы из нашего hook
func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// В какой то момент мне
var e *logrus.Entry

//Создаем структуру, потому что логер может поменяться вдруг в какой то момент жизни
type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{Entry: e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) Logger {
	return Logger{Entry: l.WithField(k, v)}
}
func init() {
	// Создаем новый экземпляр логгера
	logs := logrus.New()
	logs.SetReportCaller(true)
	//Формат отображает одну запись в журнале
	logs.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
	}
	logs.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	err := os.MkdirAll("/root/restapi/logs", 0644)
	if err != nil {
		panic(err)
	}

	// os.O_WRONLY сообщает компьютеру, что вы собираетесь только записывать в файл, а не читать
	// os.O_CREATE говорит компьютеру создать файл, если он не существует
	// os.O_APPEND сообщает компьютеру о добавлении в конец файла вместо его перезаписи или усечения
	allFile, err := os.OpenFile("/root/restapi/logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}
	//Логрус никуда ничего не пишет, потому что нам нужно писать в два места сразу.Чтобы по умолчанию логи никуда не уходили
	logs.SetOutput(io.Discard)

	logs.AddHook(&writerHook{
		[]io.Writer{allFile, os.Stdout},
		logrus.AllLevels,
	})
	logs.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(logs)
}

// kafka --info, debug
// file --error, trace
// stdout -- warning, critical
//Суть структуру writeHook- это иметь возможность распределить на каждого writer несколько уровней логирования, в разные файлы разные уровни логирования
//info.log -- info, warning, error, critical
//debug.log  -- info, warning, error, critical, debug
//trace.log -- info, warning, error, critical, debug, trace
//error.log --error, critical
// writer - это не только файл, но и elastic, logstash, kafka и т.д.ы
