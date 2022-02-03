import datetime
from jwt import encode, decode, get_unverified_header, DecodeError
from src.main.util import load_config


def extract_token(token: str) -> dict:
    """
    This function is used to validate jwt token.
    Currently, only checking few information and headers.
    if all metadata is correct, return payload in the token
    :param token:
    :return:
    """
    private_key = load_config().get("APP", {}).get("jwt_secret", None)
    if private_key is None:
        raise Exception("Please provide jwt secret in config file")
    header = get_unverified_header(token)
    EXP_DATE = datetime.datetime(year=2023, month=1, day=1).timestamp()
    exp = header.get('exp')
    if abs(EXP_DATE - float(exp)) > 1e-6:
        raise DecodeError(f"Token `{token}` is invalid")
    payload = decode(token, key=private_key, algorithms=['HS256'])
    return payload


def create_token(payload: dict) -> str:
    """
    This function is used to create jwt token based on some information.
    Currently, this function create jwt token based on secret key & payload(both from configuration file).
    Further improvement: Create token based on given payload, store it to db.
    :return: jwt token encoded by hash-256
    :rtype: str
    """
    private_key = load_config().get("APP", {}).get("jwt_secret", None)
    if private_key is None:
        raise Exception("Please provide jwt secret in config file")

    EXP_DATE = datetime.datetime(year=2023, month=1, day=1).timestamp()
    return encode(
        key=private_key,
        payload=payload,
        algorithm="HS256",
        headers={
            "exp": EXP_DATE,  # no create expired date
        }
    )
