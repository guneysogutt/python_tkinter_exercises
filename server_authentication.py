import serial

class TcpServerAuthenticator:
    def __init__(self):
        self.port: serial.Serial
        # Create port connection
        try:
            self.port = serial.Serial("/dev/cu.usbserial-0001",115200,timeout=1)    
        except serial.SerialException as e:
            print("Error: ",e)


    def send_request(self,command:str):
        # Send the message from server
        req = str.encode(command)
        self.port.write(req)

    def recieve_answer(self) -> bool:
        # Get the authorization status
        self.send_request("get_auth")
        result = ''
        # Wait for the response
        while result == '':
            result = self.read_ser(1)
        return result == '1'

    def read_ser(self,num_of_char = 1) -> str:
        # Read data from seial port
        result = self.port.read_all()
        return result.decode()

    def close_connection(self):
        self.port.close()        

