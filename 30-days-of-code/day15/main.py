#!/usr/bin/env python3

class Node:
    def __init__(self,data):
        self.data = data
        self.next = None 
class Solution: 
    def display(self,head):
        current = head
        while current:
            print(current.data,end=' ')
            current = current.next

    def insert(self,head,data):
        #Complete this method
        if head == None:
            newn = Node(data)
            head = newn
            self.tmp = head
        else:
            while(head.next!=None):
                head = head.next
            head.next = Node(data)
        return self.tmp

mylist= Solution()
T=int(input())
head=None
for i in range(T):
    data=int(input())
    head=mylist.insert(head,data)    
mylist.display(head);