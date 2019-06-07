package models

type Test struct {
	ID      int    `json:"-"`
	Name    string `json:name`
	Message string `json:message`
}

func AddTest(name string, mes string) bool {
	if name == "" || mes == "" {
		return false
	}
	db.Create(&Test{
		Name:    name,
		Message: mes,
	})
	return true
}

func QueryTestNow(name string) []Test {
	var ts = []Test{}
	db.Where("name = ?",name).First(&ts)
	//db.Order("id desc").Limit(10).Find(&ts)
	return ts
}

