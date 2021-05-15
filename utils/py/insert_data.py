import requests
import json

def load_poets_data(file):
    with open(file, 'r') as f:
        data = f.read()

    js_data = json.loads(data)
    return js_data


def insert_data(poet):
    body = {
        "descb": poet["desc"],
        "dynasty": "宋",
        "poet": poet["name"]
    }
    try:
        requests.post("http://127.0.0.1:8080/v1/poet", json.dumps(body))
    except Exception as e:
        print(e)

file  = '/Users/lex/code/chinese-poetry-master/json/authors.song.json'
js_data = load_poets_data(file)
for poet in js_data:
    insert_data(poet)