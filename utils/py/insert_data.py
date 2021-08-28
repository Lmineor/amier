# -*- encoding:utf-8 -*-

import requests
import json

def load_poets_data(file):
    with open(file, 'r') as f:
        data = f.read()

    js_data = json.loads(data)
    return js_data


def insert_poet_data(poet, dynasty):
    body = {
        "descb": poet["desc"],
        "dynasty": dynasty,
        "poet": poet["name"]
    }
    try:
        requests.post("http://127.0.0.1:8080/v1/poet", json.dumps(body))
    except Exception as e:
        print(e)

class DataLoader(object):

    @classmethod
    def loads_data(cls, file):
        with open(file, 'r') as f:
                data = f.read()
        js_data = json.loads(data)
        return js_data

    @classmethod
    def do_request(cls, url, data):
        try:
            requests.post(url, json.dumps(data))
        except Exception as e:
            print(e)

    @classmethod
    def poet_tang(cls):
        file = '/Users/lex/code/chinese-poetry-master/json/authors.tang.json'
        tang_poet_datas = DataLoader.loads_data(file)
        for poet in tang_poet_datas:
            body = {
                "descb": poet["desc"],
                "dynasty": "唐",
                "poet": poet["name"]
            }
            DataLoader.do_request("http://127.0.0.1:8080/v1/poet", body)

    @classmethod
    def poet_song(cls):
        file = '/Users/lex/code/chinese-poetry-master/json/authors.song.json'
        tang_poet_datas = DataLoader.loads_data(file)
        for poet in tang_poet_datas:
            body = {
                "descb": poet["desc"],
                "dynasty": "宋",
                "poet": poet["name"]
            }
            DataLoader.do_request("http://127.0.0.1:8080/v1/poet", body)

    @classmethod
    def poem_tang(cls):
        for index in range(0,58000, 1000):
            file = '/Users/lex/code/chinese-poetry-master/json/poet.tang.%s.json' % index
            poem_tangs = DataLoader.loads_data(file)
            for poem in poem_tangs:
                body = {
                    "poem": poem["title"],
                    "dynasty": "唐",
                    "poet": poem["author"],
                    "paragraphs": poem["paragraphs"]
                }
                DataLoader.do_request("http://127.0.0.1:8080/v1/poem", body)

    @classmethod
    def poem_song(cls):
        for index in range(0,255000, 1000):
            file = '/Users/lex/code/chinese-poetry-master/json/poet.song.%s.json' % index
            poem_tangs = DataLoader.loads_data(file)
            for poem in poem_tangs:
                body = {
                    "poem": poem["title"],
                    "dynasty": "宋",
                    "poet": poem["author"],
                    "paragraphs": poem["paragraphs"]
                }
                DataLoader.do_request("http://127.0.0.1:8080/v1/poem", body)






# 唐诗人
DataLoader.poet_tang()
print("Done poet tang")
# 宋诗人
DataLoader.poet_song()
print("Done poet song")
# 唐诗
DataLoader.poem_tang()
print("Done poem tang")
# 宋诗
DataLoader.poem_song()
print("Done poem song")