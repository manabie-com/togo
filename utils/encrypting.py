from hashids import Hashids
from constants import Encryption


hashids = Hashids(salt=Encryption.SECRET_SALT, min_length=Encryption.HASH_LENGTH)


def encrypt(id):
    return hashids.encode(id)


def decrypt(cipher_id):
    return hashids.decode(cipher_id)
