import tkinter as tk
from tkinter import ttk


class AppGUI(tk.Tk):
    def __init__(self,title,size):
        # Main setup
        super().__init__()
        self.title(title)
        self.geometry(f'{size[0]}x{size[1]}')
        self.minsize(size[0],size[1])
        self.config(bg='black')

        # widgets
        #self.tab1 = Tab(self)     

        # run
        self.mainloop() 
        

class Tab (ttk.Notebook):
    def __init__(self,parent):
        super().__init__(parent)
        self.place(x = 300, y = 10, relwidth = 0.1, relheight = 1)