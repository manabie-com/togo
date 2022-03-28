from database.models.queries import DBSession
from datetime import datetime as dt


class HandlerTask:
    def __init__(
        self,
        db: DBSession,
    ):
        self._db = db

    def add_and_assign_task(self, task):
        """

        :rtype: bool
        """
        user_id = task.get("user_id")
        date = dt.strptime(task.get("date"), "%d/%m/%Y")
        if self.ability_caculator(user_id, date):
            self.assign_task(task)
            return True
        return False

    def ability_caculator(self, user_id, date):
        user = self._db.get_user_by_id(user_id)
        assignment = self._db.get_assignment_by_date(user_id, date)
        if user and user.ability > len(assignment.tasks):
            return True
        return False

    def assign_task(self, task):
        self._db.add(task)
        self._db.remove_session()
