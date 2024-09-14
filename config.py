import mysql.connector as connector


db = connector.connect(
    host="localhost",
    user="root",
    password="123456",
    database="testDB",
)

cursor = db.cursor()