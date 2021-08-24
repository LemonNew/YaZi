import os



def get_script_path():
    return init_path()


def init_path():
    return os.getcwd() +"/scripts/"