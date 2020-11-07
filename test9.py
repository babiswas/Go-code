from collections import defaultdict

class Graph:
   def __init__(self,vertex):
       self.vertex=vertex
       self.graph=defaultdict(list)
   def add_edges(self,u,v):
       self.graph[u].append(v)

   def is_mother_vertex_util(self,visited,u):
       visited[u]=True
       for i in self.graph[u]:
          if not visited[i]:
             self.is_mother_vertex_util(visited,i)

   def is_mother_vertex(self):
       visited=[False]*self.vertex
       m=-1
       for i in range(self.vertex):
          if not visited[i]:
             self.is_mother_vertex_util(visited,i)
             m=i
       visited=[False]*self.vertex
       self.is_mother_vertex_util(visited,m)
       if False not in visited:
          print(f"{m} is the mother vertex")
       else:
          print("No mother veretx")

if __name__=="__main__":
   graph=Graph(4)
   graph.add_edges(0,2)
   graph.add_edges(2,0)
   graph.add_edges(1,2)
   graph.add_edges(0,1)
   graph.add_edges(2,3)
   graph.add_edges(3,3)
   print("The mother vertex of the graph")
   print(graph.is_mother_vertex())


       
            