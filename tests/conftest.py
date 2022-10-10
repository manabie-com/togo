import pytest
import pytest
from flask_login import AnonymousUserMixin
from flask_login import login_user, login_manager
from app import create_app, db
from app import init_routes
from app import Users
from werkzeug.security import generate_password_hash


@pytest.fixture
def client():
    main_app = create_app()
    main_app.config["TESTING"] = True
    main_app.config["SECRET_KEY"] = "secret_key"

    # create in memory db
    main_app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite://"
    main_app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = False
    main_app.config["TOKEN_REQUIRED"] = False
    main_app.config["LOGIN_DISABLED"] = True

    client = main_app.test_client()
    with main_app.app_context():
        db.create_all()

        # create sample user for the database
        hashed_password = generate_password_hash("sample_password", method="sha256")
        new_user = Users(name="sample_user", limit_per_day=1,
                         password=hashed_password)
        db.session.add(new_user)
        db.session.commit()

        # initialize the routes for the app
        init_routes(main_app)

    yield client
