# -*- encoding:"utf-8"-*-


from lib.comtroller.task import Task
from lib.flag.flag import flag
from lib.core.data import info



def yazi():
    try:
        # banner()
        info.update(flag().__dict__)
        task =Task(info)
        # task.info()
        task.run()

    except Exception as e:
        print(e)
        exit()


if __name__ == "__main__":
    yazi()