import tkinter as tk
from tkinter import messagebox
import server_authentication as auth
from helper import Helper as helper
import config


class Login(tk.Tk):
    def __init__(self,title,size):
        # Base setup
        super().__init__()
        self.title(title)
        self.geometry(f'{size[0]}x{size[1]}')
        self.minsize(size[0],size[1])
        self.resizable(False,False)

        #widgets
        self.create_widgets()
        self.is_authorized = False
        self.auth_res = False


    def create_widgets(self):
        
        title = tk.Label(text='BrandIT\nAuth Server',font=('Arial',40)).pack(pady=25)
        label = tk.Label(text="username").pack()
        self.username_entry = tk.Entry(self)
        self.username_entry.pack()
        label = tk.Label(text="password").pack()
        self.password_entry = tk.Entry(self,show="*")
        self.password_entry.pack()

        # Button setup
        self.btn_login = tk.Button(self,text='Login', command=self.login).pack(pady=5)


    def login(self):
        self.authenticator = auth.TcpServerAuthenticator()
        is_authorized = self.authenticator.recieve_answer()
        is_user_found = self.check_user()
        if is_authorized and is_user_found:
            messagebox.showinfo(title="Result", message="Succesfully logged in!")


    def check_user(self):
        # Get username and password
        username = self.username_entry.get()
        password = self.password_entry.get()

        hashed_password = helper.hashKey(password)
        connection = config.cursor

        query = f"SELECT id FROM users WHERE username='{username}' AND password='{hashed_password}';"
        connection.execute(query)
        query_result = connection.fetchall()
        connection.reset()

        # If the info is matched
        return len(query_result) > 0