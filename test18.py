class Cell:
   def __init__(self,x,y,dist):
       self.x=x
       self.y=y
       self.dist=dist


def isvalid(x,y,N):
   if x<0 or x>=N:
      return False
   elif y<0 or y>=N:
      return False
   return True


def min_steps(pos,target,N):
    dx=[2,2,-2,-2,1,-1,1,-1]
    dy=[1,-1,1,-1,2,2,-2,-2]
    visited=[[False for i in range(N)] for i in range(N)]
    queue=[]
    queue.append(Cell(pos[0],pos[1],0))
    visited[pos[0]][pos[1]]=True
    while queue:
        m=queue.pop(0)
        if m.x==target[0] and m.y==target[1]:
          return m.dist
        for i in range(8):
           X=m.x+dx[i]
           Y=m.y+dy[i]
           if isvalid(X,Y,N) and visited[X][Y]==False:
             queue.append(Cell(X,Y,m.dist+1))
             visited[X][Y]=True

if __name__=="__main__":
   pos=(0,3)
   target=(4,2)
   print(min_steps(pos,target,6))


    