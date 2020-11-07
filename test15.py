class Subset:
   def __init__(self,parent,rank):
       self.parent=parent
       self.rank=rank

class Graph:
   def __init__(self,vertex):
       self.vertex=vertex
       self.graph=[]

   def krushkal(self):
       Edges=[]
       parent=[Subset(i,0) for i in range(self.vertex)]
       def getkey(a):
          return a[2]
       self.graph=sorted(self.graph,key=getkey)
       for i in self.graph:
           s1=self.find(parent,i[0])
           s2=self.find(parent,i[1])
           if s1!=s2:
              Edges.append(i)
              self.union(parent,s1,s2)
       return Edges

   def add_edges(self,u,v,w):
       self.graph.append((u,v,w))

   def find(self,parent,i):
       while parent[i].parent!=i:
          i=parent[i].parent
       return i

   def union(self,parent,x,y):
       s1=self.find(parent,x)
       s2=self.find(parent,y)
       if parent[s1].rank>parent[s2].rank:
          parent[s1].parent=s2
          parent[s2].rank+=1
       elif parent[s2].rank<parent[s1].rank:
          parent[s2].parent=s1
          parent[s1].rank+=1
       else:
          parent[s1].parent=s2
          parent[s2].rank+=1

if __name__=="__main__":
   graph=Graph(9)
   graph.add_edges(0,1,4)
   graph.add_edges(0,7,8)
   graph.add_edges(1,7,11)
   graph.add_edges(1,2,8)
   graph.add_edges(7,6,1)
   graph.add_edges(8,6,6)
   graph.add_edges(7,8,7)
   graph.add_edges(2,8,2)
   graph.add_edges(6,5,2)
   graph.add_edges(2,5,4)
   graph.add_edges(2,3,7)
   graph.add_edges(3,5,14)
   graph.add_edges(5,4,10)
   graph.add_edges(3,4,9)
   print(graph.krushkal())

   
