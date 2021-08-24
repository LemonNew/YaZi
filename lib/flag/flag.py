import argparse


def flag():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "-u", "--url", required=True, help="Input a target url")
    parser.add_argument("-t", "--type", required=False, help="e.g. phpcms")
    parser.add_argument("-s", "--script", required=False, help="Select script")
    parser.add_argument("-c", "--check", required=False, help="check yes or no?")
    args = parser.parse_args()

    return args
