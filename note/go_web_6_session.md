## 6.`session`

1. go标准包中没有提供`session`的任何支持	

### 创建

* 定义全局session管理器

  ```go
  // session.go
  package session
  
  type Manager struct {
      cookieName  string     // private cookiename
      lock        sync.Mutex // protects session
      provider    Provider // Provider 为自定义抽象接口，用以表征 session 管理器底层存储结构
      maxLifeTime int64
  }
  
  func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
      provider, ok := provides[provideName]
      if !ok {
          return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
      }
      return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
  }
  
  type Provider interface {
      SessionInit(sid string) (Session, error)
      SessionRead(sid string) (Session, error)
      SessionDestroy(sid string) error
      SessionGC(maxLifeTime int64) // SessionGC 根据 maxLifeTime 来删除过期的数据
  }
  
  // Session 接口 /一般就这四个操作
  type Session interface {
      Set(key, value interface{}) error // set session value
      Get(key interface{}) interface{}  // get session value
      Delete(key interface{}) error     // delete session value
      SessionID() string                // back current sessionID
  }
  ```

* 创建全局唯一标识`sessionid`

  ```go
  // session.go
  package session
  
  func (manager *Manager) sessionId() string {
      b := make([]byte, 32)
      if _, err := rand.Read(b); err != nil {
          return ""
      }
      return base64.URLEncoding.EncodeToString(b)
  }
  ```

* 为每个客户关联一个 session

  ```go
  // session.go
  
  package session
  
  var provides = make(map[string]Provider)// 全局变量
  
  func Register(name string, provider Provider) {
      if provider == nil {
          panic("session: Register provider is nil")
      }
      if _, dup := provides[name]; dup {
          panic("session: Register called twice for provider " + name)
      }
      provides[name] = provider
  }
  
  // 根据cookie来创建一个session或者寻找已有seession
  func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
      manager.lock.Lock()
      defer manager.lock.Unlock()
      cookie, err := r.Cookie(manager.cookieName)
      if err != nil || cookie.Value == "" {
          sid := manager.sessionId()
          session, _ = manager.provider.SessionInit(sid)
          cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
          /* type Cookie struct {
  	// MaxAge=0 means no 'Max-Age' attribute specified.
  	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
  	// MaxAge>0 means Max-Age attribute present and given in seconds
  	MaxAge   int
  	...
  }*/ // from net/http
          http.SetCookie(w, &cookie)
      } else {
          sid, _ := url.QueryUnescape(cookie.Value)// 反向转换
          session, _ = manager.provider.SessionRead(sid)
      }
      return
  }
  ```

  

  ```go
  // example read操作
  package api
  
  func login(w http.ResponseWriter, r *http.Request) {
      sess := globalSessions.SessionStart(w, r)
      r.ParseForm()
      if r.Method == "GET" {
          t, _ := template.ParseFiles("login.gtpl")
          w.Header().Set("Content-Type", "text/html")
          t.Execute(w, sess.Get("username"))
      } else {
          sess.Set("username", r.Form["username"])
          http.Redirect(w, r, "/", 302)
      }
  }
  
  // example session的更多操作
  func count(w http.ResponseWriter, r *http.Request) {
      sess := globalSessions.SessionStart(w, r)
      createtime := sess.Get("createtime")
      if createtime == nil {
          sess.Set("createtime", time.Now().Unix())
      } else if (createtime.(int64) + 360) < (time.Now().Unix()) {
          globalSessions.SessionDestroy(w, r)// 过期 重置操作
          sess = globalSessions.SessionStart(w, r)
      }
      ct := sess.Get("countnum")
      if ct == nil {
          sess.Set("countnum", 1)
      } else {
          sess.Set("countnum", (ct.(int) + 1))
      }
      t, _ := template.ParseFiles("count.gtpl")
      w.Header().Set("Content-Type", "text/html")
      t.Execute(w, sess.Get("countnum"))
  }
  ```

  

  ```go
  // 重置操作的具体实现 session.go
  package session
  
  func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request){
      cookie, err := r.Cookie(manager.cookieName)
      if err != nil || cookie.Value == "" {
          return
      } else {
          manager.lock.Lock()
          defer manager.lock.Unlock()
          manager.provider.SessionDestroy(cookie.Value)
          expiration := time.Now()
          cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
          // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0' "from net/http"
          http.SetCookie(w, &cookie)
      }
  }
  ```

  

  ```go
  // main.go
  package main
  
  var globalSessions *session.Manager
  
  func init() {
      globalSessions, _ = NewManager("memory", "gosessionid", 3600)
  }
  
  func init() {
      go globalSessions.GC()
  }
  ```

  

  ```go
  // 处理过期session
  package session
  
  func (manager *Manager) GC() {
      manager.lock.Lock()
      defer manager.lock.Unlock()
      manager.provider.SessionGC(manager.maxLifeTime)
      time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
      /* AfterFunc 等待持续时间过去，然后在自己的 goroutine 中调用 f。它返回一个计时器，该计时器可用于通过其 Stop 方法取消调用。/type Duration int64/ "from time" */
  }
  ```

  

* session 的存储 (可以存储到内存、文件、数据库等)

  ```go
  // 基于内存的实现 memory.go
  
  package memory
  
  import (
      "container/list"
      "github.com/astaxie/session"
      "sync"
      "time"
  )
  
  var pder = &Provider{list: list.New()}
  
  type SessionStore struct {
      sid          string                      // session id唯一标示
      timeAccessed time.Time                   // 最后访问时间
      value        map[interface{}]interface{} // session里面存储的值
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
      lock     sync.Mutex               // 用来锁
      sessions map[string]*list.Element // 用来存储在内存
      list     *list.List               // 用来做 gc
  }
  
  func (pder *Provider) SessionInit(sid string) (session.Session, error) {
      pder.lock.Lock()
      defer pder.lock.Unlock()
      v := make(map[interface{}]interface{}, 0)
      newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
      element := pder.list.PushFront(newsess)
      pder.sessions[sid] = element
      return newsess, nil
  }
  
  func (pder *Provider) SessionRead(sid string) (session.Session, error) {
      if element, ok := pder.sessions[sid]; ok {
          return element.Value.(*SessionStore), nil
      } else {
          sess, err := pder.SessionInit(sid)
          return sess, err
      }
      return nil, nil
  }
  
  func (pder *Provider) SessionDestroy(sid string) error {
      if element, ok := pder.sessions[sid]; ok {
          delete(pder.sessions, sid)
          pder.list.Remove(element)
          return nil
      }
      return nil
  }
  
  func (pder *Provider) SessionGC(maxlifetime int64) {
      pder.lock.Lock()
      defer pder.lock.Unlock()
  
      for {
          element := pder.list.Back()
          if element == nil {
              break
          }
          if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
              pder.list.Remove(element)
              delete(pder.sessions, element.Value.(*SessionStore).sid)
          } else {
              break
          }
      }
  }
  
  func (pder *Provider) SessionUpdate(sid string) error {
      pder.lock.Lock()
      defer pder.lock.Unlock()
      if element, ok := pder.sessions[sid]; ok {
          element.Value.(*SessionStore).timeAccessed = time.Now()
          pder.list.MoveToFront(element)
          return nil
      }
      return nil
  }
  
  func init() {
      pder.sessions = make(map[string]*list.Element, 0)
      session.Register("memory", pder)
  }
  
  // 使用 main.go
  import (
      "github.com/astaxie/session"
      _ "github.com/astaxie/session/providers/memory"
  )
  ```

* `session`劫持防范

  * `cookieonly` 和 `token`

    * 只允许`cookie`设值，不允许URL重写
    * 在请求中加上token

    ```go
    
    h := md5.New()
    salt:="astaxie%^7&8888"
    io.WriteString(h,salt+time.Now().String())
    token:=fmt.Sprintf("%x",h.Sum(nil))
    if r.Form["token"]!=token{
        // 验证失败
    }
    sess.Set("token",token)
    ```

  * 间隔生成新的 SID

    ```go
    createtime := sess.Get("createtime")
    if createtime == nil {
        sess.Set("createtime", time.Now().Unix())
    } else if (createtime.(int64) + 60) < (time.Now().Unix()) {
        globalSessions.SessionDestroy(w, r)
        sess = globalSessions.SessionStart(w, r)
    }// 每隔六十秒就换一个id
    ```

    