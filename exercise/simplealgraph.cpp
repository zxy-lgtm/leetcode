#include<iostream>
#include<stdio.h>
#define MAX_VERTEX_NUM 20

using namespace std;

typedef int InfoType;
typedef int VertexType;

typedef struct ArcNode{
    int adjvex;
    struct ArcNode *nextarc;
    InfoType *info;
}ArcNode;

typedef struct VNode{
    VertexType data;
    ArcNode *firstarc;
}VNode,AdjList[MAX_VERTEX_NUM];

typedef struct{
    AdjList vertices;
    int vexnum,arcnum;
    int kind;
}ALGraph;

bool visited[MAX_VERTEX_NUM];

int LocateVex(ALGraph &G,VertexType v){
    for(int i = 0;i < G.vexnum;i++){
        if (G.vertices[i].data == v){
            return i;
        }
    }
    return -1;
}

void CreateUDG(ALGraph &G){
    cin>>G.vexnum>>G.arcnum;

    for(int i = 0;i < G.vexnum; i++){
        cin>>G.vertices[i].data;
        G.vertices[i].firstarc=NULL;
    }

    for(int k = 0;k < G.arcnum;k++){
        int v1,v2;
        cin>>v1>>v2;
        int i =LocateVex(G, v1);int j = LocateVex(G,v2);

        ArcNode *pi = new ArcNode;
        pi->adjvex=j;
        pi->nextarc=G.vertices[i].firstarc;
        G.vertices[i].firstarc=pi;

        ArcNode *pj = new ArcNode;
        pj->adjvex=i;
        pj->nextarc=G.vertices[j].firstarc;
        G.vertices[j].firstarc=pj;
    }
}

void Print(ALGraph &G){
    for(int i = 0;i < G.vexnum; i ++){
        cout<<i<<" : "<<G.vertices[i].data<<endl;
        ArcNode *p = new ArcNode;
        p = G.vertices[i].firstarc;
        while(p->nextarc!= NULL){
            cout<<p->adjvex<<"->";
            p = p->nextarc;
        }
        cout<<p->adjvex<<endl;
    }
}

void VisitFunc(ALGraph &G, int v){
    cout<<G.vertices[v].data<<"->";
}

int FristAdjVex(ALGraph &G, int v){
    if (G.vertices[v].firstarc != NULL){
        return G.vertices[v].firstarc->adjvex;
    }

    return -1;
}

int NextAdjVex(ALGraph &G, int v,int w){
    ArcNode *p = new ArcNode;
    p = G.vertices[v].firstarc;
    while(p->nextarc != NULL){
        if(p->adjvex == w){
            return p->nextarc->adjvex;
        }

        p = p->nextarc;
    }

    return -1;
}

void DFS(ALGraph &G,int v){
    visited[v] = true;
    VisitFunc(G, v);
    for(int w = FristAdjVex(G,v);w!= -1;w=NextAdjVex(G,v,w)){
        if(!visited[w]){
            DFS(G,w);
        }
    }
}

void DFSAL(ALGraph &G, int v){
    visited[v] = true;
    cout<<G.vertices[v].data<<"->";
    for(ArcNode *p = G.vertices[v].firstarc;p;p=p->nextarc){
        int w = p->adjvex;
        if(!visited[w]){
            DFS(G,w);
        }
    }
}

int main(){
    ALGraph G;
    CreateUDG(G);
    //DFS(G,0);
    DFSAL(G,0);
}

