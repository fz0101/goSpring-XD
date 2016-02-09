package main

import (
	"fmt"
	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
	"github.com/blang/vfs/mountfs"
	"github.com/r0cketman/goSpring-XD/rest"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	// start this whole mess
	// Make the filesystem read-only:
	var osfs vfs.Filesystem = vfs.OS()
	osfs.Mkdir("tmp", 0777)
	osfs = vfs.ReadOnly(osfs)

	// os.O_CREATE will fail and return vfs.ErrReadOnly
	// os.O_RDWR is supported but Write(..) on the file is disabled
	f, _ := osfs.OpenFile("/tmp/example.txt", os.O_RDWR, 0)

	// error shit
	_, err := f.Write([]byte("Write on readonly fs?"))
	if err != nil {
		fmt.Errorf("Filesystem is read only!\n")
	}

	//create fully writable fs
	mfs := memfs.Create()
	mfs.Mkdir("/root/FMC", 0777)

	// mount memfs inside of the /memfs
	// memfs may not exists, add check
	fs := mountfs.Create(osfs)
	fs.Mount(mfs, "/memfs")

	// create directory inside of the mem file store
	fs.Mkdir("/memfs/testdir", 0777)

	// creat test directory
	fs.Mkdir("tmp/testdir", 0777)

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/springxdsink", springxdsink),
		rest.Get("/signalfiles", getfiles),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func getfiles(w rest.ResponseWriter, r *rest.Request) {
	app := "echo"
	arg0 := "Hello World"
	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
		print(string(stdout))
	}
	w.WriteJson("Hello World")
}

func springxdsink(w rest.ResponseWriter, r *rest.Request) {

	w.WriteJson("Add Sink Stuff Here")

}
