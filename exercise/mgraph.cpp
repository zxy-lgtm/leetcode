#include<iostream>
#include<math.h>

using namespace std;

const int MAX_VERTEX_NUM=20;

typedef enum {DG,DN,AG,AB}GraphKind;

typedef struct ArcCell{
    int adj;
    string *info;
}ArcCell,AdjMatrix[MAX_VERTEX_NUM][MAX_VERTEX_NUM];

typedef struct{
    int vexs[MAX_VERTEX_NUM];
    AdjMatrix arcs;
    int vexnum,arcnum;
    GraphKind kind;
}MGraph;

typedef struct node{
    int data;
    struct node *next,*front;
}node,*n;

typedef struct{
    int len;
    node *head,*tail;
}queue;

void init_queue(queue &q){
    q.len = 0;
    node *p = new node;
    p->front = NULL;
    p->next = NULL;
    q.head = p;
    q.tail = p;
}

void en_queue(queue &q,int v){
    node *p = new node;
    p->data = v;
    p->front = q.tail;
    q.tail->next = p;
    q.tail = p;
    q.len++;
}

void del_queue(queue &q,int &v){
    node *p = new node;
    if (q.head != q.tail){
        p = q.head->next;
        q.head = q.head->next;
    }
    v=p->data;
    q.len--;
}

bool QueueEmpty_Sq(queue &q){
    return q.head == q.tail;
}

bool visited[MAX_VERTEX_NUM];

void bfs(MGraph &g,queue &q){
    for(int i = 0; i < g.vexnum; i++){
        if(!visited[i]){
            visited[i]=true;
            en_queue(q,i);
        }

        while(!QueueEmpty_Sq(q)){
            int w ;
            del_queue(q,w);
            cout<<g.vexs[w]<<"->";
            for(int j = 0; j < g.vexnum; j++){
                if(g.arcs[i][j].adj && !visited[j]){
                    visited[j] = true;
                    en_queue(q,j);
                }
            }
        }

    }
}

void BFS(MGraph &G){
    queue Q;
    init_queue(Q);
    bfs(G,Q);
}

int location(MGraph &g,int v){
    for(int i = 0; i < g.vexnum; i ++){
        if(g.vexs[i] == v){
            return i;
        }
    }
    return -1;
}

void create_graph(MGraph &g){
    cin>>g.vexnum;
    for(int i = 0; i<g.vexnum;i++){
        cin>>g.vexs[i];
    }
    for(int i = 0; i < g.vexnum; i ++){
        for(int j = 0; j < g.vexnum; j++){
            g.arcs[i][j].adj = 0;
        }
    }
    cin>>g.arcnum;
    for(int i = 0; i < g.arcnum; i++){
        int v1,v2;
        cin>>v1>>v2;
        int p1 = location(g,v1); int p2 = location(g,v2);
        g.arcs[p1][p2].adj = 1;

    }
}

void print_graph(MGraph &g){
    for(int i = 0; i < g.vexnum; i ++){
        for(int j = 0; j < g.vexnum; j++){
            cout<<g.arcs[i][j].adj<<" ";
        }
        cout<<endl;
    }
}

void dfs(MGraph &g,int v){
    visited[v] = true;
    cout<<g.vexs[v]<<"->";
    for(int i = 0; i< g.vexnum; i ++){
        for(int j = 0; j < g.vexnum; j++){
            if(g.arcs[i][j].adj && !visited[j]){
                dfs(g,j);
            }
        }
    }
}

void DFS(MGraph &g){
    for(int i = 0; i<g.vexnum;i++){
        if(!visited[i]){
            dfs(g,i);
        }
    }
}
int main(){
    MGraph g;
    create_graph(g);
    print_graph(g);
    //BFS(g);
    DFS(g);

}