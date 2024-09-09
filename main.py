import tkinter as tk
import server_authentication as auth

authenticator = auth

is_authenticated = auth.recieve_answer()

print(is_authenticated)
if is_authenticated != "1":
    print("nauuuuu")
else:
    print("eri")    



"""
root = tk.Tk()
# calculate the position of the window to center it on the screen
screen_width = root.winfo_screenwidth()
screen_height = root.winfo_screenheight()
# calculate the x-coordinate for centering  
x_coordinate = (screen_width - 320) // 2  
# calculate the y-coordinate for centering
y_coordinate = (screen_height - 320) // 2  
# set window position
root.geometry(f"320x320+{x_coordinate}+{y_coordinate}")

root.mainloop()
"""