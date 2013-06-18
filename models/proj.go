package models

import "github.com/astaxie/beego"
import "strings"
import "strconv"
import "time"

const (
	ProjNormal = iota
	ProjQuick
)

type Project struct {
	Name        string
	RD          []string
	QA          []string
	Id          int
	Type        int
	Description string
}

func names() (ret map[string]string) {
	ret = make(map[string]string, 100)
	keys, err := R.Keys("project:*:name")
	if err != nil {
		beego.Error("redis get keys error")
		return
	}
	for _, k := range keys {
		id := strings.SplitN(k, ":", 3)[1]
		ret[id] = k
	}
	return ret
}

func name2id(name string) (id string) {
	for id, k := range names() {
		if na, _ := R.Get(k); string(na) == name {
			return id
		}
	}
	return ""
}

func (pj *Project) Save() (err error) {
	if pj.Name == "" {
		beego.Warn("project name is empty")
		return ErrorInvalid
	}

	id := name2id(pj.Name)
	if id == "" {
		id64, er := R.Incr("project:count")
		if er != nil {
			beego.Error("get new id for project fail")
			return er
		}
		id = strconv.Itoa(int(id64))

		beego.Debug("add new proj name, pid:", pj.Name, id)
	}

	beego.Debug("project id:", id)

	R.Set("project:"+id+":name", []byte(pj.Name))               // Name
	R.Set("project:"+id+":desc", []byte(pj.Description))        // Desc
	R.Set("project:"+id+":type", []byte(strconv.Itoa(pj.Type))) // Type
	rdkey := "project:" + id + ":rd"
	qakey := "project:" + id + ":qa"
	R.Del(rdkey)
	for _, rd := range pj.RD {
		R.Lpush(rdkey, []byte(rd))
	}
	R.Del(qakey)
	for _, qa := range pj.QA {
		R.Lpush(qakey, []byte(qa))
	}
	return
}

func b2str(b []byte, err error) string {
	if err != nil {
		return ""
	}
	return string(b)
}

func rlist(key string, start int, end int) (res []string, err error) {
	res = make([]string, 0, 10)
	ss, er := R.Lrange(key, start, end)
	if er != nil {
		err = er
		return
	}
	for _, b := range ss {
		res = append(res, string(b))
	}
	return
}

func id2proj(id string) Project {
	var pj Project
	pj.Id, _ = strconv.Atoi(id)
	pj.Name = b2str(R.Get("project:" + id + ":name"))
	pj.Description = b2str(R.Get("project:" + id + ":desc"))
	tp, _ := R.Get("project:" + id + ":type")
	pj.Type, _ = strconv.Atoi(string(tp))

	qas, _ := rlist("project:"+id+":qa", 0, -1)
	pj.QA = qas
	rds, _ := rlist("project:"+id+":rd", 0, -1)
	pj.RD = rds
	return pj
}

func ListProject() (projs []Project, err error) {
	projs = make([]Project, 0, 100)
	for id, _ := range names() {
		var pj Project = id2proj(id)
		projs = append(projs, pj)
	}
	return
}

type Job struct {
	Project
	Content []string
}

func GetJob(proj string, start int, end int) (job Job, err error) {
	id := name2id(proj)
	beego.Trace("job id:", id)
	if id == "" {
		err = ErrorNotExisted
		return
	}
	key := "project:" + id + ":jobs"
	ds, err := R.Zrange(key, start, end)
	content := make([]string, 0, 100)
	for _, d := range ds {
		content = append(content, string(d))
	}
	beego.Trace("content:", content)
	return Job{Project: id2proj(id), Content: content}, nil
}
func ListJobs() (jobs []Job, err error) {
	projs, err := ListProject()
	if err != nil {
		return
	}
	jobs = make([]Job, 0, 1000)
	for _, pj := range projs {
		ds, er := R.Zrange("project:"+strconv.Itoa(pj.Id)+":jobs", 0, -1)
		if er != nil {
			err = er
			return
		}
		content := make([]string, 0, 20)
		for _, d := range ds {
			content = append(content, string(d))
		}
		var jb Job = Job{
			Project: pj,
			Content: content,
		}

		jobs = append(jobs, jb)
	}
	return
}

func (jb *Job) Save() (err error) {
	for _, content := range jb.Content {
		_, err = R.Zadd("project:"+strconv.Itoa(jb.Id)+":jobs", []byte(content), float64(time.Now().Unix()))
		if err != nil {
			beego.Warn("add job error:", err)
			return
		}
	}
	return
}
