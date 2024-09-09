import serial

port = serial.Serial("/dev/cu.usbserial-0001",115200,timeout=1)

def send_request(command:str): 
    req = str.encode(command)
    port.write(req)

def recieve_answer() -> str:
    send_request("get_auth")
    result = ''
    while result == '':
        result = read_ser(1)
    return result

def read_ser(num_of_char = 1) -> str:
    result = port.read_all()
    return result.decode()
