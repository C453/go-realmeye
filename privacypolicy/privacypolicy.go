package privacypolicy

import (
	"github.com/trapped/realmeye/base"
	"net/http"
)

func Serve(w http.ResponseWriter, req *http.Request) {
	b := base.Page{
		Title:       "Privacy Policy",
		Location:    "/privacy-policy",
		Description: "Privacy Policy of RealmEye.com",
		Keywords:    "Privacy Policy",
		Specific:    nil,
	}

	tem := b.Template("privacypolicy/index.gom")

	err := tem.Execute(w, b)
	if err != nil {
		panic(err)
	}
}
