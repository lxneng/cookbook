"""`as_complete` will yield the first task to complete, followed by the next quickest, and the next until all tasks are completed."""

import asyncio
from random import randrange


async def foo(n):
    s = randrange(10)
    print(f"{n} will sleep for: {s} seconds")
    await asyncio.sleep(s)
    return f"{n}!"


async def main():
    counter = 0
    tasks = [foo("a"), foo("b"), foo("c")]

    for future in asyncio.as_completed(tasks):
        n = "quickest" if counter == 0 else "next quickest"
        counter += 1
        result = await future
        print(f"the {n} result was: {result}")


asyncio.run(main())


"""
OUTPUT

(base) eric@pop-os:~/Sync/cookbook/Python$ python asyncio-as-completed.py 
a will sleep for: 0 seconds
b will sleep for: 5 seconds
c will sleep for: 6 seconds
the quickest result was: a!
the next quickest result was: b!
the next quickest result was: c!
"""