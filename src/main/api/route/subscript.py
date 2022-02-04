from flask import Blueprint

from ..middleware import credentials_validation
from ..controller.subscript import subscript, get_pricing_level

subscription_route = Blueprint('subscription-route', __name__, url_prefix='/subscription')


@subscription_route.route("/<pricing_id>", methods=["POST"])
@credentials_validation
def subscription_route_charge(pricing_id, **kwargs):
    payload = kwargs.get("payload")
    user_id = payload.get("userId")
    invoices = subscript(user_id, pricing_id)
    return {"message": f"Invoices Created '{invoices.get('invoice_id')}'"}


@subscription_route.route("/", methods=["GET"])
@credentials_validation
def subscription_route_get_pricing( **kwargs):
    payload = kwargs.get("payload")
    levels = get_pricing_level()
    print(levels)
    return {"data": levels}
