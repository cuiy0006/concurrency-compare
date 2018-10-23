import asyncio
import urllib3
import time

# python urllib3 async average time for  1000  HEAD requests
# 0.9386083841323852 s

URL = 'http://example.com'
ITERATION_CNT = 10
REQUEST_CNT = 1000


async def test_head():
    http = urllib3.PoolManager()
    futures = [
        loop.run_in_executor(
            None, lambda url: http.request('HEAD', url), URL
        ) for i in range(REQUEST_CNT)
    ]
    return await asyncio.gather(*futures)


if __name__ == "__main__":
    total = 0
    for i in range(ITERATION_CNT):
        start = time.time()
        loop = asyncio.get_event_loop()
        loop.run_until_complete(test_head())
        total += time.time() - start
    print("python urllib3 async average time for ", REQUEST_CNT, " HEAD requests")
    print(total / ITERATION_CNT, 's')