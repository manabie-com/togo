from sql_util import SqliteUtil

class TaskDaily():
    def __init__(self, table):
        self.table = table
        self.conn_db = SqliteUtil()

    def validate_data(self, data):
        if(data.get('id') is None or data.get('id') == ''):
            raise ValueError('id is invalid')
        if(data.get('project_id') is None or data.get('project_id') == ''):
            raise ValueError('project_id is invalid')
        if(data.get('task_id') is None or data.get('task_id') == ''):
            raise ValueError('task_id is invalid')
        if(data.get('task') is None or data.get('task') == ''):
            raise ValueError('task is invalid')
        if(data.get('owner') is None or data.get('owner') == ''):
            raise ValueError('owner is invalid')
        if(data.get('created_at') is None or data.get('created_at') == ''):
            raise ValueError('created_at is invalid')
        if(data.get('updated_at') is None or data.get('updated_at') == ''):
            raise ValueError('updated_at is invalid')
        # if(data.get('deadline_at') is None or data.get('deadline_at') == ''):
        #     raise ValueError('deadline_at is invalid')
        
    def insert_task_daily(self, data):
        try:
            self.validate_data(data=data)
            sql_insert = "INSERT INTO {} VALUES('{}', '{}', '{}', '{}', '{}', '{}', '{}', '{}')" \
            .format(self.table, data.get('id'), data.get('project_id'), data.get('task_id'), data.get('task'), data.get('owner'), data.get('created_at'), data.get('updated_at'), data.get('deadline_at'))
            self.conn_db.execute_sql(sql=sql_insert)
            print('Insert task daily success')
        except ValueError as e:
            print('Cannot insert task daily cause ', e)

    def get_all_tasks_daily(self):
        try:
            list_tasks = self.conn_db.get_all(self.table)
            return list_tasks
        except ValueError as e:
            print('Cannot get all tasks daily cause ', e)

    def execute_query(self, sql):
        try:
            self.conn_db.execute_sql(sql)
        except ValueError as e:
            print('Cannot execute query tasks daily cause ', e)




