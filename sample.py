class Bankaccount: 
    def __init__(self): 
        self.balance=0
        print("Hello!!!Welcome to Deposit&Withdrawl Machine") 
  
    def deposit(self): 
        amount=float(input("Enter amount to be Deposited: ")) 
        self.balance += amount 
        print("\n Amount Deposited:",amount) 
  
    def withdraw(self): 
        amount = float(input("Enter amount to be Withdrawed: ")) 
        if self.balance>=amount: 
            self.balance-=amount 
            print("\n You Withdrawed:", amount) 
        else: 
            print("\n Insufficient balance  ") 
  
    def display(self): 
        print("\n Net Available Balance=",self.balance) 
  
# Driver code 
   
# creating an object of class 
s = Bankaccount() 
   
# Calling functions with that class object 
s.deposit() 
s.withdraw() 
s.display() 
