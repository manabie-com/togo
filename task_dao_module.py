from sql_util import SqliteUtil


class Task():
    def __init__(self, table):
        self.table = table
        self.conn_db = SqliteUtil()

    def validate_data(self, data):
        if(data.get('id') is None or data.get('id') == ''):
            raise ValueError('id is invalid')
        if(data.get('task_name') is None or data.get('task_name') == ''):
            raise ValueError('task_name is invalid')
        if(data.get('project_id') is None or data.get('project_id') == ''):
            raise ValueError('project_id is invalid')
        if(data.get('created_at') is None or data.get('created_at') == ''):
            raise ValueError('created_at is invalid')
        if(data.get('updated_at') is None or data.get('updated_at') == ''):
            raise ValueError('updated_at is invalid')
        
    def insert_task(self, data):
        try:
            self.validate_data(data=data)
            sql_insert = f"INSERT INTO {self.table} VALUES('{data.get('id')}', '{data.get('task_name')}', '{data.get('project_id')}', '{data.get('created_at')}', '{data.get('updated_at')}')"
            self.conn_db.execute_sql(sql=sql_insert)
            print('Insert task success')
        except ValueError as e:
            print('Cannot insert task cause ', e)
    
    def get_all_tasks(self):
        try:
            list_tasks = self.conn_db.get_all(self.table)
            return list_tasks
        except ValueError as e: 
            print('Cannot get all tasks cause ', e)

    def execute_query(self, sql):
        try:
            self.conn_db.execute_sql(sql)
        except ValueError as e:
            print('Cannot execute query tasks cause ', e)