import sqlite3
from my_config import DB_CONNECTION_STR

class SqliteUtil():
    def __init__(self):  
        self.db_connection_str = DB_CONNECTION_STR
        self.con = sqlite3.connect(self.db_connection_str, check_same_thread=False)
        self.con.row_factory = sqlite3.Row
        self.cursorObj = self.con.cursor()

    def execute_sql(self, sql):
        self.cursorObj.execute(sql)
        self.con.commit()
    
    def drop_table(self, table_name):
        sql_drop = ("DROP TABLE IF EXISTS {}".format(table_name))
        self.cursorObj.execute(sql_drop)
        self.con.commit()

    def get_all(self, table):
        self.cursorObj.execute('SELECT * FROM {}'.format(table))
        rows = self.cursorObj.fetchall()
        rows_dict = [dict(row) for row in rows]
        return rows_dict

    def get_all_tables_name(self):
        self.cursorObj.execute('SELECT name from sqlite_master where type= "table"')
        all_tables = self.cursorObj.fetchall()
        return all_tables

    def count_all(self, table):
        rowcount = self.cursorObj.execute('SELECT * FROM {}'.format(table)).rowcount
        return rowcount

