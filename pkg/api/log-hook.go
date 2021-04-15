// Adapted from https://levelup.gitconnected.com/4-tips-for-logging-on-gcp-using-golang-and-logrus-239baf3b1ac2

package api

import (
	"cloud.google.com/go/errorreporting"
	"cloud.google.com/go/logging"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"regexp"
	"runtime"
)

type StackDriverHook struct {
	client      *logging.Client
	errorClient *errorreporting.Client
	logger      *logging.Logger
}

var logLevelMappings = map[logrus.Level]logging.Severity{
	logrus.TraceLevel: logging.Default,
	logrus.DebugLevel: logging.Debug,
	logrus.InfoLevel:  logging.Info,
	logrus.WarnLevel:  logging.Warning,
	logrus.ErrorLevel: logging.Error,
	logrus.FatalLevel: logging.Critical,
	logrus.PanicLevel: logging.Critical,
}

func NewStackDriverHook(logName, projectID string) (*StackDriverHook, error) {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	errorClient, err := errorreporting.NewClient(ctx, projectID, errorreporting.Config{
		ServiceName: projectID,
		OnError: func(err error) {
			log.Printf("Could not log error: %v", err)
		},
	})
	if err != nil {
		return nil, err
	}
	return &StackDriverHook{
		client:      client,
		errorClient: errorClient,
		logger:      client.Logger(logName),
	}, nil
}
func (sh *StackDriverHook) Close() {
	sh.client.Close()
	sh.errorClient.Close()
}
func (sh *StackDriverHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (sh *StackDriverHook) Fire(entry *logrus.Entry) error {
	payload := map[string]interface{}{
		"Message": entry.Message,
		"Data":    entry.Data,
	}
	level := logLevelMappings[entry.Level]
	sh.logger.Log(logging.Entry{Payload: payload, Severity: level})
	if int(level) >= int(logging.Error) {
		err := getError(entry)
		if err == nil {
			errData, e := json.Marshal(payload)
			if e != nil {
				fmt.Printf("Error %v", e)
			}
			fmt.Print(string(errData))
			err = fmt.Errorf(string(errData))
		}
		fmt.Println(err.Error())
		sh.errorClient.Report(errorreporting.Entry{
			Error: err,
			Stack: sh.getStackTrace(),
		})
	}
	return nil
}
func (sh *StackDriverHook) getStackTrace() []byte {
	stackSlice := make([]byte, 2048)
	length := runtime.Stack(stackSlice, false)
	stack := string(stackSlice[0:length])
	re := regexp.MustCompile("[\r\n].*logrus.*")
	res := re.ReplaceAllString(stack, "")
	return []byte(res)
}

type stackDriverError struct {
	Err         interface{}
	Code        interface{}
	Description interface{}
	Message     interface{}
	Env         interface{}
}

func (e stackDriverError) Error() string {
	return fmt.Sprintf("%v - %v - %v - %v - %v", e.Code, e.Description, e.Message, e.Err, e.Env)
}
func getError(entry *logrus.Entry) error {
	errData := entry.Data["error"]
	env := entry.Data["env"]
	code := entry.Data["ErrCode"]
	desc := entry.Data["ErrDescription"]
	msg := entry.Message
	err := stackDriverError{
		Err:         errData,
		Code:        code,
		Message:     msg,
		Description: desc,
		Env:         env,
	}
	return err
}
func (sh *StackDriverHook) Wait() {}
