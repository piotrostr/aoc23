class Solution:
    def __init__(self, input_path: str):
        with open(input_path, "r+") as f:
            self.input_contents = f.readlines()

    def main(self):
        raise NotImplementedError
