#!/usr/bin/env python
import unu
import json


def main():
    response = unu.submit(
        url='https://xtom.com/'
    )
    print(json.dumps(response, sort_keys=True, indent=2))


if __name__ == '__main__':
    main()