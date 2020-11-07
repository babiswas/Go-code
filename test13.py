class Graph:
   def __init__(self,vertex):
       self.vertex=vertex
       self.graph=[]

   def iscyclic(self):
       parent=[-1]*self.vertex
       for i in self.graph:
           x=self.find(parent,i[0])
           y=self.find(parent,i[1])
           if x==y:
              print("The graph has cycle")
              return True
           self.union(parent,x,y)
       return True

   def add_edges(self,u,v):
       self.graph.append((u,v))

   def union(self,parent,x,y):
       s1=self.find(parent,x)
       s2=self.find(parent,y)
       if s1!=s2:
          parent[s2]=s1
       
   def find(self,parent,i):
      while parent[i]!=-1:
        i=parent[i]
      return i

if __name__=="__main__":
   graph=Graph(4)
   graph.add_edges(0,2)
   graph.add_edges(1,2)
   graph.add_edges(0,1)
   graph.add_edges(2,3)
   print(graph.iscyclic())

