package tool

import (
	"reflect"
	"runtime"
	"sync"
	"time"
)

//ThreadData which is use for multiple runs
type ThreadData struct {
	wg sync.WaitGroup
	sy sync.Mutex

	//if Line <= 0 equal runtime.NumCPU()
	Line int
	//0 means ignore err ,until all is success (default);
	//1 means all runs one times,if err is happened next data also run
	//2 means all runs one times,if err is happened next data not run
	Mode int
	//if Mode == 0 ; RunTime is Mode running max time,if RunTime ==0  unlimited
	RunTime int
	//runs is record run times
	runs int
	Data interface{}
	//Back which will return data errors or used for a temp msg
	Back interface{}
	//Func is a func which used for
	Func func(interface{}) error
	//Interval means error happened sleep some duration
	Interval time.Duration
	//Err is used for return Err msg
	Err []error
}

func (thread *ThreadData) Run() {
	//log.Println("runs：", thread.runs)
	if thread.Line <= 0 {
		thread.Line = runtime.NumCPU()
	}
	l := thread.Line
	i := thread.Data
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		panic("data is not pointer")
	}
	if reflect.TypeOf(i).Elem().Kind() != reflect.Slice {
		panic("data is not a slice")
	}
	v := reflect.ValueOf(i)
	e := v.Elem()
	p := e.Len() / l
	n := e.Len() % l
	if n > 0 {
		p = p + 1
	}
	t := reflect.New(e.Type())
	t1 := t.Elem()
	for k := 0; k < p; k++ {
		s := k * l
		d := (k + 1) * l
		if d > e.Len() {
			d = e.Len()
		}
		thread.wg.Add(1)
		go func() {
			defer thread.wg.Done()
			for k1 := s; k1 < d; k1++ {
				err := thread.Func(e.Index(k1).Interface())
				if err != nil {
					thread.sy.Lock()
					thread.Err = append(thread.Err, err)
					if thread.Mode == 2 {
						thread.sy.Unlock()
						return
					}
					//add into temp
					t1.Set(reflect.Append(t1, e.Index(k1)))
					thread.sy.Unlock()
					time.Sleep(thread.Interval)
				}
			}
		}()
	}
	thread.wg.Wait()
	if t1.Len() == 0 {
		return
	}
	thread.Back = t.Interface()
	if thread.Mode == 1 || thread.Mode == 2 {
		return
	}
	thread.runs++
	if thread.Mode == 0 {
		if thread.RunTime > 0 && thread.RunTime <= thread.runs {
			return
		}
	}
	//time.Sleep(thread.Interval)
	//try again
	thread.Data = t.Interface()
	thread.Back = nil
	thread.Run()
}
