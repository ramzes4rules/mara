package mara

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (mara *Mara) writeToFile(msg string) {

	//
	var paths = strings.Split(filepath.Dir(mara.Path), "\\")

	//
	var f = fmt.Sprintf("%s", paths[0])
	for i := 1; i < len(paths); i++ {
		var dir = paths[i]
		dir = strings.Replace(dir, "{MONTH}", fmt.Sprintf("%s", time.Now().Month()), -1)
		dir = strings.Replace(dir, "{YYYY}", fmt.Sprintf("%d", time.Now().Year()), -1)
		dir = strings.Replace(dir, "{MM}", fmt.Sprintf("%02d", time.Now().Month()), -1)
		dir = strings.Replace(dir, "{DD}", fmt.Sprintf("%02d", time.Now().Day()), -1)
		f = fmt.Sprintf("%s\\%s", f, dir)
		_ = os.Mkdir(f, 0666)
	}

	// Открываем файл на запись
	file, err := os.OpenFile(fmt.Sprintf("%s\\%s", f, filepath.Base(mara.Path)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}

	// Пишем в файл
	_, err = file.WriteString(fmt.Sprintf("%s\n", msg))
	if err != nil {
		return
	}

	// Закрываем файл
	err = file.Close()
	if err != nil {
		return
	}
}
