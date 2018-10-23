import asyncio
import requests
import time

# python requests async average time for  100  HEAD requests
# 2.8183603286743164 s

URL = 'http://example.com'
ITERATION_CNT = 1
REQUEST_CNT = 100


async def test_head():
    futures = [
        loop.run_in_executor(
            None, requests.head, URL
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
    print("python requests async average time for ", REQUEST_CNT, " HEAD requests")
    print(total / ITERATION_CNT * 10, 's')