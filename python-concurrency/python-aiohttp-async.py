import asyncio
import aiohttp
import time

# python aiohttp async average time for  1000  HEAD requests
# 22.863216876983643 s

URL = 'http://example.com'
ITERATION_CNT = 1
REQUEST_CNT = 1000


async def test_head():
    async with aiohttp.ClientSession() as session:
        for i in range(REQUEST_CNT):
            await session.head(url=URL)


if __name__ == "__main__":
    total = 0
    for i in range(ITERATION_CNT):
        start = time.time()
        loop = asyncio.get_event_loop()
        loop.run_until_complete(test_head())
        total += time.time() - start
    print("python aiohttp async average time for ", REQUEST_CNT, " HEAD requests")
    print(total / ITERATION_CNT, 's')
