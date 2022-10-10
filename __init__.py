import app
from app import db

if __name__ == "__main__":
    main_app = app.create_app()
    app.init_routes(main_app)
    with main_app.app_context():
        db.create_all()
    main_app.run(debug=True)
