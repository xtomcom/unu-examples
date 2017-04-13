import json

try:
    from urllib.parse import urlencode
    from urllib.request import urlopen, Request
except ImportError:
    from urllib import urlencode
    from urllib2 import urlopen, Request


def submit(url, keyword=None, title=None, username=None, password=None):
    payload = {'action': 'shorturl', 'format': 'json', 'url': url}
    if keyword:
        payload['keyword'] = keyword
    if title:
        payload['title'] = title
    if username and password:
        payload['username'] = username
        payload['password'] = password

    request = Request('https://u.nu/api.php', urlencode(payload))
    response = urlopen(request)
    return json.loads(response.read())
