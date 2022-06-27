import os
import sys
import traceback

import six
from flask import Flask, jsonify
from flask_.error import InvalidUsage
from helpers.common import config_str_to_obj
from werkzeug.utils import import_string

DEFAULT_BP_MODULES = (
    "view",
    "controller",
)
string_types = six.string_types


class EmptyConfig:
    def __init__(self, **kwargs):
        for key, value in kwargs.items():
            setattr(self, key.upper(), value)


class ApiInit(Flask):
    def configure(self, config, blueprints=None):

        config = config or EmptyConfig()
        if isinstance(config, string_types):
            self.config.from_pyfile(config)
            config = __import__(config)
        else:
            self.config.from_object(config)

        self.config.from_envvar("FLASK_CONFIG", True)
        blueprints = blueprints or getattr(config, "BLUEPRINTS", [])
        sys.path.insert(0, os.path.join(os.path.abspath("."), "flask_/blueprints/"))

        self.add_blueprint_list(blueprints)

    def add_blueprint(self, name, kw):
        for module in self.config.get("BP_MODULES", DEFAULT_BP_MODULES):
            try:
                __import__(f"{name}.{module}", fromlist=["*"])
            except ImportError:
                self.logger.warning(f"Could not import {module} for {name}")
            except Exception:
                self.logger.error(f"Error importing {module} for {name}")
                self.logger.error(traceback.format_exc())
        blueprint = import_string(f"flask_.blueprints.{name}.view._app")
        self.register_blueprint(blueprint, **kw)

    def add_blueprint_list(self, blueprints):
        for blueprint in blueprints:
            name = blueprint
            kw = dict()
            kw.update(dict(url_prefix=f"/{name}"))
            self.add_blueprint(name, kw)

    def setup(self):
        self.configure_error_handlers()
        self.configure_extensions()

    def configure_error_handlers(self):
        @self.errorhandler(InvalidUsage)
        def handle_invalid_usage(error):

            response = jsonify(error.to_dict())
            response.status_code = error._status_code
            return response

        @self.errorhandler(404)
        def page_not_found(error):
            return jsonify({"error": "Not found."}), 404

        @self.errorhandler(405)
        def method_not_allowed_page(error):
            return jsonify({"error": "Method not allowed."}), 405

    def configure_extensions(self):
        for ext_path in self.config.get("EXTENSIONS", []):
            try:
                ext = import_string(f"flask_.extension.{ext_path}")
            except ImportError:
                raise ImportError(ext_path)

            try:
                init_kwargs = import_string(f"{ext_path}_init_kwargs")()
            except ImportError:
                init_kwargs = dict()

            init_fnc = getattr(ext, "init_app", False) or ext
            init_fnc(self, **init_kwargs)


def factory(config, app_name, blueprints=None):

    app = ApiInit(app_name)
    config = config_str_to_obj("flask_.config", config)
    app.configure(config, blueprints)
    app.setup()
    return app
