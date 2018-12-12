import requests

def lambda_handler(event, context):
    uri = 'https://www.google.co.jp'
    res = requests.get(uri)
    print(res.status_code)
    return res.status_code
