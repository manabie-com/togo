package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type App struct {
	mlock        sync.Mutex
	volumeServer string
	lock         map[string]struct{}
	volume       string
}

func (a *App) UnlockKey(key string) {
	a.mlock.Lock()
	delete(a.lock, key)
	a.mlock.Unlock()
}

func (a *App) LockKey(key string) bool {
	a.mlock.Lock()
	defer a.mlock.Unlock()
	if _, prs := a.lock[key]; prs {
		return false
	}
	a.lock[key] = struct{}{}
	return true
}

//return key and can write or not
func (a *App) CanWrite(r *http.Request) (string, bool) {
	paths := strings.Split(r.URL.Path, "/")
	//1: userId, 2: date, 3:taskname
	filePath := fmt.Sprintf("%s%s", a.volume, r.URL.Path)
	datePath := fmt.Sprintf("%s%s/%s", a.volume, paths[1], paths[2])
	key := paths[1] + paths[2]

	_, err := os.Stat(filePath)
	//Can't write because limit of maximum tasks per date
	if os.IsNotExist(err) && CountFile(datePath) >= int(GetMaxTasks(key)) {
		return "", false
	}

	return filePath, a.LockKey(filePath)
}

func (a *App) ProxyReplica(w http.ResponseWriter, r *http.Request) {
	remote := fmt.Sprintf("%s%s", a.volumeServer, r.URL.Path)
	req, err := http.NewRequest(r.Method, remote, r.Body)
	if err != nil {
		fmt.Println(err)
	}
	req.ContentLength = r.ContentLength
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (a *App) RemotePut(w http.ResponseWriter, r *http.Request) {
	key, isWrite := a.CanWrite(r)
	if key == "" {
		w.WriteHeader(403)
		return
	}
	if !isWrite {
		w.WriteHeader(409)
		return
	}

	a.ProxyReplica(w, r)
	defer a.UnlockKey(key)
}

func (a *App) RemoteDelete(w http.ResponseWriter, r *http.Request) {
	key, isWrite := a.CanWrite(r)
	if !isWrite {
		w.WriteHeader(409)
		return
	}

	a.ProxyReplica(w, r)
	defer a.UnlockKey(key)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		a.RemotePut(w, r)
		return
	} else if r.Method == "DELETE" {
		a.RemoteDelete(w, r)
		return
	}

	a.ProxyReplica(w, r)
}
