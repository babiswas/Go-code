import sys
class Graph:
   def __init__(self,vertex):
       self.vertex=vertex
       self.graph=[]

   def display(self,parent):
      for i in range(self.vertex):
         print(f"{parent[i]}--->{i}--->{self.graph[parent[i]][i]}")

   def prims_algo(self,src):
       parent=[-1]*self.vertex
       visited=[False]*self.vertex
       key=[sys.maxsize]*self.vertex
       key[src]=0
       for i in range(self.vertex-1):
          index=self.find_minm(key,visited)
          visited[index]=True
          for i in range(self.vertex):
             if key[i]>self.graph[index][i] and self.graph[index][i]>0 and visited[i]==False:
                key[i]=self.graph[index][i]
                parent[i]=index
       self.display(parent)
          
       
   def add_edges(self,u,v,w):
       self.graph.append((u,v,w))

   def find_minm(self,key,visited):
       min=sys.maxsize
       minnode=0
       for i in range(self.vertex):
          if key[i]<min and visited[i]==False:
             min=key[i]
             minnode=i
       return minnode

if __name__=="__main__":
   graph=Graph(5)
   graph.graph=[[0,2,0,6,0],[2,0,3,8,5],[0,3,0,0,7],[6,8,0,0,9],[0,5,7,9,0]]
   print("Displaying prims algorithm")
   graph.prims_algo(0)


   
