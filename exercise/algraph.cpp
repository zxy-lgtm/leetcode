#include<iostream>

using namespace std;

#define MAX_VEX_NUM 20

typedef struct Arcnode{
    int adjnum;
    struct Arcnode *nextarc;
}Anode;

typedef struct Vnoda{
    int data;
    Anode * firstarc;
}Vnode,Adjlist[MAX_VEX_NUM];

typedef struct{
    int vexnum,arcnum;
    Adjlist vertices;
}ADgraph;

int Location(ADgraph &g,int v){
    for(int i = 0; i < g.vexnum; i ++){
        if(g.vertices[i].data == v){
            return i;
        }
    }
    return -1;
}

int indegree[MAX_VEX_NUM];

void create_graph(ADgraph &g){
    cout<<"请输入图的结点总数：";
    cin>>g.vexnum;
    cout<<"请输入每一个结点的值：";
    for(int i = 0; i < g.vexnum; i ++){
        cin>>g.vertices[i].data;
        g.vertices[i].firstarc = NULL;
    }

    cout<<"请输入边的总数：";
    cin>>g.arcnum;
    cout<<"请输入每条边：";
    for(int i = 0; i< g.arcnum;i++){
        int v1,v2;
        cin>>v1>>v2;
        Anode *p = new Anode;
        int p1 = Location(g,v1); int p2 = Location(g,v2);
        indegree[p2]++;
        p->adjnum = p2;
        p->nextarc = g.vertices[p1].firstarc;
        g.vertices[p1].firstarc = p;
    }
}

bool visited[MAX_VEX_NUM];

void dfs(ADgraph &g,int v){
    visited[v] = true;
    cout<<g.vertices[v].data<<"->";
    for(Arcnode *p = g.vertices[v].firstarc;p;p=p->nextarc){
        int w = p->adjnum;
        if(!visited[w]){
            dfs(g,w);
        }
    }
}

void DFS(ADgraph &g){
    for(int i = 0; i< g.vexnum; i++){
        if(!visited[i]){
            dfs(g,i);
        }
    }
}

typedef struct node{
    int data;
    struct node *front,*next;
}qnode;

typedef struct{
    int len;
    qnode *head,*end;
}queue;

void init_queue(queue &q){
    q.len = 0;
    q.head = NULL;
    q.end = NULL;
}

void en_queue(queue &q,int v){
    qnode *p = new qnode;
    p->data = v;
    if(q.len == 0){
        q.end = p;
        q.head = p;
        q.len ++;
        return;
    }
    q.end->next = p;
    p->front = q.end;
    q.end = p;
    q.len++;
}

void del_queue(queue &q,int &v){
    if(q.len==0){
        cout<<"队列已空！";
        return;
    }
    qnode *p = new qnode;
    p = q.head;
    q.head = q.head->next;
    v=p->data;
    delete p;
    q.len--;
}

void topsort(ADgraph &g,queue &q){
    for(int i = 0; i< g.vexnum; i++){
        if(indegree[i] == 0){
            en_queue(q,i);
        }
    }

    int count = 0;

    while(q.len != 0){
        int v;
        del_queue(q,v);
        count ++;
        cout<<g.vertices[v].data<<" ";
        for(Anode *p = g.vertices[v].firstarc;p;p=p->nextarc){
            int w = p->adjnum;
            indegree[w]--;
            if (indegree[w] == 0){
                en_queue(q,w);
            }
        }
    }
    if (count < g.vexnum){
            cout<<"图有环";
        }
}

void Topsort(ADgraph &g){
    queue q;
    init_queue(q);
    topsort(g,q);
}

typedef struct{
    int len;
    node *top,*bottom;
}stack;

void init_stack(stack &s){
    s.len = 0;
    node *p = new node;
    p->front = NULL;
    p->next = NULL;
    s.top =p;
    s.bottom = p;
}

void push(stack &s,int v){
    node *p = new node;
    p->data = v;
    s.top->front = p;
    p->next = s.top;
    s.top = p;
    s.len++;
}

void pop(stack &s,int &v){
    node *p = s.top;
    if(s.len != 0){
        s.top = s.top->next;
    }
    v=p->data;
    s.len--;
    delete p;
}

void bfs(ADgraph &g,queue &q){
    for(int i = 0; i < g.vexnum;i ++){
        if(!visited[i]){
            visited[i] = true;
            en_queue(q,i);
        }
        while(q.len != 0){
            int v;
            del_queue(q,v);
            cout<<g.vertices[v].data<<"->";
            for(Anode *p = g.vertices[v].firstarc;p;p=p->nextarc){
                int w = p->adjnum;
                if(!visited[w]){
                    visited[w]=true;
                    en_queue(q,w);
                }
            }
        }
    }
}

void BFS(ADgraph &g){
    queue q;
    init_queue(q);
    bfs(g,q);
}

int main(){
    ADgraph g;
    create_graph(g);
    DFS(g);
    //BFS(g);
    //Topsort(g);
}