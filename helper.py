from hashlib import sha256


class Helper:
    def __init__():
        super().__init__()


    def hashKey(key) -> str:   
        return sha256(key.encode('utf-8')).hexdigest()