#include<iostream>
#include<algorithm>

using namespace std;

#define MAX_VEX_NUM 20
#define MAX_ARC_NUM 100

typedef struct arcnode{
    int adjnum;
    struct arcnode *nextarc;
}Arcnode;

typedef struct Vnode{
    int data;
    Arcnode * firstarc;
}ADJlist[MAX_VEX_NUM];

typedef struct{
    int vexnum,arcnum;
    ADJlist vertices;
}ALGraph;


typedef struct edge{
    int start,end,power;
}*eg,Edges[MAX_ARC_NUM];

int location(ALGraph &G,int v){
    for(int i = 0; i< G.vexnum; i ++){
        if(G.vertices[i].data == v){
            return i;
        }
    }

    return -1;
}

void create_graph(ALGraph &g,Edges &e){
    cout<<"请输入图的结点总数：";
    cin>>g.vexnum;
    cout<<"请输入每一个结点的值：";
    for(int i = 0; i < g.vexnum; i ++){
        cin>>g.vertices[i].data;
        g.vertices[i].firstarc = NULL;
    }

    cout<<"请输入边的总数：";
    cin>>g.arcnum;
    cout<<"请输入每条边以及它的权值：";
    for(int i = 0; i< g.arcnum;i++){
        int v1,v2,power;
        cin>>v1>>v2>>power;
        
        int p1 = location(g,v1); int p2 = location(g,v2);

        e[i].start = p1;e[i].end = p2;;e[i].power = power;

        Arcnode *p = new Arcnode;
        p->adjnum = p2;
        p->nextarc = g.vertices[p1].firstarc;
        g.vertices[p1].firstarc = p;

        Arcnode *q = new Arcnode;
        q->adjnum = p1;
        q->nextarc = g.vertices[p2].firstarc;
        g.vertices[p2].firstarc = q;
    }
}

bool cmp(edge e1,edge e2){
    return e1.power < e2.power;
}

void kruskal_min_tree(ALGraph &g,Edges e,edge mintree[]){
    // 为每一个顶点配置一个标记值
    int ass[g.vexnum];
    int num = 0;

    // 初始状态下，每个标记值都不同
    for(int i = 0 ; i < g.vexnum; i ++){
        ass[i] = i;
    }

    // 根据权值，对所有的边进行升序排序
    sort(e,e+g.arcnum,cmp);

    for(int i = 0; i < g.arcnum; i++){
        int start = e[i].start;
        int end = e[i].end;
        if(ass[start] != ass[end]){
            mintree[num] = e[i];

            num ++;
            int elem = ass[end];

            for(int k = 0; k < g.vexnum;k ++){
                if(ass[k] == elem){
                    ass[k] = ass[start];
                }
            }
            if(num == g.vexnum -1){
                break;
            }
        }
    }
}

void display(ALGraph &g,edge mintree[]){
    int cost = 0, i;
    cout<<"最小生成树为："<<endl;

    for(int i = 0 ; i < g.vexnum-1;i++){
        cout<<g.vertices[mintree[i].start].data<<"-"<<g.vertices[mintree[i].end].data<<"权值："<<mintree[i].power<<endl;
        cost += mintree[i].power;
    }

    cout<<"总权值为："<<cost;
}

int main(){
    ALGraph g;
    Edges e;
    create_graph(g,e);
    edge mintree[g.vexnum-1];
    kruskal_min_tree(g,e,mintree);
    display(g,mintree);
}
