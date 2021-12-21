from sql_util import SqliteUtil

class User():
    def __init__(self, table):
        self.table = table
        self.conn_db = SqliteUtil()

    def validate_data(self, data):
        if(data.get('id') is None or data.get('id') == ''):
            raise ValueError('id is invalid')
        if(data.get('username') is None or data.get('username') == ''):
            raise ValueError('username is invalid')
        if(data.get('created_at') is None or data.get('created_at') == ''):
            raise ValueError('created_at is invalid')
        if(data.get('updated_at') is None or data.get('updated_at') == ''):
            raise ValueError('updated_at is invalid')
        
    def insert_user(self, data):
        try:
            self.validate_data(data=data)
            sql_insert = f"INSERT INTO {self.table} VALUES('{data.get('id')}', '{data.get('username')}', \
             '{data.get('created_at')}', '{data.get('updated_at')}')"
            self.conn_db.execute_sql(sql=sql_insert)
            print('Insert user success')
        except ValueError as e:
            print('Cannot insert user cause ', e)
        


