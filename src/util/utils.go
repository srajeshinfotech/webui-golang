package util

import (
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func isDarwinPlatform() bool {
	return runtime.GOOS == "darwin"
}

func InitLogger(loglevel int) {

	formatter := &log.TextFormatter{
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	}

	log.SetFormatter(formatter)

	if !isDarwinPlatform() {
		logpath := "/var/log/edpd/local_ven_onboarding_ui.log"
		file, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			/* */
		} else {
			log.SetOutput(file)
		}
	}

	log.SetReportCaller(true)
	SetLogLevel(loglevel)
	log.Debugf("%s PID=%d, PPID=%d UID=%d", os.Args[0], os.Getpid(), os.Getppid(), os.Getuid())
}

func SetLogLevel(loglevel int) {
	if loglevel >= int(log.PanicLevel) && loglevel <= int(log.TraceLevel) {
		log.SetLevel(log.AllLevels[loglevel])
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func GetEpochTime() int64 {
	return time.Now().Unix()
}

// Exists reports whether the named file or directory exists.
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Remove(fn string) error {
	if _, err := os.Stat(fn); err == nil {
		return os.Remove(fn)
	}
	return nil
}
