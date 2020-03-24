package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"mj_http/src/points_mall"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime/pprof"
	"strconv"
	"strings"
)

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func GetTag() {
	tt := TagType{true, "Barak Obama", 1}
	ttType := reflect.TypeOf(tt)
	log.Printf("type : %v\n", ttType)
	for i := 0; i < ttType.NumField(); i++ {
		ixField := ttType.Field(i)
		log.Printf("index[%d] type[%v] name[%v] tag [%v]\n", ixField.Index, ixField.Type, ixField.Name, ixField.Tag)
	}

	//vType := reflect.ValueOf(tt)
	//log.Printf("value: %v \n", vType)
	//for j:=0; j<vType.NumField(); j++ {
	//	ixField := vType.Index(j)
	//	log.Printf("index[%d] type[%v] name[%v] tag [%v]\n", j, ixField.Type(), ixField.String(), ixField.NumField())
	//}
}

func GetReflect() {
	var x float64 = 3.4
	t := reflect.TypeOf(x)
	log.Printf("type[%v] name[%v] string[%v]\n", t.Name(), t.Kind(), t.String())

	v := reflect.ValueOf(x)
	log.Printf("type[%v] name[%v] value[%v]\n", v.Type(), v.Kind(), v)

	v = reflect.ValueOf(&x)
	v = v.Elem()
	if v.CanSet() {
		v.SetFloat(3.1415)
		log.Println("set value ", v.Interface())
	} else {
		log.Println("value can't set")
	}
	//fmt.Println("value:", v)
	//fmt.Println("type:", v.Type())
	//fmt.Println("kind:", v.Kind())
	//fmt.Println("value:", v.Float())
	//fmt.Println(v.Interface())
	//fmt.Printf("value is %5.2e\n", v.Interface())
	//y := v.Interface().(float64)
	//fmt.Println(y)
}

type PathError struct {
	Op   string // “open”, “unlink”, etc.
	Path string // The associated file.
	Err  error  // Returned by the system call.
}

func OpenMyFile(path string) error {
	err := PathError{
		Op:   path,
		Path: path,
		Err:  errors.New("open err")}
	return err.Err
}

func GetPanic() {
	defer func() {
		log.Println("printf log after panic")
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	log.Println("Ending the program")
}

//func GetAllFieldName() {
//	raw1 := points_mall.PayCoinReq1{}
//	t := reflect.TypeOf(raw1)
//
//	for k := 0; k < t.NumField(); k++ {
//		log.Printf("filed name : %s\n", t.Field(k).Name)
//	}
//}

func GetPBType(req *http.Request) {

	req.ParseForm()

	//map->json->pb
	params := make(map[string]interface{})
	for k, v := range req.Form {
		params[k] = v[0]

	}

	log.Printf("params %v \n", params)

	requestJson, err := json.Marshal(params)
	if err != nil {
		log.Println("failed to get josn ", err)
		return
	}

	log.Printf("json %s \n ", string(requestJson))

	//raw := &points_mall.PayCoinReq{}
	//err = json.Unmarshal([]byte(requestJson), raw)
	//if err != nil {
	//	log.Println("failed to parse json ", err)
	//	// fail
	//	return
	//}

	//log.Printf("interface %v", raw)
}

func GetPBInterface(req *http.Request) {

	req.ParseForm()

	params := make(map[string]interface{})
	for k, v := range req.Form {

		params[k] = GetFiledType(k, v[0])

	}
	log.Printf("params map[%v] \n", params)

	requestJson, err := json.Marshal(params)
	if err != nil {
		log.Println("failed to get josn ", err)
		return
	}

	log.Printf("params json[%s] \n", string(requestJson))

	raw := &points_mall.PayCoinReq1{}
	v := reflect.ValueOf(raw)
	ve := v.Elem()
	vt := ve.Type()
	log.Printf("reflect value [%v] [%v] [%v]", v, ve, vt)
	err = json.Unmarshal([]byte(requestJson), raw)
	if err != nil {
		log.Println("failed to parse json ", err)
		// fail
		return
	}

	log.Printf("params interface [%v]\n", raw)
}

func GetFiledType(key string, value string) (newVaule interface{}) {
	newVaule = value
	raw1 := points_mall.PayCoinReq1{}
	t := reflect.TypeOf(raw1)
	for k := 0; k < t.NumField(); k++ {
		// protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty
		tag := string(t.Field(k).Tag)
		if strings.Contains(tag, key) {
			switch t.Field(k).Type.Name() {
			case "int":
				fallthrough
			case "int32":
				newVaule, _ = strconv.Atoi(value)
				return

			}
		}
	}
	return
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func GenProfile() {
	GetPanic()
	flag.Parse()
	if *cpuprofile != "" {
		log.Println("start cpu profile")
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		pprof.StopCPUProfile()
		log.Println("exit main function1")
	}()
}
