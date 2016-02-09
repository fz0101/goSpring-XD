package main

//	Fernando Zavala
//	2/9/2016
//	Built for Pivotal Software
//	Post json from spring xd processors to restful api

import (
	"github.com/blang/vfs"
	"github.com/blang/vfs/memfs"
	"github.com/blang/vfs/mountfs"
	"github.com/r0cketman/goSpring-XD/rest"
	"log"
	"net/http"
)

///////////////////////////////////////////////////
////	create a virtual file system wrapper
////	create an in memory file system
////	this is where we store the physical file system
////	this is outside of the main function for variable scoping
///////////////////////////////////////////////////

// create the vfs object that access the underlying file system.
// Make the filesystem read-only:
var osfs vfs.Filesystem = vfs.OS()

func main() {

	///////////////////////////////////////////////////
	////	Here is where we create our API
	////	We build End Points
	////	Set routes
	////	Create functions that serve end points
	// 		ToDo add end point to retrieve files, to do add hypermedia for discoverable
	//		Ned to fix references to API framework, so forking doesn't cause any sillyness
	//		Fernando Zavala 2/9/2016
	///////////////////////////////////////////////////
	osfs.Mkdir("root", 0777)

	// now, let's create the in memory fs object
	mfs := memfs.Create()
	mfs.Mkdir("/root/", 0777)

	// create a vfs that supports mounts
	// add conditional check for memfs -- todo
	fs := mountfs.Create(osfs)
	fs.Mount(mfs, "/memfs")

	// create directory inside of the mem file store
	fs.Mkdir("/memfs/root", 0777)

	// added stats end point to find response time
	api := rest.NewApi()
	MWstats := &rest.StatusMiddleware{}
	api.Use(MWstats)
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("API/.status", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(MWstats.GetStatus())
		}),
		// Root is API, Integration is self explanatory, gives api more structure.
		rest.Post("API/Integration/springxdsink", springxdsink),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

///////////////////////////////////////////////////
////		functions							/////
////											/////
////											/////
////											/////
///////////////////////////////////////////////////

func springxdsink(w rest.ResponseWriter, r *rest.Request) {

	// to do... add logic to make sure that files stored properly,
	// we need to generate a unique file scheme.

	file, err := vfs.Create(osfs, "root/signal.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	if _, err := file.Write([]byte("VFS working on your filesystem")); err != nil {
		log.Fatal(err)
	}

}
