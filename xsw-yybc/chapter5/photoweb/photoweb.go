package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"text/template"

	"github.com/pkg/errors"
)

const (
	Upload_Dir   = "chapter5/photoweb/uploads"
	Template_Dir = "chapter5/photoweb/views"
	List_Dir     = 0x0001
)

var (
	templates = make(map[string]*template.Template)
)

func init() {
	//for _, tmpl := range []string{"upload", "list"} {
	//	t := template.Must(template.ParseFiles(tmpl + ".html"))
	//	templates[tmpl] = t
	//}

	fileInfoArr, err := ioutil.ReadDir(getWd() + "/" + Template_Dir)
	if err != nil {
		err = errors.Wrapf(err, "init readDir err")
		panic(err.Error())
		return
	}

	var (
		templateName, templatePath string
	)

	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		//fmt.Println("templateName = ", templateName)
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}

		templatePath = getWd() + "/" + Template_Dir + "/" + templateName
		//fmt.Println("Loading template: ", templatePath)

		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
	}

	//fmt.Printf("%+v", templates)
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//		io.WriteString(w, `
		//<h1></h1>
		//<form method="POST" action="/upload" enctype="multipart/form-data">
		//    Choose an image to upload: <input type="file" name="image">
		//    <input type="submit" value="Upload">
		//</form>
		//`)

		// 模板
		//t, err := template.ParseFiles(getWd() + "/chapter5/photoweb/uploads/upload.html")
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//t.Execute(w, nil)
		//return

		// render
		if err := renderHtml(w, getWd()+"/"+Template_Dir+"/upload.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filename := h.Filename
		defer f.Close()
		t, err := os.Create(getWd() + "/" + Upload_Dir + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}

}

func listHandle(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(getWd() + "/" + Upload_Dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//var listHtml string
	//for _, fileInfo := range fileInfoArr {
	//	imgid := fileInfo.Name()
	//	if imgid == ".keep" {
	//		continue
	//	}
	//	listHtml += fmt.Sprintf("<li><a href=\"/view?id=%s\">%s</a></li>", imgid, imgid)
	//}
	//
	//io.WriteString(w, "<h1></h1><ol>"+listHtml+"</ol>")

	// 模板
	locals := make(map[string]interface{})
	images := make([]string, 0)
	for _, fileInfo := range fileInfoArr {
		if fileInfo.Name() == ".keep" {
			continue
		}
		images = append(images, fileInfo.Name())
	}

	locals["images"] = images
	//t, err := template.ParseFiles(getWd() + "/chapter5/photoweb/uploads/list.html")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//t.Execute(w, locals)

	// render
	tmpl := getWd() + "/" + Template_Dir + "/list.html"
	if err = renderHtml(w, tmpl, locals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	//t, err := template.ParseFiles(tmpl)
	//if err != nil {
	//	err = errors.Wrapf(err, "renderHtml parseFiles err")
	//	return
	//}
	//return t.Execute(w, locals)

	return templates[tmpl].Execute(w, locals)
}

func viewHandle(w http.ResponseWriter, r *http.Request) {
	imageID := r.FormValue("id")
	imagePath := getWd() + "/" + Upload_Dir + "/" + imageID
	if exists := isExists(imagePath); exists == false {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func staticDirHandler(mux *http.ServeMux, prefix, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		//fmt.Println("dddd ", r.URL.Path, file)
		if (flags & List_Dir) == 0 {
			if exists := isExists(file); exists == false {
				http.NotFound(w, r)
				return
			}
		}

		http.ServeFile(w, r, file)
	})
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				// 或者输出自定义的 50x 错误页面

				log.Println(string(debug.Stack()))
			}
		}()

		fn(w, r)
	}
}

func getWd() string {
	wd, _ := os.Getwd()
	return wd
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/",
		getWd()+"/"+"chapter5/photoweb/public", 0)
	mux.HandleFunc("/", safeHandler(listHandle))
	mux.HandleFunc("/view", safeHandler(viewHandle))
	mux.HandleFunc("/upload", safeHandler(uploadHandle))
	err := http.ListenAndServe(":8083", mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
