import requests
import pytest

@pytest.fixture()
def TARGET():
    return {
        'HOST':'auth.api.test',
        'PORT': 8181
    } 
    



def test_apitest(TARGET):
    res = requests.request('GET', f'http://{TARGET["HOST"]}:{TARGET["PORT"]}/ping')
    
    print(res)

    assert res.status_code == 200

def test_signup():
    
    assert 1 == 1