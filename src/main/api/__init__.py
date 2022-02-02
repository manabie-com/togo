from flask import Blueprint

api = Blueprint('api', __name__, url_prefix='/v1')
# api.register_blueprint(uptime_route)
# api.register_blueprint(server_cost_route)
# api.register_blueprint(ip_counting_route)
