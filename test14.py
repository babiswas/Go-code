class Subset:
   def __init__(self,parent,rank):
       self.parent=parent
       self.rank=rank
class Graph:
   def __init__(self,vertex):
       self.vertex=vertex
       self.graph=[]

   def add_edges(self,u,v):
       self.graph.append((u,v))

   def iscyclic(self):
       parent=[Subset(i,0) for i in range(self.vertex)]
       for i in self.graph:
          x=self.find(parent,i[0])
          y=self.find(parent,i[1])
          if x==y:
             print("The graph has cycle")
             return True
          self.union(parent,x,y)
       return True

   def find(self,parent,i):
      while parent[i].parent!=i:
          i=parent[i].parent
      return i

   def union(self,parent,x,y):
       s1=self.find(parent,x)
       s2=self.find(parent,y)
       if parent[s1].rank>parent[s2].rank:
          parent[s2].parent=s1
          parent[s1].rank+=1
       elif parent[s2].rank>parent[s1].rank:
          parent[s1].parent=s2
          parent[s2].rank+=1
       else:
          parent[s2].parent=s1
          parent[s1].rank+=1

if __name__=="__main__":
   graph=Graph(4)
   graph.add_edges(0,2)
   graph.add_edges(0,1)
   graph.add_edges(2,3)
   print(graph.iscyclic())

