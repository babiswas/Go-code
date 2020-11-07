class Graph:
   def __init__(self,vertex):
       self.vertex=vertex       
       self.graph=[]

   def mst(self,parent,src):
       for i in range(self.vertex):
          print(f"{parent[i]}-->{i}-->{self.graph[parent[i]][i]}")

   def add_edges(self,u,v,w):
       self.graph.append((u,v,w))

   def find_minm_key(self,visited,key):
       min=INT_MAX
       setmin=0
       for i in range(self.vertex):
          if not visited[i] and key[i]<min:
             min=key[i] 
             setmin=i
       return setmin
       
   def prims_algo(self,src):
       parent=[-1]*self.vertex
       key=[float("inf")]*self,vertex
       key[src]=0
       visited=[False]*self.vertex
       for i in range(self.vertex-1):
          index=self.find_minm_key(visited,key)
          for i in range(self.vertex):
             if visited[i]==True and key[i]>self.graph[index][i] and self.graph[index][i]!=0:
                key[i]=self.graph[index][i]
                parent[i]=index
       self.mst(parent)

if __name__=="__main__":
   graph=Graph(5)
   graph.graph=[[0,2,0,6,0],[2,0,3,8,5],[0,3,0,0,7],[6,8,0,0,9],[0,5,7,9,0]]
   graph.prims_algo(0)
   
       