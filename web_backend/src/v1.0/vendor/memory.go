package vendor

import (
	"container/list"

	"sync"
	"time"
)

type SessionStore struct {
	sid          string                      //session id 唯一标示
	timeAccessed time.Time                   //最后访问时间
	value        map[interface{}]interface{} //session 里面存储的值
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	return nil
}
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}
func (st *SessionStore) SessionID() string {
	return st.sid
}

type Provider struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做 gc
}

func (provider *Provider) SessionInit(sid string) (Session, error) {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := provider.list.PushBack(newsess)
	provider.sessions[sid] = element
	return newsess, nil
}

func (provider *Provider) SessionRead(sid string) (Session, error) {
	if element, ok := provider.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := provider.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (provider *Provider) SessionDestroy(sid string) error {
	if element, ok := provider.sessions[sid]; ok {
		delete(provider.sessions, sid)
		provider.list.Remove(element)
		return nil
	}
	return nil
}

func (provider *Provider) SessionGC(maxLifeTime int64) {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	for {
		element := provider.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) <
			time.Now().Unix() {
			provider.list.Remove(element)
			delete(provider.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}
func (provider *Provider) SessionUpdate(sid string) error {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	if element, ok := provider.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		provider.list.MoveToFront(element)
		return nil
	}
	return nil
}

var pder = &Provider{list: list.New()}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	Register("memory", *pder)

}
