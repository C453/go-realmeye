package qanda

import (
	"github.com/trapped/realmeye/base"
	"net/http"
)

func Serve(w http.ResponseWriter, req *http.Request) {
	b := base.Page{
		Title:       "Q & A",
		Location:    "/q-and-a",
		Description: "Questions and Answers of RealmEye.com",
		Keywords:    "Q and A",
		Specific:    nil,
	}

	tem := b.Template("q-and-a/index.gom")

	err := tem.Execute(w, b)
	if err != nil {
		panic(err)
	}
}
