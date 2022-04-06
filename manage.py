import os

from flask_script import Manager

from app import create_app, db

app = create_app(os.environ.get("FLASK_CONFIG") or "default")

manager = Manager(app)

@manager.command
def recreate_db():
    db.drop_all()
    db.create_all()

@manager.command
def drop_db():
    db.drop_all()

if __name__ == "__main__":
    manager.run()