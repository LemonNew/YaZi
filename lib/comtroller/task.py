from lib.core.common import get_script_path
from lib.core.data import info
from lib.comtroller.comtroller import run_poc,load_poc


class Task:
    def __init__(self,info):
        self.url = info.url
        self.type = info.type
        self.poc_name = info.script
        self.poc_path = get_script_path() + self.type
        self.poc = load_poc(self.poc_name, [self.poc_path])


    def info(self):
        print("\033[33m[*] The target url: {} \033[0m".format(self.url))
        print("\033[33m[*] The target type: {} \033[0m".format(self.type))
        print("\033[33m[*] Try to load the script: {} \033[0m".format(
            self.poc_name))

    def run(self):

        if info.type is not None:
            if info.script is not None:
                if run_poc(self.poc, self.url):
                    self.status = "Success"
                else:
                    self.status = "Fail"